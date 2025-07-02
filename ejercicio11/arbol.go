package main

import "fmt"

type Arbol struct {
    Raiz *Nodo
}

func NuevoArbol() *Arbol {
    return &Arbol{Raiz: nil}
}

func (a *Arbol) Insertar(valor int) {
    a.Raiz = insertarNodo(a.Raiz, valor)
}

func insertarNodo(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{Valor: valor}
    }

    if valor < nodo.Valor {
        nodo.Izquierdo = insertarNodo(nodo.Izquierdo, valor)
    } else if valor > nodo.Valor {
        nodo.Derecho = insertarNodo(nodo.Derecho, valor)
    }

    return nodo
}

func (a *Arbol) MostrarHojas() {
    mostrarHojas(a.Raiz)
    fmt.Println()
}

func mostrarHojas(nodo *Nodo) {
    if nodo != nil {
        if nodo.Izquierdo == nil && nodo.Derecho == nil {
            fmt.Printf("%d ", nodo.Valor)
        }
        mostrarHojas(nodo.Izquierdo)
        mostrarHojas(nodo.Derecho)
    }
}