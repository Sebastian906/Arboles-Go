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

func (a *Arbol) Eliminar(valor int) {
    a.Raiz = eliminarNodo(a.Raiz, valor)
}

func eliminarNodo(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return nil
    }
    
    if valor < nodo.Valor {
        nodo.Izquierdo = eliminarNodo(nodo.Izquierdo, valor)
    } else if valor > nodo.Valor {
        nodo.Derecho = eliminarNodo(nodo.Derecho, valor)
    } else {
        // Caso 1: Nodo hoja
        if nodo.Izquierdo == nil && nodo.Derecho == nil {
            return nil
        }
        
        // Caso 2: Nodo con un hijo
        if nodo.Izquierdo == nil {
            return nodo.Derecho
        }
        if nodo.Derecho == nil {
            return nodo.Izquierdo
        }
        
        // Caso 3: Nodo con dos hijos
        sucesor := minimoNodo(nodo.Derecho)
        nodo.Valor = sucesor.Valor
        nodo.Derecho = eliminarNodo(nodo.Derecho, sucesor.Valor)
    }
    
    return nodo
}

func minimoNodo(nodo *Nodo) *Nodo {
    actual := nodo
    for actual.Izquierdo != nil {
        actual = actual.Izquierdo
    }
    return actual
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