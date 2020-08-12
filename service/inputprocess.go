package service

import (
	"log"
	"errors"
	"geektrust/model"
	"strings"
)

const (
	Not_Enough_Args string = "Not Enough Args"
)

//commands
const (
	ADD_CHILD        string = "ADD_CHILD"
	GET_RELATIONSHIP string = "GET_RELATIONSHIP"
)

//msg
const (
	CHILD_ADDITION_SUCCEEDED string = "child_addition_succeed"
)

//gender
const (
	MALE   string = "Male"
	FEMALE string = "Female"
)

type Request struct {
	Command   string
	Name      string
	ChildName string
	Gender    string
	Relation  string
}

//GetRequest :
func GetRequest(command string) (Request, error) {
log.Println("GetRequest")
	elements := strings.Split(command, " ")
	if len(elements) < 3 {
		return Request{}, errors.New(Not_Enough_Args)
	}
	if len(elements) == 4 {
		return Request{
			Command:   elements[0],
			Name:      elements[1],
			ChildName: elements[2],
			Gender:    elements[3],
		}, nil
	}
	return Request{
		Command:  elements[0],
		Name:     elements[1],
		Relation: elements[2],
	}, nil

}



func ProcessInput(commandLines []string) {
	
	
root:=model.Node{}

	for i := 0; i < len(commandLines); i++ {

		//get each action
		request, err := GetRequest(commandLines[i])
		if err == nil {
			log.Println("Calling ProcessCommand")
			
			returnroot:=ProcessCommand(request, &root)
			root=*returnroot
		}else{
			log.Printf("error reading command %v %v",commandLines[i],err)
		}
	}
}
