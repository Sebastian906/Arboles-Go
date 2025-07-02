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

func (a *Arbol) DatosNivel(nivel int) []int {
    lista := []int{}
    datosNivel(a.Raiz, 0, nivel, &lista)
    return lista
}

func datosNivel(nodo *Nodo, nivelActual, nivelDeseado int, lista *[]int) {
    if nodo == nil {
        return
    }

    if nivelActual == nivelDeseado {
        *lista = append(*lista, nodo.Valor)
        return
    }

    datosNivel(nodo.Izquierdo, nivelActual+1, nivelDeseado, lista)
    datosNivel(nodo.Derecho, nivelActual+1, nivelDeseado, lista)
}