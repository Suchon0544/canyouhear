package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MyFile struct {
	Path    string
	Size    int64
	Name    string
	ModTime time.Time
}

/*type empty interface {
}*/

func getDrives() (r []string) {
	var char string
	drive := fmt.Scan(&char)
	f, err := os.Open(drive + ":\\")
	if err == nil {
		d := string(drive) + ":/"
		r = append(r, d)
		f.Close()
	}
	return r
}

func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					*files = append(*files, f.Name())
				}
			}

			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files)
			}
		}
	}
}

func output(Path, Size []string) {
	file, err := os.Create("Output.txt")
	defer file.Close()
	if err != nil {
		return
	}
	for i := range Path {
		file.WriteString(fmt.Sprintf("%v\t%v\r\n", Size[i], Path[i]))
	}
}
