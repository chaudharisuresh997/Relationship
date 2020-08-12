package service

import (
	"log"
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
//Messages
const(PERSON_NOT_FOUND string="PERSON_NOT_FOUND"
CHILD_ADDITION_FAILED string="CHILD_ADDITION_FAILED")

func ProcessCommand(r Request, root *model.Node)(*model.Node) {

	switch r.Command {
	case GET_RELATIONSHIP:
	//	var obj model.Node
	 var currentPerson *model.Node
	// currentPerson=&obj
	currentPerson=Find(r.Name, root)
		mother := currentPerson.Mother

		GetRelationUsingcurrentAndMother(currentPerson, mother, r.Relation)
	
	case ADD_CHILD:
		root=AddChild(root, r)

	}
return root
}
func AddChild(root *model.Node, request Request)(*model.Node){
	log.Printf("AdChild")
	//node to find is currentNode
	var currentNode *model.Node
	var child *model.Node
	if len(root.Name) == 0 {
		log.Printf("AdChild root empty")
		root = &model.Node{
			Name:   request.Name,
			Gender: request.Gender,
			Children:make(map[string]*model.Node),
		}
		child = &model.Node{
			Name:   request.ChildName,
			Gender: request.Gender,
			Mother: root,
			
		}
		root.Children[child.Name] = child
	} else {
		//find parent first
		//then assign baby
		log.Printf("AdChild root !null")
		//
		// var obj model.Node
		
		// currentNode=&obj
		currentNode=Find(request.Name, root)//,currentNode)

		child = &model.Node{
			Name:   request.ChildName,
			Gender: request.Gender,
			Mother: currentNode,
			
		}
	
		//add child to map
		if(currentNode.Children==nil){
			currentNode.Children=make(map[string]*model.Node)
		}		
		currentNode.Children[child.Name] = child
	}
	return root
	
}

func Find(searchName string, root *model.Node)(*model.Node)  {
	if root == nil {
		fmt.Println("NIL")
		return nil
	}
	if strings.EqualFold(root.Name,searchName) {
		//*toSearchNode=*root
		return root
	} 
		if root.Children != nil {
			for i, _ := range root.Children {
				// if(strings.EqualFold(toSearchNode.Name,searchName)){
				// 	break
				// }	
				Find(searchName, root.Children[i])
				
			}
		}
		return nil
	

}
func GetRelationUsingcurrentAndMother(currentPerson *model.Node, mother *model.Node, relationship string) []*model.Node {
	var nodes []*model.Node
	switch relationship {
	case Maternal_Uncle:
		//you have mother get parent of mother 
		msMother:=mother.Mother
	if msMother!=nil {
		for _,v:= range msMother.Children {
			unclesOrAunt:=*(v)
			//FOUND
			if strings.EqualFold(unclesOrAunt.Gender,MALE) && !strings.EqualFold(mother.Name,unclesOrAunt.Name){
				fmt.Println(unclesOrAunt.Name)
				nodes=append(nodes, &unclesOrAunt)
			}
		}
	}
		return output(nodes)
		
		//get all child of mothers mother 
	case Maternal_Aunt:
		msMother:=mother.Mother
		if msMother!=nil {
		for _,v:= range msMother.Children {
			unclesOrAunt:=*(v)
			if strings.EqualFold(unclesOrAunt.Gender,FEMALE) && !strings.EqualFold(mother.Name,unclesOrAunt.Name){
				fmt.Println(unclesOrAunt.Name)
				nodes=append(nodes, &unclesOrAunt)
				
			}
		}
	}
		return output(nodes)
	case Son:
	
		for _,v:= range currentPerson.Children {
			sonorDaughter:=*(v)
			if strings.EqualFold(sonorDaughter.Gender,MALE){
				fmt.Println(sonorDaughter.Name)
				nodes=append(nodes, &sonorDaughter)
			}
		}
		return output(nodes)	
	case Daughter:
		
		for _,v:= range currentPerson.Children {
			sonorDaughter:=*(v)
			if strings.EqualFold(sonorDaughter.Gender,FEMALE){
				fmt.Println(sonorDaughter.Name)
				nodes=append(nodes, &sonorDaughter)
			}
		}
		return output(nodes)	
	case Siblings:
		for _,v:= range currentPerson.Mother.Children {
			sonorDaughter:=*(v)
			if !strings.EqualFold(sonorDaughter.Name,currentPerson.Name){
				fmt.Println(sonorDaughter.Name)
				nodes=append(nodes, &sonorDaughter)
			}
		}
		return output(nodes)	

	}
	return nil
}
func output(nodes []*model.Node)[]*model.Node{
	if(len(nodes)>0){
		return nodes	
		}else{
			//seach done but not found
			fmt.Println(PERSON_NOT_FOUND)
		}
		return nil
}