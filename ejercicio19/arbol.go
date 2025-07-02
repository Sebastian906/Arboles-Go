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

func (a *Arbol) ParesYPrimos() ([]int, []int) {
    pares := []int{}
    primos := []int{}
    
    paresYPrimos(a.Raiz, &pares, &primos)
    
    return pares, primos
}

func esPar(num int) bool {
    return num%2 == 0
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

func paresYPrimos(nodo *Nodo, pares, primos *[]int) {
    if nodo != nil {
        paresYPrimos(nodo.Izquierdo, pares, primos)
        
        if esPar(nodo.Valor) {
            *pares = append(*pares, nodo.Valor)
        }
        
        if esPrimo(nodo.Valor) {
            *primos = append(*primos, nodo.Valor)
        }
        
        paresYPrimos(nodo.Derecho, pares, primos)
    }
}