package main

import "fmt"

// Queue turini e'lon qilamiz

type Queue struct {
	data []int
}

// NewQueue funksiyasi yangi Queue obyektini yaratadi va uni qaytaradi

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty funksiyasi Queue bo'sh ekanligini tekshiradi

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek funksiyasi Queue ning boshidagi elementni qaytaradi, agar Queue bo'sh bo'lsa, 0 ni qaytaradi

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[0], nil
}

// Enqueue funksiyasi Queue ga yangi element qo'shadi

func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

// Dequeue funksiyasi Queue dan boshidagi elementni olib tashlaydi va uni qaytaradi, agar Queue bo'sh bo'lsa, 0 ni qaytaradi

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

// Size funksiyasi Queue ning o'lchamini qaytaradi

func (q *Queue) Size() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return len(q.data), nil
}

func testQueue() {
	q := NewQueue()

	// Queue bo'sh ekanligini tekshirish

	fmt.Println("Is queue empty?", q.IsEmpty()) // true

	// Element qo'shish

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Queue bo'sh emasligini tekshirish

	fmt.Println("Is queue empty?", q.IsEmpty()) // false

	// Boshidagi elementni ko'rish

	if value, err := q.Peek(); err == nil {
		fmt.Println("Peek:", value) // 10
	} else {
		fmt.Println(err)
	}

	// Elementlarni olib tashlash

	if value, err := q.Dequeue(); err == nil {
		fmt.Println("Dequeue:", value) // 10
	} else {
		fmt.Println(err)
	}

	if value, err := q.Dequeue(); err == nil {
		fmt.Println("Dequeue:", value) // 20
	} else {
		fmt.Println(err)
	}

	// Queue o'lchamini tekshirish

	if size, err := q.Size(); err == nil {
		fmt.Println("Queue size:", size) // 1
	} else {
		fmt.Println(err)
	}

	// Boshidagi elementni ko'rish

	if value, err := q.Peek(); err == nil {
		fmt.Println("Peek:", value) // 30
	} else {
		fmt.Println(err)
	}

	// Oxirgi elementni olib tashlash

	if value, err := q.Dequeue(); err == nil {
		fmt.Println("Dequeue:", value) // 30
	} else {
		fmt.Println(err)
	}

	// Queue bo'sh ekanligini tekshirish

	fmt.Println("Is queue empty?", q.IsEmpty()) // true

	// Bo'sh Queue dan element olib tashlashga harakat qilish
	
	if value, err := q.Dequeue(); err == nil {
		fmt.Println("Dequeue:", value)
	} else {
		fmt.Println(err) // queue is empty
	}
}

func main() {
	testQueue()
}
