package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql_server_gqlgen/graph/generated"
	"graphql_server_gqlgen/graph/model"
)

// AddBook is the resolver for the add_book field.
func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.PostStatus, error) {
	panic(fmt.Errorf("not implemented: AddBook - add_book"))
}

// UpdateBook is the resolver for the update_book field.
func (r *mutationResolver) UpdateBook(ctx context.Context, input *model.UpdateInput) (*model.PutStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateBook - update_book"))
}

// DeleteBook is the resolver for the delete_book field.
func (r *mutationResolver) DeleteBook(ctx context.Context, bookID string) (*model.DeleteStatus, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - delete_book"))
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.GetBookResult, error) {
	panic(fmt.Errorf("not implemented: Book - book"))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	bookGenre := model.BookGenre("FICTION")
	bookID := "1"
	mail := "logesh@mail.com"
	var authorList = []*model.Author{{Name: "Logesh", Mail: &mail}}
	reVal := &model.Book{Title: "Empty title", BookID: &bookID, Genre: &bookGenre, Authors: authorList}
	var returnList = []*model.Book{reVal}
	return returnList, nil
}

// Getbooks is the resolver for the getbooks field.
func (r *queryResolver) Getbooks(ctx context.Context, getgenre *model.BookGenre) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Getbooks - getbooks"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
