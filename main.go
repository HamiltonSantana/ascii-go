package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

//Define the intensity of the character base on grayscale (these is a rough aproximation of
// the real deal
const grayscale = "@#09867543&2%$1|!;:.=-*^' "

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
    indx := int(grayPoint/10)

    //print out the colored character and background to the relate intensity char
    return fmt.Sprintf("\033[48;2;%d;%d;%dm%s\033[30m",cr, cg, cb,string(grayscale[indx]))
}

//process to render the image
func renderImg(img image.Image) {
    for i := img.Bounds().Min.Y;  i < img.Bounds().Max.Y; i++ {
        for j := img.Bounds().Min.X;  j < img.Bounds().Max.X; j++ {
            r,g,b,a := img.At(j,i).RGBA()
            fmt.Print(colorASCII(r,g,b,a))
        }
        //correct the trailing color from the last 'print'
        fmt.Println("\033[0m")
    }
}

func main() {
    file, err := os.Open("img.jpg")
    img, err1 := jpeg.Decode(file)
    defer file.Close()
    if err != nil {
        fmt.Print(err.Error())
        return
    }
    if err1 != nil {
        fmt.Print(err1.Error())
    }

    renderImg(img)

    fmt.Println()
}
