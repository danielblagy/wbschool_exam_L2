Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Программа выведет:
2
1

В test возращаемый аргумент имеет имя x, следовательно defer изменит x когда функция будет возвращать
значение. В anotherTest, аргумент возращаемый неименованный, мы возвращаем x и потом он изменяется
в defer, то есть изменения не сохранятся в возвращаемом значении.

```
