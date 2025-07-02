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
        return &Nodo{Valor: valor, Repetidos: 0}
    }

    if valor < nodo.Valor {
        nodo.Izquierdo = insertarNodo(nodo.Izquierdo, valor)
    } else if valor > nodo.Valor {
        nodo.Derecho = insertarNodo(nodo.Derecho, valor)
    } else {
        nodo.Repetidos++
    }

    return nodo
}

func (a *Arbol) ObtenerRepeticiones(valor int) int {
    nodo := buscarNodo(a.Raiz, valor)
    if nodo != nil {
        return nodo.Repetidos + 1 // +1 porque cuenta la existencia inicial
    }
    return 0
}

func buscarNodo(nodo *Nodo, valor int) *Nodo {
    if nodo == nil || nodo.Valor == valor {
        return nodo
    }
    
    if valor < nodo.Valor {
        return buscarNodo(nodo.Izquierdo, valor)
    }
    return buscarNodo(nodo.Derecho, valor)
}