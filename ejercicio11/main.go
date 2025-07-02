package main

func main() {
    arbol := NuevoArbol()
    numeros := []int{50, 30, 70, 20, 40, 60, 80, 10, 45, 55, 65}
    
    for _, num := range numeros {
        arbol.Insertar(num)
    }
    
    println("Hojas del árbol:")
    arbol.MostrarHojas()
}