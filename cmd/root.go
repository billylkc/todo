package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Global variable
var (
	cfgFile  string
	todoFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Get things done.",
	Long:  `Get things done.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	var err error
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")
	rootCmd.PersistentFlags().StringVar(&todoFile, "todo", "", "todo file (default is $HOME/todo.txt)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	todoFile, err = checkFilePath()
	if err != nil {
		fmt.Println(err)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".todo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// checkFilePath checks if the todo.txt is created under root folder
// Create it if it doesnt exist
func checkFilePath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, "todo.txt")

	if _, err := os.Stat(path); os.IsNotExist(err) {

		file, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer file.Close()

		// Create file
		fmt.Printf("\nNo existing TODO file. \nNew TODO file created  - %s \n\n", path)
		return "", err
	} else {
		// Normal case
		// fmt.Printf("\nUsing TODO file  - %s \n\n", path)
	}

	return path, nil
}
