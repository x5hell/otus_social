package font

import (
	"component/handler"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
)

const DefaultFontFilePath = "github.com/golang/freetype/testdata/luximr.ttf"
const defaultDpi = 72
const DefaultFontSize = 14.0

func createFont(fontFilePath string) (font *truetype.Font, err error) {
	fontBytes, err := ioutil.ReadFile(fontFilePath)
	if err != nil {
		handler.ErrorLog(err)
		return font, err
	}
	font, err = freetype.ParseFont(fontBytes)
	handler.ErrorLog(err)
	return font, err
}

func CreateDefaultFont() (font *truetype.Font, err error) {
	return createFont(DefaultFontFilePath)
}

func CreateFontContext(font *truetype.Font, size float64) (fontContext *freetype.Context) {
	fontContext = freetype.NewContext()
	fontContext.SetFont(font)
	fontContext.SetFontSize(size)
	fontContext.SetDPI(defaultDpi)
	return fontContext
}
