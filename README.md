# go-mrz-scanner

Utility package to validate and parse the MRZ string.

## Common
For more information regarding the format of the MRZ string, refer to the
[**Machine-readable passport**](https://en.wikipedia.org/wiki/Machine-readable_passport)

## Installation
To install this package, run:
```shell
go get -u github.com/okieraised/go-mrz-scanner
```

## Usage
Below is an example to parse the MRZ components:
```go
package main

import (
	"fmt"
	go_mrz_parser "github.com/okieraised/go-mrz-scanner"
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
	fmt.Println(result)
}
```


