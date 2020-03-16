package graphic

import (
	"component/font"
	"component/handler"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

const CanvasGraphicWidthOffsetPercent = 10
const CanvasGraphicHeightOffsetPercent = 10

const LineTypeSolid = "solid"
const LineTypeDotted = "dotted"
const LineTypeDashed = "dashed"

type FloatPoint struct {
	X float64
	Y float64
	Description string
}

type IntPoint struct {
	X int
	Y int
}

type Data2D struct {
	XName string
	YName string
	Data []FloatPoint
	XSize int
	YSize int
}

type CanvasInfo struct {
	CanvasWidth     int
	CanvasHeight    int
	GraphicWidth    int
	GraphicHeight   int
	XOffset         int
	YOffset         int
	MinPoint		FloatPoint
	MaxPoint		FloatPoint
	MinGraphicPoint FloatPoint
	MaxGraphicPoint FloatPoint
	Paletted        *image.Paletted
	Zero            FloatPoint
	XFactor         float64
	YFactor         float64
}

type LineType struct {
	Type string
	PeriodSize int
	DashSize int
	Color uint8
}

func  (canvasInfo *CanvasInfo) initGraphic () {
	canvasInfo.XOffset = canvasInfo.CanvasWidth * CanvasGraphicWidthOffsetPercent / 100
	canvasInfo.YOffset = canvasInfo.CanvasHeight * CanvasGraphicHeightOffsetPercent / 100
	canvasInfo.GraphicWidth = canvasInfo.CanvasWidth - canvasInfo.XOffset * 2
	canvasInfo.GraphicHeight = canvasInfo.CanvasHeight - canvasInfo.YOffset * 2
	canvasInfo.XFactor = canvasInfo.getXFactor()
	canvasInfo.YFactor = canvasInfo.getYFactor()
	canvasInfo.Zero = canvasInfo.getFloatZero()
}

func (canvasInfo *CanvasInfo) getXFactor () float64 {
	return math.Abs((float64(canvasInfo.CanvasWidth) - float64(canvasInfo.XOffset) * 2) /
		(canvasInfo.MaxGraphicPoint.X - canvasInfo.MinGraphicPoint.X))
}

func (canvasInfo *CanvasInfo) getYFactor () float64 {
	return math.Abs((float64(canvasInfo.CanvasHeight) - float64(canvasInfo.YOffset) * 2) /
		(canvasInfo.MaxGraphicPoint.Y - canvasInfo.MinGraphicPoint.Y))
}

func (canvasInfo *CanvasInfo) getZero() (zero IntPoint) {
	minX := canvasInfo.XOffset
	maxX := canvasInfo.XOffset + canvasInfo.GraphicWidth
	minY := canvasInfo.YOffset
	maxY := canvasInfo.YOffset + canvasInfo.GraphicHeight

	if canvasInfo.MaxGraphicPoint.X <= 0 {
		zero.X = maxX
	} else if canvasInfo.MinGraphicPoint.X >= 0 {
		zero.X = minX
	} else {
		zero.X = int(math.Abs(canvasInfo.MinGraphicPoint.X) * canvasInfo.XFactor) +
			canvasInfo.XOffset
	}
	if canvasInfo.MaxGraphicPoint.Y <= 0 {
		zero.Y = minY
	} else if canvasInfo.MinGraphicPoint.Y >= 0 {
		zero.Y = maxY
	} else {
		zero.Y = int(math.Abs(canvasInfo.MaxGraphicPoint.Y) * canvasInfo.YFactor) +
			canvasInfo.YOffset
	}

	return zero
}

func (canvasInfo *CanvasInfo) getFloatZero() (zero FloatPoint) {
	zeroInt := canvasInfo.getZero()
	zero.X = float64(zeroInt.X)
	zero.Y = float64(zeroInt.Y)
	return zero
}

func Draw2DGraphic(graphicData Data2D, fileName string) (err error) {
	const (
		whiteIndex = 0
		blackIndex = 1
	)
	palette := color.Palette{color.White, color.Black}
	rect := image.Rect(0, 0, graphicData.XSize, graphicData.YSize)
	img := image.NewPaletted(rect, palette)
	minPoint, maxPoint := getMinMaxFromData2D(graphicData)
	minGraphicPoint, maxGraphicPoint := getMinMaxGraphic(minPoint, maxPoint)
	canvasInfo := CanvasInfo{
		CanvasWidth:     graphicData.XSize,
		CanvasHeight:    graphicData.YSize,
		MinPoint: 		 minPoint,
		MaxPoint:		 maxPoint,
		MaxGraphicPoint: maxGraphicPoint,
		MinGraphicPoint: minGraphicPoint,
		Paletted:        img,
	}
	canvasInfo.initGraphic()
	drawGraphAxis(canvasInfo, graphicData, blackIndex)
	Draw2DGraphicLines(canvasInfo, graphicData, blackIndex)
	file, err := os.Create(fileName)
	err = png.Encode(file, img)
	handler.ErrorLog(err)
	return err
}

func Draw2DGraphicLines(canvasInfo CanvasInfo, graphicData Data2D, colorIndex uint8) {
	var graphicPoint, previousGraphicPoint FloatPoint
	for pointNumber, point := range graphicData.Data {
		graphicPoint = calculateGraphicPoint(canvasInfo, point)
		DrawPointCoordinates(canvasInfo, point, colorIndex)
		if pointNumber > 0 {
			drawCrossPoint(canvasInfo.Paletted, previousGraphicPoint, colorIndex)
			drawCrossPoint(canvasInfo.Paletted, graphicPoint, colorIndex)
			drawLine(canvasInfo, previousGraphicPoint, graphicPoint, colorIndex)
		}
		previousGraphicPoint = graphicPoint
	}
}

func DrawPointCoordinates(canvasInfo CanvasInfo, point FloatPoint, colorIndex uint8)  {
	graphicPoint := calculateGraphicPoint(canvasInfo, point)

	x1 := FloatPoint{
		X: canvasInfo.Zero.X,
		Y: graphicPoint.Y,
	}
	x2 := FloatPoint{
		X: graphicPoint.X,
		Y: graphicPoint.Y,
	}
	y1 := FloatPoint{
		X: graphicPoint.X,
		Y: canvasInfo.Zero.Y,
	}
	y2 := FloatPoint{
		X: graphicPoint.X,
		Y: graphicPoint.Y,
	}
	drawDottedLine(canvasInfo, x1, x2, colorIndex)
	drawDottedLine(canvasInfo, y1, y2, colorIndex)

	fontSize := font.DefaultFontSize
	xText := FloatPoint{
		X: graphicPoint.X - fontSize / 4,
		Y: canvasInfo.Zero.Y - fontSize * 1.5,
		Description: fmt.Sprint(point.X),
	}
	yText := FloatPoint{
		X: canvasInfo.Zero.X + fontSize / 2,
		Y: graphicPoint.Y - fontSize / 2,
		Description: fmt.Sprint(point.Y),
	}

	if len(point.Description) > 0 {
		yText.Description += "/" + point.Description
	}
	drawText(canvasInfo.Paletted, xText, fontSize, colorIndex)
	drawText(canvasInfo.Paletted, yText, fontSize, colorIndex)
}

func calculateGraphicPoint(canvasInfo CanvasInfo, point FloatPoint) (graphicPoint FloatPoint) {
	graphicPoint.X = point.X * canvasInfo.XFactor + canvasInfo.Zero.X
	graphicPoint.Y = canvasInfo.Zero.Y - point.Y * canvasInfo.YFactor
	return graphicPoint
}

func drawCrossPoint(palette *image.Paletted, point FloatPoint, colorIndex uint8)  {
	palette.SetColorIndex(int(point.X) + 1, int(point.Y), colorIndex)
	palette.SetColorIndex(int(point.X) - 1, int(point.Y), colorIndex)
	palette.SetColorIndex(int(point.X), int(point.Y) + 1, colorIndex)
	palette.SetColorIndex(int(point.X), int(point.Y) - 1, colorIndex)
	palette.SetColorIndex(int(point.X) + 2, int(point.Y), colorIndex)
	palette.SetColorIndex(int(point.X) - 2, int(point.Y), colorIndex)
	palette.SetColorIndex(int(point.X), int(point.Y) + 2, colorIndex)
	palette.SetColorIndex(int(point.X), int(point.Y) - 2, colorIndex)
}

func getMinMaxFromData2D(graphicData Data2D) (pMin, pMax FloatPoint) {
	xMin, xMax, yMin, yMax := math.Inf(+1), math.Inf(-1), math.Inf(+1), math.Inf(-1)

	for _, point := range graphicData.Data {
		if point.X < xMin {
			xMin = point.X
		}
		if point.X > xMax {
			xMax = point.X
		}
		if point.Y < yMin {
			yMin = point.Y
		}
		if point.Y > yMax {
			yMax = point.Y
		}
	}

	return FloatPoint{X: xMin, Y: yMin}, FloatPoint{X: xMax, Y: yMax}
}

func getMinMaxGraphic(minPoint, maxPoint FloatPoint) (pMin, pMax FloatPoint) {
	if minPoint.X > 0 {
		minPoint.X = 0
	}
	if minPoint.Y > 0 {
		minPoint.Y = 0
	}
	if maxPoint.X < 0 {
		minPoint.X = 0
	}
	if maxPoint.Y < 0 {
		minPoint.Y = 0
	}
	return minPoint, maxPoint
}

func drawLine(canvasInfo CanvasInfo, p1, p2 FloatPoint, colorIndex uint8) {
	lineType := LineType{
		Type:       LineTypeSolid,
		Color:      colorIndex,
	}

	if math.Abs(p2.X - p1.X) >= math.Abs(p2.Y - p1.Y) {
		drawLineX(canvasInfo.Paletted, p1, p2, lineType)
	} else {
		drawLineY(canvasInfo.Paletted, p1, p2, lineType)
	}
}

func drawDashedLine(canvasInfo CanvasInfo, p1, p2 FloatPoint, colorIndex uint8) {
	lineType := LineType{
		Type:       LineTypeDashed,
		Color:      colorIndex,
		PeriodSize: 10,
		DashSize:	5,
	}
	if math.Abs(p2.X - p1.X) >= math.Abs(p2.Y - p1.Y) {
		drawLineX(canvasInfo.Paletted, p1, p2, lineType)
	} else {
		drawLineY(canvasInfo.Paletted, p1, p2, lineType)
	}
}

func drawDottedLine(canvasInfo CanvasInfo, p1, p2 FloatPoint, colorIndex uint8) {
	lineType := LineType{
		Type:       LineTypeDotted,
		Color:      colorIndex,
		PeriodSize: 10,
	}
	if math.Abs(p2.X - p1.X) >= math.Abs(p2.Y - p1.Y) {
		drawLineX(canvasInfo.Paletted, p1, p2, lineType)
	} else {
		drawLineY(canvasInfo.Paletted, p1, p2, lineType)
	}
}



func drawLineX(palette *image.Paletted, p1, p2 FloatPoint, lineType LineType)  {
	k := float64(0)
	if p1.X != p2.X && p1.Y != p2.Y {
		k = (p2.Y - p1.Y) / (p2.X - p1.X)
	}
	n := p1.Y - k * p1.X
	xMin := int(p1.X)
	xMax := int(p2.X)
	if xMin > xMax {
		xMin, xMax = xMax, xMin
	}

	for x := xMin; x <= xMax; x++ {
		y := k * float64(x) + n
		if checkDrawLinePoint(float64(x), lineType){
			palette.SetColorIndex(x, int(y), lineType.Color)
		}

	}
}

func drawLineY(palette *image.Paletted, p1, p2 FloatPoint, lineType LineType)  {
	k := float64(0)
	if p1.X != p2.X && p1.Y != p2.Y {
		k = (p2.X - p1.X) / (p2.Y - p1.Y)
	}
	n := p1.X - k * p1.Y

	yMin := int(p1.Y)
	yMax := int(p2.Y)
	if yMin > yMax {
		yMin, yMax = yMax, yMin
	}

	for y := yMin; y <= yMax; y++ {
		x := k * float64(y) + n
		if checkDrawLinePoint(float64(y), lineType){
			palette.SetColorIndex(int(x), y, lineType.Color)
		}
	}
}

func checkDrawLinePoint(cord float64, lineType LineType) bool {
	switch lineType.Type {
		case LineTypeDotted:
			return int(cord) % lineType.PeriodSize == 0
		case LineTypeDashed:
			return int(cord) % lineType.PeriodSize < lineType.DashSize
		default:
			return true
	}
}


func drawGraphAxis(canvasInfo CanvasInfo, graphicData Data2D, colorIndex uint8) {

	minX := float64(canvasInfo.XOffset)
	maxX := float64(canvasInfo.XOffset + canvasInfo.GraphicWidth)
	minY := float64(canvasInfo.YOffset)
	maxY := float64(canvasInfo.YOffset + canvasInfo.GraphicHeight)

	drawLine(
		canvasInfo,
		FloatPoint{X: minX, Y: canvasInfo.Zero.Y},
		FloatPoint{X: maxX, Y: canvasInfo.Zero.Y},
		colorIndex,
	)

	drawLine(
		canvasInfo,
		FloatPoint{X: canvasInfo.Zero.X, Y: minY},
		FloatPoint{X: canvasInfo.Zero.X, Y: maxY},
		colorIndex,
	)

	fontSize := font.DefaultFontSize

	zero := FloatPoint{
		X: canvasInfo.Zero.X - fontSize,
		Y: canvasInfo.Zero.Y + fontSize / 2,
		Description: "0",
	}
	drawText(canvasInfo.Paletted, zero, fontSize, colorIndex)

	xName := FloatPoint{
		X: maxX - fontSize/1.5 * float64(len(graphicData.XName)),
		Y: maxY,
		Description: graphicData.XName,
	}
	drawText(canvasInfo.Paletted, xName, fontSize, colorIndex)

	for letterNumber, letter := range strings.Split(graphicData.YName, "") {
		yNameLetter := FloatPoint{
			X: minX - fontSize,
			Y: minY + float64(letterNumber) * fontSize,
			Description: letter,
		}
		drawText(canvasInfo.Paletted, yNameLetter, fontSize, colorIndex)
	}
}

func drawText(palette *image.Paletted, txt FloatPoint, fontSize float64, colorIndex uint8) {
	f, _ := font.CreateDefaultFont()
	colorType := image.NewUniform(palette.Palette[colorIndex])
	fontContext := font.CreateFontContext(f, fontSize)
	fontContext.SetDst(palette)
	fontContext.SetClip(palette.Bounds())
	fontContext.SetSrc(colorType)
	fontHeight := int(fontContext.PointToFixed(fontSize))>>6
	point := freetype.Pt(int(txt.X), int(txt.Y) + fontHeight)
	_, err := fontContext.DrawString(txt.Description, point)
	handler.ErrorLog(err)
}