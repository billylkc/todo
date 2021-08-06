# TODO CLI

Inspired by the todo-cli package by [zwbetz-gh](https://github.com/zwbetz-gh).
This is a golang implementation of the original program. This implementation uses plain text file (default ~/todo.txt) instead of sqlite to minimize the dependencies.

Get things done via command line

## Install
   (TBA)

## Usage

```
						Usage:
todo <command>

Commands:
h, help                 Show this help
l, list [-s]            List undone todos. If -s given, list only the undone (simplified) tasks.
a, add <todo> [-d]      Add a todo. If -d mmdd is provided, the date would be the input date, instead of current date.
bk, backup [--done]     Archive the tasks to other file. Usually ~/todo.txt.bk. If done flag is provided, only the finished task would be backup.
e, edit <id> <todo>     Edit a todo. If -d mmdd is provided, the date would be the input date, instead of current date
d, done <id>            Done a todo
u, undone <id>          Undone a todo
s, summarize <id>       Summarize the man-day of todo tasks using simple weightings.
r, remove <id>          Remove a todo
wipe [--done]           Wipe all todos. If done flag is provided, only the finished task would be wiped.
```

## Examples

#### Add Task

Add task adds to the todo list. Date can also be provided to add the task with a specific date. `todo a task -d 0822`

```
$ todo a Eat Apple

  ---
  Added task - Eat Apple

  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Eat Apple

  ---


$ todo a Eat Banana -d 0807

  ---
  Added task - Eat Banana

  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Eat Apple
  2       2021-08-07      [ ]     Eat Banana

  ----
```

#### Edit Task
```
$ todo e 1 "Eat Many Apples"

  ----
  Edited:
	Eat Banana

	To be:
	  Eat Many Bananas


	  id      date            done    todo
	  --      ----            ----    ----
	  1       2021-08-06      [ ]     Eat Many Apples
	  2       2021-08-07      [ ]     Eat Banana
  ---
```

#### Done / Undone
```
$ todo d 2

  ---
  Done:
	Eat Banana

    id      date            done    todo
    --      ----            ----    ----
	1       2021-08-06      [ ]     Eat Many Apples
	2       2021-08-07      [x]     Eat Banana

  ---

$ todo u 2

  ---
  Done:
	Eat Banana

    id      date            done    todo
    --      ----            ----    ----
	1       2021-08-06      [ ]     Eat Many Apples
	2       2021-08-07      [ ]     Eat Banana

  ---
```

#### Remove
```
$ todo r 2

  ---
  Removed:
	 Eat Banana

	 id      date            done    todo
	 --      ----            ----    ----
	 1       2021-08-06      [ ]     Eat Many Apples

  ---
```

#### Wipe

```
$ todo wipe

  ---
  All todo tasks removed.

  id      date            done    todo
  --      ----            ----    ----

  ---

```

#### List

```
$ todo a Task 1
$ todo a Task 2
$ todo a Task 3
$ todo a Task 4 -d 0807
$ todo d 4
$ todo l

  ---
  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Task 1
  2       2021-08-06      [ ]     Task 2
  3       2021-08-06      [ ]     Task 3
  4       2021-08-07      [x]     Task 4

  ---

$ todo l -s

  ---
  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Task 1
  2       2021-08-06      [ ]     Task 2
  3       2021-08-06      [ ]     Task 3

  ---

```

#### Summarize
```
$ todo s

  ---
  id      time            done    todo
  --      ----            ----    ----
  1       1.0 days         [x]    Task 4
  2       0.3 days         [ ]    Task 1
  3       0.3 days         [ ]    Task 2
  4       0.3 days         [ ]    Task 3

  Total of 2 days
	  [2021-08-06 2021-08-07]

  ---

```


## Reference
Inspired by [todo-cli](https://github.com/zwbetz-gh/todo-cli).
