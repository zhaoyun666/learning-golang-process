package main

import (
    "image/color"
    "fmt"
    "io"
    "math/rand"
    "image/gif"
    "image"
    "math"
    "time"
    "os"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = iota // first color in palette
    blackIndex = iota // next color in palette
)

func main() {
    fmt.Println(whiteIndex, blackIndex)
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5 // number of complete x oscillator revolutions
        res = 0.001 // angular resolution
        size = 100 // image canvas convers [-size..+size]
        nframes = 64 // number of animation frames
        delay = 8 // delay between fremes in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    animate := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase diffrence
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size + 1, 2*size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size + int(x*size + 0.5), size + int(y*size + 0.5), blackIndex)
        }
        phase += 0.1
        animate.Delay = append(animate.Delay, delay)
        animate.Image = append(animate.Image, img)
    }
    gif.EncodeAll(out, &animate)
}