package main

import (
	"container/list"
	"fmt"
)

// 리스트 형태
func main_list() {
	// 새로운 리스트 객체를 생성
	l := list.New()
	// 제일 마지막에 값을 추가
	element4 := l.PushBack(4)
	// 제일 처음에 값을 추가
	element2 := l.PushFront(1)
	// 지정한 위치 바로 이전에 값을 추가
	l.InsertBefore(3, element4)
	// 지정한 위치 바로 이후에 값을 추가
	l.InsertAfter(2, element2)

	fmt.Println("From Front")

	// 앞에서부터 값에 접근
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	//결과
	//From Front
	//1 2 3 4

	fmt.Println()
	fmt.Println("From Back")

	// 뒤에서 부터 값에 접근
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	//결과
	//From Back
	//4 3 2 1

}

// 큐 자료 구조
// 리스트를 사용해서 구현한다.
type Queue struct {
	// lit 패키지의 List 타입의 포인터인 v 를 의미한다.
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return nil
	}

	return q.v.Remove(front)
}

func main_queue() {
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Println(v)
		v = queue.Pop()
	}
}

// 스택 자료구조
func main3() {
	stack := list.New()

	// 스택에 요소 추가
	// 뒤에 넣는다.
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	// 스택에서 요소 꺼내기
	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		fmt.Println(element.Value)
	}
}
