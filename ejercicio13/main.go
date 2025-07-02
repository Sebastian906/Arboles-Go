package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 30, 20, 50, 80, 50, 20}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Repeticiones de 50:", arbol.ObtenerRepeticiones(50))
    fmt.Println("Repeticiones de 30:", arbol.ObtenerRepeticiones(30))
    fmt.Println("Repeticiones de 20:", arbol.ObtenerRepeticiones(20))
    fmt.Println("Repeticiones de 70:", arbol.ObtenerRepeticiones(70))
    fmt.Println("Repeticiones de 80:", arbol.ObtenerRepeticiones(80))
    fmt.Println("Repeticiones de 100:", arbol.ObtenerRepeticiones(100))
}