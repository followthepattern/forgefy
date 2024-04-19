package main

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/io"
)

func main() {
	yaml := `
version: 0
product_name: "test product"
apps:
  - name: backend1
    type: backend
  - name: frontend1
    type: frontend
features:
  - name: user
    fields:
      - name: ID
        type: string
      - name: FirstName
        type: string
      - name: LastName
        type: string
      - name: Age
        type: int
  - name: product
    fields:
      - name: ID
        type: string
      - name: Name
        type: string
      - name: Type
        type: string
      - name: Price
        type: int
  - name: task
  - name: image
`
	f := forgefy.New()
	fw := io.NewFileWriter("output")

	f.Forge(yaml, fw)
}
