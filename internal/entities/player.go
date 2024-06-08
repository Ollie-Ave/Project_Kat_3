package entities

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	idleAnimationName = "idle"
)

func NewPlayer(initialPosition rl.Vector2) (engine_entities.EntityUpdatable, error) {
	animations, err := loadAnimationData()

	if err != nil {
		return nil, err
	}

	return &player{
		Position: initialPosition,
		AnimationMeta: animationMeta{
			CurrentAnimationFrame:    0,
			CurrentAnimation:         idleAnimationName,
			TimeSinceLastFrameChange: 0,
			TimeBetweenFrameChange:   150,
		},
		Animations: animations,
	}, nil
}

func loadAnimationData() (map[string]animationData, error) {
	animations := map[string]animationData{
		idleAnimationName: {
			TexturePath: "_Idle.png",
			FrameCount:  10,
		},
	}

	for _, animation := range animations {
		const playerSpritePath = "entities/player"
		playerSpriteFullPath := fmt.Sprintf(
			"%s/%s/%s",
			engine_shared.AssetsPath,
			playerSpritePath,
			animation.TexturePath)

		_, err := os.Stat(playerSpriteFullPath)

		if err != nil {
			return nil, err
		}

		animation.Texture = rl.LoadTexture(playerSpriteFullPath)

		animation.FrameSize = rl.NewVector2(
			float32(animation.Texture.Width/int32(animation.FrameCount)),
			float32(animation.Texture.Height))

		animations[idleAnimationName] = animation
	}

	return animations, nil
}

type player struct {
	Position      rl.Vector2
	AnimationMeta animationMeta

	Animations map[string]animationData
}

type animationMeta struct {
	CurrentAnimationFrame int
	CurrentAnimation      string

	TimeSinceLastFrameChange int
	TimeBetweenFrameChange   int
}

type animationData struct {
	Texture rl.Texture2D

	TexturePath string
	FrameSize   rl.Vector2
	FrameCount  int
}

func (p *player) Update() {
	p.handleAnimation()
}

func (p *player) handleAnimation() {
	p.AnimationMeta.TimeSinceLastFrameChange += int(rl.GetFrameTime() * 1000)

	if p.AnimationMeta.TimeSinceLastFrameChange >= p.AnimationMeta.TimeBetweenFrameChange {
		p.AnimationMeta.CurrentAnimationFrame++

		p.AnimationMeta.TimeSinceLastFrameChange = 0
	}

	if p.AnimationMeta.CurrentAnimationFrame >= p.Animations[p.AnimationMeta.CurrentAnimation].FrameCount {
		p.AnimationMeta.CurrentAnimationFrame = 0
	}
}

func (p *player) Render() {
	currentAnimation := p.Animations[p.AnimationMeta.CurrentAnimation]
	textureSourceRec := rl.NewRectangle(
		currentAnimation.FrameSize.X*float32(p.AnimationMeta.CurrentAnimationFrame+1),
		0,
		currentAnimation.FrameSize.X,
		currentAnimation.FrameSize.Y)

	rl.DrawTextureRec(
		currentAnimation.Texture,
		textureSourceRec,
		p.Position,
		rl.White)
}
