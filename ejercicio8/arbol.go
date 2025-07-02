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

func (a *Arbol) NumerosPrimos() []int {
    lista := []int{}
    numerosPrimos(a.Raiz, &lista)
    return lista
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

func numerosPrimos(nodo *Nodo, lista *[]int) {
    if nodo != nil {
        numerosPrimos(nodo.Izquierdo, lista)
        if esPrimo(nodo.Valor) {
            *lista = append(*lista, nodo.Valor)
        }
        numerosPrimos(nodo.Derecho, lista)
    }
}