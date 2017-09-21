package Prototype

type BookPrototype struct {
	title    string
	category string
}

func (this *BookPrototype) GetTitle() string {
	return this.title
}
func (this *BookPrototype) SetTitle(title string) {
	this.title = title
}
func (this *BookPrototype) Clone() *BookPrototype {
	newObj := *this
	return &newObj
}

type BarBookPrototype struct {
	category string
	BookPrototype
}

type FooBookPrototype struct {
	BookPrototype
}
