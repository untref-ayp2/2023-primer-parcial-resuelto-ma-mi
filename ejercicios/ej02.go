package main

// 2. Dado un diccionario que contiene las notas de los estudiantes, escriba una función
// que devuelva la nota final (promedio de notas de cada alumno).
// Ej {"Perez" : [4,4,8,6], "Sánchez": [7,5,7,7], "Flores": [4,9,8]}.
// Debe devolver {"Perez": 5.50, "Sánchez": 6.50, "Flores": 7.0}

func NotasFinales(notas map[string][]int) map[string]float64 {
	promedios := make(map[string]float64)
	for nombre, notas := range notas {
		promedioParaNombre := 0.0
		for _, nota := range notas {
			promedioParaNombre += float64(nota)
		}
		promedioParaNombre /= float64(len(notas))
		promedios[nombre] = promedioParaNombre
	}
	return promedios
}
