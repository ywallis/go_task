package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ywallis/go_task/database"
)

func (a *appConfig) createTask(name string) (task, error){

	queryParams := database.CreateTaskParams{
		Name:      name,
		CreatedAt: time.Now(), 
		UpdatedAt: time.Now(),
	}
	dbtask, err := a.db.CreateTask(context.Background(), queryParams)
	if err != nil {
		return task{}, fmt.Errorf("Error while creating task: %s\n", err)
	}
	output := task{
		id:   int(dbtask.ID),
		name: dbtask.Name,
	}
	return output, nil
}
