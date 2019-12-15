package image_test

import (
	"testing"

	"github.com/jklapacz/aoc/image"
	"github.com/stretchr/testify/assert"
)

const exampleInput = "123456789012"

func TestCountLayers(t *testing.T) {
	width := 3
	height := 2

	assert.Equal(t, 2, image.CountLayers(image.StringToPixels(exampleInput), width, height))
}

func TestStringToPixels(t *testing.T) {
	input := "1234"
	assert.Equal(t, image.Pixels{1, 2, 3, 4}, image.StringToPixels(input))
}

func TestStringToLayers(t *testing.T) {
	layers := image.PixelsToLayers(image.StringToPixels(exampleInput), 3, 2)
	assert.Equal(t, 2, len(layers))
}

func TestLayerWithLeastZeros(t *testing.T) {
	input := "11102200"
	layers := image.PixelsToLayers(image.StringToPixels(input), 2, 2)
	assert.Equal(t, 0, image.LayerWithLeastZeros(layers))
}

func TestChecksum(t *testing.T) {
	smallInput := "11220112"
	assert.Equal(t, 4, image.Checksum(smallInput, 2, 2))
	assert.Equal(t, 1, image.Checksum(exampleInput, 3, 2))
}

func TestImage_Decode(t *testing.T) {
	input := "0222112222120000"
	i := image.CreateImage(input, 2, 2)
	assert.Equal(t, image.Pixels{0, 1, 1, 0}, i.Decode())
	i.Print()
	i.ToPNG()
}
