package main

import (
	"encoding/xml"
	"fmt"
)

// Go offers built-in support for XML and XML-like formats

// xml.Name dictates the name of the XML element
// id,attr means that the field is an XML attribute, rather than an element
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func xmls() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emit XML // Marshal Indent makes the output more readable
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Add generic XML header
	fmt.Println(xml.Header + string(out))

	// Unmarshal XML
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

	// parent>child>plant tells the encoder to 
	// nest all plants under <parent><child>
	// could name them whatever we want
	type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }

	nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}