package main

type PlaceholderFormat interface {
	ReplacePlaceholders(sql string) (string, error)
}
type Builder struct{}
type StatementBuilderType Builder

func (b StatementBuilderType) PlaceholderFormat(f PlaceholderFormat) StatementBuilderType {
	return (struct{}{}).(StatementBuilderType)
}

var EmptyBuilder = Builder{}
var StatementBuilder = StatementBuilderType(EmptyBuilder).PlaceholderFormat()

func main() {
	StatementBuilder.PlaceholderFormat(struct{})
}
