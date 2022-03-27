// Copyright (c) 2022 Bi Anlatsana
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package goht

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
	generated := "<" + element.Key

	for key, elem := range element.Props {
		generated += " " + key + "=\"" + elem + "\""
	}

	generated += ">" + element.Value + "</" + element.Key + ">"
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

// Mark document as HTML
func Html(elements ...DomElement) string {
	return "<!DOCTYPE html>\n" + DomList(elements...)
}
