package engine_entities

import (
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileColliderHeight = 8
	tileColliderWidth  = 8
)

type PhysicsHandler interface {
	HandlePhysics(*physicsHandlerData, *rl.Vector2, *rl.Vector2)
	IsTouchingGround(hitbox *rl.Rectangle) bool
}

type physicsHandlerImpl struct {
	EntityManager EntityManager
}

func NewPhysicsHandler(entityManager EntityManager) PhysicsHandler {
	return &physicsHandlerImpl{
		EntityManager: entityManager,
	}
}

type physicsHandlerData struct {
	Hitbox           *rl.Rectangle
	HitboxDimensions *rl.Rectangle
}

func NewPhysicsHandlerData(hitbox *rl.Rectangle, hitboxDimensions *rl.Rectangle) *physicsHandlerData {
	return &physicsHandlerData{
		Hitbox:           hitbox,
		HitboxDimensions: hitboxDimensions,
	}
}

func (p *physicsHandlerImpl) HandlePhysics(handlerData *physicsHandlerData, velocityMut *rl.Vector2, positionMut *rl.Vector2) {
	p.handleGravityForce(velocityMut)

	p.handleTileMapCollisions(
		handlerData,
		velocityMut,
		positionMut,
	)
}

func (p *physicsHandlerImpl) IsTouchingGround(hitbox *rl.Rectangle) bool {
	levelCollider := p.EntityManager.GetLevelCollider()

	collisionData, tileWidth, tileHeight := levelCollider.GetLayerCollisionData(engine_shared.GroundLayerName)

	for y, collisionRow := range collisionData {
		for x, shouldCollide := range collisionRow {
			if shouldCollide {
				tileHitbox := rl.NewRectangle(
					float32(x*tileWidth),
					float32(y*tileHeight),
					float32(tileWidth),
					float32(tileHeight),
				)

				if p.isPlayerBottomTouchingTile(hitbox, &tileHitbox) {
					return true
				}
			}
		}
	}

	return false
}

func (p *physicsHandlerImpl) handleGravityForce(velocityMut *rl.Vector2) {
	const gravity = +0.5

	velocityMut.Y -= gravity
}

func (p *physicsHandlerImpl) handleTileMapCollisions(
	handlerData *physicsHandlerData,
	velocityMut *rl.Vector2,
	positionMut *rl.Vector2,
) {
	levelCollider := p.EntityManager.GetLevelCollider()

	collisionData, tileWidth, tileHeight := levelCollider.GetLayerCollisionData(engine_shared.GroundLayerName)

	for y, collisionRow := range collisionData {
		for x, shouldCollide := range collisionRow {
			if shouldCollide {
				tileHitbox := rl.NewRectangle(
					float32(x*tileWidth),
					float32(y*tileHeight),
					float32(tileWidth),
					float32(tileHeight),
				)

				p.handleCollisionOnBottomSide(handlerData, tileHitbox, velocityMut, positionMut)
				p.handleCollisionOnTopSide(handlerData, tileHitbox, velocityMut, positionMut)
				p.handleCollisionOnLeftSide(handlerData, tileHitbox, velocityMut, positionMut)
				p.handleCollisionOnRightSide(handlerData, tileHitbox, velocityMut, positionMut)
			}
		}
	}
}

func (p *physicsHandlerImpl) handleCollisionOnBottomSide(
	handlerData *physicsHandlerData,
	tileHitbox rl.Rectangle,
	velocityMut *rl.Vector2,
	positionMut *rl.Vector2,
) {

	entityCollider := rl.NewRectangle(
		handlerData.Hitbox.X,
		handlerData.Hitbox.Y+handlerData.Hitbox.Height-tileColliderHeight,
		handlerData.Hitbox.Width,
		tileColliderHeight,
	)

	tileCollider := rl.NewRectangle(
		tileHitbox.X,
		tileHitbox.Y-tileColliderHeight,
		tileHitbox.Width,
		tileColliderHeight,
	)

	if os.Getenv(engine_shared.DebugModeEnvironmentVariable) == "true" {
		rl.DrawRectangleLines(
			int32(tileCollider.X),
			int32(tileCollider.Y),
			int32(tileCollider.Width),
			int32(tileCollider.Height),
			rl.Purple,
		)
	}

	if rl.CheckCollisionRecs(entityCollider, tileCollider) && velocityMut.Y < 0 {
		velocityMut.Y = 0

		positionMut.Y = tileHitbox.Y - (handlerData.HitboxDimensions.Height + handlerData.HitboxDimensions.Y)
	}
}

