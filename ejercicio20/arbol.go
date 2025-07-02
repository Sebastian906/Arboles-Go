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

func (a *Arbol) Pares() []int {
    lista := []int{}
    pares(a.Raiz, &lista)
    return lista
}

func pares(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        pares(nodo.Izquierdo, lista)
        if nodo.Valor%2 == 0 {
            *lista = append(*lista, nodo.Valor)
        }
        pares(nodo.Derecho, lista)
    }
}

func (a *Arbol) CrearArbolPares() *Arbol {
    arbolPares := NuevoArbol()
    listaPares := a.Pares()
    
    for _, num := range listaPares {
        arbolPares.Insertar(num)
    }
    
    return arbolPares
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