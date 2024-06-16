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

func NewPlayer(
	initialPosition rl.Vector2,
	entityMangaer engine_entities.EntityManager,
	physicsHandler engine_entities.PhysicsHandler,
) (engine_entities.EntityUpdater, error) {

	animations, err := loadAnimationData()

	if err != nil {
		return nil, err
	}

	playerHitboxDimensions := rl.NewRectangle(
		40,
		40,
		32,
		40,
	)

	return &player{
		EntityManager:    entityMangaer,
		HitboxDimensions: playerHitboxDimensions,
		Position:         initialPosition,
		PhysicsHandler:   physicsHandler,
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
	EntityManager engine_entities.EntityManager

	Position      rl.Vector2
	AnimationMeta animationMeta

	Animations map[string]animationData

	Velocity         rl.Vector2
	HitboxDimensions rl.Rectangle

	PhysicsHandler engine_entities.PhysicsHandler
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

func (p *player) Update() error {
	p.handleAnimation()

	p.handleMovementInput(&p.Velocity)
	p.handlePhysics(&p.Velocity, &p.Position)

	p.Position.X += p.Velocity.X
	p.Position.Y -= p.Velocity.Y

	return nil
}

func (p *player) GetHitbox() rl.Rectangle {
	return rl.NewRectangle(
		p.HitboxDimensions.X+p.Position.X,
		p.HitboxDimensions.Y+p.Position.Y,
		p.HitboxDimensions.Width,
		p.HitboxDimensions.Height)
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

	if os.Getenv(engine_shared.DebugModeEnvironmentVariable) != "true" {
		return
	}

	playerHitbox := p.GetHitbox()

	rl.DrawRectangleLines(
		int32(playerHitbox.X),
		int32(playerHitbox.Y),
		int32(playerHitbox.Width),
		int32(playerHitbox.Height),
		rl.Red)
}

func (p *player) handleMovementInput(velocityMut *rl.Vector2) {
	const speed = 5

	velocityMut.X = 0

	if rl.IsKeyDown(rl.KeyD) {
		velocityMut.X = speed
	} else if rl.IsKeyDown(rl.KeyA) {
		velocityMut.X = -speed
	}

	hitbox := p.GetHitbox()

	if rl.IsKeyDown(rl.KeySpace) && p.PhysicsHandler.IsTouchingGround(&hitbox) {
		velocityMut.Y = 10
	}
}

func (p *player) handlePhysics(velocityMut *rl.Vector2, positionMut *rl.Vector2) {
	playerHitbox := p.GetHitbox()

	handlerData := engine_entities.NewPhysicsHandlerData(&playerHitbox, &p.HitboxDimensions)

	p.PhysicsHandler.HandlePhysics(handlerData, velocityMut, positionMut)
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
