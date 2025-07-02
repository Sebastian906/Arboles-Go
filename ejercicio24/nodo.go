package main

type Nodo struct {
    Vehiculo  *Vehiculo
    Izquierdo *Nodo
    Derecho   *Nodo
}