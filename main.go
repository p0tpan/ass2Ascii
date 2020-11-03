package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"time"
	"github.com/qeesung/image2ascii/convert"
	"io/ioutil"
	"log"
)

func lsImages()  {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files[0].Name())
}

func assRandomizer() string {
	assList := []string{"./images/ass2.jpg", "./images/2487pgaxqaw51.jpg", "./7CwaBKErbKumOQ5YT3nQzCHSQMCHbjKYxDg-YFxOS2E.jpg", "./b_8v6vTWWDFAtPThzp_mFs8H5Yr-ArYQZfLZbuffI4I.png", "./cmwaqatpz0w51.jpg", "./CvqjdJ7pDBm5VuZ0vR-fw-N2iEvopw9OX7mOXDGv8dw.png", "./d46ossbim4v51.jpg", "./l2nbyfih2nw51.jpg", "./lrmtjoka7bo51.jpg", "./mf0kyq6r01851.jpg", "./p7pke6lwotv51.jpg", "./r9s0so5k7qw51.jpg", "./WjYU5lmgtoCS-x2rDcYlsfXvSH1DjK6T97fIzQj9K_8.png", "./y8mh1dmlyew51.jpg", "./yh1022ntbpw51.jpg", "./zscpi0hgikv51.jpg"}
	rand.Shuffle(len(assList), func(i, j int) {
		assList[i], assList[j] = assList[j], assList[i]
	})
	fmt.Println(assList[0])
	return assList[0]
}

func assPrinter() {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	convertOptions.Colored = false

	// Create the image converter
	converter := convert.NewImageConverter()
	// Prints that ass
	fmt.Print(converter.ImageFile2ASCIIString(assRandomizer(), &convertOptions))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//assPrinter()
	lsImages()
}
