package visual

import (
	"fmt"

	"github.com/TweepCoding/truffle"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	Texture       *sdl.Texture
	Width, Height int32
}

func NewImage(Path string) (Image, error) {

	Result := Image{}
	Surface, err := img.Load(Path)
	if err != nil {
		return Image{}, fmt.Errorf("Error creating Image from %s: %s", Path, err.Error())
	}

	Result.Width, Result.Height = Surface.W, Surface.H
	Result.Texture, err = truffle.Renderer.CreateTextureFromSurface(Surface)

	if err != nil {
		return Image{}, fmt.Errorf("Error creating Image from %s: %s", Path, err.Error())
	}

	Surface.Free()
	return Result, nil
}
