package postgres

import (
	"database/sql"
	pb "discovery_servcie/genproto"
	strorage "discovery_servcie/help"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type DiscoveryRepository struct {
	Db *sql.DB
}

func NewDiscoveryRepository(db *sql.DB) *DiscoveryRepository {
	return &DiscoveryRepository{Db: db}
}

//	func (repo *CompositionRepository) CreateCompositionMetadata(metadata *models.CompositionMetadata) (interface{}, error) {
//		return repo.Db.Exec("INSERT INTO composition_metadata (composition_id, genre, tags, listen_count, like_count) VALUES ($1, $2, $3, $4, $5)",
//			metadata.CompositionID, metadata.Genre, pq.Array(metadata.Tags), metadata.ListenCount, metadata.LikeCount)
//	}
//
//	func (repo *CompositionRepository) GetCompositionMetadataByID(id int) (*models.CompositionMetadata, error) {
//		row := repo.Db.QueryRow("SELECT composition_id, genre, tags, listen_count, like_count FROM composition_metadata WHERE composition_id = $1", id)
//
//		var metadata models.CompositionMetadata
//		err := row.Scan(&metadata.CompositionID, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
//		if err != nil {
//			return nil, err
//		}
//		return &metadata, nil
//	}
//
//	func (repo *CompositionRepository) UpdateCompositionMetadata(metadata *models.CompositionMetadata) (interface{}, error) {
//		return repo.Db.Exec("UPDATE composition_metadata SET genre = $1, tags = $2, listen_count = $3, like_count = $4 WHERE composition_id = $5",
//			metadata.Genre, pq.Array(metadata.Tags), metadata.ListenCount, metadata.LikeCount, metadata.CompositionID)
//	}
//
//	func (repo *CompositionRepository) DeleteCompositionMetadata(id int) (interface{}, error) {
//		return repo.Db.Exec("update  composition_metadata set deleted_at=$1 WHERE composition_id = $2", id, time.Now())
//	}
func (repo *DiscoveryRepository) GetCompositionMetadata(filter *pb.GetDiscoveryRequest) (*pb.DiscoveriesResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	queryFilter := ""

	if filter.Genre != "" {
		params["genre"] = filter.Genre
		queryFilter += " AND genre = :genre"
	}
	if len(filter.Tags) > 0 {
		params["tags"] = pq.Array(filter.Tags)
		queryFilter += " AND tags @> :tags"
	}
	if filter.ListenCount > 0 {
		params["listen_count"] = filter.ListenCount
		queryFilter += " AND listen_count = :listen_count"
	}
	if filter.LikeCount > 0 {
		params["like_count"] = filter.LikeCount
		queryFilter += " AND like_count = :like_count"
	}

	if filter.LimitOffset.Limit > 0 {
		params["limit"] = filter.LimitOffset.Limit
		limit = ` LIMIT :limit`
	}
	if filter.LimitOffset.Offset > 0 {
		params["offset"] = filter.LimitOffset.Offset
		offset = ` OFFSET :offset`
	}

	query := "SELECT composition_id, genre, tags, listen_count, like_count FROM composition_metadata WHERE deleted_at is null "
	query = query + queryFilter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compositions []*pb.DiscoveryResponse
	for rows.Next() {
		var metadata pb.DiscoveryResponse
		err := rows.Scan(&metadata.CompositionId, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, &metadata)
	}

	return &pb.DiscoveriesResponse{Discoveries: compositions}, nil
}

func (repo *DiscoveryRepository) GetCompositionTrending(void *pb.Void) (*pb.DiscoveriesResponse, error) {
	rows, err := repo.Db.Query("SELECT   composition_id, genre, tags, listen_count, like_count  from composition_metadata order by like_count desc ")
	if err != nil {
		fmt.Println("++++++++++++++++", err)
		return nil, err
	}
	var compositions []*pb.DiscoveryResponse
	for rows.Next() {
		var metadata pb.DiscoveryResponse
		err := rows.Scan(&metadata.CompositionId, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, &metadata)
	}

	return &pb.DiscoveriesResponse{Discoveries: compositions}, nil
}
func (repo *DiscoveryRepository) GetCompositionRecommend(void *pb.Void) (*pb.DiscoveriesResponse, error) {
	rows, err := repo.Db.Query("SELECT   composition_id, genre, tags, listen_count, like_count  composition_metadata order by like_count desc ")
	if err != nil {
		return nil, err
	}
	var compositions []*pb.DiscoveryResponse
	for rows.Next() {
		var metadata pb.DiscoveryResponse
		err := rows.Scan(&metadata.CompositionId, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, &metadata)
	}

	return &pb.DiscoveriesResponse{Discoveries: compositions}, nil
}
func (repo *DiscoveryRepository) GetDiscoveriesByGenre(genre *pb.GetGenre) (*pb.DiscoveriesResponse, error) {
	if genre == nil || genre.Genre == "" {
		return nil, errors.New("invalid genre parameter")
	}

	query := `
        SELECT composition_id, genre, tags, listen_count, like_count
        FROM composition_metadata
        WHERE genre = $1
    `

	rows, err := repo.Db.Query(query, genre.Genre)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var compositions []*pb.DiscoveryResponse
	for rows.Next() {
		var metadata pb.DiscoveryResponse
		err := rows.Scan(&metadata.CompositionId, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, &metadata)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.DiscoveriesResponse{Discoveries: compositions}, nil
}

func (repo *DiscoveryRepository) DeleteCompositionLike(compositionId *pb.LikeRequest) (*pb.Void, error) {

	_, err := repo.Db.Exec("delete from user_interactions WHERE user_id = $1 AND composition_id = $2  AND interaction_type = 'like' ", compositionId.UserId, compositionId.CompositionId)
	if err != nil {
		return nil, err
	}
	_, err = repo.Db.Exec("update composition_metadata set like_count=like_count-1 where composition_id=$1 ", compositionId.CompositionId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo *DiscoveryRepository) CreateCompositionLike(compositionId *pb.LikeRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into user_interactions (user_id, composition_id, interaction_type) values ($1, $2, 'like')", compositionId.UserId, compositionId.CompositionId)
	if err != nil {
		return nil, err
	}

	_, err = repo.Db.Exec("update composition_metadata SET like_count = like_count + 1 where composition_id = $1", compositionId.CompositionId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (repo *DiscoveryRepository) CreateCompositionListen(compositionId *pb.LikeRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into user_interactions (user_id, composition_id, interaction_type)values ($1, $2, 'listen')", compositionId.UserId, compositionId.CompositionId)
	if err != nil {
		return nil, err
	}
	_, err = repo.Db.Exec("update composition_metadata set listen_count=listen_count+1 where composition_id=$1 ", compositionId.CompositionId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
