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
var ADD = "-a"
var REMOVE = "-r"
var ARGS = "args"

func main() {
	/**
	arg 0 is always the name of this file, ignore
	arg 1: 'react', 'js', ...
	arg 2: 'es5', 'es6', ...
	arg 3: name(s) of class(es) with TODO optional -a (add) and -r (remove) flags for elements to add or remove
		react classes:
			* state, props, defaultprops (e.g. react es6 Cats -a props -r state Dogs Animals)
		js classes:
			* constructor arguments, e.g: js es5 Cats -a -args legTotal hasWhiskers Dogs Animals
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

func makeES5Class(name string) string {
	usestrict := "'use strict';\n"
	constructor := fmt.Sprintf("function %s() {\n}\n", name)
	constructoradd := fmt.Sprintf("%s.prototype.constructor = %s;\n", name, name)
	export := fmt.Sprintf("module.exports = %s;", name)
	return fmt.Sprintf("%s\n%s\n%s\n%s", usestrict, constructor, constructoradd, export)
}

func makeES6Class(name string) string {
	return fmt.Sprintf("export default class %s {\n\tconstructor() {\n\t}\n}", name)
}

func makeES6ReactClass(name string) string {
	render := "\trender() {\n\t\treturn (\n\t\t);\n\t}\n"
	state := "\t\tthis.state = {\n\t\t};\n"
	super := "\t\tsuper();\n\n"
	constructor := fmt.Sprintf("\tconstructor() {\n%s%s\t}\n\n", super, state)
	class := fmt.Sprintf("export default class %s extends Component {\n%s%s}", name, constructor, render)
	return fmt.Sprintf("import React, {PropTypes, Component} from 'react';\n\n\n%s", class)
}

func makeES5ReactClass(name string) string {
	usestrict := "'use strict';\n"
	imprt := "var React = require('react');\n\n\n"

	render := "\trender: function() {\n\t\treturn (\n\t\t);\n\t}\n"
	getInitialState := "\tgetInitialState: function() {\n\t\treturn {\n\t\t};\n\t},\n\n"

	class := fmt.Sprintf("var %s = React.createClass({\n%s%s});\n\n", name, getInitialState, render)
	export := fmt.Sprintf("module.exports = %s;", name)
	return fmt.Sprintf("%s%s%s%s", usestrict, imprt, class, export)
}

func (f *File) createContent() (content string) {
	if f.language == JS {
		if f.version == ES5 {
			content += makeES5Class(f.name)
		}

		if f.version == ES6 {
			content += makeES6Class(f.name)
		}
	}

	if f.language == REACT {
		if f.version == ES5 {
			content += makeES5ReactClass(f.name)
		}

		if f.version == ES6 {
			content += makeES6ReactClass(f.name)
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
