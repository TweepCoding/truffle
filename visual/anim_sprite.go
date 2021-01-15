package visual

import (
	"fmt"
	"time"

	"github.com/TweepCoding/truffle"
	"github.com/TweepCoding/truffle/node"
	"github.com/veandco/go-sdl2/sdl"
)

/*
The AnimatedSprite node is a node that consists of showing an image to the window, based on it's
position.
*/
type AnimatedSprite struct {
	Parent           node.Node
	Children         []node.Node
	Name             string
	CurrentAnimation Animation
	Animations       map[string]Animation
	SinceLastFrame   time.Time
	X, Y             float64
	DrawFunction     func() error
	UpdateFunction   func(float64) error
}

var (
	_ node.Node        = (*AnimatedSprite)(nil)
	_ node.DrawUpdater = (*AnimatedSprite)(nil)
	_ node.Positioner  = (*AnimatedSprite)(nil)
)

// Creates a new AnimatedSprite node. The X and Y values determine their initial position relative to their Parent
func NewAnimatedSprite(DefaultAnimation Animation, X, Y float64) (*AnimatedSprite, error) {

	Result := &AnimatedSprite{}
	Result.X, Result.Y, Result.Name = X, Y, "AnimatedSprite"
	Result.Animations, Result.CurrentAnimation = make(map[string]Animation), DefaultAnimation
	Result.SinceLastFrame = time.Now()
	Result.UpdateFunction, Result.DrawFunction = func(Delta float64) error { return nil }, func() error { return nil }

	Result.AddAnimation("default", DefaultAnimation)
	return Result, nil
}

func (AnimatedSprite *AnimatedSprite) Draw() error {
	CurrentImage := AnimatedSprite.CurrentAnimation.Sprites[AnimatedSprite.CurrentAnimation.CurrentFrame]

	x, y := int32(AnimatedSprite.X)-(CurrentImage.Width/2), int32(AnimatedSprite.Y)-(CurrentImage.Height/2)

	if Positionable, ok := AnimatedSprite.Parent.(node.Positioner); ok {
		x, y = x+int32(Positionable.GetX()), y+int32(Positionable.GetY())
	}

	truffle.Renderer.Copy(
		CurrentImage.Texture,
		&sdl.Rect{X: 0, Y: 0, W: CurrentImage.Width, H: CurrentImage.Height},
		&sdl.Rect{X: x, Y: y, W: CurrentImage.Width, H: CurrentImage.Height})

	return AnimatedSprite.DrawFunction()
}

func (AnimatedSprite *AnimatedSprite) Update(Delta float64) error {
	if time.Since(AnimatedSprite.SinceLastFrame) >= time.Duration(float64(time.Second)/float64(AnimatedSprite.CurrentAnimation.FPS)) {
		AnimatedSprite.CurrentAnimation.CurrentFrame++
		AnimatedSprite.SinceLastFrame = time.Now()
	}
	return AnimatedSprite.UpdateFunction(Delta)
}

func (AnimatedSprite *AnimatedSprite) OnDraw(Draw func() error) {
	AnimatedSprite.DrawFunction = Draw
}

func (AnimatedSprite *AnimatedSprite) OnUpdate(Update func(float64) error) {
	AnimatedSprite.UpdateFunction = Update
}

func (AnimatedSprite *AnimatedSprite) AddAnimation(Name string, Animation Animation) {
	AnimatedSprite.Animations[Name] = Animation
}

func (AnimatedSprite *AnimatedSprite) SetAnimation(Name string) {
	AnimatedSprite.CurrentAnimation = AnimatedSprite.Animations[Name]
}

// Default Node behaivour

func (AnimatedSprite *AnimatedSprite) AddChild(Child node.Node) {
	AnimatedSprite.Children = append(AnimatedSprite.Children, Child)
	Child.SetParent(AnimatedSprite)
}

func (AnimatedSprite *AnimatedSprite) RemoveChild(Child node.Node) error {
	var i int = -1
	for index, AnimatedSprite := range AnimatedSprite.Children {
		if AnimatedSprite == Child {
			i = index
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("Error while removing child from %s AnimatedSprite: Could not find AnimatedSprite to remove", AnimatedSprite.Name)
	}

	//Removes element by shifting the element to delete to last position, then just cutting the array
	AnimatedSprite.Children[len(AnimatedSprite.Children)-1], AnimatedSprite.Children[i] = AnimatedSprite.Children[i], AnimatedSprite.Children[len(AnimatedSprite.Children)-1]
	AnimatedSprite.Children = AnimatedSprite.Children[:len(AnimatedSprite.Children)-1]
	return nil
}

func (AnimatedSprite *AnimatedSprite) GetParent() node.Node {
	return AnimatedSprite.Parent
}

func (AnimatedSprite *AnimatedSprite) SetParent(Parent node.Node) {
	AnimatedSprite.Parent = Parent
}

func (AnimatedSprite *AnimatedSprite) GetChildren() []node.Node {
	return AnimatedSprite.Children
}

func (AnimatedSprite *AnimatedSprite) GetName() string {
	return AnimatedSprite.Name
}

// Default Positionable behaivour

func (AnimatedSprite *AnimatedSprite) GetX() float64 {
	return AnimatedSprite.X
}

func (AnimatedSprite *AnimatedSprite) GetY() float64 {
	return AnimatedSprite.Y
}

func (AnimatedSprite *AnimatedSprite) SetX(x float64) {
	AnimatedSprite.X = x
}

func (AnimatedSprite *AnimatedSprite) SetY(y float64) {
	AnimatedSprite.Y = y
}
