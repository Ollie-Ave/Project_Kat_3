package levels

import rl "github.com/gen2brain/raylib-go/raylib"

type levelData struct {
	Layers   []*layer
	TileSets []*tileSet
}

type layer struct {
	Id int

	Data  []int
	Image string

	Height int
	Width  int

	Name      string
	LayerType string `json:"type"`
}

type tileSet struct {
	FirstGid int

	Image   string
	Texture rl.Texture2D

	TileHeight int
	TileWidth  int

	Columns int
}

type tile struct {
	Texture    rl.Texture2D
	Height     int
	Width      int
	TextureRec rl.Rectangle
}
