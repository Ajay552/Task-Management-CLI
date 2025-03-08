# CLI Tasks

A simple command-line interface (CLI) Task Management application written in Go. This tool allows users to add, list, delete, and mark tasks as complete using a CSV file for persistence.

## Features

- Add new tasks
- List all tasks in a tabular format
- Mark tasks as completed
- Delete tasks
- Stores tasks in a CSV file for persistence

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Ajay552/Task-Management-CLI.git
   cd cli-tasks
   ```
2. Build the project:
   ```sh
   go build -o cli-tasks
   ```
3. Run the program:
   ```sh
   ./cli-tasks
   ```

## Usage

Once the application is running, you can use the following commands:

### Add a Task
```sh
$ task add "Buy groceries"
```

### List Tasks
```sh
$ task list
```

### Mark Task as Completed
```sh
$ task complete <task_id>
```

### Delete a Task
```sh
$ task delete <task_id>
```

### Exit the CLI
```sh
$ exit
```

## File Storage
The tasks are stored in a `tasks.csv` file with the following format:
```
Id, Name, Status, Created
1, Buy groceries, false, 03:04 1/2/2024
2, Read a book, true, 10:15 2/3/2024
```

## Contributing
Contributions are welcome! Feel free to fork the repository and submit a pull request.

## License
This project is licensed under the MIT License.

## Author
Ajay S

