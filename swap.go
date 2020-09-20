package main

import "fmt"

func main() {
	
	a := 9
	b := 8
	fmt.Println("a =", a, "b =", b)
	
	a = a + b  // сохраняем сумму a и b в a
	b = a - b  // вычитаем из суммы b, это будет прежднее значение a. в a все еще хранится сумма
	a = a - b 
	
	fmt.Println("a =", a, "b =", b)
}
