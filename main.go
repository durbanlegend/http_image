package main
 
import (
	"bytes"
	"image"
	"image/color"
	"image/png"
    "net/http"
)

type Image struct{
	Width, Height int
	colr uint8	
}
 
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}
 
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}
 
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr+uint8(x ^ y), r.colr+uint8(x * y), r.colr+uint8((x + y) / 5), 255}
}
 
func (m *Image) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type","image/png")
    w.Write(buf.Bytes())
}

func main() {
	m := Image{1020, 255, 0}
    http.Handle("/image", &m)
    http.ListenAndServe("localhost:4001", nil)
}
