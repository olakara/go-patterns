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

func (s *Subject) Detach(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
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
