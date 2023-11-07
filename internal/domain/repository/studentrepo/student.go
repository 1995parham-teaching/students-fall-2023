package studentrepo

import (
	"context"
	"errors"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
)

var ErrStudentIDDuplicate = errors.New("student id already exists")

type GetCommand struct {
	ID           *uint64
	FirstName    *string
	LastName     *string
	EntranceYear *int
}

type Repository interface {
	Add(ctx context.Context, model model.Student) error
	Get(ctx context.Context, cmd GetCommand) []model.Student
}
