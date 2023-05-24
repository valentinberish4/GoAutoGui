package keyboard

import (
	"fmt"
	"time"
)

func press(key rune, duration time.Time) error {
	fmt.Println("darwin")
	return nil
}

func write(input string) error {
	return nil
}
