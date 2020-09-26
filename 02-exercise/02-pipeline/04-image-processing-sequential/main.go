package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
)

// Image processing - sequential
// Input - directory with images.
// output - thumbnail images
func main() {
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}
	start := time.Now()

	err := walkFiles(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

// walfiles - take diretory path as input
// does the file walk
// generates thumbnail images
// saves the image to thumbnail directory.
func walkFiles(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// filter out error
		if err != nil {
			return err
		}

		// check if it is file
		if !info.Mode().IsRegular() {
			return nil
		}

		// check if it is image/jpeg
		contentType, _ := getFileContentType(path)
		if contentType != "image/jpeg" {
			return nil
		}

		// process the image
		thumbnailImage, err := processImage(path)
		if err != nil {
			return err
		}

		// save the thumbnail image to disk
		err = saveThumbnail(path, thumbnailImage)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// processImage - takes image file as input
// return pointer to thumbnail image in memory.
func processImage(path string) (*image.NRGBA, error) {

	// load the image from file
	srcImage, err := imaging.Open(path)
	if err != nil {
		return nil, err
	}

	// scale the image to 100px * 100px
	thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)

	return thumbnailImage, nil
}

// saveThumbnail - save the thumnail image to folder
func saveThumbnail(srcImagePath string, thumbnailImage *image.NRGBA) error {
	filename := filepath.Base(srcImagePath)
	dstImagePath := "thumbnail/" + filename

	// save the image in the thumbnail folder.
	err := imaging.Save(thumbnailImage, dstImagePath)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", srcImagePath, dstImagePath)
	return nil
}

// getFileContentType - return content type and error status
func getFileContentType(file string) (string, error) {

	out, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
