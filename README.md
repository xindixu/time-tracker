# Todo Time Tracker

## How to use
1. Clone the repo
2. Run `make build`
3. Link current directory to path `export PATH=/Users/xindixu/Code/todo-time-tracker:$PATH`
4. Run `ttt help` to see the available commands
```
$ ttt help
Simple CLI tool to help you keep track of the tasks to do and the time spent on a task

Usage:
  ttt [command]

Available Commands:
  add         Add a task or a list of tasks
  cleanup     Remove all completed tasks
  completion  Generate the autocompletion script for the specified shell
  current     Get the tasking that is currently being worked on
  delete      Delete a task or a list of tasks
  done        Mark a task or a list of tasks as completed
  help        Help about any command
  list        List out all of added tasks
  log         Show the log of sessions spent on tasks
  start       Start the timer for a task
  stop        Stop the timer for the current task

Flags:
  -h, --help     help for ttt
  -t, --toggle   Help message for toggle

Use "ttt [command] --help" for more information about a command.
```

### Task related commands
- Add tasks (`add`) 

  Single | Multiple
  --- | ---
  `ttt add yoga` | `ttt add -b jogging swimming`
  `ttt add yoga for 1 hr` | `ttt add -b "jogging for 1 hr" "swimming for 1 hr"`

- Mark tests as complete (`done`)

  Single | Multiple
  --- | ---
  `ttt done yoga` | `ttt done -b jogging swimming`
  `ttt done yoga for 1 hr` | `ttt done -b "jogging for 1 hr" "swimming for 1 hr"`

- Delete tasks (`delete`)

  Single | Multiple
  --- | ---
  `ttt delete yoga` | `ttt delete -b jogging swimming`
  `ttt delete yoga for 1 hr` | `ttt delete -b "jogging for 1 hr" "swimming for 1 hr"`

- List tasks (`list`)

  Incomplete Only | All
  --- | ---
  `ttt list` | `ttt list -a`


## Library used
- BoltDB
- Cobra
