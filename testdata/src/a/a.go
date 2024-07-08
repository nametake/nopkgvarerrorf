package a

import "fmt"

var ErrNotFound = fmt.Errorf("not found") // want "do not use Errorf at the package level"

func f() error {
	return fmt.Errorf("not found")
}
