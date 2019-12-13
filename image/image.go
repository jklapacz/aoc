package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
)

type Pixels = []int

type Layers = []Pixels

// CountLayers takes a stream of Pixels and the image dimensions returning the count of layers
func CountLayers(pixels Pixels, width, height int) int {
	layerSize := width * height
	totalLength := len(pixels)
	return totalLength / layerSize
}

func PixelsToLayers(pixels Pixels, width, height int) Layers {
	layerCount := CountLayers(pixels, width, height)
	layerLength := width * height

	layers := make(Layers, layerCount)
	layerIdx := 0
	for offset := 0; offset < len(pixels); offset += layerLength {
		currentLayer := pixels[offset : offset+layerLength]
		layers[layerIdx] = Pixels(currentLayer)
		layerIdx++
	}
	return layers
}

func StringToPixels(input string) Pixels {
	pixels := make(Pixels, len(input))
	for pixelIdx, c := range input {
		pixelVal, err := strconv.ParseInt(string(c), 10, 32)
		if err != nil {
			log.Fatal("bad input!")
		}
		pixels[pixelIdx] = int(pixelVal)
	}
	return pixels
}

func CountZeros(layer Pixels) int {
	return generateDigitCount(layer)[0]
}

func CountOnes(layer Pixels) int {
	return generateDigitCount(layer)[1]
}

func CountTwos(layer Pixels) int {
	return generateDigitCount(layer)[2]
}

func LayerWithLeastZeros(layers Layers) int {
	var minLayerIdx int
	currentZeroMin := int(math.MaxInt32)
	for layerIdx, layer := range layers {
		zeroCount := CountZeros(layer)
		if zeroCount < currentZeroMin {
			minLayerIdx = layerIdx
			currentZeroMin = zeroCount
		}
	}
	return minLayerIdx
}

type digitCount = map[int]int

func generateDigitCount(layer Pixels) digitCount {
	counts := make(digitCount, 10)
	for _, digitValue := range layer {
		counts[digitValue]++
	}
	return counts
}

func checksum(layers Layers) int {
	minLayer := layers[LayerWithLeastZeros(layers)]
	return CountOnes(minLayer) * CountTwos(minLayer)
}

func Checksum(input string, width, height int) int {
	layers := PixelsToLayers(StringToPixels(input), width, height)
	return checksum(layers)
}

type Image struct {
	layers        Layers
	width, height int
	decoded       Pixels
}

func CreateImage(input string, width, height int) *Image {
	return &Image{
		PixelsToLayers(StringToPixels(input), width, height),
		width,
		height,
		make(Pixels, width*height),
	}
}

func (i *Image) Decode() Pixels {
	var currentPixel int
	for xIdx := 0; xIdx < i.width; xIdx++ {
		for yIdx := 0; yIdx < i.height; yIdx++ {
			//fmt.Printf("Decoding: (%v, %v)[%v]\n", xIdx, yIdx, currentPixel)
			i.decoded[currentPixel] = i.DecodeLayersAt(currentPixel)
			currentPixel++
		}
	}
	//fmt.Println(i.decoded)
	return i.decoded
}

func (i *Image) DecodeLayersAt(pixel int) int {
	pixelVal := -1
	for _, layer := range i.layers {
		//fmt.Printf("at value %v, layer %v, pixel is %v\n", pixel, layerIdx, layer[pixel])
		switch layer[pixel] {
		case 0:
			if pixelVal < 0 {
				pixelVal = 0
			}
		case 1:
			if pixelVal < 0 {
				pixelVal = 1
			}
		}
	}
	return pixelVal
}

func (i *Image) Print() {
	var currentPixel int
	for x := 0; x < i.width+2; x++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	for y := 0; y < i.height; y++ {
		fmt.Printf("|")
		for x := 0; x < i.width; x++ {
			if i.decoded[currentPixel] == 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}

			currentPixel++

		}
		fmt.Printf("|\n")
	}
	for x := 0; x < i.width+2; x++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

}

func (i *Image) ToPNG() {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{i.width, i.height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	var currentPixel int
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			switch i.decoded[currentPixel] {
			case 0:
				img.Set(x, y, color.Black)
			case 1:
				img.Set(x, y, color.White)
			}
			currentPixel++
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
