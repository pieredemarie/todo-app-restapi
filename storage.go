package main 

import (
	"errors"
	"sync"
)

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Storage interface {
	Insert(t *Todo)
	Get(id int) (Todo, error)
	Update(id int, t Todo)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Todo
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Todo),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(t *Todo) {
	s.Lock()

	t.ID = s.counter
	s.data[t.ID] = *t

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Todo, error) {
	s.Lock()
	defer s.Unlock()

	todo, alright := s.data[id]
	if !alright {
		return todo,errors.New("Todo Task not found")
	}

	return todo, nil
}

func (s *MemoryStorage) Update(id int, t Todo) {
	s.Lock()
	s.data[id] = t 
	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data,id)
	s.Unlock()
}