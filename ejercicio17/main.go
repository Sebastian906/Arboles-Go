package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 30, 20, 50, 80, 50, 20, 30, 50}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    fmt.Println("Ocurrencias de 50:", arbol.ContarOcurrencias(50))
    fmt.Println("Ocurrencias de 30:", arbol.ContarOcurrencias(30))
    fmt.Println("Ocurrencias de 20:", arbol.ContarOcurrencias(20))
    fmt.Println("Ocurrencias de 70:", arbol.ContarOcurrencias(70))
    fmt.Println("Ocurrencias de 80:", arbol.ContarOcurrencias(80))
    fmt.Println("Ocurrencias de 100:", arbol.ContarOcurrencias(100))
}