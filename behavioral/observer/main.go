package main

func main() {
	subject := Subject{}
	observer1 := ConcreteObserver{id: "Observer 1"}
	observer2 := ConcreteObserver{id: "Observer 2"}

	subject.Attach(&observer1)
	subject.Attach(&observer2)

	subject.SetMessage("Hello World!")

	subject.Detach(&observer1)

	subject.SetMessage("Hello World Again!")
}
