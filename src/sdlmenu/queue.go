package main 

import (
	"errors"
	"fmt"
)


type Queue struct {
	Elements []string
	Size int
}

func (q *Queue) Enqueue(elem string) {
	if (q.GetLength() == q.Size) {
		fmt.Println("Overflow at given value: ", elem)
		return;
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		fmt.Println("Underflow")
		return "";
	}

	element := q.Elements[0]

	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}

	q.Elements = q.Elements[1:]
	return element

}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (string, error){
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}
	return q.Elements[0], nil;
}


