package main

type ArbolParesImpares struct {
    Raiz *Nodo
}

func NuevoArbolParesImpares() *ArbolParesImpares {
    return &ArbolParesImpares{Raiz: nil}
}

func (a *ArbolParesImpares) Insertar(valor int) {
    a.Raiz = insertarParesImpares(a.Raiz, valor)
}

func insertarParesImpares(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{Valor: valor}
    }

    if valor%2 == 0 { // Par va a la izquierda
        nodo.Izquierdo = insertarParesImpares(nodo.Izquierdo, valor)
    } else { // Impar va a la derecha
        nodo.Derecho = insertarParesImpares(nodo.Derecho, valor)
    }

    return nodo
}

func (a *ArbolParesImpares) InOrden() []int {
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