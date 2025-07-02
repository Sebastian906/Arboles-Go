package main

import "fmt"

func main() {
    arbol := NuevoArbolParesImpares()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 7, 13, 2, 11, 15}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Recorrido InOrden del árbol pares-impares:", arbol.InOrden())
}