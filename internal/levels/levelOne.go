package levels

import rl "github.com/gen2brain/raylib-go/raylib"

func initLevelOne() *levelOne {
	return &levelOne{}
}

type levelOne struct {
}

func (l *levelOne) Render() {
	rl.DrawRectangle(10, 10, 100, 100, rl.Red)
}
