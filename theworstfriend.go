package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func corrupter(fpath string) {
	ffile, _ := ioutil.ReadFile(fpath)
	fcontent := strings.Split(string(ffile), "\n")
	fcontent = append(fcontent[:0], fcontent[1:]...)
	info, _ := os.Stat(fpath)
	mode := info.Mode()
	ioutil.WriteFile(fpath, []byte(strings.Join(fcontent,"\n")), mode)
}


func twfirend(path string) {
	files, _ := ioutil.ReadDir(path)
	count:=0
	for _, fname := range files {
		fpath := path+fname.Name()
		corrupter(fpath)
		fmt.Println( "-> ", fname.Name())
		count+=1
	}
	fmt.Println("- total files : ",count)
}

func main() {
	args := os.Args
	path := args[1]
	twfirend(path)
}
