package main

import (
	"fmt"
	"math/rand"
	"time"
)

// С помощью функции asChan в программе создаются два канала: a и b.
// В них горутины с задержкой в 1-1000 мс передают значения и закрывают канал
// В функции merge эти значения передаются в третий канал.
// Когда один из каналов передает все свои значения в канал c, то переменная,
// ссылающаяся на него устанавливается в nil.
// Когда оба канала (a и b) закрыты, канал c закрывается.
// Далее в main мы читаем из канала c значения и выводим их в консоль в одну строчку

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v)
	}
}
