package student

import (
	"errors"
)

type Student struct {
	Name  string
	Age   int
	Grade string
	ID    int
}

type Manager struct {
	students []Student
}

func NewManager() *Manager {
	return &Manager{students: []Student{}}
}

func (m *Manager) AddStudent(s Student) {
	m.students = append(m.students, s)
}

func (m *Manager) RemoveStudentByID(id int) error {
	for i, s := range m.students {
		if s.ID == id {
			m.students = append(m.students[:i], m.students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}

func (m *Manager) GetStudentByID(id int) (Student, error) {
	for _, s := range m.students {
		if s.ID == id {
			return s, nil
		}
	}
	return Student{}, errors.New("student not found")
}

func (m *Manager) UpdateStudentGrade(id int, newGrade string) error {
	for i := range m.students {
		if m.students[i].ID == id {
			m.students[i].Grade = newGrade
			return nil
		}
	}
	return errors.New("student not found")
}

func (m *Manager) UpdateStudentAge(id int, newAge int) error {
	for _, s := range m.students {
		if s.ID == id {
			s.Age = newAge
			return nil
		}
	}
	return errors.New("student not found")
}

func (m *Manager) ListStudents() []Student {
	return m.students
}
