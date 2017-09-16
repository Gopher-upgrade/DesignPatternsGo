package AbstractFactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var f CreatorFactory
		f = new(JsonFactory)
		j := f.CreateText()
		jsonRender := "我是一个 Json \n"
		j.Production(&jsonRender)

		f = new(HtmlFactory)
		h := f.CreateText()
		htmlRender := "我是一个 Html \n"
	    h.Production(&htmlRender)
}
