package diccionario

import (
	TDALista "ABB/Lista-enlazada"
	TDAPila "ABB/pila"
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
	if abb.Cantidad() == 0 {
		return false
	}
	nodo, _ := (abb.buscar(clave, abb.raiz))
	return abb.cmp(nodo.clave, clave) == 0
}

func (abb abb[K, V]) Obtener(clave K) V {
	if abb.Cantidad() == 0 {
		panic("La clave no pertenece al diccionario")
	}

	nodo, _ := (abb.buscar(clave, abb.raiz))
	if abb.cmp(nodo.clave, clave) != 0 {
		panic("La clave no pertenece al diccionario")
	}
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

func (abb abb[K, V]) buscarMayor(nodo *nodoAbb[K, V], clave *K) *nodoAbb[K, V] {
	if nodo.derecho == nil {
		return nodo
	} else if clave != nil {
		if abb.cmp(nodo.derecho.clave, *clave) > 0 {
			return nodo
		}
	}
	return abb.buscarMayor(nodo.derecho, clave)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.Cantidad() == 0 {
		panic("La clave no pertenece al diccionario")
	}

	nodo, padre := abb.buscar(clave, abb.raiz)
	if abb.cmp(nodo.clave, clave) != 0 {
		panic("La clave no pertenece al diccionario")
	}

	clave_retornar := nodo.dato
	if nodo.izquierdo != nil && nodo.derecho != nil {
		reemplazo := abb.buscarMayor(nodo.izquierdo, nil)
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

	} else { //un hijo, abuelo apunta al nieto
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
		}
	}
	abb.cantidad--
	return clave_retornar
}

func (abb abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.raiz.iterar(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(K, V) bool) {
	if nodo == nil {
		return
	}
	nodo.izquierdo.iterar(visitar)
	visitar(nodo.clave, nodo.dato)
	nodo.derecho.iterar(visitar)
}

func (iter *iterAbb[K, V]) apilarHijos(nodo *nodoAbb[K, V], desde *K) {
	if nodo == nil {
		return
	}

	iter.pila.Apilar(nodo)
	if iter.desde == nil {
		iter.apilarHijos(nodo.izquierdo, desde)
	} else {
		if iter.cmp(nodo.clave, *desde) < 0 {
			iter.Siguiente()
		} else {
			iter.apilarHijos(nodo.izquierdo, desde)
		}
	}

}

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := abb.IteradorRango(nil, nil)

	return iter
}

func (iter iterAbb[K, V]) HaySiguiente() bool {
	//Verifico si la pila est?? vac??a o el tope es mayor a hasta
	if iter.hasta != nil {
		if !iter.pila.EstaVacia() {
			return iter.cmp(iter.pila.VerTope().clave, *iter.hasta) <= 0
		}
	}
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.pila.Desapilar()
	iter.apilarHijos(elemento.derecho, iter.desde)
	return elemento.clave
}

func (iter iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iter.pila.VerTope()
	return elemento.clave, elemento.dato
}

func (abb abb[K, V]) elementoMedio(raiz *nodoAbb[K, V], desde *K) *nodoAbb[K, V] {
	if raiz == nil {
		return nil
	}
	if desde != nil {
		if abb.cmp(raiz.clave, *desde) < 0 {
			return abb.elementoMedio(raiz.derecho, desde)
		}
	}
	return raiz
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde, iter.hasta, iter.cmp = desde, hasta, abb.cmp

	if abb.raiz == nil {
		return iter
	}
	nodo := abb.elementoMedio(abb.raiz, desde)
	if nodo != nil {
		iter.apilarHijos(nodo, iter.desde)
	}
	return iter
}

func (abb abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(K, V) bool, cmp func(K, K) int) {

	if nodo == nil {
		return
	}

	if desde == nil {
		nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp)
	} else if cmp(nodo.clave, *desde) >= 0 {
		nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp)
	}

	if desde == nil && hasta == nil {
		visitar(nodo.clave, nodo.dato)
	} else if desde == nil {
		if cmp(nodo.clave, *hasta) <= 0 {
			visitar(nodo.clave, nodo.dato)
		}
	} else if hasta == nil {
		if cmp(nodo.clave, *desde) >= 0 {
			visitar(nodo.clave, nodo.dato)
		}
	} else if cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0 {
		visitar(nodo.clave, nodo.dato)
	}

	if hasta == nil {
		nodo.derecho.iterarRango(desde, hasta, visitar, cmp)
	} else if cmp(nodo.clave, *hasta) <= 0 {
		nodo.derecho.iterarRango(desde, hasta, visitar, cmp)
	}
}

