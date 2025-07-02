package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 7, 13, 2, 11, 15}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    pares, primos := arbol.ParesYPrimos()
    fmt.Println("Números pares:", pares)
    fmt.Println("Números primos:", primos)
}