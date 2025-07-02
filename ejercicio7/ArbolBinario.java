public class ArbolBinario {
    private Nodo raiz;

    public ArbolBinario() {
        this.raiz = null;
    }

    public void insertar(int valor) {
        raiz = insertarNodo(raiz, valor);
    }

    private Nodo insertarNodo(Nodo nodo, int valor) {
        if (nodo == null) {
            return new Nodo(valor);
        }

        if (valor < nodo.valor) {
            nodo.izquierdo = insertarNodo(nodo.izquierdo, valor);
        } else if (valor > nodo.valor) {
            nodo.derecho = insertarNodo(nodo.derecho, valor);
        }

        return nodo;
    }

    public int sumaDesdeNivel(int nivel) {
        return sumaDesdeNivel(raiz, 0, nivel);
    }

    private int sumaDesdeNivel(Nodo nodo, int nivelActual, int nivelDeseado) {
        if (nodo == null) {
            return 0;
        }

        if (nivelActual >= nivelDeseado) {
            return nodo.valor + 
                   sumaDesdeNivel(nodo.izquierdo, nivelActual + 1, nivelDeseado) + 
                   sumaDesdeNivel(nodo.derecho, nivelActual + 1, nivelDeseado);
        } else {
            return sumaDesdeNivel(nodo.izquierdo, nivelActual + 1, nivelDeseado) + 
                   sumaDesdeNivel(nodo.derecho, nivelActual + 1, nivelDeseado);
        }
    }

    public void inOrden() {
        inOrden(raiz);
        System.out.println();
    }

    private void inOrden(Nodo nodo) {
        if (nodo != null) {
            inOrden(nodo.izquierdo);
            System.out.print(nodo.valor + " ");
            inOrden(nodo.derecho);
        }
    }
}