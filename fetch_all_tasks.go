package main

import (
	"context"
	"fmt"
)

func (a *appConfig) getAllTasks() []task{

	tasks, err := a.db.GetAllTasks(context.Background())
	if err != nil {
		fmt.Printf("Error while fetching tasks: %s\n", err)
	}
	output := []task{}

	for _, item := range tasks {
		taskit := task{
			id:   int(item.ID),
			name: item.Name,
		}
		output = append(output, taskit)
	}
	return output
}
