package main

type ArbolPrimos struct {
    Raiz *Nodo
}

func NuevoArbolPrimos() *ArbolPrimos {
    return &ArbolPrimos{Raiz: nil}
}

func (a *ArbolPrimos) Insertar(valor int) {
    a.Raiz = insertarPrimos(a.Raiz, valor)
}

func insertarPrimos(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{Valor: valor}
    }

    if esPrimo(valor) { // Primo va a la derecha
        nodo.Derecho = insertarPrimos(nodo.Derecho, valor)
    } else { // No primo va a la izquierda
        nodo.Izquierdo = insertarPrimos(nodo.Izquierdo, valor)
    }

    return nodo
}

func esPrimo(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func (a *ArbolPrimos) InOrden() []int {
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