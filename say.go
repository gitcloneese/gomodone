package gomodone

import (
	"errors"
	"fmt"
)

func SayHi(name, lang, greeter string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s, %s!", name, greeter), nil
	case "pt":
		return fmt.Sprintf("Oi, %s, %s!", name, greeter), nil
	case "es":
		return fmt.Sprintf("Hola, %s, %s!", name, greeter), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s, %s!", name, greeter), nil
	default:
		return "", errors.New("unkown language ...")
	}
}
