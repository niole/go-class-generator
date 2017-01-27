package main

import (
	"fmt"
	"log"
	"os"
)

var REACT = "react"
var JS = "js"
var ES5 = "es5"
var ES6 = "es6"
var JSSUFFIX = ".js"
var REACTSUFFIX = ".jsx"

func main() {
	/**
	arg 0 is always the name of this file, ignore
	arg 1: 'react', 'js', ...
	arg 2: 'es5', 'es6', ...
	arg 3: name of class
	*/

	args := os.Args

	language := args[1]
	version := args[2]
	classNames := args[3:] //can make as many classes as you want

	for _, name := range classNames {
		file := File{name, version, language}
		fileHandle := file.createFile()
		content := file.createContent()
		file.writeContent(content, fileHandle)
	}
}

type File struct {
	name     string
	version  string
	language string
}

func (f *File) createFile() *os.File {
	suffix := JSSUFFIX
	if f.language == REACT {
		suffix = REACTSUFFIX
	}

	file, err := os.Create(f.name + suffix)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func makeES5Class(name string) (content string) {
	content += "'use strict';\n\n"
	content += fmt.Sprintf("function %s() {\n}\n\n", name)
	content += fmt.Sprintf("%s.prototype.constructor = %s;\n\n", name, name)
	content += fmt.Sprintf("%s.prototype.get%s = function() {\n};\n\n", name, name)
	content += fmt.Sprintf("%s.prototype.set%s = function() {\n};\n\n", name, name)
	return
}

func (f *File) createContent() (content string) {
	if f.language == JS {
		if f.version == ES5 {
			content += makeES5Class(f.name)
		}

		if f.version == ES6 {
		}
	}

	if f.language == REACT {
		if f.version == ES5 {
		}

		if f.version == ES6 {
		}
	}

	return
}

func (f *File) writeContent(content string, file *os.File) *os.File {
	_, err := file.WriteString(content)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
