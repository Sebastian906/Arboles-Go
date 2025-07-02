package main

import "fmt"

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    for nivel := 0; nivel < 4; nivel++ {
        fmt.Printf("Datos en el nivel %d: %v\n", nivel, arbol.DatosNivel(nivel))
    }
}