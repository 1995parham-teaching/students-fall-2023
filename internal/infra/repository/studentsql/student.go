package studentsql

import (
	"context"
	"errors"
	"time"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"gorm.io/gorm"
)

type StudentDTO struct {
	model.Student `gorm:"embedded"`

	ID        uint64         `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Add(ctx context.Context, model model.Student) error {
	// nolint: exhaustruct
	if err := r.db.WithContext(ctx).Create(StudentDTO{ID: model.ID, Student: model}).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return studentrepo.ErrStudentIDDuplicate
		}

		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, cmd studentrepo.GetCommand) []model.Student {
	var studentDTOs []StudentDTO

	var condition StudentDTO

	if cmd.ID != nil {
		condition.ID = *cmd.ID
	}

	if cmd.FirstName != nil {
		condition.FirstName = *cmd.FirstName
	}

	if cmd.LastName != nil {
		condition.LastName = *cmd.LastName
	}

	if cmd.EntranceYear != nil {
		condition.EntranceYear = *cmd.EntranceYear
	}

	if err := r.db.WithContext(ctx).Where(&condition).Find(&studentDTOs).Error; err != nil {
		return nil
	}

	students := make([]model.Student, len(studentDTOs))

	for index, dto := range studentDTOs {
		students[index] = dto.Student
		students[index].ID = dto.ID
	}

	return students
}
