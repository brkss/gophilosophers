package main 

import (
  "fmt"
  "os"
  "log"
)


func main(){

  argv := os.Args

  // check args validity 
  if len(argv) != 5 && len(argv) != 6 {
    log.Fatal("Invalid Args !")
  }



  fmt.Println("args : ", argv)

}
