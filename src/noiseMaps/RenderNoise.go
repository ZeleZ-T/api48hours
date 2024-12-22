package noiseMap

import (
	"errors"
	"github.com/mazznoer/colorgrad"
	"image"
	"image/color"
	"math"
)

func MapPlainImage(worldMap *WorldMap) (image.Image, error) {
	MapData := worldMap.MapData

	colorSet := map[string]color.Color{
		"white":       color.White,
		"gray":        color.RGBA{R: 130, G: 130, B: 130, A: 255},
		"orange":      color.RGBA{R: 253, G: 174, B: 97, A: 255},
		"lightOrange": color.RGBA{R: 254, G: 224, B: 139, A: 255},
		"darkGreen":   color.RGBA{R: 102, G: 194, B: 165, A: 255},
		"green":       color.RGBA{R: 171, G: 221, B: 164, A: 255},
		"lightGreen":  color.RGBA{R: 230, G: 245, B: 152, A: 255},
		"yellow":      color.RGBA{R: 255, G: 255, B: 191, A: 255},
		"blue":        color.RGBA{R: 50, G: 136, B: 189, A: 255},
		"darkBlue":    color.RGBA{R: 94, G: 79, B: 162, A: 255},
	}

	width := len(MapData[0])
	height := len(MapData)

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for i := 0; i < len(MapData); i++ {
		for k := 0; k < len(MapData[i]); k++ {
			cord := MapData[i][k]
			var pixel color.Color
			switch {
			case cord < -0.7:
				pixel = colorSet["darkBlue"]
			case -0.7 <= cord && cord < -0.5:
				pixel = colorSet["blue"]
			case -0.5 <= cord && cord < -0.35:
				pixel = colorSet["yellow"]
			case -0.35 <= cord && cord < -0.15:
				pixel = colorSet["lightGreen"]
			case -0.15 <= cord && cord < 0:
				pixel = colorSet["green"]
			case 0 <= cord && cord < 0.2:
				pixel = colorSet["darkGreen"]
			case 0.2 <= cord && cord < 0.55:
				pixel = colorSet["gray"]
			case 0.55 <= cord:
				pixel = colorSet["white"]
			default:
				return nil, errors.New("image generation failed")
			}
			img.Set(k, i, pixel)
		}
	}
	return img, nil
}

func MapGradientImage(worldMap *WorldMap) (image.Image, error) {
	gradient, _ := colorgrad.NewGradient().
		Colors(
			colorgrad.Rgb8(77, 66, 155, 255),
			colorgrad.Rgb8(94, 79, 162, 255),
			colorgrad.Rgb8(50, 136, 189, 255),
			colorgrad.Rgb8(255, 255, 191, 255),
			colorgrad.Rgb8(230, 245, 152, 255),
			colorgrad.Rgb8(171, 221, 164, 255),
			colorgrad.Rgb8(102, 194, 165, 255),
			colorgrad.Rgb8(150, 150, 150, 255),
			colorgrad.Rgb8(150, 150, 150, 255),
			colorgrad.Rgb8(255, 255, 255, 255),
			colorgrad.Rgb8(255, 255, 255, 255),
		).Build()
	gradient.Sharp(100, 0)

	width := len(worldMap.MapData[0])
	height := len(worldMap.MapData)

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for i := 0; i < len(worldMap.MapData); i++ {
		for k := 0; k < len(worldMap.MapData[i]); k++ {
			cord := (worldMap.MapData[i][k] + 1) / 2
			if cord < 0 {
				cord = 0
			} else if cord > 0.9 {
				cord = 1
			}
			pixel := gradient.At(math.Round(cord*10) / 10)
			img.Set(k, i, pixel)
		}
	}
	return img, nil
}
