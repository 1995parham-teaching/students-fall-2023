package studentmem_test

import (
	"context"
	"testing"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/repository/studentmem"
	"github.com/stretchr/testify/suite"
)

type StudentsInMemorySuite struct {
	suite.Suite

	repo *studentmem.Repository
}

func (suite *StudentsInMemorySuite) SetupTest() {
	suite.repo = studentmem.New()

	var _ studentrepo.Repository = suite.repo
}

func (suite *StudentsInMemorySuite) TestAdd() {
	require := suite.Require()

	require.NoError(suite.repo.Add(context.Background(), model.Student{
		ID:           9231058,
		FirstName:    "Parham",
		LastName:     "Alvani",
		Courses:      nil,
		EntranceYear: 2013,
	}))
}

func (suite *StudentsInMemorySuite) TestGet() {
	require := suite.Require()

	require.NoError(suite.repo.Add(context.Background(), model.Student{
		ID:           9231058,
		FirstName:    "Parham",
		LastName:     "Alvani",
		Courses:      nil,
		EntranceYear: 2013,
	}))

	require.NoError(suite.repo.Add(context.Background(), model.Student{
		ID:           9631025,
		FirstName:    "Elahe",
		LastName:     "Dastan",
		Courses:      nil,
		EntranceYear: 2017,
	}))

	suite.Run("find all students", func() {
		st := suite.repo.Get(context.Background(), studentrepo.GetCommand{
			ID:           nil,
			FirstName:    nil,
			LastName:     nil,
			EntranceYear: nil,
		})

		require.Len(st, 2)
	})
}

func TestStudentsInMemory(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(StudentsInMemorySuite))
}
