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

func (a *Arbol) Altura() int {
    return altura(a.Raiz)
}

func altura(nodo *Nodo) int {
    if nodo == nil {
        return 0
    }
    
    alturaIzq := altura(nodo.Izquierdo)
    alturaDer := altura(nodo.Derecho)
    
    if alturaIzq > alturaDer {
        return alturaIzq + 1
    }
    return alturaDer + 1
}