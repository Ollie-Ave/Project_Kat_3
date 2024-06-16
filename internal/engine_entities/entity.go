package engine_entities

import rl "github.com/gen2brain/raylib-go/raylib"

type EntityUpdater interface {
	Update() error
}

type Collider interface {
	GetHitbox() rl.Rectangle
}
