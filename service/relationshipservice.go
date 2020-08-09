package service

import (
	"fmt"
	"geektrust/model"
	"strings"
)

const (
	Maternal_Uncle string = "Maternal-Uncle"
	Maternal_Aunt  string = "Maternal-Aunt"
	Son            string = "Son"
	Daughter       string = "Daughter"
	Siblings       string = "siblings"
)

func ProcessCommand(r Request, root model.Node) {

	switch r.Command {
	case GET_RELATIONSHIP:
		currentPerson := Find(r.Name, &root)
		mother := currentPerson.Mother
		GetRelationUsingcurrentAndMother(currentPerson, mother, r.Relation)
	case ADD_CHILD:
		AddChild(root, r)

	}

}
func AddChild(root model.Node, request Request) {
	var currentNode *model.Node
	if len(root.Name) == 0 {
		root = model.Node{
			Name:   request.Name,
			Gender: request.Gender,
		}
		currentNode = &root
	} else {
		//find parent first
		//then assign baby
		currentNode = Find(request.Name, &root)
	}

	child := model.Node{
		Name:   request.ChildName,
		Gender: request.Gender,
		Mother: currentNode,
	}

	//add child to map
	currentNode.Children[child.Name] = &child
}

func Find(searchName string, root *model.Node) *model.Node {
	if root == nil {
		fmt.Println("NIL")
		return nil
	}
	if strings.EqualFold(root.Name, searchName) {
		return root
	} else {
		if root.Children != nil {
			for _, v := range root.Children {

				return Find(searchName, v)
				// if strings.EqualFold(*&(v).Name, searchName) {
				// 	return v
				// }
				//else process again its child
			}
		}
	}
	return &model.Node{}
}
func GetRelationUsingcurrentAndMother(currentPerson *model.Node, mother *model.Node, relationship string) *model.Node {
	
	switch relationship {
	case Maternal_Uncle:
		//you have mother get parent of mother 
		msMother:=mother.Mother
		for _,v:= range msMother.Children {
			unclesOrAunt:=*(v)
			if strings.EqualFold(unclesOrAunt.Gender,MALE){
				fmt.Println(unclesOrAunt.Name)
			}
		}
		//get all child of mothers mother 
	case Maternal_Aunt:
		msMother:=mother.Mother
		for _,v:= range msMother.Children {
			unclesOrAunt:=*(v)
			if strings.EqualFold(unclesOrAunt.Gender,FEMALE){
				fmt.Println(unclesOrAunt.Name)
			}
		}
	case Son:
	case Daughter:
	case Siblings:


	}
}
