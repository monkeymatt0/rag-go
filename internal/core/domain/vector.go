package domain

import "errors"

var (
	ErrCollectionNotFound = errors.New("collection not found")
	ErrVectorNotFound = errors.New("vector not found")
	ErrInvalidQuery = errors.New("invalid query vector")
	ErrEmptyCollection = errors.New("collection name cannot be empty")
)

/*
Right now I just need these errors, since for now the types provided
from the qdrant SDK are enough.
*/