package main

import "fmt"

type Nodo struct {
	valor     int
	izquierda *Nodo
	derecha   *Nodo
}

type ABB struct {
	raiz *Nodo
}

func (t *ABB) Buscar(valor int) bool {
	return buscarNodo(t.raiz, valor)
}

func buscarNodo(nodo *Nodo, valor int) bool {
	if nodo == nil {
		return false
	}
	if valor == nodo.valor {
		return true
	} else if valor < nodo.valor {
		return buscarNodo(nodo.izquierda, valor)
	} else {
		return buscarNodo(nodo.derecha, valor)
	}
}

func (t *ABB) InOrden() {
	inOrdenNodo(t.raiz)
	fmt.Println()
}

func inOrdenNodo(nodo *Nodo) {
	if nodo != nil {
		inOrdenNodo(nodo.izquierda)
		fmt.Printf("%d ", nodo.valor)
		inOrdenNodo(nodo.derecha)
	}
}

func (t *ABB) Eliminar(valor int) {
    t.raiz = eliminarNodo(t.raiz, valor)
}

func eliminarNodo(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return nil
    }
    
    if valor < nodo.valor {
        nodo.izquierda = eliminarNodo(nodo.izquierda, valor)
    } else if valor > nodo.valor {
        nodo.derecha = eliminarNodo(nodo.derecha, valor)
    } else {
        // Nodo a eliminar encontrado
        if nodo.izquierda == nil {
            return nodo.derecha
        } else if nodo.derecha == nil {
            return nodo.izquierda
        }
        
        // Nodo con dos hijos
        min := encontrarMinimo(nodo.derecha)
        nodo.valor = min.valor
        nodo.derecha = eliminarNodo(nodo.derecha, min.valor)
    }
    return nodo
}

func encontrarMinimo(nodo *Nodo) *Nodo {
    for nodo.izquierda != nil {
        nodo = nodo.izquierda
    }
    return nodo
}

func (t *ABB) Insertar(valor int) {
	t.raiz = insertarNodo(t.raiz, valor)
}

func insertarNodo(nodo *Nodo, valor int) *Nodo {
    if nodo == nil {
        return &Nodo{valor: valor}
    }
    if valor < nodo.valor {
        nodo.izquierda = insertarNodo(nodo.izquierda, valor)
    } else if valor > nodo.valor {
        nodo.derecha = insertarNodo(nodo.derecha, valor)
    }
    // Si valor == nodo.valor, no hacemos nada (duplicado ignorado)
    // O podrías retornar un error o incrementar un contador
    return nodo
}

func (t *ABB) Altura() int {
    return alturaNodo(t.raiz)
}

func alturaNodo(nodo *Nodo) int {
    if nodo == nil {
        return 0
    }
    alturaIzq := alturaNodo(nodo.izquierda)
    alturaDer := alturaNodo(nodo.derecha)
    if alturaIzq > alturaDer {
        return alturaIzq + 1
    }
    return alturaDer + 1
}

func (t *ABB) Tamano() int {
    return tamanoNodo(t.raiz)
}

func tamanoNodo(nodo *Nodo) int {
    if nodo == nil {
        return 0
    }
    return 1 + tamanoNodo(nodo.izquierda) + tamanoNodo(nodo.derecha)
}