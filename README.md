# TODO CLI

Inspired by the todo-cli package by [zwbetz-gh](https://github.com/zwbetz-gh).
This is a golang implementation of the original program. The implementation uses plain text file instead of sqlite to minimize the dependencies.

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
a, add <todo>           Add a todo
bk, backup [--done]     Archive the tasks to other file. Usually ~/todo.txt.bk. If done flag is provided, only the finished task would be backup.
e, edit <id> <todo>     Edit a todo
d, done <id>            Done a todo
u, undone <id>          Undone a todo
s, summarize <id>       Summarize the man-day of todo tasks using simple weightings.
r, remove <id>          Remove a todo
wipe [--done]           Wipe all todos. If done flag is provided, only the finished task would be wiped.
```

## Examples

#### Add
```
$ todo add Eat Apple

  Added task - Eat Apple

  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Eat Apple


$ todo add Eat Banana -d 0807

  Added task - Eat Banana

  id      date            done    todo
  --      ----            ----    ----
  1       2021-08-06      [ ]     Eat Apple
  2       2021-08-07      [ ]     Eat Banana
```

#### Edit Task
```
$ todo e 2 "Buy a lot of eggs"

Edited:
  Buy eggs
To be:
  Buy a lot of eggs

$ todo l

id  date        todo
--  ----        ----
1   2019-09-07  Buy milk
2   2019-09-07  Buy a lot of eggs
3   2019-09-07  Refill Millie's tennis balls

$ todo d 2

Done:
  Buy a lot of eggs

$ todo l

id  date        done  todo
--  ----        ----  ----
1   2019-09-07  [ ]   Buy milk
2   2019-09-07  [x]   Buy a lot of eggs
3   2019-09-07  [ ]   Refill Millie's tennis balls

$ todo u 2

Undone:
  Buy a lot of eggs

id  date        done  todo
--  ----        ----  ----
1   2019-09-07  [ ]   Buy milk
2   2019-09-07  [ ]   Buy a lot of eggs
3   2019-09-07  [ ]   Refill Millie's tennis balls


$ todo r 2

Removed:
  Buy a lot of eggs

id  date        todo
--  ----        ----
1   2019-09-07  Buy milk
2   2019-09-07  Refill Millie's tennis balls

$ todo wipe
Wiped todos

```

## Reference
Inspired by [todo-cli](https://github.com/zwbetz-gh/todo-cli).
