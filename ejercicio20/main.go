package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 7, 13, 2, 11, 15}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    arbolPares := arbol.CrearArbolPares()
    fmt.Println("Números pares en el árbol original:", arbol.Pares())
    fmt.Println("Recorrido InOrden del árbol de pares:", arbolPares.InOrden())
}