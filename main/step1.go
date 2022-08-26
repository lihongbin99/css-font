package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Tag struct {
	tag    string
	offset uint32
	length uint32
}

var (
	fontData []byte
	font     *truetype.Font
	tags     = make(map[string]Tag, 0)

	fontDPI  = 96 // 屏幕每英寸的分辨率
	fontSize = 96 // 字体尺寸
	maxCount = 30
)

func step1() {
	// 解析范围
	lines := make([]string, 0)
	lineCount := 0
	line := ""
	for i := 0; i <= 65535; i++ {
		if hasChar(i) {
			lineCount++
			line += string(rune(i))
			if lineCount >= maxCount {
				lines = append(lines, line)
				lineCount = 0
				line = ""
			}
		}
	}
	if lineCount > 0 {
		lines = append(lines, line)
	}

	_ = os.Mkdir(path, 0700)
	//for i := 0; i <= len(lines); i++ {
	//	doMain(lines[i:i+1], i)
	//}
	base := 20
	for i := 0; i <= len(lines)/base; i++ {
		if len(lines) > 0 {
			end := i*base + base
			if end > len(lines) {
				end = len(lines)
			}
			doMain(lines[i*base:end], i)
		}
	}
}

func doMain(lines []string, i int) {
	imgFile, _ := os.Create(fmt.Sprintf("%s/word_%d.jpg", path, i))
	defer func() {
		_ = imgFile.Close()
	}()

	dy := int(float64(len(lines)*fontSize) * 1.7)
	dx := int(float64(maxCount*fontSize) * 1.35)
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	// 画背景
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
		}
	}

	c := freetype.NewContext()
	c.SetDPI(float64(fontDPI))
	c.SetFont(font)
	c.SetFontSize(float64(fontSize))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	for i, line := range lines {
		if _, err := c.DrawString(line, freetype.Pt(0, (i+1)*int(c.PointToFixed(float64(fontSize))>>6)+((i+1)*25))); err != nil {
			panic(err)
		}
	}

	// 以PNG格式保存文件
	if err := jpeg.Encode(imgFile, img, nil); err != nil {
		panic(err)
	}
}
