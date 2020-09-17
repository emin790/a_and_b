package main

import "fmt"

func swap(a,b *int){					//функция принимает указатели a и b
	*a = *b
	*b = *a
}

func main() {
	
	a := 111
	b := 222
	
	fmt.Println("До замены:")
	fmt.Println("a =", a, "b =", b)
	
	swap(&a,&b) 						//передаем в функцию адреса a и b
	
	fmt.Println("После замены:")
	fmt.Println("a =", a, "b =", b)

}
