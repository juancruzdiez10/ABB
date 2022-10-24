package diccionario

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
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
	if abb.cmp(raiz.clave, nodo.clave) == 0 {
		return nil
	}
	var busc *nodoAbb[K, V]
	if abb.cmp(raiz.clave, nodo.clave) > 0 {
		busc = abb.buscarPadre(nodo, raiz.izquierdo)
		if busc == nil {
			return raiz
		}
	} else {
		busc = abb.buscarPadre(nodo, raiz.derecho)
		if busc == nil {
			return raiz
		}
	}
	return busc
}
func (abb abb[K, V]) buscarMayor(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.derecho == nil {
		return nodo
	}
	return abb.buscarMayor(nodo.derecho)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo := abb.buscar(clave, abb.raiz)
	clave_retornar := nodo.dato

	if nodo.izquierdo == nil && nodo.derecho == nil {
		padre := abb.buscarPadre(nodo, abb.raiz)
		if padre == nil {
			abb.raiz = nil
		} else {
			if abb.cmp(padre.clave, nodo.clave) > 0 {
				padre.izquierdo = nil
			} else {
				padre.derecho = nil
			}
		}

		abb.cantidad--
		return clave_retornar
	}

	if nodo.izquierdo == nil || nodo.derecho == nil { //un hijo, abuelo apunta al nieto
		padre := abb.buscarPadre(nodo, abb.raiz)
		var enlace **nodoAbb[K, V]

		if padre == nil {
			if nodo.izquierdo != nil {
				abb.raiz = nodo.izquierdo
			} else {
				abb.raiz = nodo.derecho
			}
		} else {
			if abb.cmp(padre.clave, nodo.clave) > 0 {
				enlace = &padre.izquierdo
			} else {
				enlace = &padre.derecho
			}

			if nodo.izquierdo != nil {
				*enlace = nodo.izquierdo
			} else {
				*enlace = nodo.derecho
			}
		}
		abb.cantidad--
	}
	if nodo.izquierdo != nil && nodo.derecho != nil {
		reemplazo := abb.buscarMayor(nodo.izquierdo)
		reemplazo_dato, reemplazo_clave := reemplazo.dato, reemplazo.clave

		abb.Borrar(reemplazo.clave)

		nodo.clave, nodo.dato = reemplazo_clave, reemplazo_dato
	}

	return clave_retornar
}
