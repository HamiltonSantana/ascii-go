package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

//Define the intensity of the character base on grayscale (these is a rough aproximation of
// the real deal
const grayscale = "@#09867543&2%$1|!;:.=-*^' "

var _ png.UnsupportedError
var _ jpeg.UnsupportedError
var _ gif.GIF
//convert the point (r, g or b) to the scale of 0 to 255
// the decimal part are truncate
func convertPoint(p uint32, op uint32) int {
    return int(float64(p)/float64(op) * 255)
}

func colorASCII(r uint32, g uint32, b uint32, a uint32) string {
    //[38 use to change character color
    //[48 use to change background color
    cr := convertPoint(r, a)
    cg := convertPoint(g,a)
    cb := convertPoint(b,a)

    // base convertion, sum all colors divede by thre
    //grayPoint := (cr + cg + cb) / 3

    // Convert to grayscale using Luma formula
    grayPoint := ( 0.2126 * float64(cr) + 0.7152 * float64(cg) + 0.0722 * float64(cb))

    //divede the number to 10 to fit the range of grayscale (char intensity)

    var indx int
    if int(grayPoint/10) < 0 {
        indx =  25
    } else {
        indx = int(grayPoint/10)
    }

    //print out the colored character and background to the relate intensity char
    charact := string(grayscale[indx])
    return fmt.Sprintf("\033[48;2;%d;%d;%dm%s%s\033[30m",cr, cg, cb, charact, charact)
}

//process to render the image
func renderImg(img image.Image) {
    var i, j int
    for j = 0;  j < img.Bounds().Dy(); j++ {
        for i := 0;  i < img.Bounds().Dx(); i++ {
            r,g,b,a := img.At(int(i),j).RGBA()
            fmt.Print(colorASCII(r,g,b,a))
        }
        //correct the trailing color from the last 'print'
        fmt.Println("\033[0m")
    }
    fmt.Println(i," ", j)
}

func main() {
    pathImg := os.Args[1]

    file, err := os.Open(pathImg)
    
    defer file.Close()
    if err != nil {
        fmt.Print(err.Error())
        return
    }

    foo, _, bar := image.Decode(file)

    if bar != nil {
        fmt.Println(bar)
    }

    renderImg(foo)

    fmt.Println()
}
