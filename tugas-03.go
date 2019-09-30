package main

import "fmt"

func main(){
  var buah = []string{"apel", "jeruk", "anggur", "melon"}
  buah = append(buah, "pepaya")
  fmt.Println("Jumlah Element : ", len(buah))
  fmt.Println("Isi Element : ", buah)
  fmt.Println("Element ke - : 0 ", (buah[0]))
  fmt.Println("Element ke - : 1 ", (buah[1]))
  fmt.Println("Element ke - : 2 ", (buah[2]))
  fmt.Println("Element ke - : 3 ", (buah[3]))
  fmt.Println("Element ke - : 4 ", (buah[4]))
}
