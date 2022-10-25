package diccionario_test

import (
	TDADiccionario "abb"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000}

func TestDiccionarioVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(6) })
}
func TestUnElement(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(1, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(2))
	require.EqualValues(t, 10, dic.Obtener(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })
}

func TestDiccionarioGuardar(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}
func TestDiccionarioBorrar(t *testing.T) {

	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	require.False(t, dic.Pertenece(4))
	dic.Guardar(4, 4) //raiz
	dic.Guardar(0, 0) //raiz.izquierdo
	dic.Guardar(1, 1)
	dic.Guardar(3, 3)
	dic.Guardar(2, 2)
	dic.Guardar(8, 8) //raiz.derecho
	dic.Guardar(6, 6)
	dic.Guardar(5, 5)
	dic.Guardar(7, 7)

	require.True(t, dic.Pertenece(1))
	require.EqualValues(t, 1, dic.Borrar(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(1) })
	require.EqualValues(t, 8, dic.Cantidad())
	require.False(t, dic.Pertenece(1))

	require.True(t, dic.Pertenece(4))
	require.EqualValues(t, 4, dic.Borrar(4))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(4) })
	require.EqualValues(t, 7, dic.Cantidad())
	require.False(t, dic.Pertenece(4))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })

	require.True(t, dic.Pertenece(8))
	require.EqualValues(t, 8, dic.Borrar(8))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(8) })
	require.EqualValues(t, 6, dic.Cantidad())
	require.False(t, dic.Pertenece(8))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(8) })

	require.True(t, dic.Pertenece(6))
	require.EqualValues(t, 6, dic.Borrar(6))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(6) })
	require.EqualValues(t, 5, dic.Cantidad())
	require.False(t, dic.Pertenece(6))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(6) })

	require.True(t, dic.Pertenece(3))
	require.EqualValues(t, 3, dic.Borrar(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(3) })
	require.EqualValues(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })

	require.True(t, dic.Pertenece(7))
	require.EqualValues(t, 7, dic.Borrar(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(7) })
	require.EqualValues(t, 3, dic.Cantidad())
	require.False(t, dic.Pertenece(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(7) })

	require.True(t, dic.Pertenece(5))
	require.EqualValues(t, 5, dic.Borrar(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(5) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })

	require.True(t, dic.Pertenece(2))
	require.EqualValues(t, 2, dic.Borrar(2))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(2) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(2))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })

}
func TestReemplazoDato(t *testing.T) {
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}
func TestBorrarYGuardar(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(150, "arbol")
	dic.Borrar(150)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(150))
	dic.Guardar(150, "binario")
	require.True(t, dic.Pertenece(150))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "binario", dic.Obtener(150))
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	claves := make([]string, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el hash
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que devuelva los valores correctos
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}

		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que borre y devuelva los valores correctos
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}
	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

//hacerlo aleatorio
/*func BenchmarkDiccionario(b *testing.B) {
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}*/

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	require.EqualValues(t, primero, iter.Siguiente())
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	require.EqualValues(t, segundo, iter.Siguiente())

	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	fmt.Println(tercero)
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	require.EqualValues(t, tercero, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()

	iter3 := dic.Iterador()
	primero := iter3.Siguiente()
	segundo := iter3.Siguiente()

	tercero := iter3.Siguiente()

	require.False(t, iter3.HaySiguiente())

	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	require.EqualValues(t, clave1, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)

	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}
