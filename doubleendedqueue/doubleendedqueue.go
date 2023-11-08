package doubleendedqueue

type DoubleEndedQueue struct {
	elements []string
}

func NewDoubleEndedQueue() *DoubleEndedQueue {
	return &DoubleEndedQueue{elements: []string{}}
}

func (deque *DoubleEndedQueue) PushFront(value string) {
	deque.elements = append([]string{value}, deque.elements...)
}

func (deque *DoubleEndedQueue) PushBack(value string) {
	deque.elements = append(deque.elements, value)
}

func (deque *DoubleEndedQueue) PopFront() string {
	if len(deque.elements) == 0 {
		return "" // Valor padrão para fila vazia
	}
	value := deque.elements[0]
	deque.elements = deque.elements[1:]
	return value
}

func (deque *DoubleEndedQueue) PopBack() string {
	if len(deque.elements) == 0 {
		return "" // Valor padrão para fila vazia
	}
	value := deque.elements[len(deque.elements)-1]
	deque.elements = deque.elements[:len(deque.elements)-1]
	return value
}

func (deque *DoubleEndedQueue) IsEmpty() bool {
	return len(deque.elements) == 0
}
