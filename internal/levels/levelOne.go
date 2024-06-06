package levels

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	assetsPath  = "/home/oliver/Code/Project_Kat_3/assets"
	levelPath   = "levels/export"
	tileSrcPath = "tileSets/src"
)

func initLevelOne() (*levelOne, error) {
	levelOneFilePath := fmt.Sprintf("%s/%s/testLevel.json", assetsPath, levelPath)

	levelData, err := loadLevelData(&levelOneFilePath)

	if err != nil {
		return nil, err
	}

	return &levelOne{
		levelData: levelData,
	}, nil
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

func loadLevelData(filePath *string) (*levelData, error) {
	var levelData *levelData

	levelFile, err := os.ReadFile(*filePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(levelFile, &levelData)

	if err != nil {
		return nil, err
	}

	for _, tileSet := range levelData.TileSets {
		tileTexurePath := fmt.Sprintf("%s/%s/%s", assetsPath, tileSrcPath, tileSet.Image)

		_, err := os.Stat(tileTexurePath)

		if err != nil {
			return nil, err
		}

		tileSet.Texture = rl.LoadTexture(tileTexurePath)
	}

	return levelData, nil
}

type levelOne struct {
	levelData *levelData
}

func (l *levelOne) Render() {
	const tileLayer = "tilelayer"

	slices.SortFunc(l.levelData.Layers, func(a, b *layer) int {
		return cmp.Compare(a.Id, b.Id)
	})

	for _, layer := range l.levelData.Layers {
		if layer.LayerType == tileLayer {
			l.renderTileLayer(layer)
		}
	}
}

func (l *levelOne) renderTileLayer(layer *layer) {
	x := -1
	y := 0

	for _, tileId := range layer.Data {
		x, y = getNewTilePosition(x, y, layer.Width)

		if tileId == 0 {
			continue
		}

		tile := l.getTileData(tileId)

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

func (l *levelOne) getTileData(tileId int) *tile {
	tileSet := l.getTileSetById(tileId)
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

func (l *levelOne) getTileSetById(id int) *tileSet {
	var returnValue *tileSet

	slices.SortFunc(l.levelData.TileSets, func(a, b *tileSet) int {
		return cmp.Compare(a.FirstGid, b.FirstGid)
	})

	for _, tileSet := range l.levelData.TileSets {
		if tileSet.FirstGid > id {
			return returnValue
		}

		returnValue = tileSet
	}

	return returnValue
}

func (l *levelOne) getTilePositionByTileId(id int, tileSet *tileSet) (int, int) {
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
