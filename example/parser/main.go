package main

import (
	"fmt"
	go_mrz_parser "github.com/okieraised/go-mrz-scanner"
	"github.com/okieraised/go-mrz-scanner/constants"
	"log"
)

func main() {

	input := []string{
		"I<UTOD231458907<<<<<<<<<<<<<<<",
		"7408122F1204159UTO<<<<<<<<<<<6",
		"ERIKSSON<<ANNA<MARIA<<<<<<<<<<",
	}

	parser := go_mrz_parser.NewMRZLineParser(input)
	result, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Fields[constants.MRZFieldGender].GetRawValue())

	names := result.Fields[constants.MRZFieldName].GetValue().([]string)
	fmt.Println(names)
}
