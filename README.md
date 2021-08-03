# TODO CLI

Inspired by the todo-cli package by [zwbetz-gh](https://github.com/zwbetz-gh).
This is a golang implementation of the original program.

Get things done via command line

## Install
   (TBA)

## Usage

```
						Usage:
todo <command>

Commands:
h, help                 Show this help
l, list [-a]            List undone todos. If -a given, list all todos
a, add <todo>           Add a todo
e, edit <id> <todo>     Edit a todo
d, done <id>            Done a todo
u, undone <id>          Undone a todo
r, remove <id>          Remove a todo
wipe                    Wipe all todos
```

## Examples

```
$ todo a "Buy milk"

Added:
  Buy milk

$ todo a "Buy eggs"
Added:
  Buy eggs

$ todo a "Refill Millie's tennis balls"
Added:
  Refill Millie's tennis balls

$ todo l
id  date        todo
--  ----        ----
1   2019-09-07  Buy milk
2   2019-09-07  Buy eggs
3   2019-09-07  Refill Millie's tennis balls

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
