package main

import (
	"fmt"
	"os"
)

// Программа выведет false,
// тк оператор "==", при сравнении интерфейсов, сравнивает не только его значение, но и тип
// (структура интерфейса: type, value. Для примера, var a interface{}; fmt.Println(reflect.TypeOf(a));
// выведет <nil>, а сравнение a == nil, в данном случае, выведет в консоль true, тк, и тип и значение
// пустого интерфейса равны nil) и в нашем случае тип переменной err равен *fs.PathError,
// что не совпадает с nil-значением типа nil.
// Но если нам требуется узнать является ли значение переменной nil'ом, то в данном случае подойдет запись ->
// -> reflect.ValueOf(err).IsNil()

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
