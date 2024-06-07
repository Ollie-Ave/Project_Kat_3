package levels

import (
	"cmp"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newLevelRenderer() levelRenderer {
	return &levelRendererImpl{}
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

type levelRenderer interface {
	Render(levelData *levelData)
}

type levelRendererImpl struct {
}

func (l *levelRendererImpl) Render(levelData *levelData) {
	const tileLayer = "tilelayer"
	const imageLayer = "imagelayer"

	for _, layer := range levelData.Layers {
		if layer.LayerType == tileLayer {
			l.renderTileLayer(layer, levelData)
		} else if layer.LayerType == imageLayer {
			l.renderImageLayer(layer)
		}
	}
}

func (l *levelRendererImpl) renderImageLayer(layer *layer) {
	position := rl.NewVector2(0, 200)
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
}

func (l *levelRendererImpl) renderTileLayer(layer *layer, levelData *levelData) {
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

func (l *levelRendererImpl) getTileData(tileId int, levelData *levelData) *tile {
	tileSet := l.getTileSetById(tileId, levelData)
	tileX, tileY := l.getTilePositionByTileId(tileId, tileSet)

	return &tile{
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

func (l *levelRendererImpl) getTileSetById(id int, levelData *levelData) *tileSet {
	var returnValue *tileSet

	slices.SortFunc(levelData.TileSets, func(a, b *tileSet) int {
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

func (l *levelRendererImpl) getTilePositionByTileId(id int, tileSet *tileSet) (int, int) {
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