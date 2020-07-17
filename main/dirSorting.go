package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type fl struct {
	name string
	size int64
}

func main() {
	fls, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	files := make([]fl, len(fls))
	for i := 0; i < len(files); i++ {
		files[i].name = fls[i].Name()
		if fls[i].IsDir() {
			files[i].size = dirSize(fls[i].Name(), ".")
		} else {
			files[i].size = fls[i].Size()
		}
	}

	for i := 1; i < len(files); i++ {
		if files[i].size < files[i-1].size {
			b := files[i]
			g := i
			for ; g > 0 && b.size < files[g-1].size; g-- {
				files[g] = files[g-1]
			}
			files[g] = b
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Printf("%s %d \n", files[i].name, files[i].size)
	}
}

func dirSize(dir string, path string) int64 {
	path = path + "/" + dir
	fls, err := ioutil.ReadDir(path)
	var s int64
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(fls); i++ {
		if fls[i].IsDir() {
			s += dirSize(fls[i].Name(), path)
		} else {
			s += fls[i].Size()
		}
	}
	return s
}
