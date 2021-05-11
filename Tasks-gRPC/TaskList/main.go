package main

import (
	"fmt"
	"log"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World!")	

	submitLab := &TaskList{
		TaskName: "Submit lab",
		TaskDate: "05.25",
	}
	
	data, err := proto.Marshal(submitLab)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}
	fmt.Println(data)
	
	newSubmission := &TaskList{}
	err = proto.Unmarshal(data, newSubmission)
	if err != nil {
		log.Fatal("UnMarshalling error", err)		
	}
	
	fmt.Println(newSubmission.GetTaskName)
	fmt.Println(newSubmission.GetTaskDate)
}
