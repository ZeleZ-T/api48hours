package models

import "github.com/KEINOS/go-noise"

type MapCreationParams struct {
	Seed            int64
	Height          int
	Width           int
	Smoothness      float64
	WaterSmoothness float64
	NoiseType       noise.Algo
}

type WorldMap struct {
	Seed    int64
	Height  int
	Width   int
	MapData map[int]map[int]float64
}

type MapSaveParams struct {
	UserEmail      string
	Name           string
	CreationParams MapCreationParams
}

type MapUpdateParams struct {
	UserEmail string
	MapName   string
	NewName   string
}
