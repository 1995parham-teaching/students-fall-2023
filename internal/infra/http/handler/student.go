package handler

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/1995parham-teaching/students-fall-2023/internal/domain/model"
	"github.com/1995parham-teaching/students-fall-2023/internal/domain/repository/studentrepo"
	"github.com/1995parham-teaching/students-fall-2023/internal/infra/http/request"
	"github.com/labstack/echo/v4"
)

type Student struct {
	repo studentrepo.Repository
}

func NewStudent(repo studentrepo.Repository) *Student {
	return &Student{
		repo: repo,
	}
}

func (s *Student) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	students := s.repo.Get(c.Request().Context(), studentrepo.GetCommand{
		ID:           &id,
		FirstName:    nil,
		LastName:     nil,
		EntranceYear: nil,
	})
	if len(students) == 0 {
		return echo.ErrNotFound
	}

	if len(students) > 1 {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, students[0])
}

func (s *Student) Get(c echo.Context) error {
	var idPtr *uint64
	idStr := c.QueryParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err == nil {
		idPtr = &id
	}

	var fnPtr *string
	fn := c.QueryParam("first_name")
	if fn != "" {
		fnPtr = &fn
	}

	students := s.repo.Get(c.Request().Context(), studentrepo.GetCommand{
		ID:           idPtr,
		FirstName:    fnPtr,
		LastName:     nil,
		EntranceYear: nil,
	})
	if len(students) == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, students)
}

func (s *Student) Create(c echo.Context) error {
	var req request.StudentCreate

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	// we have the filled request
	if err := req.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	id := rand.Uint64() % 1_000_000
	if err := s.repo.Add(c.Request().Context(), model.Student{
		ID:           id,
		FirstName:    req.Name,
		LastName:     req.Family,
		EntranceYear: 0,
		Courses:      []model.Course{},
	}); err != nil {
		if errors.Is(err, studentrepo.ErrStudentIDDuplicate) {
			return echo.ErrBadRequest
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (s *Student) Register(g *echo.Group) {
	g.GET("", s.Get)
	g.POST("", s.Create)
	g.GET(":id", s.GetByID)
}
