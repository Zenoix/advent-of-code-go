package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
	"time"
)

const startMessagePrefix = "Creating"

const endMessagePrefix = "Created"

const messageFormatString = "skeleton for year %d day %d in %s"

func getRootDir() (rootPath string) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find root directory")
	}

	rootPath = filepath.Join(filepath.Dir(filename), "..")
	return
}

func checkFileDoesNotExist(path string) {
	_, err := os.Stat(filepath.Join(path))
	if err == nil {
		log.Fatalf("File %s already exists", path)
	}
}

func createSkeleton(day, year int) {
	if day < 1 || day > 25 {
		log.Fatalf("Invalid value for --day, must be between 1 and 25 inclusive, got %d", day)
	}

	if year < 2015 {
		log.Fatalf("Invalid value for --year, must be 2015 or later, got %d", year)
	}

	rootDir := getRootDir()
	targetDir := fmt.Sprintf("%s/%d/day%d", rootDir, year, day)

	log.Printf(startMessagePrefix+" "+messageFormatString, year, day, targetDir)

	err := os.MkdirAll(targetDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	temp, err := template.ParseFS(
		os.DirFS(filepath.Join(rootDir, "skeleton")),
		"aoc-template/*",
	)
	if err != nil {
		log.Fatal(err)
	}

	mainFilePath := filepath.Join(targetDir, "main.go")
	makeFilePath := filepath.Join(targetDir, "Makefile")

	checkFileDoesNotExist(mainFilePath)
	checkFileDoesNotExist(makeFilePath)

	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		log.Fatal(err)
	}
	makeFile, err := os.Create(makeFilePath)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Year, Day int
	}{
		Year: year,
		Day:  day,
	}
	err = temp.ExecuteTemplate(mainFile, "main.go", data)
	if err != nil {
		log.Fatal(err)
	}
	err = temp.ExecuteTemplate(makeFile, "Makefile", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(endMessagePrefix+" "+messageFormatString, year, day, targetDir)

}

func main() {
	log.SetFlags(0)
	now := time.Now()

	day := flag.Int("day", now.Day(), "Challenge day")
	year := flag.Int("year", now.Year(), "Year")
	flag.Parse()

	createSkeleton(*day, *year)
}
