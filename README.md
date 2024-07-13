# nopkgvarerrrof

nopkgvarerrrof issues a warning if Errorf is used as a package variable definition in Go.

Errors defined as package variables should be created with the New function.

When using third-party error packages or creating your own error package, it is common to accumulate a stack trace with Errorf. In such cases, using Errorf for package variables may result in incorrect stack traces.


```console
go install github.com/nametake/nopkgvarerrorf/cmd/nopkgvarerrorf@latest
```

## Usage

```console
go vet -vettool=`which nopkgvarerrorf` .
```

## Example

```go
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
```
