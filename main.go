package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// variables declaration
	var directory string
	var DoMoveFiles bool

	// flags declaration using flag package
	flag.StringVar(&directory, "d", ".", "Specify directory where wallpapers are stored. Default is current directory.")
	flag.BoolVar(&DoMoveFiles, "m", false, "Specify if files should be moved to new directories where they are sorted by portrait. Default is false")

	flag.Parse() // after declaring flags we need to call it

	// recursively get all files in directory
	files, err := filepath.Glob(directory + "/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// open file
		img, _, err := image.Decode(openFile(file))
		if err != nil {
			log.Fatal(err)
		}

		// get image dimensions
		width := img.Bounds().Max.X
		height := img.Bounds().Max.Y

		// check if image is portrait
		if width < height {
			// move file to portrait directory
			if DoMoveFiles {
				moveFile(file, directory+"/portrait")
			}
		} else if width > height {
			// move file to landscape directory
			if DoMoveFiles {
				moveFile(file, directory+"/landscape")
			}
		} else {
			// move file to square directory
			if DoMoveFiles {
				moveFile(file, directory+"/square")
			}
		}
	}
}

func openFile(file string) *os.File {
	// open file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func moveFile(file string, directory string) {
	// move file to new directory
	CopyFile(file, directory)

	// delete original file
	DeleteFile(file)
}

func CopyFile(file string, directory string) {
	// get file name
	fileName := filepath.Base(file)

	// if directory doesn't exist, create it
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.MkdirAll(directory, 0755)
	}

	// open file
	src, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// create destination file
	dst, err := os.Create(directory + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	// copy file
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func DeleteFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}
