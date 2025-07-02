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

func (a *Arbol) AlturaHasta(valor int) int {
    return alturaHasta(a.Raiz, valor, 0)
}

func alturaHasta(nodo *Nodo, valor, alturaActual int) int {
    if nodo == nil {
        return -1 // No encontrado
    }
    
    if nodo.Valor == valor {
        return alturaActual
    }
    
    if valor < nodo.Valor {
        return alturaHasta(nodo.Izquierdo, valor, alturaActual+1)
    } else {
        return alturaHasta(nodo.Derecho, valor, alturaActual+1)
    }
}