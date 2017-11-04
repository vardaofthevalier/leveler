package main

import (
    "fmt"
    "go/importer"
)

func main() {
    pkg, err := importer.Default().Import("leveler/endpoints")
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return
    }
    for _, declName := range pkg.Scope().Names() {
        fmt.Println(declName)
    }
}