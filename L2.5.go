package main

// В результате работы программы в консоль выведется "error".
// Что происходит в программе: создается кастомный тип customError
// для него определяется метод Error, после чего данный тип
//начинает имплементировать встроенный интерфейс error.
// Далее функция test возвращает значение типа *customError, но равное nil.
// В main мы присваиваем это значение переменной err типа error.
// В интерфейсе при этом сохраняется пара: (type = *main.customError, value = nil).
// При сравнении с nil учитываются оба поля (и type, и value).
// Так как type отличается от nil, результат сравнения err != nil будет true,
// из-за чего в консоль и будет выведено "error".

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
