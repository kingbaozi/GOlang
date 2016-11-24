package demo

import (
	"fmt"
)

type Base struct{}

type Voodoo struct {
	Base
}

func (Base) Magic() {
	fmt.Println("Auuuuuuu what the fuck!!!")
}
func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

func (Voodoo) Magic() {
	fmt.Println("Ha this voodooooooo")
}
