package services

import (
	"agmc-day-8/internal/dto"
	"agmc-day-8/internal/models"
)

func (s *services) CreateBook(book *models.Book) (*dto.Response, error) {
	bookMapping := &models.Book{
		Title:  book.Title,
		Writer: book.Writer,
	}

	err := s.repo.CreateBook(bookMapping)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    201,
		Status:  "success",
		Message: "Success to create book",
		Data:    err,
	}

	return result, err
}

func (s *services) UpdateBook(book *models.Book, id int) (*dto.Response, error) {
	bookMapping := &models.Book{
		Title:  book.Title,
		Writer: book.Writer,
	}

	err := s.repo.UpdateBook(bookMapping, id)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to update book",
		Data:    err,
	}

	return result, nil
}

func (s *services) DeleteBook(id int) (*dto.Response, error) {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to delete book",
		Data:    err,
	}

	return result, nil
}

func (s *services) GetBookById(id int) (*dto.Response, error) {
	book, err := s.repo.GetBookById(id)
	if err != nil {
		return nil, err
	}

	bookMapping := &dto.BookMapping{
		Title:     book.Title,
		Writer:    book.Writer,
		ID:        book.ID,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get book",
		Data:    bookMapping,
	}
	return result, nil
}

func (s *services) GetAllBooks(keywords string) (*dto.Response, error) {
	books, err := s.repo.GetAllBooks(keywords)
	if err != nil {
		return nil, err
	}

	var arrBooks []dto.BookMapping
	for _, book := range books {
		bookMapping := dto.BookMapping{
			Title:     book.Title,
			Writer:    book.Writer,
			ID:        book.ID,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}

		arrBooks = append(arrBooks, bookMapping)
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get all books",
		Data:    arrBooks,
	}
	return result, nil
}
