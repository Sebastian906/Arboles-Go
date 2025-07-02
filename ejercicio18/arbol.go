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

func (a *Arbol) Intercambiar(valor1, valor2 int) bool {
    nodo1 := buscarNodo(a.Raiz, valor1)
    nodo2 := buscarNodo(a.Raiz, valor2)
    
    if nodo1 == nil || nodo2 == nil {
        return false
    }
    
    nodo1.Valor, nodo2.Valor = nodo2.Valor, nodo1.Valor
    return true
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