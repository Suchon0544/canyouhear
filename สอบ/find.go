package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getDrives() (r []string) {
	fmt.Print("Directory : ")
	var drive string
	fmt.Scan(&drive)
	f, err := os.Open(drive + ":\\")
	if err == nil {
		d := string(drive) + ":/"
		r = append(r, d)
		f.Close()
	}
	return
}

func findFileFromExtension(extension []string, dir string, files *[]string) {
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
				findFileFromExtension(extension, path, files)
			}
		}
	}
}

func main() {
	drives := getDrives()
	files := []string{}
	for _, d := range drives {
		findFileFromExtension([]string{}, d, &files)
	}
	fmt.Println(len(files))
	fmt.Println(drives)
}
