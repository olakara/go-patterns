package main

type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
	message   string
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) SetMessage(message string) {
	s.message = message
	s.Notify()
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.message)
	}
}
