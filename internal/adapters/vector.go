package adapters

import (
	"context"
	"customrag/internal/core/domain"

	"github.com/qdrant/go-client/qdrant"
)

type VectorRepository struct {
	client *qdrant.Client
}

func NewVectorRepository(host string, port int) (*VectorRepository, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: host,
		Port: port,
		SkipCompatibilityCheck: true, // This flag disable compatibility check for now it's ok -> TODO: Fix this warning, this regards the version of qdrant you have on docker compared to the client you are using.
	})

	if err != nil {
		return nil, err
	}

	return &VectorRepository{client: client}, nil
}

func(vr *VectorRepository) Search(
	ctx context.Context,
	collectionName string,
	query []float32,
	limit uint64,
) ([]*qdrant.ScoredPoint, error) {
	searchRequest := &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query: qdrant.NewQueryDense(query),
		WithPayload: qdrant.NewWithPayload(true),
		WithVectors: qdrant.NewWithVectors(false),
		Limit: &limit,
		// Offset: offset, TODO: Update the interface to include also an offset
	}

	result, err := vr.client.Query(ctx, searchRequest)
	if err != nil {
		return nil, domain.ErrInvalidQuery // TODO: Map errors to the ones of the domain
	}
	return result, nil
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

func(vr *VectorRepository) DeleteCollection(
	ctx context.Context, 
	collectionName string,
) (error) {
	if collectionName == "" || collectionName == " " {
		return domain.ErrEmptyCollection
	}

	err := vr.client.DeleteCollection(ctx, collectionName)
	if err != nil {
		return err
	}

	return nil
}


func(vr *VectorRepository) CreateData(
	ctx context.Context,
	collectionName string,
	points []*qdrant.PointStruct,
) (error) {
	request := qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: points,
	}
	_, err := vr.client.Upsert(ctx, &request)
	if err != nil {
		return domain.ErrInvalidQuery
	}

	return nil
}


func(vr *VectorRepository) DeleteData(
	ctx context.Context,
	collectionName string,
	ids []uint64,
) (error) {

	var qids []*qdrant.PointId
	for _, id := range(ids) {
		qids = append(qids, &qdrant.PointId{
			PointIdOptions: &qdrant.PointId_Num{
				Num: id,
			},
		})
	}

	deleteRequest := &qdrant.DeletePointVectors{
		CollectionName: collectionName,
		PointsSelector: &qdrant.PointsSelector{
			PointsSelectorOneOf: &qdrant.PointsSelector_Points{
				Points: &qdrant.PointsIdsList{
					Ids: qids,
				},
			},
		},
	}

	if len(ids) == 0 || collectionName == "" || collectionName == " "{
		return domain.ErrEmptyCollection
	}

	_, err := vr.client.DeleteVectors(ctx, deleteRequest)
	if err != nil {
		return domain.ErrInvalidQuery
	}

	return nil
}