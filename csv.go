package main
//
// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// 	"strconv"
// )
//
// const filePath string = "tasks.csv"
//
// func writeToFile(tasklist *[]task) error {
// 	f, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
//
// 	writer := csv.NewWriter(f)
// 	defer writer.Flush()
//
// 	header := []string{"name", "completed"}
// 	if err := writer.Write(header); err != nil {
// 		return err
// 	}
// 	for _, task := range *tasklist {
// 		record := []string{task.name, fmt.Sprintf("%t", task.completed)}
// 		writer.Write(record)
// 	}
//
// 	return nil
// }
//
// func readFromFile() ([]task, error) {
// 	taskList := []task{}
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	reader := csv.NewReader(f)
// 	reader.FieldsPerRecord = -1
//
// 	data, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, nil
// 	}
//
// 	for _, item := range data[1:] {
// 		completed, err := strconv.ParseBool(item[1])
// 		if err != nil {
// 			return nil, err
// 		}
// 		outputTask := task{item[0], completed}
// 		taskList = append(taskList, outputTask)
// 	}
// 	return taskList, nil
// 	}
// func createTask(name string, taskList *[]task) {
// 	newTask := task{
// 		name:      name,
// 		completed: false}
// 	*taskList = append(*taskList, newTask)
// }
