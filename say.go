package gomodone

import (
	"fmt"
	"errors"
)

func SayHi(name, lang string) (string, error) {
	switch lang {
		case "en" :
			return fmt.Sprintf("Hi, %s!", name), nil
		case "pt" :
			return fmt.Sprintf("Oi, %s!", name), nil
		case "es" :
			return fmt.Sprintf("Hola, %s!", name), nil
		case "fr" :
			return fmt.Sprintf("Bonjour, %s!", name), nil
		default :
			return "", errors.New("unkown language ...")
	}
}
