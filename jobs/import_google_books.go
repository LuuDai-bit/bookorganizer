package jobs

import (
	"book-organizer/models"
	"book-organizer/services"

	"github.com/go-co-op/gocron/v2"
)

type ImportGoogleBook struct{}

func (i *ImportGoogleBook) Perform(s gocron.Scheduler) {
	_, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(23, 30, 0),
			),
		),
		gocron.NewTask(i.importGoogleBook),
	)

	if err != nil {
		// Write to log
		return
	}
}

func (i *ImportGoogleBook) importGoogleBook() {
	fetchNewBook := new(services.FetchNewBook)
	userModel := new(models.UserModel)
	categories, err := userModel.GetFavoriteCategories()
	if err != nil {
		// Write to log
		return
	}

	for _, category := range categories {
		books, err := fetchNewBook.FetchGoogleBooks(category)
		if err != nil {
			// Write to log
			return
		}

		formatedBooks := i.flattenGoogleBooks(books)
		categoryBook := new(models.CategoryBook)
		err = categoryBook.Create(category, formatedBooks)
		if err != nil {
			// Write to log
			return
		}
	}
}

func (i *ImportGoogleBook) flattenGoogleBooks(books []services.GoogleBook) []models.GoogleBook {
	var formatedBooks []models.GoogleBook
	for _, book := range books {
		var formatedBook models.GoogleBook
		formatedBook.ID = book.ID
		formatedBook.Kind = book.Kind
		formatedBook.Title = book.VolumeInfo.Title
		formatedBook.SubTitle = book.VolumeInfo.SubTitle
		formatedBook.Authors = book.VolumeInfo.Authors
		formatedBook.Description = book.VolumeInfo.Description
		formatedBook.Categories = book.VolumeInfo.Categories

		formatedBooks = append(formatedBooks, formatedBook)
	}

	return formatedBooks
}
