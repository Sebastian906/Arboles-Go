package main

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

func (a *Arbol) Nivel() int {
    return nivel(a.Raiz)
}

func nivel(nodo *Nodo) int {
    if nodo == nil {
        return 0
    }
    
    izq := nivel(nodo.Izquierdo)
    der := nivel(nodo.Derecho)
    
    if izq > der {
        return izq + 1
    }
    return der + 1
}