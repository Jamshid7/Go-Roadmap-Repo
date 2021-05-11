package main

import (
	"fmt"
	"log"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World!")	

	mike := &ContactList{
		Name: "Mike",
		PhoneNumber: 12345,
	}
	
	data, err := proto.Marshal(mike)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}
	fmt.Println(data)
	
	newMike := &ContactList{}
	err = proto.Unmarshal(data, newMike)
	if err != nil {
		log.Fatal("UnMarshalling error", err)		
	}
	
	fmt.Println(newMike.GetName)
	fmt.Println(newMike.GetPhoneNumber)
}
