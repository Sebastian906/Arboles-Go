package main

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    println("Recorrido InOrden:")
    arbol.InOrden()
    
    println("\nRecorrido PreOrden:")
    arbol.PreOrden()
    
    println("\nRecorrido PostOrden:")
    arbol.PostOrden()
}