func (p *physicsHandlerImpl) handleCollisionOnTopSide(
	handlerData *physicsHandlerData,
	tileHitbox rl.Rectangle,
	velocity *rl.Vector2,
	positionMut *rl.Vector2,
) {

	entityCollider := rl.NewRectangle(
		handlerData.Hitbox.X,
		handlerData.Hitbox.Y,
		handlerData.Hitbox.Width,
		tileColliderHeight,
	)

	tileCollider := rl.NewRectangle(
		tileHitbox.X,
		tileHitbox.Y+tileHitbox.Height,
		tileHitbox.Width,
		tileColliderHeight,
	)

	bottomSideCollides := rl.CheckCollisionRecs(entityCollider, tileCollider)

	if bottomSideCollides && velocity.Y > 0 {
		velocity.Y = 0

		positionMut.Y = tileCollider.Y - handlerData.HitboxDimensions.Y
	}
}

func (p *physicsHandlerImpl) handleCollisionOnLeftSide(
	handlerData *physicsHandlerData,
	tileHitbox rl.Rectangle,
	velocity *rl.Vector2,
	positionMut *rl.Vector2,
) {

	entityCollider := rl.NewRectangle(
		handlerData.Hitbox.X,
		handlerData.Hitbox.Y,
		tileColliderWidth,
		handlerData.Hitbox.Height,
	)

	tileCollider := rl.NewRectangle(
		tileHitbox.X+tileHitbox.Width-tileColliderWidth,
		tileHitbox.Y,
		tileColliderWidth,
		tileHitbox.Height,
	)

	bottomSideCollides := rl.CheckCollisionRecs(entityCollider, tileCollider)

	if bottomSideCollides && velocity.X < 0 {
		velocity.X = 0

		positionMut.X = tileHitbox.X + tileHitbox.Width - handlerData.HitboxDimensions.X - tileColliderWidth
	}
}

func (p *physicsHandlerImpl) handleCollisionOnRightSide(
	handlerData *physicsHandlerData,
	tileHitbox rl.Rectangle,
	velocity *rl.Vector2,
	positionMut *rl.Vector2,
) {

	entityCollider := rl.NewRectangle(
		handlerData.Hitbox.X+handlerData.Hitbox.Width,
		handlerData.Hitbox.Y,
		tileColliderWidth,
		handlerData.Hitbox.Height,
	)

	tileCollider := rl.NewRectangle(
		tileHitbox.X-tileColliderWidth,
		tileHitbox.Y,
		tileColliderWidth,
		tileHitbox.Height,
	)

	bottomSideCollides := rl.CheckCollisionRecs(entityCollider, tileCollider)

	if bottomSideCollides && velocity.X > 0 {
		velocity.X = 0

		positionMut.X = tileHitbox.X - handlerData.HitboxDimensions.Width - handlerData.HitboxDimensions.X - tileColliderWidth
	}
}

func (p *physicsHandlerImpl) isPlayerBottomTouchingTile(playerHitbox *rl.Rectangle, tileHitbox *rl.Rectangle) bool {

	playerCollider := rl.NewRectangle(
		playerHitbox.X,
		playerHitbox.Y+playerHitbox.Height-tileColliderHeight,
		playerHitbox.Width,
		tileColliderHeight,
	)

	tileCollider := rl.NewRectangle(
		tileHitbox.X,
		tileHitbox.Y-tileColliderHeight,
		tileHitbox.Width,
		tileColliderHeight,
	)

	return rl.CheckCollisionRecs(playerCollider, tileCollider)
}
