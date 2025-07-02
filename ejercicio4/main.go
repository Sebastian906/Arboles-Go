package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    elemento := 40
    altura := arbol.AlturaHasta(elemento)
    if altura != -1 {
        fmt.Printf("Altura hasta el elemento %d: %d\n", elemento, altura)
    } else {
        fmt.Printf("Elemento %d no encontrado\n", elemento)
    }
    
    elemento = 90
    altura = arbol.AlturaHasta(elemento)
    if altura != -1 {
        fmt.Printf("Altura hasta el elemento %d: %d\n", elemento, altura)
    } else {
        fmt.Printf("Elemento %d no encontrado\n", elemento)
    }
}