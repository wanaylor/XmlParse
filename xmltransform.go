package main

import (
	"fmt"
	"encoding/xml"
    "os"
)

type Library struct{
    Books []Book `xml:"book"`
}

type Book struct{
    Author string `xml:"author"`
    Title string `xml:"title"`
    Genre string `xml:"genre"`
    Price float64 `xml:"price"`
    Publish_date string `xml:"publish_date"`
    Description string `xml:"description"`
}

func main (){
    fmt.Println("Here we go!");
    var lib Library;
    dat, err := os.ReadFile("import_dbs/library.xml")
    if err != nil{
        panic("could not read file")
    }
    xml.Unmarshal(dat, &lib)
    fmt.Println(lib)
}
