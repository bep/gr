// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var altDoc = map[string]string{
	"key": `Key adds an optional, unique identifier. 
When your component shuffles around during render passes, it might be destroyed 
and recreated due to the diff algorithm. Assigning it a key that persists makes 
sure the component stays.`,
	"ref": "Ref adds an ref to a component, see http://facebook.github.io/react/docs/more-about-refs.html",
	"dangerouslySetInnerHTML": `DangerouslySetInnerHTML Provides the ability to insert raw HTML, 
mainly for cooperating with DOM string manipulation libraries.`,
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
// - https://facebook.github.io/react/docs/tags-and-attributes.html
// - http://facebook.github.io/react/docs/special-non-dom-attributes.html

package attr

import "github.com/bep/gr"
`)

	words := strings.Fields(source)
	sort.Strings(words)

	for _, w := range words {
		funcName := strings.Title(w)
		funcName = replacements.Replace(funcName)
		docString := fmt.Sprintf("%s creates an HTML attribute for '%s'.", funcName, w)

		if alt, ok := altDoc[w]; ok {
			docString = strings.Replace(alt, "\n", "\n// ", -1)
		}

		funcBody := fmt.Sprintf(`
// %s
func %s(v string) gr.Modifier {
	return gr.Prop("%s", v)
}
`, docString, funcName, w)

		fmt.Fprintf(file, "%s", funcBody)

	}

}
