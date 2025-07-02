public class Main {
    public static void main(String[] args) {
        ArbolBinario arbol = new ArbolBinario();
        int[] numeros = {50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85};
        
        for (int num : numeros) {
            arbol.insertar(num);
        }
        
        System.out.println("Recorrido InOrden del árbol:");
        arbol.inOrden();
        
        int nivel = 2;
        int suma = arbol.sumaDesdeNivel(nivel);
        System.out.println("Suma desde el nivel " + nivel + ": " + suma);
        
        nivel = 1;
        suma = arbol.sumaDesdeNivel(nivel);
        System.out.println("Suma desde el nivel " + nivel + ": " + suma);
        
        nivel = 0;
        suma = arbol.sumaDesdeNivel(nivel);
        System.out.println("Suma desde el nivel " + nivel + ": " + suma);
    }
}