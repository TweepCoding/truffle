package truffle

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Error initializing SDL: " + err.Error())
		return
	}
}
func Stop() {
	sdl.Quit()
}
