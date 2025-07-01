package main

import "fmt"

func main() {
    arbol := &ABB{}
    arbol.Insertar(5)
    arbol.Insertar(3)
    arbol.Insertar(8)
    arbol.Insertar(1)
    arbol.Insertar(4)
    arbol.Insertar(7)
    arbol.Insertar(9)
    
    fmt.Print("Recorrido in-orden: ")
    arbol.InOrden()
    
    fmt.Printf("Altura del árbol: %d\n", arbol.Altura())
    fmt.Printf("Tamaño del árbol: %d\n", arbol.Tamano())
    
    fmt.Printf("¿Existe el 4? %t\n", arbol.Buscar(4))
    fmt.Printf("¿Existe el 10? %t\n", arbol.Buscar(10))
}