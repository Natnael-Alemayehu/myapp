package book

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

func (f *Book) ToDto() *DTO {
	return &DTO{
		ID:            f.ID.String(),
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: f.PublishedDate.Format("2006-01-02"),
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}
}

type Form struct {
	Title         string `json:"title" validate:"required,max=256"`
	Author        string `json:"author" validate:"required,alphaspace,max=256"`
	PublishedDate string `json:"published_date" validate:"required,datetime=2006-01-02"`
	ImageURL      string `json:"image_url" validate:"url"`
	Description   string `json:"description"`
}

func (f *Form) ToModel() *Book {
	pubDate, _ := time.Parse("2006-01-02", f.PublishedDate)

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}
}

type Book struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book

func (bs Books) ToDto() []*DTO {
	dtos := make([]*DTO, len(bs))
	for i, v := range bs {
		dtos[i] = v.ToDto()
	}
	return dtos
}
