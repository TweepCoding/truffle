package node

/*
A node is a part of the engine. The engine functions by creating a tree with various nodes. These
nodes can have one parent, and multiple children. The children will allow inheritance of abilities
to their parent nodes.

An example would be a "Player node", which have 2 children nodes: "Sprite node", and "Movement node",
which allow the player to both have a sprite, and also be able to move.
*/
type Node interface {
	AddChild(Node)
	RemoveChild(Node) error
	GetParent() Node
	SetParent(Node)
	GetChildren() []Node
	GetName() string
}

func ForEveryChild(Execute func(Node) error, Parent Node) error {
	for _, Child := range Parent.GetChildren() {
		if err := Execute(Child); err != nil {
			return err
		}
		ForEveryChild(Execute, Child)
	}
	return nil
}
