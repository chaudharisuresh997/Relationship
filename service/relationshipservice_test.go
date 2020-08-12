package service

import (
	"geektrust/model"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	r := model.Node{
		Name:   "seeta",
		Gender: FEMALE,
		Children: make(map[string]*model.Node),
	}

	r.Children = make(map[string]*model.Node)
	c1 := &model.Node{
		Name:     "Geeta",
		Gender:   FEMALE,
		Children: make(map[string]*model.Node),
	}
	r.Children["Geeta"] = c1
	c1_1 := &model.Node{
		Name:   "Geeta_kiBeti",
		Gender: FEMALE,
	}
	c1.Children["Geeta_kiBeti"] = c1_1

	// //negative case
	// c1_1_negative := &model.Node{
	// 	Name:   "Geeta_kiBeti_NOTFOUND",
	// 	Gender: FEMALE,
	// }

	type args struct {
		searchName string
		root       *model.Node
	}
	tests := []struct {
		name string
		args args
		
	}{
		{name: "find child", args: args{searchName: "Geeta", root: &r}},
		{name: "find child", args: args{searchName: "Geeta", root: &r}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 _ = Find(tt.args.searchName, tt.args.root)
				
			
		})
	}
}

func TestGetRelationUsingcurrentAndMother(t *testing.T) {
	r := model.Node{
		Name:   "seeta",
		Gender: FEMALE,
		Children: make(map[string]*model.Node),
	}
	//seeta daughter1
	seeta_d1:=model.Node{
		Name:"seetad1",
		Gender:FEMALE,
		Mother:&r,
	}

	//seeta daughter2
	seeta_d2:=model.Node{
		Name:"seetad2",
		Gender:FEMALE,
		Mother:&r,
	}

	//seeta son 1
	seeta_s1:=model.Node{
		Name:"seetas1",
		Gender:MALE,
		Mother:&r,
	}
	
	r.Children["seeta_d1"]=&seeta_d1
	r.Children["seeta_d2"]=&seeta_d2
	//adding son
	r.Children["seeta_s1"]=&seeta_s1
	//seeta grand daughter d1_1
	seeta_d1_1:=model.Node{
		Name:"seetad1_1",
		Gender:FEMALE,
		Mother:&seeta_d1,
	}

	type args struct {
		currentPerson *model.Node
		mother        *model.Node
		relationship  string
	}
	tests := []struct {
		name string
		args args
		want []*model.Node
	}{
		{name:"Find maternal aunt",args:args{currentPerson:&seeta_d1_1,mother:&seeta_d1,relationship:Maternal_Aunt},
		want:[]*model.Node{&seeta_d2}},
		{name:"Find maternal uncle",args:args{currentPerson:&seeta_d1_1,mother:&seeta_d1,relationship:Maternal_Uncle},
		want:[]*model.Node{&seeta_s1}},
		{name:"Son",args:args{currentPerson:&r,mother:nil,relationship:Son},
		want:[]*model.Node{&seeta_s1}},
		{name:"Daughter",args:args{currentPerson:&r,mother:nil,relationship:Daughter},
		want:[]*model.Node{&seeta_d1,&seeta_d2}},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRelationUsingcurrentAndMother(tt.args.currentPerson, tt.args.mother, tt.args.relationship); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRelationUsingcurrentAndMother() = %v, want %v", got, tt.want)
			}
		})
	}
}
