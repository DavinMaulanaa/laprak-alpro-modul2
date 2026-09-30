package main
import (
"fmt"
"math"
)
func main(){
var (
jari_jari, luas float64
)
fmt.Println("=====WELCOME=====")
fmt.Println("Masukkan nilai jari jari lingkaran:")
fmt.Scanf("%f", &jari_jari)
luas = math.Pi * jari_jari * jari_jari
fmt.Printf("%.1f\n",luas)
}
