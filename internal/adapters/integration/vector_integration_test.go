package integration

import (
	"context"
	"customrag/internal/adapters"
	"os"
	"testing"

	"github.com/qdrant/go-client/qdrant"
	"github.com/stretchr/testify/require"
)

const(
	testHost = "localhost"
	testPort = 6334
	testCollection = "testCollection"
)

var(
	data = []*qdrant.PointStruct{
	{
		Id: &qdrant.PointId{
			PointIdOptions: &qdrant.PointId_Num{
				Num: 1,
			},
		},
		Vectors: &qdrant.Vectors{
			VectorsOptions: &qdrant.Vectors_Vector{
				Vector: &qdrant.Vector{
					Data: []float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0},
				},
			},
		},
		Payload: map[string]*qdrant.Value{
			"title": {
				Kind: &qdrant.Value_StringValue{
					StringValue: "Test Document 1",
				},
			},
			"content": {
				Kind: &qdrant.Value_StringValue{
					StringValue: "This is test content for document 1",
				},
			},
		},
	},
}
)

func setupTestVectorRepository() (*adapters.VectorRepository, context.Context, error) {
	repo, err :=  adapters.NewVectorRepository(testHost, testPort)
	ctx := context.Background()
	repo.DeleteCollection(ctx, testCollection)
	return repo, ctx, err
}

// TODO: Write integration tests
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateCollection_Success(t *testing.T) {
	// Setup
	repo, ctx, err := setupTestVectorRepository()
	require.NoError(t, err, "Failed to setup the repo")

	// Creating collection
	err2 := repo.CreateCollection(ctx, testCollection, 516)
	require.NoError(t, err2, "Failed to create collection")
}

func TestCreateCollection_DuplicatedAndVoidCollectionName_Fail(t *testing.T) {
	repo, ctx, err := setupTestVectorRepository()
	require.NoError(t, err, "Failed to setup the repo")
	
	// Creating 2 times same collection must result in an error
	err2 := repo.CreateCollection(ctx, testCollection, 516)
	require.NoError(t, err2, "Failed to create the first collection")
	// This call has to fail
	err3 := repo.CreateCollection(ctx, testCollection, 516)
	require.Error(t, err3, "Properly failed as expected")
	// This call has to fail
	err4 := repo.CreateCollection(ctx, "", 516)
	require.Error(t, err4)
	// This call has to fail
	err5 := repo.CreateCollection(ctx, " ", 516)
	require.Error(t, err5)

}

func TestDeleteCollection_DeleteNotExistingCollection_Fail(t *testing.T) {
	repo, ctx, err := setupTestVectorRepository()
	require.NoError(t, err)

	err2 := repo.DeleteCollection(ctx, testCollection)
	require.Error(t, err2)
}

func TestCreationAndSearch_Success(t *testing.T) {
	repo, ctx, err := setupTestVectorRepository()
	require.NoError(t, err)

	err2 := repo.CreateCollection(ctx, testCollection, 10)
	require.NoError(t, err2)

	err3 := repo.CreateData(ctx, testCollection, data)
	require.NoError(t, err3)

	queryVector := []float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	_, err4 := repo.Search(ctx, testCollection, queryVector, 1)
	require.NoError(t, err4, "Failed to search")
}