package entities

import (
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PlayerIdleAnimation = iota
	PlayerRunAnimation
	PlayerJumpAnimation
	PlayerFallAnimation
)

func NewPlayer(
	initialPosition rl.Vector2,
	entityMangaer engine_entities.EntityManager,
	physicsHandler engine_entities.PhysicsHandler,
	animationHandler engine_entities.AnimationHandler,
) (engine_entities.EntityUpdater, error) {
	animations := map[int]engine_entities.AnimationData{
		PlayerIdleAnimation: engine_entities.NewAnimationData(
			"PlayerIdle.png",
			10,
			150,
		),
		PlayerRunAnimation: engine_entities.NewAnimationData(
			"PlayerRun.png",
			8,
			85,
		),
		PlayerJumpAnimation: engine_entities.NewAnimationDataEx(
			"PlayerJump.png",
			3,
			150,
			false,
		),
	}

	err := animationHandler.LoadAnimations(animations)

	if err != nil {
		return nil, err
	}

	playerHitboxDimensions := rl.NewRectangle(
		24,
		16,
		46,
		64,
	)

	return &player{
		EntityManager:    entityMangaer,
		HitboxDimensions: playerHitboxDimensions,
		Position:         initialPosition,
		PhysicsHandler:   physicsHandler,
		AnimationHandler: animationHandler,
	}, nil
}

type player struct {
	EntityManager engine_entities.EntityManager

	Position         rl.Vector2
	AnimationHandler engine_entities.AnimationHandler

	Velocity         rl.Vector2
	HitboxDimensions rl.Rectangle

	PhysicsHandler engine_entities.PhysicsHandler
}

func (p *player) Update() error {

	p.handleMovementInput(&p.Velocity)
	p.handlePhysics(&p.Velocity, &p.Position)
	p.handleAnimations()

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
	p.AnimationHandler.RenderAnimationFrame(p.Position)

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
	const speed = 300

	velocityMut.X = 0
	deltaTime := rl.GetFrameTime()

	if rl.IsKeyDown(rl.KeyD) {
		velocityMut.X = speed * deltaTime
	} else if rl.IsKeyDown(rl.KeyA) {
		velocityMut.X = -speed * deltaTime
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

func (p *player) handleAnimations() {
	playerHitbox := p.GetHitbox()

	if !p.PhysicsHandler.IsTouchingGround(&playerHitbox) {
		p.AnimationHandler.SetAnimation(PlayerJumpAnimation)
	} else if p.Velocity.X == 0 {
		p.AnimationHandler.SetAnimation(PlayerIdleAnimation)
	} else if p.Velocity.X < 0 {
		p.AnimationHandler.SetAnimation(PlayerRunAnimation)
	} else if p.Velocity.X > 0 {
		p.AnimationHandler.SetAnimation(PlayerRunAnimation)
	}

	if p.Velocity.X < 0 {
		p.AnimationHandler.SetDirection("left")
	} else if p.Velocity.X > 0 {
		p.AnimationHandler.SetDirection("right")
	}

	p.AnimationHandler.HandleAnimation()
}
