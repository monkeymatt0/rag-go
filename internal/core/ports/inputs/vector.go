package inputs

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

type VectorInterface interface {


	/* Search will be used to search information and in general to fetch relevant infos
		params:
			collectionName: name of the collection,
			query: The query the user sent,
			limit: Result limit
	*/
	Search(ctx context.Context, collectionName string, query []float32, limit uint64) ([]*qdrant.ScoredPoint, error) // Search from a vector DB

	
	/* CreateCollection is used to create a new collection
		params:
			collectionName: name of the collection,
			vectorSize: Unsigned integer indicating how is big the vector
	*/
	CreateCollection(ctx context.Context, collectionName string, vectorSize uint64) (error) // Create a new collection


	/* Delete will be used to delete a collection of a given name
		params:
			collectionName: name of the collection,
	*/
	DeleteCollection(ctx context.Context, collectionName string) (error) // To delete a complete collection


	/* CreateData will be used to create a vector within a specific collection
		params:
			collectionName: name of the collection,
			points: The points that will be created
	*/
	CreateData(ctx context.Context, collectionName string, points []*qdrant.PointStruct) (error)


	/* Search wil be used to search information and in general to fetch relevant infos
		params:
			collectionName: name of the collection,
			ids: IDs of the record to delete
	*/
	DeleteData(ctx context.Context, collectionName string, ids []uint64) (error)
}