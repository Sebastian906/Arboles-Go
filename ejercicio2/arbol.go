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

func (a *Arbol) InOrden() {
    inOrden(a.Raiz)
    fmt.Println()
}

func inOrden(nodo *Nodo) {
    if nodo != nil {
        inOrden(nodo.Izquierdo)
        fmt.Printf("%d ", nodo.Valor)
        inOrden(nodo.Derecho)
    }
}

func (a *Arbol) PreOrden() {
    preOrden(a.Raiz)
    fmt.Println()
}

func preOrden(nodo *Nodo) {
    if nodo != nil {
        fmt.Printf("%d ", nodo.Valor)
        preOrden(nodo.Izquierdo)
        preOrden(nodo.Derecho)
    }
}

func (a *Arbol) PostOrden() {
    postOrden(a.Raiz)
    fmt.Println()
}

func postOrden(nodo *Nodo) {
    if nodo != nil {
        postOrden(nodo.Izquierdo)
        postOrden(nodo.Derecho)
        fmt.Printf("%d ", nodo.Valor)
    }
}