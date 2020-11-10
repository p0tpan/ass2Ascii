package ass2Ascii

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"github.com/qeesung/image2ascii/convert"
	"io/ioutil"
	"log"
	"path/filepath"
)

func lsImages() []string {
	//a function that will scan the ./images directory and will list them within a slice of strings.
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	var listOfFiles []string

	for _, file := range files {
		listOfFiles = append(listOfFiles, file.Name())
	}
	return listOfFiles
}

func assRandomizer() string {
	//a function that randomizes the listOfFiles slice and returns the first index of that randomized slice.
	listOfFiles := lsImages()
	rand.Shuffle(len(listOfFiles), func(i, j int) {
		listOfFiles[i], listOfFiles[j] = listOfFiles[j], listOfFiles[i]
	})
	return filepath.Join("./images/", listOfFiles[0])
}

func AssPrinter() {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	convertOptions.Colored = true

	// Create the image converter
	converter := convert.NewImageConverter()
	// Prints that ass
	fmt.Print(converter.ImageFile2ASCIIString(assRandomizer(), &convertOptions))
}
