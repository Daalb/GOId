package main
import (
  "fmt"
  "regexp"
  "log"
  "os"
  "bufio"
   

)

func comprobar(a string){
	r, _ := regexp.Compile("^[-+]?[0-9]+.?[0-9]+([eE][-+]?[0-9]+)?$")
	if r.MatchString(a) == true{
		fmt.Println(a,"Es un número real")
	}else{
		fmt.Println(a,"No es un número real")
	} 
}

func main() {

  //Abro el archivo
  file, err := os.Open("Leer.txt")

  //Manejo de error mientras se abre
  if err != nil {
	  log.Fatalf("Error when opening file: %s", err)
  }

  fileScanner := bufio.NewScanner(file)

  //Leer linea por linea  
  for fileScanner.Scan() {
	 comprobar(fileScanner.Text())

  }

  //Manejo de error mientras se lee
  if err := fileScanner.Err(); err != nil {
	  log.Fatalf("Error while reading file: %s", err)
  }

  file.Close()
}