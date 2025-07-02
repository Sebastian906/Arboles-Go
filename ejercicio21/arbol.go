package main

type Arbol struct {
    Raiz *Nodo
}

func NuevoArbol() *Arbol {
    return &Arbol{Raiz: nil}
}

func (a *Arbol) InsertarConRepeticiones(valor int) {
    a.Raiz = insertarConRepeticiones(a.Raiz, valor)
}

func insertarConRepeticiones(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{Valor: valor, Repetidos: 0}
    }

    if valor < nodo.Valor {
        nodo.Izquierdo = insertarConRepeticiones(nodo.Izquierdo, valor)
    } else if valor > nodo.Valor {
        nodo.Derecho = insertarConRepeticiones(nodo.Derecho, valor)
    } else {
        nodo.Repetidos++
    }

    return nodo
}

func (a *Arbol) ObtenerRepeticiones(valor int) int {
    nodo := buscarNodo(a.Raiz, valor)
    if nodo != nil {
        return nodo.Repetidos
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

func (a *Arbol) InOrden() []int {
    lista := []int{}
    inOrden(a.Raiz, &lista)
    return lista
}

func inOrden(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        inOrden(nodo.Izquierdo, lista)
        *lista = append(*lista, nodo.Valor)
        inOrden(nodo.Derecho, lista)
    }
}