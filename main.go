package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	fmt.Println("Welcome to CLI Tasks")
	initializeTaskStore()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input", err)
			continue
		}

		cmd = strings.TrimSpace(cmd)
		cmdChunks := strings.Split(cmd, " ")

		cmd = cmdChunks[0]
		if cmd == "exit" {
			os.Exit(0)
		}

		if len(cmdChunks) <= 1 {
			fmt.Println("Please enter a task method")
			return
		}

		if cmd == "task" {
			handleTask(cmdChunks)
		} else {
			fmt.Printf("%s: command not found \n", cmd)
		}
	}

}

func handleTask(cmdChunks []string) {
	if len(cmdChunks) <= 1 {
		fmt.Println("Please enter a task method")
		return
	}

	method := cmdChunks[1]

	if len(cmdChunks) <= 2 && (method == "add" || method == "delete" || method == "complete") {
		fmt.Println("Please provide a task ID or name")
		return
	}

	switch method {
	case "add":
		task := strings.Join(cmdChunks[2:], " ")
		handleAddTask(task)
	case "list":
		displayTask()
	case "delete":
		deleteTaskByid(cmdChunks[2])
	case "complete":
		handleTaskCompletedById(cmdChunks[2])
	default:
		fmt.Printf("%s: method not found \n", method)
	}
}

func initializeTaskStore() {
	fileName := "tasks.csv"

	// File already exists
	if _, err := os.Stat(fileName); err == nil {
		return
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("error initilizing db file")
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Id", "Name", "Status", "Created"}
	writer.Write(header)
}

func handleAddTask(task string) {
	if task == "" {
		fmt.Println("Task name cannot be empty")
		return
	}

	fileName := "tasks.csv"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	id := getTaskId()
	currentTime := time.Now()
	formattedTime := currentTime.Format("03:04 1/2/2006")

	data := make([]string, 4)
	data[0] = strconv.Itoa(id)
	data[1] = task
	data[2] = "false"
	data[3] = formattedTime

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(data)

	fmt.Println("Task successfully created")
}

func getTaskId() int {
	fileName := "tasks.csv"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 1
	}

	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return 1
	}

	if len(tasks) <= 1 {
		return 1
	}

	lastRow := tasks[len(tasks)-1]

	id, err := strconv.Atoi(lastRow[0])

	if err != nil {
		return 0
	}

	return id + 1
}

func displayTask() {
	fileName := "tasks.csv"
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()

	if len(tasks) <= 1 {
		fmt.Println("No tasks available.")
		return
	}

	if err != nil {
		fmt.Println("An error occurred while getting the task list")
		return
	}

	tabWriter := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	defer tabWriter.Flush()

	fmt.Fprintln(tabWriter, "ID\tTask\tStatus\tCreated")
	fmt.Fprintln(tabWriter, "--\t----\t------\t-------")
	for i, row := range tasks {
		if i == 0 {
			continue
		}
		status := "Pending"
		if row[2] == "true" {
			status = "Completed"
		}
		fmt.Fprintln(tabWriter, row[0]+"\t"+row[1]+"\t"+status+"\t"+row[3])
	}
}

func handleTaskCompletedById(taskId string) {
	fileName := "tasks.csv"

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	isUpdated := false

	for i, task := range tasks {
		if len(task) > 0 && task[0] == taskId {
			if task[2] == "true" {
				fmt.Printf("Task %s Already Completed \n", taskId)
				return
			}

			tasks[i][2] = "true"
			isUpdated = true
		}
	}

	if !isUpdated {
		fmt.Printf("Task %s not found", taskId)
		return
	}

	file, err = os.Create(fileName)

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(tasks)
	fmt.Printf("Task %s status succesfully updated \n", taskId)
}

func deleteTaskByid(taskId string) {
	fileName := "tasks.csv"

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	taskLength := len(tasks)

	tasks = deleteRecord(tasks, taskId)

	// task not found
	if taskLength == len(tasks) {
		return
	}

	file, err = os.Create(fileName)

	if err != nil {
		fmt.Println("An Error Occured")
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(tasks)

}

func deleteRecord(records [][]string, taskId string) [][]string {
	isDeleted := false
	for i, record := range records {
		if len(record) > 0 && record[0] == taskId {
			// return append(records[:i], records[i+1:]...)
			records = slices.Delete(records, i, i+1)
			isDeleted = true
			break
		}
	}

	if !isDeleted {
		fmt.Printf("Task %s not found \n", taskId)
	} else {
		fmt.Printf("Task %s deleted successfully \n", taskId)
	}

	return records
}
