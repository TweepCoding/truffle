package visual

import (
	"fmt"

	"github.com/TweepCoding/truffle"
	"github.com/TweepCoding/truffle/node"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

/*
The Sprite node is a node that consists of showing an image to the window, based on it's
position.
*/
type Sprite struct {
	Parent              node.Node
	Children            []node.Node
	Name                string
	Texture             *sdl.Texture
	Width, Height, X, Y float64
	DrawFunction        func() error
	UpdateFunction      func(float64) error
}

var (
	_ node.Node        = (*Sprite)(nil)
	_ node.DrawUpdater = (*Sprite)(nil)
	_ node.Positioner  = (*Sprite)(nil)
)

// Creates a new Sprite node. The X and Y values determine their initial position relative to their Parent
func NewSprite(Path string, X, Y float64) (*Sprite, error) {

	SpriteNode := &Sprite{}

	Surface, err := img.Load(Path)

	if err != nil {
		return nil, fmt.Errorf("Error loading image from Path %s", Path)
	}

	SpriteNode.X, SpriteNode.Y = X, Y
	SpriteNode.Width, SpriteNode.Height, SpriteNode.Name = float64(Surface.W), float64(Surface.H), "Sprite"
	SpriteNode.Texture, err = truffle.Renderer.CreateTextureFromSurface(Surface)
	SpriteNode.UpdateFunction, SpriteNode.DrawFunction = func(Delta float64) error { return nil }, func() error { return nil }

	Surface.Free()

	return SpriteNode, nil
}

func (Sprite *Sprite) Draw() error {
	x, y := Sprite.X-(Sprite.Width/2), Sprite.Y-(Sprite.Height/2)

	if Positionable, ok := Sprite.Parent.(node.Positioner); ok {
		x, y = x+Positionable.GetX(), y+Positionable.GetY()
	}

	truffle.Renderer.Copy(
		Sprite.Texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(Sprite.Width), H: int32(Sprite.Height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(Sprite.Width), H: int32(Sprite.Height)},
	)

	return Sprite.DrawFunction()
}

func (Sprite *Sprite) Update(Delta float64) error {
	return Sprite.UpdateFunction(Delta)
}

func (Sprite *Sprite) OnDraw(Draw func() error) {
	Sprite.DrawFunction = Draw
}

func (Sprite *Sprite) OnUpdate(Update func(float64) error) {
	Sprite.UpdateFunction = Update
}

// Default Node behaivour

func (Sprite *Sprite) AddChild(Child node.Node) {
	Sprite.Children = append(Sprite.Children, Child)
	Child.SetParent(Sprite)
}

func (Sprite *Sprite) RemoveChild(Child node.Node) error {
	var i int = -1
	for index, Sprite := range Sprite.Children {
		if Sprite == Child {
			i = index
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("Error while removing child from %s Sprite: Could not find Sprite to remove", Sprite.Name)
	}

	//Removes element by shifting the element to delete to last position, then just cutting the array
	Sprite.Children[len(Sprite.Children)-1], Sprite.Children[i] = Sprite.Children[i], Sprite.Children[len(Sprite.Children)-1]
	Sprite.Children = Sprite.Children[:len(Sprite.Children)-1]
	return nil
}

func (Sprite *Sprite) GetParent() node.Node {
	return Sprite.Parent
}

func (Sprite *Sprite) SetParent(Parent node.Node) {
	Sprite.Parent = Parent
}

func (Sprite *Sprite) GetChildren() []node.Node {
	return Sprite.Children
}

func (Sprite *Sprite) GetName() string {
	return Sprite.Name
}

// Default Positionable behaivour

func (Sprite *Sprite) GetX() float64 {
	return Sprite.X
}

func (Sprite *Sprite) GetY() float64 {
	return Sprite.Y
}

func (Sprite *Sprite) SetX(x float64) {
	Sprite.X = x
}

func (Sprite *Sprite) SetY(y float64) {
	Sprite.Y = y
}

// Default Measurable behaivour

func (Sprite *Sprite) GetW() float64 {
	return Sprite.Width
}

func (Sprite *Sprite) GetH() float64 {
	return Sprite.Height
}

func (Sprite *Sprite) SetW(Width float64) {
	Sprite.Width = Width
}

func (Sprite *Sprite) SetH(Height float64) {
	Sprite.Height = Height
}
