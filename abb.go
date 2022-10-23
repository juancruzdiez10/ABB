package diccionario

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V] {
	arbol := new(abb[K, V])
	arbol.cmp = funcion_cmp
	return arbol
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = valor
	return nodo

}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	nuevoNodo := crearNodo[K, V](clave, valor)
	if abb.cantidad == 0 {
		abb.raiz = nuevoNodo
	} else {
		nodo := abb.buscar(clave, abb.raiz)
		if abb.cmp(nodo.clave, clave) == 0 {
			nodo = nuevoNodo
			return
		}
		if abb.cmp(nodo.clave, clave) > 0 {
			nodo.izquierdo = nuevoNodo
		}
		if abb.cmp(nodo.clave, clave) < 0 {
			nodo.derecho = nuevoNodo
		}
	}
	abb.cantidad++
}

func (abb abb[K, V]) Pertenece(clave K) bool {
	if abb.cantidad == 0 {
		return false
	}
	return abb.cmp((abb.buscar(clave, abb.raiz)).clave, clave) == 0
}

func (abb abb[K, V]) Obtener(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	return (abb.buscar(clave, abb.raiz)).dato
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb abb[K, V]) buscar(clave K, raiz *nodoAbb[K, V]) *nodoAbb[K, V] {
	if abb.cmp(raiz.clave, clave) == 0 {
		return raiz
	}
	if abb.cmp(raiz.clave, clave) > 0 {
		if raiz.izquierdo == nil {
			return raiz
		}
		return abb.buscar(clave, raiz.izquierdo)
	}
	if raiz.derecho == nil {
		return raiz
	}
	return abb.buscar(clave, raiz.derecho)
}

func (abb abb[K, V]) buscarPadre(nodo *nodoAbb[K, V], raiz *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.clave == raiz.clave {
		return nil
	}
	if abb.cmp(raiz.clave, nodo.clave) > 0 {
		if raiz.izquierdo == nil {

		}
	}
	if raiz.izquierdo.clave == nodo.clave || raiz.derecho.clave == nodo.clave {
		return raiz
	}

}

func (abb *abb[K, V]) Borrar(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo := abb.buscar(clave, abb.raiz)
	//no tiene hijos, padre apunta a nil
	if nodo.izquierdo == nil && nodo.derecho == nil {
		nodo = nil
	}
	//un hijo, abuelo apunta al nieto

	//dos hijos, se reemplaza por el mas chico de la derecha o el mayor de la izquierda

}
