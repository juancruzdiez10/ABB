package diccionario

import (
	TDAPila "abb/pila"
)

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

type iterAbb[K comparable, V any] struct {
	actual *nodoAbb[K, V]
	pila   TDAPila.Pila[*nodoAbb[K, V]]
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
	nuevoNodo := crearNodo(clave, valor)
	if abb.cantidad == 0 {
		abb.raiz = nuevoNodo
	} else {
		nodo, _ := abb.buscar(clave, abb.raiz)
		if abb.cmp(nodo.clave, clave) == 0 {
			nodo.dato = nuevoNodo.dato
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
	nodo, _ := (abb.buscar(clave, abb.raiz))
	return abb.cmp(nodo.clave, clave) == 0
}

func (abb abb[K, V]) Obtener(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo, _ := (abb.buscar(clave, abb.raiz))
	return nodo.dato
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb abb[K, V]) buscar(clave K, raiz *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if abb.cmp(raiz.clave, clave) == 0 {
		return raiz, nil
	}
	var nodo, padre *nodoAbb[K, V]
	if abb.cmp(raiz.clave, clave) > 0 {
		if raiz.izquierdo == nil {
			return raiz, raiz
		}
		nodo, padre = abb.buscar(clave, raiz.izquierdo)
	} else {
		if raiz.derecho == nil {
			return raiz, raiz
		}
		nodo, padre = abb.buscar(clave, raiz.derecho)
	}
	if padre == nil {
		padre = raiz
	}
	return nodo, padre
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
	nodo, padre := abb.buscar(clave, abb.raiz)
	clave_retornar := nodo.dato
	if nodo.izquierdo != nil && nodo.derecho != nil {

		reemplazo := abb.buscarMayor(nodo.izquierdo)
		reemplazo_dato, reemplazo_clave := reemplazo.dato, reemplazo.clave
		abb.Borrar(reemplazo.clave)
		nodo.clave, nodo.dato = reemplazo_clave, reemplazo_dato
		return clave_retornar //return aca para no restar la cantidad todavia

	} else if padre == nil {

		if nodo.izquierdo != nil {
			abb.raiz = nodo.izquierdo
		} else if nodo.derecho != nil {
			abb.raiz = nodo.derecho
		} else {
			abb.raiz = nil
		}

	} else if nodo.izquierdo == nil && nodo.derecho == nil {

		if abb.cmp(padre.clave, nodo.clave) > 0 {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}

	} else if nodo.izquierdo == nil || nodo.derecho == nil { //un hijo, abuelo apunta al nieto

		var enlace **nodoAbb[K, V]
		if abb.cmp(padre.clave, nodo.clave) > 0 {
			enlace = &padre.izquierdo
		} else {
			enlace = &padre.derecho
		}
		if nodo.izquierdo != nil {
			*enlace = nodo.izquierdo
		} else {
			*enlace = nodo.derecho
		} //Creeria q este seria el unico para modificar, pero no encuentro manera mas corta que esta

	}
	abb.cantidad--
	return clave_retornar
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	//con recorrido preorder
	iter := new(iterAbb[K, V])
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.raiz != nil {
		pila.Apilar(abb.raiz)
		iter.actual = pila.VerTope()
	}
	iter.pila = pila
	return iter
}

func (iter iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.pila.Desapilar()
	if elemento.derecho != nil {
		iter.pila.Apilar(elemento.derecho)
	}
	if elemento.izquierdo != nil {
		iter.pila.Apilar(elemento.izquierdo)
	}
	return elemento.clave
}

func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.pila.VerTope()
	return elemento.clave, elemento.dato
}
