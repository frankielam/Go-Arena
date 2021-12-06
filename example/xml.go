package main

import (
    "fmt"
    "encoding/xml"
)

type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id int `xml:"id,attr"`
    Name string `xml:"name"`
    Origin []string `xml:"origin"`
}

func (p Plant) String() string {
   return fmt.Sprintf("Plant xml=%v, id=%v, name=%v, origin=%v ", p.XMLName, p.Id, p.Name, p.Origin)
}


func main() {
    plant := &Plant{
        Id: 1,
	Name: "Flower"}
    plant.Origin = []string{"China", "USA"}
    fmt.Println(plant)

    out, _ := xml.MarshalIndent(plant, " ", "  ")
    fmt.Println(string(out))

    fmt.Println(xml.Header)

    var p Plant
    err := xml.Unmarshal(out, &p)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    fmt.Println(p)

    tree := &Plant{Id: 2, Name: "tree", Origin: []string{"GZ", "SH"}}
    fmt.Println(tree)
    
    type Nesting struct {
      XMLName xml.Name `xml:"nesting"`
      Plants []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{plant, tree}

    out, _ = xml.MarshalIndent(nesting, "  ", "  ")
    fmt.Printf(string(out))
}
