package io

import (
	"bufio"
	"github.com/gammazero/deque"
	"os"
	"strings"
)

// Reader : Функция чтения команд из консоли и добавления их в очередь
func Reader() (*deque.Deque[[]string], error) {
	deq := newDeque[[]string]()
	var split []string
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		return deq, scanner.Err()
	}
	line := scanner.Text()
	if strings.Contains(line, "|") {
		split = strings.Split(line, "|")
		for _, s := range split {
			args := strings.Fields(strings.TrimSpace(s))
			deq.PushBack(args)
		}
	} else {
		args := strings.Fields(strings.TrimSpace(line))
		deq.PushBack(args)
	}

	return deq, scanner.Err()
}

func newDeque[T any]() *deque.Deque[T] {
	var d deque.Deque[T]
	return &d
}
