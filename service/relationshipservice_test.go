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

	//negative case
	c1_1_negative := &model.Node{
		Name:   "Geeta_kiBeti_NOTFOUND",
		Gender: FEMALE,
	}

	type args struct {
		searchName string
		root       *model.Node
	}
	tests := []struct {
		name string
		args args
		want *model.Node
	}{
		{name: "find child", args: args{searchName: "Geeta", root: &r}, want: c1},
		{name: "find child", args: args{searchName: "Geeta", root: &r}, want: c1_1_negative},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.searchName, tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
