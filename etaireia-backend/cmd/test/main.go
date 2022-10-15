package main

import (
	"log"

	"github.com/piprate/json-gold/ld"
)

func main() {

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// // expanding remote document

	// expanded, err := proc.Expand("http://json-ld.org/test-suite/tests/expand-0002-in.jsonld", options)
	// if err != nil {
	// 	log.Println("Error when expanding JSON-LD document:", err)
	// 	return
	// }

	// log.Println(expanded...)

	// expanding in-memory document

	doc := map[string]interface{}{
		"@context":  "http://schema.org/",
		"@type":     "Person",
		"name":      "Jane Doe",
		"jobTitle":  "Professor",
		"telephone": "(425) 123-4567",
		"url":       "http://www.janedoe.com",
	}

	expanded, _ := proc.Expand(doc, options)

	log.Println(expanded)

}
