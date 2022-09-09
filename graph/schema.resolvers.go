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

func nextBookID(bookstore []*model.Book) (string, error) {
	fmt.Println("Finding next Book ID")
	if len(bookstore) < 1 {
		fmt.Println("Next BOOKID: 1")
		return "1", nil
	}
	max := *bookstore[0].BookID
	for _, book := range bookstore {
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

// AddBook is the resolver for the add_book field.
func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.PostStatus, error) {
	log.Println("Adding BOOK to the BOOKSTORE")
	a := input.Authors
	var authorList = []*model.Author{}
	for _, auth := range a {
		authorList = append(authorList, &model.Author{Name: auth.Name, Mail: auth.Mail})
	}
	bookID, err := nextBookID(r.BOOKSTORE)
	if err != nil {
		fmt.Println(err)
		des := "Internal errror"
		return &model.PostStatus{Iserror: true, Description: &des}, nil
	}
	gBook := model.Book{Title: input.Title, Genre: input.Genre, Authors: authorList, BookID: &bookID}
	r.BOOKSTORE = append(r.BOOKSTORE, &gBook)
	log.Println("Successfully added the book.")
	des := "Successfully added"
	return &model.PostStatus{Iserror: false, Description: &des, BookID: &bookID}, nil
}

// UpdateBook is the resolver for the update_book field.
func (r *mutationResolver) UpdateBook(ctx context.Context, input *model.UpdateInput) (*model.PutStatus, error) {
	var bookId = input.BookID
	log.Printf("Updating the book (BookID : %s) from the BOOKSTORE\n", bookId)
	for index, book := range r.BOOKSTORE {
		if *book.BookID == bookId {
			r.BOOKSTORE = append(r.BOOKSTORE[:index], r.BOOKSTORE[index+1:]...)
			uAuthorList := []*model.Author{}
			var uTitle string
			var uGenre *model.BookGenre
			if input.Authors != nil {
				for _, a := range input.Authors {
					uAuthorList = append(uAuthorList, &model.Author{Name: *a.Name, Mail: a.Mail})
				}

			} else {
				uAuthorList = book.Authors
			}
			if input.Genre != nil {
				uGenre = input.Genre
			} else {
				uGenre = book.Genre
			}
			if input.Title != nil {
				uTitle = *input.Title
			} else {
				uTitle = book.Title
			}
			uBook := model.Book{BookID: &bookId, Title: uTitle, Genre: uGenre, Authors: uAuthorList}
			r.BOOKSTORE = append(r.BOOKSTORE, &uBook)
			des := "Updation success"
			log.Println(des)
			return &model.PutStatus{Iserror: false, Description: &des}, nil
		}
	}
	des := "failed to find the book id"
	log.Println(des)
	return &model.PutStatus{Iserror: true, Description: &des}, errors.New(des)

}

// DeleteBook is the resolver for the delete_book field.
func (r *mutationResolver) DeleteBook(ctx context.Context, bookID string) (*model.DeleteStatus, error) {
	log.Printf("Deleting the book (BookID : %s) from the BOOKSTORE\n", bookID)
	for index, book := range r.BOOKSTORE {
		if *book.BookID == bookID {
			r.BOOKSTORE = append(r.BOOKSTORE[:index], r.BOOKSTORE[index+1:]...)
			des := "deletion success"
			log.Println(des)
			return &model.DeleteStatus{Iserror: false, Description: &des}, nil
		}
	}
	des := "faild to find the book id to delete"
	log.Println(des)
	return &model.DeleteStatus{Iserror: true, Description: &des}, errors.New(des)

}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.GetBookResult, error) {
	log.Println("Getting Book ", bookID, " from BOOKSTORE")
	for _, book := range r.BOOKSTORE {
		if *book.BookID == bookID {
			return &model.GetBookResult{Isexists: true, Book: book}, nil
		}
	}
	return &model.GetBookResult{Isexists: false}, errors.New("bookid not found")

}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	log.Println("Getting all Books from BOOKSTORE")
	return r.BOOKSTORE, nil
}

// Getbooks is the resolver for the getbooks field.
func (r *queryResolver) Getbooks(ctx context.Context, getgenre *model.BookGenre) ([]*model.Book, error) {
	log.Println("Getting books which are ", getgenre, " from BOOKSTORE")
	var returnList = []*model.Book{}
	for _, book := range r.BOOKSTORE {
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
