package truffle

import (
	"fmt"

	"github.com/TweepCoding/truffle/node"
)

type EmptyNode struct {
	Parent   node.Node
	Children []node.Node
	Name     string
}

var _ node.Node = (*EmptyNode)(nil)

func NewEmptyNode() *EmptyNode {
	return &EmptyNode{Name: "EmptyNode"}
}

// Default node behaivour

func (EmptyNode *EmptyNode) AddChild(Child node.Node) {
	EmptyNode.Children = append(EmptyNode.Children, Child)
	Child.SetParent(EmptyNode)
}

func (EmptyNode *EmptyNode) RemoveChild(Child node.Node) error {
	var i int = -1
	for index, EmptyNode := range EmptyNode.Children {
		if EmptyNode == Child {
			i = index
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("Error while removing child from %s EmptyNode: Could not find EmptyNode to remove", EmptyNode.Name)
	}

	//Removes element by shifting the element to delete to last position, then just cutting the array
	EmptyNode.Children[len(EmptyNode.Children)-1], EmptyNode.Children[i] = EmptyNode.Children[i], EmptyNode.Children[len(EmptyNode.Children)-1]
	EmptyNode.Children = EmptyNode.Children[:len(EmptyNode.Children)-1]
	return nil
}

func (EmptyNode *EmptyNode) GetParent() node.Node {
	return EmptyNode.Parent
}

func (EmptyNode *EmptyNode) SetParent(Parent node.Node) {
	EmptyNode.Parent = Parent
}

func (EmptyNode *EmptyNode) GetChildren() []node.Node {
	return EmptyNode.Children
}

func (EmptyNode *EmptyNode) GetName() string {
	return EmptyNode.Name
}
