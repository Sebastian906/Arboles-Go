package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 7, 13, 2, 11}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Números primos en el árbol:", arbol.NumerosPrimos())
}