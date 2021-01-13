package truffle

import (
	"fmt"
	"sync"
	"time"
	"github.com/TweepCoding/truffle/node"
	"github.com/veandco/go-sdl2/sdl"
)

/*
The Root Root is the top Root of the game tree, and it's the only Root that doesn't have a Root.
This Root can also be accessed as a global variable by any other Roots, this is done in this way
to allow complex traversal behaivor in a simpler manner.
*/
type Root struct {
	Window   *sdl.Window
	FPS      float64
	Name     string
	Parent   node.Node
	Children []node.Node
	Mutex    *sync.RWMutex
}

var (
	RootNode *Root         = nil
	Renderer *sdl.Renderer = nil
	_        node.Node          = (*Root)(nil)
)

/*
NewRoot will create a new Root Root, which can be accessed globally by any package using engine.RootRoot and engine.RootValue. Take note, this will be
the only Root that will be public, since it's required to avoid convoluted behaivor with some Roots.
*/
func NewRoot(width, height int32, fps float64) error {
	if RootNode != nil {
		return fmt.Errorf("Error creating new Root Node: A Root Root has already been made")
	}

	var err error

	RootNode = &Root{}

	RootNode.Window, err = sdl.CreateWindow("Trufflemania", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("Error creating SDL2 window: %s", err.Error())
	}

	Renderer, err = sdl.CreateRenderer(RootNode.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("Error creating SDL2 renderer: %s", err.Error())
	}

	RootNode.FPS, RootNode.Mutex, RootNode.Name = fps, &sync.RWMutex{}, "Root"

	return nil
}

/*
The mainloop will block until the program is closed. While it blocks, it will also manage
the draw() and update() of each Root on the tree, which all Roots should have. In case a Root
is misconfigured, and doesn't have these functions, it will crash.
*/
func (Root *Root) MainLoop() {
	defer Renderer.Destroy()
	defer Root.Window.Destroy()

	Delta := 1.0

	for {
		FrameStart := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		Renderer.SetDrawColor(255, 255, 255, 255)
		Renderer.Clear()

		node.ForEveryChild(func(Child node.Node) error {
			if DrawUpdate, ok := Child.(node.DrawUpdater); ok {
				if err := DrawUpdate.Draw(); err != nil {
					return err
				}
				if err := DrawUpdate.Update(Delta); err != nil {
					return err
				}
			}
			return nil
		}, Root)

		Renderer.Present()

		Delta = time.Since(FrameStart).Seconds() * Root.FPS
	}
}

func (Root *Root) PrintTree() {
	node.ForEveryChild(func(Child node.Node) error {

		return nil
	}, Root)
}

// Default Node behaivour

func (Root *Root) AddChild(Child node.Node) {
	Root.Children = append(Root.Children, Child)
	Child.SetParent(Root)
}

func (Root *Root) RemoveChild(Child node.Node) error {
	var i int = -1
	for index, Root := range Root.Children {
		if Root == Child {
			i = index
			break
		}
	}
	if i == -1 {
		return fmt.Errorf("Error while removing child from %s Root: Could not find Root to remove", Root.Name)
	}

	//Removes element by shifting the element to delete to last position, then just cutting the array
	Root.Children[len(Root.Children)-1], Root.Children[i] = Root.Children[i], Root.Children[len(Root.Children)-1]
	Root.Children = Root.Children[:len(Root.Children)-1]
	return nil
}

func (Root *Root) GetParent() node.Node {
	return Root.Parent
}

func (Root *Root) SetParent(Parent node.Node) {
	Root.Parent = Parent
}

func (Root *Root) GetChildren() []node.Node {
	return Root.Children
}

func (Root *Root) GetName() string {
	return Root.Name
}
