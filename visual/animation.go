package visual

import (
	"fmt"
	"io/ioutil"
)

type Animation struct {
	Sprites           []Image
	FPS, CurrentFrame int
}

func NewAnimation(Source string, FPS int) (Animation, error) {
	Result := Animation{}
	Result.FPS = FPS
	Files, err := ioutil.ReadDir(Source)

	if err != nil {
		return Animation{}, fmt.Errorf("Error loading animation from %s: %s", Source, err.Error())
	}

	for Number, File := range Files {
		Image, err := NewImage(Source + File.Name())
		if err != nil {
			return Animation{}, fmt.Errorf("Error loading frame number %d from animation: %s", Number, err.Error())
		}
		Result.Sprites = append(Result.Sprites, Image)
	}
	return Result, nil
}
