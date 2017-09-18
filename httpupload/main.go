package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	// Show memory before uploading file
	printMem()

	targetURL := "http://127.0.0.1/upload"
	fieldname := "uploadfile"
	filename := "./main.go"

	err := postFile(targetURL, fieldname, filename)
	if err != nil {
		fmt.Println(err)
	}

	// Show memory after uploading file
	printMem()
}

func postFile(targetURL, fieldname, filename string) error {

	// Open file handle
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file")
	}
	defer file.Close()

	// Create a pipe for providing Reader interface connected with multipart writer
	bodyReader, bodyWriter := io.Pipe()
	multiWriter := multipart.NewWriter(bodyWriter)

	// Create a go routine to write multipart data
	go func() {
		part, err := multiWriter.CreateFormFile("uploadfile", filepath.Base(filename))
		if err != nil {
			bodyWriter.CloseWithError(err)
			return
		}

		if _, err := io.Copy(part, file); err != nil {
			bodyWriter.CloseWithError(err)
			return
		}

		bodyWriter.CloseWithError(multiWriter.Close())
	}()

	// Start to post file
	resp, err := http.Post(targetURL, multiWriter.FormDataContentType(), bodyReader)
	if err != nil {
		fmt.Println("resp error:", err)
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

func printMem() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("MemAlloc", getMemoryHuman(mem.Alloc))
	fmt.Println("MemSys", getMemoryHuman(mem.Sys))
}

var (
	GB uint64 = 1024 * 1024 * 1024
	MB uint64 = 1024 * 1024
	KB uint64 = 1024
)

func getMemoryHuman(m uint64) string {
	if m > GB {
		return fmt.Sprintf("%0.3fG", float64(m)/float64(GB))
	} else if m > MB {
		return fmt.Sprintf("%0.3fM", float64(m)/float64(MB))
	} else if m > KB {
		return fmt.Sprintf("%0.3fK", float64(m)/float64(KB))
	}

	return fmt.Sprintf("%d", m)
}
