package services

import (
	"errors"
	"maca/entities"
	"maca/inputs"
	"maca/repositories"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type bookS struct {
	r repositories.BookR
}

func NewBookS(r repositories.BookR) *bookS {
	return &bookS{r}
}

type BookS interface {
	Create(input inputs.CreateBook) (entities.Book, error)
	ReadAll(IsCompleted string) ([]entities.Book, error)
	ReadById(id string) (entities.Book, error)
	ReadSearch(input inputs.SearchBook) ([]entities.Book, error)
	UpdateIsCompleted(id string) (entities.Book, error)
	Delete(id string) (entities.Book, error)
	DeleteAll() ([]entities.Book, error)
}

func (s *bookS) Create(input inputs.CreateBook) (entities.Book, error) {
	data, err := s.r.Create(
		entities.Book{
			Id:          uuid.NewString(),
			Title:       input.Title,
			Author:      input.Author,
			Published:   input.Published,
			IsCompleted: input.IsCompleted,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})

	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *bookS) ReadAll(IsCompleted string) ([]entities.Book, error) {
	if IsCompleted == "" {
		datas, err := s.r.ReadAll()
		if err != nil {
			return datas, err
		}

		return datas, nil
	}

	isCompleted, err := strconv.ParseBool(IsCompleted)
	datas, err := s.r.ReadByIsCompleted(isCompleted)
	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (s *bookS) ReadById(id string) (entities.Book, error) {
	data, err := s.r.ReadById(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *bookS) ReadSearch(input inputs.SearchBook) ([]entities.Book, error) {
	datas, err := s.r.ReadSearch(input.Search)
	if err != nil {
		return datas, err
	}

	if input.Search == "" {
		return datas, errors.New("No data")
	}

	return datas, nil
}

func (s *bookS) UpdateIsCompleted(id string) (entities.Book, error) {
	data, err := s.r.ReadById(id)
	if err != nil {
		return data, err
	}

	data.IsCompleted = !data.IsCompleted
	data.UpdatedAt = time.Now()

	dataUpdate, err := s.r.Update(data)
	if err != nil {
		return dataUpdate, err
	}

	return dataUpdate, nil
}

func (s *bookS) Delete(id string) (entities.Book, error) {
	data, err := s.r.Delete(id)
	if err != nil {
		return data, err
	}

	if id == "" {
		return data, errors.New("There's no book with this id")
	}

	return data, nil
}

func (s *bookS) DeleteAll() ([]entities.Book, error) {
	datas, err := s.r.DeleteAll()
	if err != nil {
		return datas, err
	}

	return datas, nil
}
