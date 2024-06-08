package engine_levels

import (
	"cmp"
	"slices"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewLevelRenderer(entityManager engine_entities.EntityManager) LevelRenderer {
	return &levelRendererImpl{
		entityManager: entityManager,
	}
}

func getNewTilePosition(x, y int, maxX int) (int, int) {
	if x < (maxX - 1) {
		x++
	} else {
		x = 0
		y++
	}

	return x, y
}

type LevelRenderer interface {
	Render(levelData *LevelData)
}

type levelRendererImpl struct {
	entityManager engine_entities.EntityManager
}

func (l *levelRendererImpl) Render(levelData *LevelData) {

	for _, layer := range levelData.Layers {
		if layer.LayerType == engine_shared.TileLayer {
			l.renderTileLayer(layer, levelData)
		} else if layer.LayerType == engine_shared.ImageLayer {
			l.renderImageLayer(layer)
		}
	}
}

func (l *levelRendererImpl) renderImageLayer(layer *Layer) error {
	cameraEntity, err := l.entityManager.GetCamera()

	if err != nil {
		return err
	}

	camera := cameraEntity.GetCamera()

	const backgroundSpeed = 7.5
	xPos := -(camera.Offset.X / backgroundSpeed) * layer.Parallaxx

	position := rl.NewVector2(xPos, 200)
	rotation := 0
	scale := 1.0

	for index := 0; index < 4; index++ {
		rl.DrawTextureEx(
			layer.ImageTexture,
			position,
			float32(rotation),
			float32(scale),
			rl.White)

		position.X += float32(layer.ImageTexture.Width)
	}

	position = rl.NewVector2(0, 0)

	if layer.Name == "Sky" {
		for index := 0; index < 4; index++ {
			for index := 0; index < 4; index++ {
				rl.DrawTextureEx(
					layer.ImageTexture,
					position,
					float32(rotation),
					float32(scale),
					rl.White)

				position.X += float32(layer.ImageTexture.Width)
			}

			position.Y += float32(layer.ImageTexture.Height)
			position.X = 0
		}
	}

	return nil
}

func (l *levelRendererImpl) renderTileLayer(layer *Layer, levelData *LevelData) {
	x := -1
	y := 0

	for _, tileId := range layer.Data {
		x, y = getNewTilePosition(x, y, layer.Width)

		if tileId == 0 {
			continue
		}

		tile := l.getTileData(tileId, levelData)

		tilePosition := rl.NewVector2(
			float32(x*tile.Width),
			float32(y*tile.Height))

		rl.DrawTextureRec(
			tile.Texture,
			tile.TextureRec,
			tilePosition,
			rl.White)
	}
}

func (l *levelRendererImpl) getTileData(tileId int, levelData *LevelData) *Tile {
	tileSet := l.getTileSetById(tileId, levelData)
	tileX, tileY := l.getTilePositionByTileId(tileId, tileSet)

	return &Tile{
		Texture: tileSet.Texture,
		Height:  tileSet.TileHeight,
		Width:   tileSet.TileWidth,
		TextureRec: rl.NewRectangle(
			float32(tileX),
			float32(tileY),
			float32(tileSet.TileWidth),
			float32(tileSet.TileHeight)),
	}
}

func (l *levelRendererImpl) getTileSetById(id int, levelData *LevelData) *TileSet {
	var returnValue *TileSet

	slices.SortFunc(levelData.TileSets, func(a, b *TileSet) int {
		return cmp.Compare(a.FirstGid, b.FirstGid)
	})

	for _, tileSet := range levelData.TileSets {
		if tileSet.FirstGid > id {
			return returnValue
		}

		returnValue = tileSet
	}

	return returnValue
}

func (l *levelRendererImpl) getTilePositionByTileId(id int, tileSet *TileSet) (int, int) {
	x, y := 0, 0

	for i := tileSet.FirstGid; i < id; i++ {
		x++

		if x == tileSet.Columns {
			x = 0
			y++
		}
	}

	return x * tileSet.TileWidth, y * tileSet.TileHeight
}