func (ab abb[K, V]) EsCompleto() bool {
	return ab.raiz.EsCompletoAux()
}

func (ab *nodoAbb[K, V]) EsCompletoAux() bool {
	if ab == nil {
		return true
	}
	ok := false
	izq := ab.izquierdo.EsCompletoAux()
	der := ab.derecho.EsCompletoAux()
	if ab.izquierdo != nil && ab.derecho != nil {
		if ab.izquierdo.Altura() == ab.derecho.Altura() {
			ok = true
		}
	} else if ab.derecho == nil && ab.izquierdo == nil {
		ok = true
	}
	return ok && izq && der
}

func (ab *nodoAbb[K, V]) Altura() int {
	if ab == nil {
		return 0
	}
	return max(ab.izquierdo.Altura(), ab.derecho.Altura()) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (ab *abb[K, V]) Mayores(cadena K) TDALista.Lista[K] {
	lista := TDALista.CrearListaEnlazada[K]()
	ab.raiz.MayoresAUX(cadena, ab.cmp, lista)
	return lista
}

func (ab *nodoAbb[K, V]) MayoresAUX(cadena K, cmp func(K, K) int, lista TDALista.Lista[K]) {
	if ab == nil {
		return
	}
	if cmp(ab.clave, cadena) > 0 {
		ab.izquierdo.MayoresAUX(cadena, cmp, lista)
		lista.InsertarUltimo(ab.clave)
	}
	ab.derecho.MayoresAUX(cadena, cmp, lista)
}

func (ab *abb[K, int]) suma() int {
	return ab.raiz.sumaAux()
}

func (ab *nodoAbb[K, int]) sumaAux() int {
	if ab == nil {
		return 0
	}
	valor1 := ab.izquierdo.sumaAux()
	valor := ab.dato
	valor2 := ab.derecho.sumaAux()

	return valor1 + valor2 + valor
}

func (ab *abb[K, V]) dosHijos() int {
	return ab.raiz.dosHijosAux()
}

func (ab *nodoAbb[K, V]) dosHijosAux() int {
	if ab == nil {
		return 0
	}
	valor := 0
	if ab.izquierdo != nil && ab.derecho != nil {
		valor = 1
	}
	izq := ab.izquierdo.dosHijosAux()
	der := ab.derecho.dosHijosAux()
	return izq + der + valor
}

func (arbol *abb[K, V]) Invertir() {
	arbol.raiz.InvertirAUX()
}

func (arbol *nodoAbb[K, V]) InvertirAUX() {
	if arbol == nil {
		return
	}
	arbol.izquierdo.InvertirAUX()
	arbol.derecho.InvertirAUX()
	*arbol.izquierdo, *arbol.derecho = *arbol.derecho, *arbol.izquierdo
}

func (ab abb[K, V]) Quiebres() int {
	return ab.raiz.QuiebresAux()
}

func (ab *nodoAbb[K, V]) QuiebresAux() int {
	if ab == nil {
		return 0
	}
	suma := 0
	if ab.izquierdo != nil && ab.derecho == nil {
		if ab.izquierdo.izquierdo == nil && ab.izquierdo.derecho != nil {
			suma++
		}
	} else if ab.derecho != nil && ab.izquierdo == nil {
		if ab.derecho.izquierdo != nil && ab.derecho.derecho == nil {
			suma++
		}
	}
	izq := ab.izquierdo.QuiebresAux()
	der := ab.derecho.QuiebresAux()

	return suma + izq + der
}

func (abb *abb[K, V]) recorrido() TDALista.Lista[K] {
	lista := TDALista.CrearListaEnlazada[K]()
	abb.raiz.recorridoAUX(lista)
	return lista
}

func (abb *nodoAbb[K, V]) recorridoAUX(lista TDALista.Lista[K]) {
	if abb == nil {
		return
	}

	lista.InsertarUltimo(abb.clave)
	abb.izquierdo.recorridoAUX(lista)
	abb.derecho.recorridoAUX(lista)
}

/*type ab struct {
	dato string
	izq  *ab
	der  *ab
}

func (ab *ab) armarArbol(pre, in []string) *ab {
	ab.recursividad(pre, in, ab)
	return ab
}

func (ab *ab) recursividad(pre, in []string, nodo *ab) {
	if len(pre)==0{
		return
	} else if len(pre)==1{
		ab.dato = pre[0]
		return
	}

	valor := pre[1]
	if valor


}*/
