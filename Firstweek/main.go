package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
);

type Task struct { // struct where we bind the different type of data type in one object.
	 ID int
	 Description string
	 Completed bool
}

var tasks []Task; // slice which is used to add the collection dynamically

func addTask(description string){
    newID := len(tasks)+1;
	newTask := Task{
		ID: newID,
		Description:  description,
		Completed: false,
	}

	tasks = append(tasks, newTask);
	fmt.Printf("Added task - %d - %s\n",newTask.ID,newTask.Description);
}


func listOfTask(){
	fmt.Print("Your List of Tasks");

	if len(tasks) == 0{
		fmt.Print("Opps - No Task is available right away");
		return;
	}

	for _, task := range tasks{
		status := "Not Completed";
		if task.Completed {
			status = "Completed";
		}
		fmt.Printf("ID - %d | Description - %s | Status - %s \n",task.ID,task.Description,status);
	}
}

   func completedTask(Id int){
	 for i := range tasks{
		if tasks[i].ID == Id {
		tasks[i].Completed = true
		fmt.Print("Your task is completed. the task id is -",tasks[i].ID);
		return
		}
	 }
	 fmt.Print("Oops your task is not present, please add the task");
   }

   func deleteTask(Id int){

	for i := range tasks{
		if tasks[i].ID == Id{
			tasks = append(tasks[:i] ,tasks[i+1:]...)
			fmt.Print("Task is successfully deleted..");
			return;
		}
	}
	fmt.Print("Task is not present...");
   }

func main(){
	fmt.Print("Welcome to To-Do Cli!..");

	fmt.Print("Available Commond : add[Description] ,list, complete[id], delete[id],exit.");

	reader := bufio.NewReader(os.Stdin);
	for {
		fmt.Print("> ");
		input , _ := reader.ReadString('\n');
		input = strings.TrimSpace(input);

		// parts := strings.Fields(input)
		parts := strings.SplitN(input, " ", 2)
		if len(input) == 0{
			continue;
		}

		command :=  parts[0];
		
		var args []string

		switch command {
		case "add"  :
			if len(args) == 0{
				fmt.Print("Please provide the desciption for the task.")
				continue;
			}
			description := strings.Join(args, " ")
			addTask(description);

		case "list" :
			listOfTask();

		case "Completed":
			if len(args) == 0 {
				fmt.Print("Please Provide the task id ");
				continue;
			}
			id,err := strconv.Atoi(args[0])
			if err == nil {
				fmt.Print("Invalid Id. Please provide the valid id");
				continue;
			}
			completedTask(id);

		case "delete" :
			if len(args) == 0{
				fmt.Print("Oops id is not present");
				continue;
			}

			id,err := strconv.Atoi(args[0]);
			if err == nil{
				fmt.Print("Invalid id. Please provide the valid id")
				continue;
			}
			deleteTask(id);

		case "exit" : 
		     fmt.Print("Existing To-Do cli");
			 return;

		case "default":
			fmt.Print("Invalid command : please provide the valid command");

		}
	}
}