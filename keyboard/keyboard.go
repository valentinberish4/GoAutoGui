package keyboard

import "time"

func Press(key rune, duration time.Time) error {
	return press(key, duration)
}

func Write(input string) error {
	return write(input)
}
