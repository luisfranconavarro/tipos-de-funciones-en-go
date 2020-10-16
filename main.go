package main

import "fmt"

func fibonacci(x int) int{
	if x == 0{
		return 0
	}
	if x == 1{
		return 1
	}else{
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func masGrande(datos ...int) int{
	var aux int
	for i := range datos {
		if i+1 < len(datos){
			if datos[i] > datos[i+1]{
				aux = datos[i]
			}
		}
	}
	return aux
}

func generadorImpares() func() uint {
	i := uint(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func intercambia(a *int, b*int){
	aux := *a
	*a = *b
	*b = aux
}

func main()  {
	impares := generadorImpares()
	a, b := -5, 8

	fmt.Println(masGrande(5,40,1,3,9,-15,19,22,6,7))
	
	fmt.Println(impares())
	fmt.Println(impares())
	fmt.Println(impares())
	fmt.Println(impares())
	
	fmt.Println(a, b)
	intercambia(&a, &b)
	fmt.Println(a, b)

	fmt.Println(fibonacci(10))

}