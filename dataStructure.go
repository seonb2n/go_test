package main

import (
	"container/list"
	"fmt"
)

// 리스트 형태
func main() {
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
