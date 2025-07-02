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

func (a *Arbol) InOrdenLista() []int {
    lista := []int{}
    inOrdenLista(a.Raiz, &lista)
    return lista
}

func inOrdenLista(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        inOrdenLista(nodo.Izquierdo, lista)
        *lista = append(*lista, nodo.Valor)
        inOrdenLista(nodo.Derecho, lista)
    }
}

func (a *Arbol) PreOrdenLista() []int {
    lista := []int{}
    preOrdenLista(a.Raiz, &lista)
    return lista
}

func preOrdenLista(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        *lista = append(*lista, nodo.Valor)
        preOrdenLista(nodo.Izquierdo, lista)
        preOrdenLista(nodo.Derecho, lista)
    }
}

func (a *Arbol) PostOrdenLista() []int {
    lista := []int{}
    postOrdenLista(a.Raiz, &lista)
    return lista
}

func postOrdenLista(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        postOrdenLista(nodo.Izquierdo, lista)
        postOrdenLista(nodo.Derecho, lista)
        *lista = append(*lista, nodo.Valor)
    }
}