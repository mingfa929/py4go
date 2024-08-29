package jinja2
import (
	"testing"
	"fmt"
)

func TestParser(t *testing.T) {
	parser := NewParser()
	input := parser.Parse("{{ foo }}")
	fmt.Println(input)
	// ast := parser.Parse("{{ foo }}")
	// if ast == nil {
	// 	t.Fatal("parser.Parse returned nil")
	// }
	// if ast.Type != "Root" {
	// 	t.Fatalf("expected Root node, got %s", ast.Type)
	// }
}