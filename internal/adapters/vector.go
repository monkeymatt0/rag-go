package adapters

import (
	"context"
	"customrag/internal/core/domain"

	"github.com/qdrant/go-client/qdrant"
)

type VectorRepository struct {
	client *qdrant.Client
}

func(vr *VectorRepository) NewVectorRepository(host string, port int) (*VectorRepository, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: host,
		Port: port,
	})

	if err != nil {
		return nil, err
	}

	return &VectorRepository{client: client}, nil
}

func(vr *VectorRepository) CreateCollection(
	ctx context.Context, 
	collectionName string, 
	vectorSize uint64,
) (error) {

	if collectionName == "" || collectionName == " " {
		return domain.ErrEmptyCollection
	}

	collectionConfig := &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &qdrant.VectorsConfig{
			Config: &qdrant.VectorsConfig_Params{
				Params: &qdrant.VectorParams{
					Size:	vectorSize,
					Distance: qdrant.Distance_Cosine,
				},
			},
		},
	}

	err := vr.client.CreateCollection(ctx, collectionConfig)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Finish to implement the VectorInterface + write tests