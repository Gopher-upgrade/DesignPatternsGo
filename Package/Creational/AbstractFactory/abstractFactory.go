package AbstractFactory

import (
	"fmt"
)


type BaseFactory interface {
	Production(text *string)
}

type CreatorFactory interface {
	CreateText() BaseFactory
}

type JsonFactory struct {
}

type HtmlFactory struct {
}

type Result struct {
}

func (r *Result) Production(text *string) {
	fmt.Printf("依赖于抽象而不依赖于具体 %v", *text)
}

func (J *JsonFactory) CreateText() BaseFactory {
	return new(Result)
}

func (H *HtmlFactory) CreateText() BaseFactory {
	return new(Result)
}
