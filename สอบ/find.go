package find

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		/*f := fmt.Scan()*/
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
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

/*func main() {
	drives := getDrives()
	files := []string{}
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg", ".gif"}, d, &files)
	}
	fmt.Println(len(files))
	fmt.Println(drives)
}*/
