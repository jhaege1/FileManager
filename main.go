package main

import (
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
type microsoftInstaller struct{}
type compressedFile struct{}
type videos struct{}

func main() {
	dir := "C:\\Users\\Jeroen\\Downloads\\"

	pdf := portableDocumentFormat{}
	img := image{}
	exe := executable{}
	msi := microsoftInstaller{}
	zip := compressedFile{}
	vid := videos{}

	createDirectories(pdf)
	createDirectories(img)
	createDirectories(exe)
	createDirectories(msi)
	createDirectories(zip)
	createDirectories(vid)

	currentFileLocations := returnFileLocations(dir)

	for _, currentFileLocation := range currentFileLocations {
		fileName := currentFileLocation[strings.LastIndex(currentFileLocation, "\\")+1:]
		moveFiles(pdf, currentFileLocation, fileName)
		moveFiles(img, currentFileLocation, fileName)
		moveFiles(exe, currentFileLocation, fileName)
		moveFiles(msi, currentFileLocation, fileName)
		moveFiles(zip, currentFileLocation, fileName)
		moveFiles(vid, currentFileLocation, fileName)
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

func (msi microsoftInstaller) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Executables\\"
}

func (zip compressedFile) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Zipped files\\"
}

func (vid videos) getLocation() string {
	return "C:\\Users\\Jeroen\\Downloads\\Videos\\"
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

func (msi microsoftInstaller) getExtension() string {
	return ".msi"
}

func (zip compressedFile) getExtension() string {
	return ".zip"
}

func (vid videos) getExtension() string {
	return ".mp4"
}

func returnFileLocations(dir string) []string {
	fileLocations := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath, err := filepath.Abs(dir + fileName)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
	}
}

func createDirectories(f folder) {
	err := os.MkdirAll(f.getLocation(), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}
