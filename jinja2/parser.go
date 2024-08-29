package jinja2

import (
  "fmt"

	python "github.com/mingfa929/py4go"
)

type Parser struct {
  
}

func NewParser() *Parser {
  return &Parser{}
}

func version() {
	fmt.Printf("Go >> Python version:\n%s\n", python.Version())
}

func (p *Parser) Parse(template string) ([]string) {
  fmt.Println("Parsing template")
	python.PrependPythonPath(".")

	python.Initialize()
	defer python.Finalize()

	version()
	fmt.Println()

	parser, _ := python.Import("parser")
	defer parser.Release()

	parse, _:=parser.GetAttr("parse")
	defer parse.Release()

	r,_:= parse.Call(template)
	defer r.Release()
	rr,_ := r.ToList()
	fmt.Println(rr)

	return rr
}