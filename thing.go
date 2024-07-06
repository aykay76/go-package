package thing

import "fmt"

type Thing struct {
	Name string
}

func (thing Thing) Hello() {
	fmt.Println(thing.Name)
}

func init() {
	var t Thing
	t.Hello()
}
