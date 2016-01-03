package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	directory := "/Users/geoffrey/Downloads/Fargo - Season 2/"
	files, _ := ioutil.ReadDir(directory)

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".srt" {
			file := directory + f.Name()
			cleanSubtitle(file)
			fmt.Println(f.Name() + " Done!!!")
		}
	}
}

func cleanSubtitle(file string) {

	dat, err := ioutil.ReadFile(file)
	check(err)
	str := string(dat)
	// fmt.Print(str)

	// Match _ line
	re1, err := regexp.Compile(`\d+\s*.*\n_\s*`)
	// Match mark-ups
	re2, err := regexp.Compile(`[^\S\n]*</?\w*>[^\S\n]*`)
	// Match last spaces
	re3, err := regexp.Compile(`\s*\z`)
	// Match Ads
	re4, err := regexp.Compile(`\d+\s*\n.*-->.*\n(.+\n)?.*(addic7ed\.com|\.org)\n+`)

	check(err)

	// str = re1.ReplaceAllString(str, "")
	str = re1.ReplaceAllString(str, "")
	str = re2.ReplaceAllString(str, "")
	str = re3.ReplaceAllString(str, "")
	str = re4.ReplaceAllString(str, "")

	// result := re3.FindAllStringSubmatch(str, -1)

	f, err := os.Create(file)
	check(err)
	defer f.Close()
	f.WriteString(str)
	f.Sync()
}
