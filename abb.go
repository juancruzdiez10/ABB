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
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	cmp   func(K, K) int
	desde *K
	hasta *K
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

func (iter *iterAbb[K, V]) apilarHijos(nodo *nodoAbb[K, V], desde *K) {
	if nodo == nil {
		return
	}
	if iter.desde == nil {
		iter.pila.Apilar(nodo)
		iter.apilarHijos(nodo.izquierdo, desde)
	} else {
		if iter.cmp(nodo.clave, *desde) > 0 {
			iter.pila.Apilar(nodo)
			iter.apilarHijos(nodo.izquierdo, desde)
		}
	}

}

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	//con recorrido preorder
	iter := new(iterAbb[K, V])
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.pila = pila
	iter.cmp = abb.cmp

	if abb.raiz != nil {
		iter.pila.Apilar(abb.raiz)
		iter.apilarHijos(abb.raiz.izquierdo, iter.desde)
	}

	return iter
}

func (iter iterAbb[K, V]) HaySiguiente() bool {
	//Verifico si la pila está vacía o el tope es mayor a hasta
	if iter.hasta == nil {
		return !iter.pila.EstaVacia()
	} else {
		return !iter.pila.EstaVacia() && iter.cmp(iter.pila.VerTope().clave, *iter.hasta) < 0
	}
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.pila.Desapilar()
	if elemento.derecho != nil {
		iter.apilarHijos(elemento.derecho, iter.desde)
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

func (abb abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.raiz.iterar(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(K, V) bool) {
	//recorrido in orden
	if nodo == nil {
		return
	}

	nodo.izquierdo.iterar(visitar)
	if visitar(nodo.clave, nodo.dato) {
		nodo.derecho.iterar(visitar)
	}
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.pila = pila
	iter.desde = desde
	iter.hasta = hasta
	if abb.raiz != nil {
		iter.apilarHijos(abb.raiz, iter.desde)
	}
	return iter
}
