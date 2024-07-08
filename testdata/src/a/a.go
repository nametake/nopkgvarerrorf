package a

import "fmt"

var ErrNotFound = fmt.Errorf("not found") // want "do not use Errorf at the package level"

func f() error {
	return fmt.Errorf("not found")
}

func g() {
	err := fmt.Errorf("not found")
	fmt.Println(err)
}

var ErrNotFound2 = Errorf() // want "do not use Errorf at the package level"

func Errorf() error {
	return fmt.Errorf("not found")
}
