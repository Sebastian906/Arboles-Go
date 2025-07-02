package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 10, 45, 55, 65}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Cantidad de nodos en el árbol:", arbol.ContarNodos())
}