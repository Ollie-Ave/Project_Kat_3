package levels

import (
	"testing"
)

func TestGetTileSetByIdWith12(t *testing.T) {
	testGetTileSetById(12, 10, t)
}

func TestGetTileSetByIdWith5(t *testing.T) {
	testGetTileSetById(5, 1, t)
}

func TestGetTileSetByIdWith15(t *testing.T) {
	testGetTileSetById(15, 15, t)
}

func TestGetTileSetByIdWith14(t *testing.T) {
	testGetTileSetById(14, 10, t)
}

func TestGetTileSetByIdWith20(t *testing.T) {
	testGetTileSetById(20, 15, t)
}

func testGetTileSetById(tileSetId, expectedFirstGid int, t *testing.T) {
	level := &levelOne{
		levelData: &levelData{
			Layers: []*layer{},
			TileSets: []*tileSet{
				{FirstGid: 1},
				{FirstGid: 10},
				{FirstGid: 15}},
		},
	}

	tileSet := level.getTileSetById(tileSetId)

	if tileSet == nil || tileSet.FirstGid != expectedFirstGid {
		t.Fatalf("Expected tileSet.FirstGid to be %v, got %v", expectedFirstGid, tileSet.FirstGid)
	}
}

func TestGetNewTilePositionWhenShouldIncreaseWith10(t *testing.T) {
	testGetNewTilePosition(10, true, t)
}

func TestGetNewTilePositionWhenShouldNotIncreaseWith7(t *testing.T) {
	testGetNewTilePosition(7, false, t)
}

func TestGetNewTilePositionWhenShouldNotIncreaseWith0(t *testing.T) {
	testGetNewTilePosition(0, false, t)
}

func testGetNewTilePosition(x int, yShouldIncrease bool, t *testing.T) {
	x, y := getNewTilePosition(x, 0, 10)

	if yShouldIncrease && (y != 1 || x != 0) {
		t.Fatalf("Expected y to be 1 and x to be 0, got y: %v and x: %v", y, x)
	} else if !yShouldIncrease && y != 0 {
		t.Fatalf("Expected y to be 0, got %v", y)
	}
}
