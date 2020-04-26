Todo is a CLI for managing your TODOs

Installation

1. Install go and setup your workspace

2. $ go get -u github.com/abhinavdwivedi440/todo


Usage:
  todo [command]

Available Commands:
  add         Adds a task to your task list.
  do          Marks a task as complete
  help        Help about any command
  list        Lists all of your tasks.

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.


Example:

$ todo add review talk proposal
Added "review talk proposal" to your task list.

$ todo add clean dishes
Added "clean dishes" to your task list.

$ todo list
You have the following tasks:
1. review talk proposal
2. clean dishes

$ todo do 1
Marked "1" as completed.

$ todo list
You have the following tasks:
1. clean dishes