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
  - name: product
  - name: task
  - name: image
`
	f := forgefy.New()
	fw := io.NewFileWriter("output")

	f.Forge(yaml, fw)
}
