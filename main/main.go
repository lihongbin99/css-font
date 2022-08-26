package main

import (
	"css-font/common/utils"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"os"
)

var (
	fontFile string
	path     string
)

func init() {
	flag.StringVar(&fontFile, "f", fontFile, "font file")
	flag.StringVar(&path, "p", path, "path")
	flag.Parse()
}

func main() {
	var err error = nil
	if fontData, err = os.ReadFile(fontFile); err != nil {
		panic(err)
	}
	initFont()
	save("xxxx")

	//step1()
	step2()
}

func initFont() {
	var err error = nil
	if font, err = freetype.ParseFont(fontData); err != nil {
		panic(err)
	}

	if scalerType := utils.UInt32(fontData); scalerType != 65536 {
		panic(fmt.Errorf("scalerType error: %d", scalerType))
	}

	numTables := utils.UInt16(fontData[4:])

	var i uint16 = 0
	for ; i < numTables; i++ {
		index := i*16 + 12
		tag := string(fontData[index : index+4])
		offset := utils.UInt32(fontData[index+8:])
		length := utils.UInt32(fontData[index+12:])

		tags[tag] = Tag{tag, offset, length}
		//fmt.Printf("tag: %s, offset: %d, length: %d\n", tag, offset, length)
	}
}

func hasChar(i int) bool {
	return font.Index(rune(i)) != 0
}

func save(tagName string) {
	if tag, ok := tags[tagName]; ok {
		file, err := os.Create(tagName)
		if err != nil {
			panic(err)
		}
		_, _ = file.Write(fontData[tag.offset : tag.offset+tag.length])
		_ = file.Close()
	}
}
