package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphql_server_gqlgen/graph/generated"
	"graphql_server_gqlgen/graph/model"
	"log"
	"strconv"
)

var BOOKSTORE = []*model.Book{}

func nextBookID() (string, error) {
	fmt.Println("Finding next Book ID")
	if len(BOOKSTORE) < 1 {
		fmt.Println("Next BOOKID: 1")
		return "1", nil
	}
	max := *BOOKSTORE[0].BookID
	for _, book := range BOOKSTORE {
		cBookID, err := strconv.Atoi(*book.BookID)
		if err != nil {
			return "0", err
		}
		maxVar, err := strconv.Atoi(max)
		if err != nil {
			return "0", err
		}
		if cBookID > maxVar {
			max = *book.BookID
		}
	}
	maxInt, _ := strconv.Atoi(max)
	max = strconv.Itoa(maxInt + 1)
	fmt.Println("Next BOOKID : ", max)
	return max, nil
}

func updateBookStore(bookid string, updatedBook *model.UpdateInput) bool {
	for index, book := range BOOKSTORE {
		if *book.BookID == bookid {
			BOOKSTORE = append(BOOKSTORE[:index], BOOKSTORE[index+1:]...)
			uAuthorList := []*model.Author{}
			var uTitle string
			var uGenre *model.BookGenre
			if updatedBook.Authors != nil {
				for _, a := range updatedBook.Authors {
					uAuthorList = append(uAuthorList, &model.Author{Name: *a.Name, Mail: a.Mail})
				}

			} else {
				uAuthorList = book.Authors
			}
			if updatedBook.Genre != nil {
				uGenre = updatedBook.Genre
			} else {
				uGenre = book.Genre
			}
			if updatedBook.Title != nil {
				uTitle = *updatedBook.Title
			} else {
				uTitle = book.Title
			}
			uBook := model.Book{BookID: &bookid, Title: uTitle, Genre: uGenre, Authors: uAuthorList}
			BOOKSTORE = append(BOOKSTORE, &uBook)
			return true
		}
	}
	return false
}

// AddBook is the resolver for the add_book field.
func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.PostStatus, error) {
	log.Println("Adding BOOK to the BOOKSTORE")
	fmt.Printf("%v\n", input)
	a := input.Authors
	var authorList = []*model.Author{}
	for _, auth := range a {
		authorList = append(authorList, &model.Author{Name: auth.Name, Mail: auth.Mail})
	}
	bookID, err := nextBookID()
	if err != nil {
		fmt.Println(err)
		des := "Internal errror"
		return &model.PostStatus{Iserror: true, Description: &des}, nil
	}
	gBook := model.Book{Title: input.Title, Genre: input.Genre, Authors: authorList, BookID: &bookID}
	BOOKSTORE = append(BOOKSTORE, &gBook)
	log.Println("Successfully added the book.")
	des := "Successfully added"
	fmt.Printf("BOOKSTORE : %v\n", &BOOKSTORE)
	return &model.PostStatus{Iserror: false, Description: &des, BookID: &bookID}, nil
}

// UpdateBook is the resolver for the update_book field.
func (r *mutationResolver) UpdateBook(ctx context.Context, input *model.UpdateInput) (*model.PutStatus, error) {
	var bookId = input.BookID
	res := updateBookStore(bookId, input)
	des := "Updation failed"
	if res {
		des = "Updation success"
		return &model.PutStatus{Iserror: false, Description: &des}, nil
	}
	return &model.PutStatus{Iserror: true, Description: &des}, errors.New("book id not found")

}

// DeleteBook is the resolver for the delete_book field.
func (r *mutationResolver) DeleteBook(ctx context.Context, bookID string) (*model.DeleteStatus, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - delete_book"))
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.GetBookResult, error) {
	for _, book := range BOOKSTORE {
		if *book.BookID == bookID {
			return &model.GetBookResult{Isexists: true, Book: book}, nil
		}
	}
	return &model.GetBookResult{Isexists: false}, errors.New("bookid not found")

}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return BOOKSTORE, nil
}

// Getbooks is the resolver for the getbooks field.
func (r *queryResolver) Getbooks(ctx context.Context, getgenre *model.BookGenre) ([]*model.Book, error) {
	var returnList = []*model.Book{}
	for _, book := range BOOKSTORE {
		if *book.Genre == *getgenre {
			returnList = append(returnList, book)
		}
	}
	return returnList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
