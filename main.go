package main

import (
	"flag"
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
	file := flag.String("file", "log.txt", "Nome do arquivo de log")
	flag.Parse()
	c, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	s := "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + c + "\n"

	f, err := os.OpenFile(dir+"/"+*file, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

}
