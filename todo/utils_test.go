package todo

import (
	"testing"
	"time"
)

func TestParseInputDate(t *testing.T) {
	type args struct {
		in string
	}
	year := time.Now().Format("2006")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Normal case",
			args:    args{"08-04"},
			want:    year + "-08-04",
			wantErr: false,
		},
		{
			name:    "mmdd",
			args:    args{"0804"},
			want:    year + "-08-04",
			wantErr: false,
		},
		{
			name:    "dd/mm",
			args:    args{"04/08"},
			want:    year + "-08-04",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInputDate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInputDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInputDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
