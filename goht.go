// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package goht

import "fmt"

// Props type
type Props map[string]string

// Document element struct
type DomElement struct {
	Key, Value string
	Props      Props
}

// Document element list
type DomElementList []DomElement

// Convert document element to string
func (element DomElement) C() string {
	generated := "<" + element.Key + element.Props.Build()

	if element.Key == "area" || element.Key == "base" || element.Key == "br" || element.Key == "col" || element.Key == "embed" || element.Key == "hr" || element.Key == "img" || element.Key == "input" || element.Key == "link" || element.Key == "meta" || element.Key == "param" || element.Key == "source" || element.Key == "track" || element.Key == "wbr" {
		generated += " />"
	} else {
		generated += element.Value + "</" + element.Key + ">"
	}

	return generated
}

// Convert document element list to string
func (elementsList DomElementList) C() string {
	generated := ""

	for _, elem := range elementsList {
		generated += elem.C()
	}

	return generated
}

// Build props
func (p Props) Build() string {
	generated := ""
	for key, elem := range p {
		generated += " " + key + "=\"" + elem + "\""
	}

	return generated
}

// Mark document as HTML
func Html(props Props, elements ...DomElement) string {
	return "<!DOCTYPE html>\n" + Dom("html", DomList(elements...), props)
}

// Encode string to prevent XSS attack
func Encode(body string) string {
	generated := ""

	for _, char := range body {
		generated += fmt.Sprintf("&#%d;", char)
	}

	return generated
}
