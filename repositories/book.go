package repositories

import (
	"maca/entities"

	"gorm.io/gorm"
)

type bookR struct {
	db *gorm.DB
}

func NewBookR(db *gorm.DB) *bookR {
	return &bookR{db}
}

type BookR interface {
	Create(data entities.Book) (entities.Book, error)
	ReadAll() ([]entities.Book, error)
	ReadById(id string) (entities.Book, error)
	ReadByIsCompleted(isCompleted bool) ([]entities.Book, error)
	ReadSearch(input string) ([]entities.Book, error)
	Update(data entities.Book) (entities.Book, error)
	Delete(id string) (entities.Book, error)
	DeleteAll() ([]entities.Book, error)
}

func (r *bookR) Create(data entities.Book) (entities.Book, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *bookR) ReadAll() ([]entities.Book, error) {
	var data []entities.Book
	if err := r.db.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *bookR) ReadById(id string) (entities.Book, error) {
	var data entities.Book
	if err := r.db.Where("id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *bookR) ReadByIsCompleted(isCompleted bool) ([]entities.Book, error) {
	var datas []entities.Book
	if err := r.db.Where("is_completed = ?", isCompleted).Order("created_at desc").Find(&datas).Error; err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *bookR) ReadSearch(input string) ([]entities.Book, error) {
	var datas []entities.Book
	if err := r.db.Where("title = ? or author = ? or published = ?", input, input, input).Find(&datas).Error; err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *bookR) Update(data entities.Book) (entities.Book, error) {
	if err := r.db.Save(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *bookR) Delete(id string) (entities.Book, error) {
	var data entities.Book
	if err := r.db.Where("id = ?", id).Delete(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *bookR) DeleteAll() ([]entities.Book, error) {
	if err := r.db.Exec("DELETE FROM books").Error; err != nil {
		return []entities.Book{}, err
	}

	return []entities.Book{}, nil
}
