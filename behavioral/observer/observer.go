package main

type ConcreteObserver struct {
	id string
}

func (co *ConcreteObserver) Update(message string) {
	println(co.id, "received:", message)
}
