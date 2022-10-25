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
	dic.Guardar(2, 2) //raiz.izquierdo
	dic.Guardar(1, 1)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8) //raiz.derecho
	dic.Guardar(6, 6)
	dic.Guardar(5, 5)
	dic.Guardar(7, 7)

	require.True(t, dic.Pertenece(1))
	require.EqualValues(t, 1, dic.Borrar(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(1) })
	require.EqualValues(t, 7, dic.Cantidad())
	require.False(t, dic.Pertenece(1))

	require.True(t, dic.Pertenece(4))
	require.EqualValues(t, 4, dic.Borrar(4))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(4) })
	require.EqualValues(t, 6, dic.Cantidad())
	require.False(t, dic.Pertenece(4))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })

	require.True(t, dic.Pertenece(8))
	require.EqualValues(t, 8, dic.Borrar(8))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(8) })
	require.EqualValues(t, 5, dic.Cantidad())
	require.False(t, dic.Pertenece(8))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(8) })

	require.True(t, dic.Pertenece(6))
	require.EqualValues(t, 6, dic.Borrar(6))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(6) })
	require.EqualValues(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(6))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(6) })

	require.True(t, dic.Pertenece(3))
	require.EqualValues(t, 3, dic.Borrar(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(3) })
	require.EqualValues(t, 3, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })

	require.True(t, dic.Pertenece(7))
	require.EqualValues(t, 7, dic.Borrar(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(7) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(7) })

	require.True(t, dic.Pertenece(5))
	require.EqualValues(t, 5, dic.Borrar(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(5) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })

	require.True(t, dic.Pertenece(2))
	require.EqualValues(t, 2, dic.Borrar(2))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(2) })
	require.EqualValues(t, 0, dic.Cantidad())
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
