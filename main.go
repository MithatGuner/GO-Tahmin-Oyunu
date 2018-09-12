package main

import ("fmt"
        "bufio"
        "os")

func main() {
  tara := bufio.NewScanner(os.Stdin)
  fmt.Println("Lütfen Aklından Bir Sayı Tut 1 ile 100 Arasında")
  fmt.Println("Hazır Olduğuna Enter Tuşuna Bas")
  tara.Scan()

  tahmin := 50
  for{
    fmt.Println("Tahmin Ettiğim Sayı :", tahmin)
    fmt.Println("Yüksek (a)")
    fmt.Println("Düşük (b)")
    fmt.Println("Doğru (c)")
    tara.Scan()
    cikti := tara.Text()

    if(cikti == "a"){
      tahmin--
    }else if(cikti == "b"){
      tahmin++
    }else if(cikti == "c"){
      fmt.Println("Ben Kazandım")
      break
    }else{
      fmt.Println("Hatalı Giriş, Tekrar Dene")
    }
  }
}
