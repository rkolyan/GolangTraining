package main

import (
	"image"
	"image/png"
	"os"
)

type NormalMap image.Image
type HeightMap image.Image
type TextureMap image.Image

type TextureGenerator interface {
	SetRect(r image.Rectangle)
	createNormalMap() NormalMap
	createHeightMap() HeightMap
	createTextureMap() TextureMap
}

func generateTextures(generator TextureGenerator) (TextureMap, NormalMap, HeightMap) {
	return generator.createTextureMap(), generator.createNormalMap(), generator.createHeightMap()
}

func main() {
	generator := ConcreteTextureGenerator{}
	generator.SetRect(image.Rect(0, 0, 800, 800))
	tm, nm, hm := generateTextures(generator)
	file, err := os.Create("texture.png")
	if err != nil {
		return
	}
	png.Encode(file, tm)
	file.Close()

	file, err = os.Create("normal.png")
	if err != nil {
		return
	}
	png.Encode(file, nm)
	file.Close()

	file, err = os.Create("height.png")
	if err != nil {
		return
	}
	png.Encode(file, hm)
	file.Close()
}
