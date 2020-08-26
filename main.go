package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(ex)
	file := flag.String("f", "log.txt", "Log file name")
	view := flag.Bool("v", false, "View log file")
	flag.Parse()

	if *view {
		b, err := ioutil.ReadFile(dir + "/" + *file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	c, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	s := "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + c + "\n"

	f, err := os.OpenFile(dir+"/"+*file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

}
