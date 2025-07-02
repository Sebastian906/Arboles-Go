package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 90}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Número mayor del árbol:", arbol.Maximo())
}