package main

import "fmt"

func main() {
    arbol := NuevoArbolVehiculos()
    
    vehiculos := []*Vehiculo{
        NuevoVehiculo("ABC123", "Toyota", "Corolla"),
        NuevoVehiculo("XYZ789", "Honda", "Civic"),
        NuevoVehiculo("DEF456", "Ford", "Fiesta"),
        NuevoVehiculo("GHI789", "Chevrolet", "Spark"),
    }
    
    for _, v := range vehiculos {
        arbol.Insertar(v)
    }
    
    fmt.Println("Vehículos ordenados por placa:")
    for _, v := range arbol.InOrden() {
        fmt.Printf("Placa: %s, Marca: %s, Modelo: %s\n", v.Placa, v.Marca, v.Modelo)
    }
}