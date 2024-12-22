package noiseMap

import (
	"errors"
	"github.com/KEINOS/go-noise"
	"math"
)

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

func NullMap() WorldMap {
	return WorldMap{Seed: 0, Height: 0, Width: 0, MapData: nil}
}

func MapCreation(params MapCreationParams) (WorldMap, error) {

	output := make(map[int]map[int]float64)

	genNoise, err := noise.New(params.NoiseType, params.Seed)

	if err == nil {
		for y := 0; y < params.Height; y++ {
			yy := float64(y) / params.Smoothness
			line := make(map[int]float64)

			for x := 0; x < params.Width; x++ {
				xx := float64(x) / params.Smoothness

				n := genNoise.Eval64(xx, yy)
				if params.WaterSmoothness > 0 {
					n += waterBorder(float64(x), float64(params.Width), params.WaterSmoothness) +
						waterBorder(float64(y), float64(params.Height), params.WaterSmoothness)
				}
				line[x] = n
			}
			output[y] = line
		}
		return WorldMap{
				params.Seed,
				params.Height,
				params.Width,
				output},
			nil
	}
	return NullMap(), errors.New("noise generation failed")
}

func waterBorder(x float64, width float64, waterSmooth float64) float64 {
	return -(math.Pow(math.Abs(x-(width/2))*(2/width), waterSmooth))
}
