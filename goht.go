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
func (de DomElement) C() string {
	generated := "<" + de.Key

	for key, elem := range de.Props {
		generated += " " + key + "=\"" + elem + "\""
	}

	generated += ">" + de.Value + "</" + de.Key + ">"
	return generated
}

// Convert document element list to string
func (del DomElementList) C() string {
	generated := ""

	for _, elem := range del {
		generated += elem.C()
	}

	return generated
}
