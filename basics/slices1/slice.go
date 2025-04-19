package main
import "fmt"

func main() {
	selfie := []byte{ /* ~100 MB pixel data */ }
	invertColors(selfie)
}

func invertColors(colors []byte) {
    fmt.Println("slices1/slices.go")   
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
}
