package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tr "github.com/snakesel/libretranslate"
)

func main() {
	translate := tr.New(tr.Config{
		Url: "http://127.0.0.1:5000",
		Key: "",
	})

	for _, v := range os.Args[1:] {
		content, err := ioutil.ReadFile(v) // the file is inside the local directory
		if err != nil {
			fmt.Println("Err")
		}
		fmt.Println(string(content))

		trtext, err := translate.Translate(string(content), "en", "zh")
		if err == nil {
			fmt.Println(trtext)
		} else {
			fmt.Println(err.Error())
		}

		if err := os.WriteFile(v+".transtxt", []byte(trtext), 0666); err != nil {
			log.Fatal(err)
		}

	}
	// you can use "auto" for source language
	// so, translator will detect language

	// Detect the language of the text
	/*
		conf, lang, err := translate.Detect("Nächster Stil")
		if err == nil {
			fmt.Printf("%s (%f)\n", lang, conf)
		} else {
			fmt.Println(err.Error())
		}*/
}
