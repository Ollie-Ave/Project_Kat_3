package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PastTimePeriod        = "Past"
	FutureTimePeriod      = "Future"
	SpawnerLayerIndicator = "Spawner"
)

type LevelHandler interface {
	Initialise() error
	Update()
	GetTimePeriod() string
	engine_shared.Renderable
}

type LevelData struct {
	PastPeriod        *LevelTimePeriod
	FuturePeriod      *LevelTimePeriod
	CurrentTimePeriod *LevelTimePeriod
}

type LevelTimePeriod struct {
	Layers   []*Layer
	TileSets []*TileSet

	Width  int
	Height int

	TileWidth  int
	TileHeight int
}

type Layer struct {
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

type TileSet struct {
	FirstGid int

	Image   string
	Texture rl.Texture2D

	TileHeight int
	TileWidth  int

	Columns int
}

type Tile struct {
	Texture    rl.Texture2D
	Height     int
	Width      int
	TextureRec rl.Rectangle
}
