// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var elemNameMap = map[string]string{
	"a":    "Anchor",
	"abbr": "Abbreviation",
}

func main() {
	b, err := ioutil.ReadFile("htmlattributes.source.txt")
	if err != nil {
		log.Fatal(err)
	}

	replacements := strings.NewReplacer("Html", "HTML", "Http", "HTTP", "Href", "HRef", "Id", "ID")

	source := string(b)

	file, err := os.Create("htmlattributes.autogen.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `//go:generate go run generate.go

// Package attr defines markup to create HTML attributes supported by Facebook React.
//
// Created from "HTML Attributes" as defined by Facebook in
// https://facebook.github.io/react/docs/tags-and-attributes.html

package attr

import "github.com/bep/gr"
`)

	words := strings.Fields(source)

	for _, w := range words {
		funcName := strings.Title(w)
		funcName = replacements.Replace(funcName)
		funcBody := fmt.Sprintf(`
// %s creates an HTML attribute for '%s'.
func %s(v string) gr.Modifier {
	return gr.Prop("%s", v)
}
`, funcName, w, funcName, w)

		fmt.Fprintf(file, "%s", funcBody)

	}

}
