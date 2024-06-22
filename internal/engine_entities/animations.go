package engine_entities

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimationHandler interface {
	RenderAnimationFrame(position rl.Vector2)
	HandleAnimation()
	LoadAnimations(animations map[int]AnimationData) error
	SetDirection(direction string) error
	SetAnimation(animation int)
}

type AnimationHandlerImpl struct {
	AnimationMeta *animationMeta

	Animations map[int]AnimationData
}

func NewAnimationHandler(defaultAnimation int) AnimationHandler {
	animationMeta := &animationMeta{
		CurrentAnimationFrame:    0,
		CurrentAnimation:         defaultAnimation,
		TimeSinceLastFrameChange: 0,
	}

	return &AnimationHandlerImpl{
		AnimationMeta: animationMeta,
	}
}

type animationMeta struct {
	CurrentAnimationFrame int
	CurrentAnimation      int

	TimeSinceLastFrameChange int

	Direction bool
}

type AnimationData struct {
	Texture rl.Texture2D

	TexturePath string
	FrameSize   rl.Vector2
	FrameCount  int

	TimeBetweenFrameChange int

	Loop bool
}

func NewAnimationData(
	texturePath string,
	frameCount int,
	timeBetweenFrameChange int,
) AnimationData {
	return AnimationData{
		TexturePath:            texturePath,
		FrameCount:             frameCount,
		TimeBetweenFrameChange: timeBetweenFrameChange,
		Loop:                   true,
	}
}

func NewAnimationDataEx(
	texturePath string,
	frameCount int,
	timeBetweenFrameChange int,
	loop bool,
) AnimationData {
	return AnimationData{
		TexturePath:            texturePath,
		FrameCount:             frameCount,
		TimeBetweenFrameChange: timeBetweenFrameChange,
		Loop:                   loop,
	}
}

func (a *AnimationHandlerImpl) RenderAnimationFrame(position rl.Vector2) {
	currentAnimation := a.Animations[a.AnimationMeta.CurrentAnimation]

	textureSourceRec := rl.NewRectangle(
		currentAnimation.FrameSize.X*float32(a.AnimationMeta.CurrentAnimationFrame),
		0,
		currentAnimation.FrameSize.X,
		currentAnimation.FrameSize.Y)

	if a.AnimationMeta.Direction {
		textureSourceRec = rl.NewRectangle(
			textureSourceRec.X,
			textureSourceRec.Y,
			-textureSourceRec.Width,
			textureSourceRec.Height,
		)
	}

	rl.DrawTextureRec(
		currentAnimation.Texture,
		textureSourceRec,
		position,
		rl.White,
	)
}

func (a *AnimationHandlerImpl) HandleAnimation() {
	a.AnimationMeta.TimeSinceLastFrameChange += int(rl.GetFrameTime() * 1000)
	fmt.Println(a.AnimationMeta.CurrentAnimationFrame)

	if a.AnimationMeta.TimeSinceLastFrameChange >= a.Animations[a.AnimationMeta.CurrentAnimation].TimeBetweenFrameChange {
		if a.AnimationMeta.CurrentAnimationFrame < a.Animations[a.AnimationMeta.CurrentAnimation].FrameCount-1 {
			a.AnimationMeta.CurrentAnimationFrame++
		}

		a.AnimationMeta.TimeSinceLastFrameChange = 0
	}

	if a.Animations[a.AnimationMeta.CurrentAnimation].Loop &&
		a.AnimationMeta.CurrentAnimationFrame == a.Animations[a.AnimationMeta.CurrentAnimation].FrameCount-1 {

		a.AnimationMeta.CurrentAnimationFrame = 0
	}
}

func (a *AnimationHandlerImpl) LoadAnimations(animations map[int]AnimationData) error {

	for key, animation := range animations {
		const playerSpritePath = "entities/player"
		playerSpriteFullPath := fmt.Sprintf(
			"%s/%s/%s",
			engine_shared.AssetsPath,
			playerSpritePath,
			animation.TexturePath)

		_, err := os.Stat(playerSpriteFullPath)

		if err != nil {
			return err
		}

		animation.Texture = rl.LoadTexture(playerSpriteFullPath)

		animation.FrameSize = rl.NewVector2(
			float32(animation.Texture.Width/int32(animation.FrameCount)),
			float32(animation.Texture.Height))

		animations[key] = animation
	}

	a.Animations = animations

	return nil
}

func (a *AnimationHandlerImpl) SetDirection(direction string) error {
	if direction == "left" {
		a.AnimationMeta.Direction = true

		return nil
	} else if direction == "right" {
		a.AnimationMeta.Direction = false

		return nil
	}

	return fmt.Errorf("invalid direction: %s", direction)
}

func (a *AnimationHandlerImpl) SetAnimation(animation int) {
	if animation == a.AnimationMeta.CurrentAnimation {
		return
	}

	a.AnimationMeta.CurrentAnimation = animation
	a.AnimationMeta.CurrentAnimationFrame = 0
}
