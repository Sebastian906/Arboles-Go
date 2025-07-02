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
