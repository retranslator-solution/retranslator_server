package models

import "encoding/xml"

type Resources struct {
	XMLName xml.Name `xml:"resources"`
	Strings interface{}
}

type Name struct {
	Name string `xml:"name,attr"`
}

type Text struct {
	Text string `xml:",innerxml"`
}

type StringResource struct {
	Name
	Text
	XMLName xml.Name `xml:"string"`
}

type ArrayItem struct {
	Text
	XMLName xml.Name `xml:"item"`
}

type ArrayResource struct {
	Name
	XMLName xml.Name `xml:"string-array"`
	Items   []ArrayItem
}

type PluralItem struct {
	Text
	XMLName  xml.Name `xml:"item"`
	Quantity string   `xml:"quantity,attr"`
}

type PluralResource struct {
	Name
	XMLName xml.Name `xml:"plurals"`
}
