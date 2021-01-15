package physics

import (
	"fmt"

	"github.com/TweepCoding/truffle"
	"github.com/TweepCoding/truffle/node"
)

/*
Collisionbox is a type of node
*/
type CollisionBox struct {
	Width, Height     int32
	Parent            node.Node
	Children          []node.Node
	Name              string
	UpdateFunction    func(float64) error
	DrawFunction      func() error
	CollisionFunction func(node.Collisioner)
}

var (
	_ node.Node        = (*CollisionBox)(nil)
	_ node.DrawUpdater = (*CollisionBox)(nil)
	_ node.Measurer    = (*CollisionBox)(nil)
	_ node.Collisioner = (*CollisionBox)(nil)
)

func NewCollisionBox(Width, Height int32) (*CollisionBox, error) {

	Result := &CollisionBox{}

	Result.Width, Result.Height, Result.Name = Width, Height, "CollisionBox"
	Result.UpdateFunction, Result.DrawFunction, Result.CollisionFunction = func(Delta float64) error { return nil }, func() error { return nil }, func(Collisioner node.Collisioner) {}

	return Result, nil
}

func (CollisionBox *CollisionBox) Draw() error {
	return CollisionBox.DrawFunction()
}

func (CollisionBox *CollisionBox) Update(Delta float64) error {
	node.ForEveryChild(func(Child node.Node) error {
		if Child.GetName() == "CollisionBox" {
			Pos := CollisionBox.GetParent().(node.Positioner)
			Pos2 := Child.GetParent().(node.Positioner)
			Col2 := Child.(node.Collisioner)
			x1, y1, w1, h1 := int32(Pos.GetX()), int32(Pos.GetY()), CollisionBox.GetW()/2, CollisionBox.GetW()/2
			x2, y2, w2, h2 := int32(Pos2.GetX()), int32(Pos2.GetY()), Col2.GetW()/2, Col2.GetH()/2
			if (y1+h1) > (y2-h2) && (y1-h1) < (y2+h2) && (x1+w1) > (x2-w2) && (x1-w1) < (x2+w2) {
				CollisionBox.CollisionFunction(Col2)
			}
		}
		return nil
	}, truffle.RootNode)
	return CollisionBox.UpdateFunction(Delta)
}

func (CollisionBox *CollisionBox) OnCollision(Collision func(node.Collisioner)) {
	CollisionBox.CollisionFunction = Collision
}

// Default "OnUpdate" and "OnDraw" behaivour

func (CollisionBox *CollisionBox) OnUpdate(Update func(float64) error) {
	CollisionBox.UpdateFunction = Update
}

func (CollisionBox *CollisionBox) OnDraw(Draw func() error) {
	CollisionBox.DrawFunction = Draw
}

// Default Node behaivour

func (CollisionBox *CollisionBox) AddChild(Child node.Node) {
	CollisionBox.Children = append(CollisionBox.Children, Child)
	Child.SetParent(CollisionBox)
}

func (CollisionBox *CollisionBox) RemoveChild(Child node.Node) error {
	var i int = -1
	for index, CollisionBox := range CollisionBox.Children {
		if CollisionBox == Child {
			i = index
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("Error while removing child from %s CollisionBox: Could not find CollisionBox to remove", CollisionBox.Name)
	}

	//Removes element by shifting the element to delete to last position, then just cutting the array
	CollisionBox.Children[len(CollisionBox.Children)-1], CollisionBox.Children[i] = CollisionBox.Children[i], CollisionBox.Children[len(CollisionBox.Children)-1]
	CollisionBox.Children = CollisionBox.Children[:len(CollisionBox.Children)-1]
	return nil
}

func (CollisionBox *CollisionBox) GetParent() node.Node {
	return CollisionBox.Parent
}

func (CollisionBox *CollisionBox) SetParent(Parent node.Node) {
	CollisionBox.Parent = Parent
}

func (CollisionBox *CollisionBox) GetChildren() []node.Node {
	return CollisionBox.Children
}

func (CollisionBox *CollisionBox) GetName() string {
	return CollisionBox.Name
}

// Default Measurable behaivour

func (CollisionBox *CollisionBox) GetW() int32 {
	return CollisionBox.Width
}

func (CollisionBox *CollisionBox) GetH() int32 {
	return CollisionBox.Height
}

func (CollisionBox *CollisionBox) SetW(Width int32) {
	CollisionBox.Width = Width
}

func (CollisionBox *CollisionBox) SetH(Height int32) {
	CollisionBox.Height = Height
}
