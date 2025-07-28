package studentmem

import (
	"context"
	"sync"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
)

type Repository struct {
	// You can modify the map by utilizing sync.Map instead of supplying it with a lock.
	// The inclusion of a lock in this instance is solely intended for demonstration purposes.
	students map[uint64]model.Student
	lock     sync.RWMutex
}

func New() *Repository {
	return &Repository{
		students: make(map[uint64]model.Student),
		lock:     sync.RWMutex{},
	}
}

func (r *Repository) Add(_ context.Context, model model.Student) error {
	r.lock.RLock()

	if _, ok := r.students[model.ID]; ok {
		return studentrepo.ErrStudentIDDuplicate
	}

	r.lock.RUnlock()

	r.lock.Lock()
	r.students[model.ID] = model
	r.lock.Unlock()

	return nil
}

// nolint: cyclop
func (r *Repository) Get(_ context.Context, cmd studentrepo.GetCommand) []model.Student {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var students []model.Student

	if cmd.ID != nil {
		student, ok := r.students[*cmd.ID]
		if !ok {
			return nil
		}

		students = []model.Student{student}
	} else {
		for _, student := range r.students {
			students = append(students, student)
		}
	}

	for i := 0; i < len(students); i++ {
		if cmd.FirstName != nil {
			if students[i].FirstName != *cmd.FirstName {
				students = append(students[:i], students[i+1:]...)
				i--

				continue
			}
		}

		if cmd.LastName != nil {
			if students[i].LastName != *cmd.LastName {
				students = append(students[:i], students[i+1:]...)
				i--

				continue
			}
		}

		if cmd.EntranceYear != nil {
			if students[i].EntranceYear != *cmd.EntranceYear {
				students = append(students[:i], students[i+1:]...)
				i--

				continue
			}
		}
	}

	return students
}
