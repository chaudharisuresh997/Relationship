package service

import (
	"errors"
	"geektrust/model"
	"strings"
)

const (
	Not_Enough_Args string = "Not Enough Args"
)

//commands
const (
	ADD_CHILD        string = "add_child"
	GET_RELATIONSHIP string = "get_relationship"
)

//msg
const (
	CHILD_ADDITION_SUCCEEDED string = "child_addition_succeed"
)

//gender
const (
	MALE   string = "male"
	FEMALE string = "female"
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

var root model.Node

func ProcessInput(commandLines []string) {
	for i := 0; i < len(commandLines); i++ {

		//get each action
		request, err := GetRequest(commandLines[i])
		if err == nil {
			ProcessCommand(request, root)
		}
	}
}
