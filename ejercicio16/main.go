package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("InOrden antes de eliminar:", arbol.InOrden())
    
    // Eliminar nodo hoja (20)
    arbol.Eliminar(20)
    fmt.Println("InOrden después de eliminar 20 (hoja):", arbol.InOrden())
    
    // Eliminar nodo con un hijo (30)
    arbol.Eliminar(30)
    fmt.Println("InOrden después de eliminar 30 (un hijo):", arbol.InOrden())
    
    // Eliminar nodo con dos hijos (50)
    arbol.Eliminar(50)
    fmt.Println("InOrden después de eliminar 50 (dos hijos):", arbol.InOrden())
}