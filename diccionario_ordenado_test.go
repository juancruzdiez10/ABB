package diccionario_test

import (
	TDADiccionario "abb"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(6) })
}
func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(1, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(2))
	require.EqualValues(t, 10, dic.Obtener(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })
}
func TestDiccionarioBorrar(t *testing.T) {

	clave1 := 4
	clave2 := 1
	clave3 := 3
	valor1 := 4
	valor2 := 1
	valor3 := 3
	claves := []int{clave1, clave2, clave3}
	valores := []int{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(8, 8)
	dic.Guardar(5, 5)

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, 1, dic.Borrar(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(1) })
	require.EqualValues(t, 6, dic.Cantidad())
	require.False(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.True(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(8))
}
