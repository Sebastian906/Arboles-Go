package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("InOrden como lista:", arbol.InOrdenLista())
    fmt.Println("PreOrden como lista:", arbol.PreOrdenLista())
    fmt.Println("PostOrden como lista:", arbol.PostOrdenLista())
}