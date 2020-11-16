package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	url := "^(https?|ftp|mailto)://(www.)?[a-zA-Z0-9-._]{2,256}\\.[a-z]{2,4}\\/?[^\\s]*$"
	r, _ := regexp.Compile(url)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese una url para verificar.")
	for scanner.Scan() {
		var f = scanner.Text()
		if r.MatchString(f) == true {
			fmt.Println("Es una url Válida")
		} else {
			fmt.Println("No es una url Válida")
		}
		fmt.Println("Ingrese otra url, en otro casso oprima Ctrl+C para salir.")
	}
}
