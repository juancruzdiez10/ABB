package diccionario_test

import (
	TDADiccionario "abb"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{100000, 300000, 500000}

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

	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	require.False(t, dic.Pertenece("azul"))
	require.False(t, dic.Pertenece("azul"))
	dic.Guardar("azul", "mar")
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("azul"))
	require.True(t, dic.Pertenece("azul"))
	require.EqualValues(t, "mar", dic.Obtener("azul"))
	require.EqualValues(t, "mar", dic.Obtener("azul"))

	require.False(t, dic.Pertenece("verde"))
	require.False(t, dic.Pertenece("amarillo"))
	dic.Guardar("verde", "arbol")
	require.True(t, dic.Pertenece("azul"))
	require.True(t, dic.Pertenece("verde"))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "mar", dic.Obtener("azul"))
	require.EqualValues(t, "arbol", dic.Obtener("verde"))

	require.False(t, dic.Pertenece("amarillo"))
	dic.Guardar("amarillo", "girasol")
	require.True(t, dic.Pertenece("azul"))
	require.True(t, dic.Pertenece("verde"))
	require.True(t, dic.Pertenece("amarillo"))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, "mar", dic.Obtener("azul"))
	require.EqualValues(t, "arbol", dic.Obtener("verde"))
	require.EqualValues(t, "girasol", dic.Obtener("amarillo"))
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

	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar("Gato", "miau")
	dic.Guardar("Perro", "guau")
	require.True(t, dic.Pertenece("Gato"))
	require.True(t, dic.Pertenece("Perro"))
	require.EqualValues(t, "miau", dic.Obtener("Gato"))
	require.EqualValues(t, "guau", dic.Obtener("Perro"))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar("Gato", "miu")
	dic.Guardar("Perro", "baubau")
	require.True(t, dic.Pertenece("Gato"))
	require.True(t, dic.Pertenece("Perro"))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener("Gato"))
	require.EqualValues(t, "baubau", dic.Obtener("Perro"))
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
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	claves := make([]int, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = rand.Intn(500000000000000000) //pa q no se repitan
		claves[i] = valores[i]
		dic.Guardar(claves[i], valores[i])
	}

	ok := true
	for i := 0; i < dic.Cantidad(); i++ {
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

	// Verifica que borre y devuelva los valores correctos
	cantidad := dic.Cantidad()
	for i := 0; i < cantidad; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	} //la idea es q al guardar elementos no se repitan, xq sino al borrar estariamos eliminando elementos q ya borramos antes
	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

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
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	require.EqualValues(t, tercero, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
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
func TestInternoStrings(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })

	dic.Guardar("perro", 2)
	dic.Guardar("ballena", 5)
	dic.Guardar("gato", 7)
	dic.Guardar("yegua", 7)
	dic.Guardar("sapo", 3)
	dic.Guardar("foca", 8)
	dic.Guardar("colibri", 15)

	multiplicador := 1
	ptrMult := &multiplicador

	dic.Iterar(func(clave string, dato int) bool {
		if *ptrMult > 100 {
			return false
		}
		*ptrMult *= dato
		return true
	})
	//se pasa de 100 xq esa es la ultima iteracion, despues corta correctamente
	require.EqualValues(t, 600, multiplicador)
}
func TestIteradorRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })
	dic.Guardar("perro", 2)
	dic.Guardar("ballena", 5)
	dic.Guardar("gato", 7)
	dic.Guardar("yegua", 7)
	dic.Guardar("sapo", 3)
	dic.Guardar("foca", 8)
	dic.Guardar("colibri", 15)
	clave1 := "delfin"
	clave2 := "vaca"

	iter := dic.IteradorRango(&clave1, &clave2)

	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "foca", iter.Siguiente())
	require.EqualValues(t, "gato", iter.Siguiente())
	require.EqualValues(t, "perro", iter.Siguiente())
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, "sapo", iter.Siguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorInternoRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(12, 12)
	dic.Guardar(17, 17)
	dic.Guardar(3, 3)
	dic.Guardar(7, 7)
	clave1 := 6
	clave2 := 15

	sumador := 0
	ptrSumador := &sumador
	dic.IterarRango(&clave1, &clave2, func(clave int, dato int) bool {
		*ptrSumador += dato
		return true
	})

	require.EqualValues(t, 44, sumador)
}

func TestInternoStringsRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int { return strings.Compare(a, b) })

	dic.Guardar("perro", 2)
	dic.Guardar("ballena", 5)
	dic.Guardar("gato", 7)
	dic.Guardar("yegua", 7)
	dic.Guardar("sapo", 3)
	dic.Guardar("foca", 8)
	dic.Guardar("colibri", 15)
	clave1 := "colibri"
	clave2 := "vaca"

	multiplicador := 1
	ptrMult := &multiplicador

	dic.IterarRango(&clave1, &clave2, func(clave string, dato int) bool {
		if *ptrMult > 100 {
			return false
		}
		*ptrMult *= dato
		return true
	})
	require.EqualValues(t, 120, multiplicador)
}

func TestInternoRangoInt(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	dic.Guardar(5, 5)
	dic.Guardar(3, 3)
	dic.Guardar(7, 7)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(6, 6)
	dic.Guardar(8, 8)
	clave1 := 5
	clave2 := 99

	contador := 0
	multiplicador := 1
	ptrMult := &multiplicador
	ptrCont := &contador
	dic.IterarRango(&clave1, &clave2, func(clave int, valor int) bool {
		if *ptrCont >= 3 {
			return false
		}
		*ptrCont++
		*ptrMult *= clave
		return true

	})
	require.EqualValues(t, 210, multiplicador)
	require.EqualValues(t, 3, contador)
}

func TestExternoRangoVariosElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(77, 7)
	dic.Guardar(65, 65)
	dic.Guardar(9, 6)
	dic.Guardar(7, 3)
	dic.Guardar(2, 2)
	dic.Guardar(3, 77)
	dic.Guardar(9, 9)
	dic.Guardar(23, 1)
	require.False(t, dic.Pertenece(66))
	dic.Guardar(44, 12)
	require.EqualValues(t, 77, dic.Obtener(3))
	require.EqualValues(t, 9, dic.Borrar(9))
	dic.Guardar(7, 4)
	clave1 := 5
	clave2 := 50

	iter := dic.IteradorRango(&clave1, &clave2)
	c1, v1 := iter.VerActual()
	require.EqualValues(t, 7, c1)
	require.EqualValues(t, 4, v1)
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 7, iter.Siguiente())
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 23, iter.Siguiente())
	c1, _ = iter.VerActual()
	require.EqualValues(t, 44, c1)
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 44, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	require.EqualValues(t, 4, dic.Borrar(7))
	require.EqualValues(t, 12, dic.Borrar(44))
	require.EqualValues(t, 1, dic.Borrar(23))
	require.EqualValues(t, 2, dic.Borrar(2))
	require.EqualValues(t, 65, dic.Borrar(65))
	require.EqualValues(t, 7, dic.Borrar(77))
	require.EqualValues(t, 77, dic.Borrar(3))
	//voy a ver q vacie todo

	iter2 := dic.Iterador()
	require.False(t, iter2.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	claves := make([]int, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el hash
	for i := 0; i < n; i++ {
		claves[i] = rand.Intn(500000000000)
		valores[i] = claves[i]
		dic.Guardar(claves[i], valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave int
	var valor int

	for i = 0; i < dic.Cantidad(); i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if !dic.Pertenece(clave) {
			ok = false
			break
		}
		valor = v1
		if dic.Obtener(clave) != valor {
			ok = false
			break
		}
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}
