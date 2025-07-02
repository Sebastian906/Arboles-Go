package main

type Vehiculo struct {
    Placa  string
    Marca  string
    Modelo string
}

func NuevoVehiculo(placa, marca, modelo string) *Vehiculo {
    return &Vehiculo{
        Placa:  placa,
        Marca:  marca,
        Modelo: modelo,
    }
}