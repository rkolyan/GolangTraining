package awesomeProject

import (
	"image"
	"image/color"
	"math"
)

type ConcreteTextureGenerator struct {
	rgba     *image.RGBA
	normal   *image.RGBA
	bump     *image.Gray
	vertices [][]float64
}

func subf(x float64) float64 {
	return (math.Sin(x) - math.Sin(x*math.Sin(x)) + math.Cos(x+x*math.Cos(x))) / 3
}

func f(x, y float64) float64 {
	z := subf(x) + subf(y)/2
	if z < -1 {
		z = -1
	} else if z > 1 {
		z = 1
	}
	return z
}

func (generator ConcreteTextureGenerator) SetRect(r image.Rectangle) {
	generator.rgba = image.NewRGBA(r)
	generator.vertices = make([][]float64, r.Max.Y+1)
	for i := range generator.vertices {
		generator.vertices[i] = make([]float64, r.Max.X+1)
	}
	for y := range generator.vertices {
		for x := range generator.vertices[y] {
			generator.vertices[y][x] = f(float64(x), float64(y))
		}
	}
}

func (generator ConcreteTextureGenerator) createTextureMap() TextureMap {
	for x := generator.rgba.Rect.Min.X; x < generator.rgba.Rect.Max.X; x++ {
		for y := generator.rgba.Rect.Min.Y; y < generator.rgba.Rect.Max.Y; y++ {
			generator.rgba.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
		}
	}
	return generator.rgba
}

type Point3D struct {
	x, y, z float64
}

func (generator ConcreteTextureGenerator) createNormalMap() NormalMap {
	generator.normal = image.NewRGBA(generator.rgba.Rect)
	max_y := generator.normal.Rect.Max.Y
	max_x := generator.normal.Rect.Min.X
	vector1 := Point3D{x: 0, y: 0, z: 0}
	vector2 := Point3D{0, 0, 0}
	vector_result1 := Point3D{0, 0, 0}
	vector_result2 := Point3D{0, 0, 0}
	normal_vector := Point3D{0, 0, 0}
	var norma = float64(0)
	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			vector1.x = 1
			vector1.y = 0
			vector1.z = generator.vertices[y][x+1] - generator.vertices[y][x]
			vector2.x = 0
			vector2.y = 1
			vector2.z = generator.vertices[y+1][x] - generator.vertices[y][x]
			vector_result1.x = vector1.y*vector2.z - vector2.y*vector1.z
			vector_result1.y = vector1.z*vector2.x - vector2.z*vector1.x
			vector_result1.z = vector1.x*vector2.y - vector2.x*vector1.y
			vector1.x = -1
			vector1.y = 0
			vector1.z = generator.vertices[y+1][x] - generator.vertices[y+1][x+1]
			vector2.x = 0
			vector2.y = -1
			vector2.z = generator.vertices[y][x+1] - generator.vertices[y+1][x+1]
			vector_result2.x = vector1.y*vector2.z - vector2.y*vector1.z
			vector_result2.y = vector1.z*vector2.x - vector2.z*vector1.x
			vector_result2.z = vector1.x*vector2.y - vector2.x*vector1.y
			normal_vector.x = (vector_result1.x + vector_result2.x) / 2
			normal_vector.y = (vector_result1.y + vector_result2.y) / 2
			normal_vector.z = (vector_result1.z + vector_result2.z) / 2
			norma = math.Sqrt(normal_vector.x*normal_vector.x + normal_vector.y*normal_vector.y + normal_vector.z*normal_vector.z)
			normal_vector.x = normal_vector.x / norma
			normal_vector.y = normal_vector.y / norma
			normal_vector.z = normal_vector.z / norma
			normal_vector.x = (normal_vector.x + 1) / 2
			normal_vector.y = (normal_vector.y + 1) / 2
			normal_vector.z = (normal_vector.z + 1) / 2
			generator.normal.Set(x, y, color.RGBA{uint8(255 * normal_vector.x), uint8(255 * normal_vector.y), uint8(255 * normal_vector.z), 255})
		}
	}
	return generator.normal
}

func (generator ConcreteTextureGenerator) createHeightMap() HeightMap {
	generator.bump = image.NewGray(generator.rgba.Rect)
	return generator.bump
}
