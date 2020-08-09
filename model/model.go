package model

type Node struct {
	Name     string
	Gender   string
	Children map[string]*Node
	Mother   *Node
}
