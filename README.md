inside () are how the commands should be typed. NOTE: can be changed later.
TODO list Requirements
- List all TODO's ($ todo-go) Outputs all todo's created.
- List/ADD a TODO list ($ todo-go -a) Outputs "todo-name created." if todo doesn't exist, create it. If it already exists, Output it.
- ADD task to TODO ($ todo-go -) Outputs todo with the new task.
 -a is for adding a task, -p is for specifying priority.
 NOTE: Todo list is by default is sorted by priority. (low to high, 1 to 4) (low, medium, high, very high)
- REMOVE TODO ($ todo-go todo-name -r) Outputs "todo-name removed." if no argument for -r is passed, entire todo is removed.
- REMOVE TODO task ($ todo-go todo-name -r task-name) Outputs todo without the task. argument for -r is passed, only task is removed.
- MARK task as DONE ($ todo-go todo-name -c task-name) Outputs the todo with the task-name marked as completed.
- MARK completed as uncompleted ($ todo-go todo-name -u task-name) Outputs the todo with the previous completed task
unmarked as completed. If provided task wasn't completed, Output the todo.

#### Special characters such as , # $ should be valid

Each TODO should have a unique identifier
Text of a Task should have maximum length
NOTE: recycle bin.
NOTE: Design should be simple empty squares for uncompleted task, and ? for completed.

todo-item {
id,
name,
is_completed,
is_deleted,

new_item (id, name, is_completed, is_deleted),
to_string(), // serialize to csv file
from_string(), // deserialize from csv file
}
