package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

// Pipeline
// walkfile ----------> processImage -----------> saveImage
//            (paths)                  (results)

type result struct {
	srcImagePath   string
	thumbnailImage *image.NRGBA
	err            error
}

// Image processing - Pipeline
// Input - directory with images.
// output - thumbnail images
func main() {
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}
	start := time.Now()
	err := setupPipeLine(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func setupPipeLine(root string) error {
	done := make(chan struct{})
	defer close(done)

	// do the file walk
	paths, errc := walkFiles(done, root)

	// process the image
	results := processImage(done, paths)

	// save thumbnail images
	for r := range results {
		if r.err != nil {
			return r.err
		}
		saveThumbnail(r.srcImagePath, r.thumbnailImage)
	}

	// check for error on the channel, from walkfiles stage.
	if err := <-errc; err != nil {
		return err
	}
	return nil
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {

	// create output channels
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

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

			// send file path to next stage
			select {
			case paths <- path:
			case <-done:
				return fmt.Errorf("walk cancelled")
			}
			return nil
		})
	}()
	return paths, errc
}

func processImage(done <-chan struct{}, paths <-chan string) <-chan result {
	results := make(chan result)
	var wg sync.WaitGroup

	thumbnailer := func() {
		for srcImagePath := range paths {
			srcImage, err := imaging.Open(srcImagePath)
			if err != nil {
				select {
				case results <- result{srcImagePath, nil, err}:
				case <-done:
					return
				}
			}
			thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)

			select {
			case results <- result{srcImagePath, thumbnailImage, err}:
			case <-done:
				return
			}
		}
	}

	const numThumbnailer = 5
	for i := 0; i < numThumbnailer; i++ {
		wg.Add(1)
		go func() {
			thumbnailer()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
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
