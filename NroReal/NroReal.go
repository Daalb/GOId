package main
import (
  "fmt"
  "regexp"
  "os"
  "bufio"
)

func comprobar(a string){
	r, _ := regexp.Compile("^[-+]?[0-9]+(.[0-9]+([eE][-+]?[0-9]+)?)?$")
	if r.MatchString(a) == true{
		fmt.Println(a,"Es un número real")
	}else{
		fmt.Println(a,"No es un número real")
	} 
}


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Ingrese un número.")
  for scanner.Scan() {
    comprobar(scanner.Text())
    fmt.Println("Ingrese otro número u oprima Ctrl+C para salir.")
  }
}