// Copyright (c) 2022 Bi Anlatsana
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package goht

type Props map[string]string

type DomElement struct {
	Key, Value string
	Props      Props
}

type DomList []DomElement

func (de DomElement) C() string {
	generated := "<" + de.Key

	for key, elem := range de.Props {
		generated += " " + key + "=\"" + elem + "\""
	}

	generated += ">" + de.Value + "</" + de.Key + ">"
	return generated
}

func (dl DomList) C() string {
	generated := ""

	for _, elem := range dl {
		generated += elem.C()
	}

	return generated
}
