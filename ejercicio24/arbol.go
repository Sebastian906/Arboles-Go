package main

type ArbolVehiculos struct {
    Raiz *Nodo
}

func NuevoArbolVehiculos() *ArbolVehiculos {
    return &ArbolVehiculos{Raiz: nil}
}

func (a *ArbolVehiculos) Insertar(vehiculo *Vehiculo) {
    a.Raiz = insertarVehiculo(a.Raiz, vehiculo)
}

func insertarVehiculo(nodo *Nodo, vehiculo *Vehiculo) *Nodo {
    if nodo == nil {
        return &Nodo{Vehiculo: vehiculo}
    }

    if vehiculo.Placa < nodo.Vehiculo.Placa {
        nodo.Izquierdo = insertarVehiculo(nodo.Izquierdo, vehiculo)
    } else if vehiculo.Placa > nodo.Vehiculo.Placa {
        nodo.Derecho = insertarVehiculo(nodo.Derecho, vehiculo)
    }

    return nodo
}

func (a *ArbolVehiculos) InOrden() []*Vehiculo {
    lista := []*Vehiculo{}
    inOrden(a.Raiz, &lista)
    return lista
}

func inOrden(nodo *Nodo, lista *[]*Vehiculo) {
    if nodo != nil {
        inOrden(nodo.Izquierdo, lista)
        *lista = append(*lista, nodo.Vehiculo)
        inOrden(nodo.Derecho, lista)
    }
}