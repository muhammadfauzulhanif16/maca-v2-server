package formatters

import (
	"maca/entities"
	"time"
)

type bookFormat struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Published   string    `json:"published"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func BookFormat(data entities.Book) bookFormat {
	return bookFormat{
		Id:          data.Id,
		Title:       data.Title,
		Author:      data.Author,
		Published:   data.Published,
		IsCompleted: data.IsCompleted,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func BooksFormat(datas []entities.Book) []bookFormat {
	datasFormat := []bookFormat{}

	for _, data := range datas {
		dataFormat := BookFormat(data)
		datasFormat = append(datasFormat, dataFormat)
	}

	return datasFormat
}
