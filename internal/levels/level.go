package levels

import rl "github.com/gen2brain/raylib-go/raylib"

type levelData struct {
	Layers   []*layer
	TileSets []*tileSet

	Width     int
	TileWidth int
}

type layer struct {
	Id int

	Data []int

	ImagePath    string `json:"image"`
	ImageTexture rl.Texture2D

	Offsetx float32
	Offsety float32

	Height int
	Width  int

	Parallaxx float32

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
