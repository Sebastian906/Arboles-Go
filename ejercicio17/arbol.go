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

func (a *Arbol) ContarOcurrencias(valor int) int {
    return contarOcurrencias(a.Raiz, valor)
}

func contarOcurrencias(nodo *Nodo, valor int) int {
    if nodo == nil {
        return 0
    }
    
    contador := 0
    if nodo.Valor == valor {
        contador = 1 + nodo.Repetidos
    }
    
    if valor <= nodo.Valor {
        contador += contarOcurrencias(nodo.Izquierdo, valor)
    }
    if valor >= nodo.Valor {
        contador += contarOcurrencias(nodo.Derecho, valor)
    }
    
    return contador
}