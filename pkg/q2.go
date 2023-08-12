package pkg

import (
	"fmt"
	"log"
)

// create and throw a custom error
type DivisionByZero struct {
	s string
}

func NewDivisionByZero(s string) error {
	return &DivisionByZero{s}
}

func (d DivisionByZero) Error() string {
	return fmt.Sprintf("can divide by zero: %s", d.s)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, NewDivisionByZero(fmt.Sprintf("%d / %d", a, b))
	}

	return a / b, nil
}

func Q2() {
	_, err := div(1, 0)
	if err != nil {
		if e, ok := err.(*DivisionByZero); ok {
			log.Printf("specific error: %s\n", e)
			return
		}
		log.Printf("generic error: %s\n", err)
	}
}
