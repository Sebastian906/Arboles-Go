package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("InOrden antes de intercambiar:", arbol.InOrden())
    
    if arbol.Intercambiar(30, 70) {
        fmt.Println("InOrden después de intercambiar 30 y 70:", arbol.InOrden())
    } else {
        fmt.Println("No se encontraron ambos nodos")
    }
}