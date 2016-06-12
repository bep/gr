// +build ignore

/*
Copyright 2016 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com> All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	"defaultValue": `DefaultValue can be used to initialize an uncontrolled React component with a non-empty value.

See https://facebook.github.io/react/docs/forms.html`,
}

var altType = map[string]string{
	"key": "interface{}",
	"ref": "interface{}",
	"dangerouslySetInnerHTML": "interface{}",
}

func main() {
	b, err := ioutil.ReadFile("htmlattributes.source.txt")
	if err != nil {
		log.Fatal(err)
	}

	replacements := strings.NewReplacer(
		"Html", "HTML", "Http", "HTTP",
		"Href", "HRef", "Id", "ID",
		"Wmode", "WMode")

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
		propType := "interface{}"
		if alt, ok := altDoc[w]; ok {
			docString = strings.Replace(alt, "\n", "\n// ", -1)
		}

		if alt, ok := altType[w]; ok {
			propType = alt
		}

		funcBody := fmt.Sprintf(`
// %s
func %s(v %s) gr.Modifier {
	return gr.Prop("%s", v)
}
`, docString, funcName, propType, w)

		fmt.Fprintf(file, "%s", funcBody)

	}

}
