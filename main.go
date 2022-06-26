package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type doc interface {
	getLocation() string
	getExtension() string
}

type folder interface {
	getLocation() string
}

type portableDocumentFormat struct{}
type image struct{}
type executable struct{}
type msi struct{}
type zip struct{}

func main() {
	dir := "C:\\Users\\Jeroen\\Downloads\\"

	pdf := portableDocumentFormat{}
	img := image{}
	exe := executable{}
	msi := msi{}
	zip := zip{}

	createDirectories(pdf)
	createDirectories(img)
	createDirectories(exe)
	createDirectories(msi)
	createDirectories(zip)

	currentFileLocations := returnFileLocations(dir)

	for _, currentFileLocation := range currentFileLocations {
		fileName := currentFileLocation[strings.LastIndex(currentFileLocation, "\\")+1:]
		moveFiles(pdf, currentFileLocation, fileName)
		moveFiles(img, currentFileLocation, fileName)
		moveFiles(exe, currentFileLocation, fileName)
		moveFiles(msi, currentFileLocation, fileName)
		moveFiles(zip, currentFileLocation, fileName)
	}
}

func (pdf portableDocumentFormat) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\PDF\\"
}

func (img image) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Images\\"
}

func (exe executable) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Executables\\"
}

func (msi msi) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Executables\\"
}

func (zip zip) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Zipped files\\"
}

func (pdf portableDocumentFormat) getExtension() string {
	return ".pdf"
}

func (img image) getExtension() string {
	return ".jpg"
}

func (exe executable) getExtension() string {
	return ".exe"
}

func (msi msi) getExtension() string {
	return ".msi"
}

func (zip zip) getExtension() string {
	return ".zip"
}

func returnFileLocations(dir string) []string {
	fileLocations := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath, err := filepath.Abs(dir + fileName)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fileLocations = append(fileLocations, filePath)
	}

	return fileLocations
}

func moveFiles(d doc, currentFileLocation string, fileName string) {
	if strings.HasSuffix(currentFileLocation, d.getExtension()) {
		docDir := d.getLocation()
		err := os.Rename(currentFileLocation, docDir+fileName)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func createDirectories(f folder) {
	err := os.MkdirAll(f.getLocation(), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}
