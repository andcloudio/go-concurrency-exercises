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

type result struct {
	srcImagePath   string
	thumbnailImage *image.NRGBA
	err            error
}

type StageError struct {
	Stage string
	Err   error
}

func (s *StageError) Error() string {
	return fmt.Sprintf("stage = %s: error = %v", s.Stage, s.Err)
}

// Image processing - sequential
// Input - directory with images.
// output - thumbnail images
func main() {
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}
	start := time.Now()

	err := setupPipeline(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func setupPipeline(root string) error {

	done := make(chan struct{})
	//defer close(done)

	// first stage walk the files
	paths, errc := walkFiles(done, root)
	select {
	case err := <-errc:
		if err != nil {
			return &StageError{
				Stage: "walkfiles",
				Err:   err,
			}
		}
	default:
	}

	// process the files

	result := processImage(done, paths)

	// save the files
	count := 0

	for img := range result {
		if img.err != nil {
			return &StageError{
				Stage: "processImage",
				Err:   img.err,
			}
		}
		count++
		if count == 5 {
			close(done)
		}
		saveThumbnail(img)

	}

	return nil
}

// walfiles - take diretory path as input
// does the file walk
// generates thumbnail images
// saves the image to thumbnail directory.
func walkFiles(done chan struct{}, root string) (<-chan string, <-chan error) {
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

			select {
			case <-done:
				return fmt.Errorf("walk cancelled")
			case paths <- path:
			}

			return nil
		})
	}()

	return paths, errc
}

// processImage - takes image file as input
// return pointer to thumbnail image in memory.
func processImage(done chan struct{}, paths <-chan string) <-chan result {
	out := make(chan result)
	count := 5
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				// load the image from file
				srcImage, err := imaging.Open(path)
				if err != nil {
					select {
					case out <- result{path, nil, err}:
					case <-done:
						return
					}
				}

				// scale the image to 100px * 100px
				thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)
				select {
				case out <- result{path, thumbnailImage, nil}:
				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

// saveThumbnail - save the thumnail image to folder
func saveThumbnail(img result) error {
	filename := filepath.Base(img.srcImagePath)
	dstImagePath := "thumbnail/" + filename

	// save the image in the thumbnail folder.
	err := imaging.Save(img.thumbnailImage, dstImagePath)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", img.srcImagePath, dstImagePath)
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
