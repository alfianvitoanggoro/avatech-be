package avapay

import "fmt"

var (
	Name      string
	PayStatus bool
)

func PayerName(name string) {
	fmt.Println("Hello, ", name)
}
