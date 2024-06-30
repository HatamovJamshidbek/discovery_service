package postgres

import (
	"database/sql"
	strorage "discovery_servcie/help"
	"discovery_servcie/models"
	"github.com/lib/pq"
)

type CompositionRepository struct {
	Db *sql.DB
}

func NewCompositionRepository(db *sql.DB) *CompositionRepository {
	return &CompositionRepository{Db: db}
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
func (repo *CompositionRepository) GetCompositionMetadata(filter *models.CompositionMetadata) (*[]models.CompositionMetadata, error) {
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

	if filter.Limit > 0 {
		params["limit"] = filter.Limit
		limit = ` LIMIT :limit`
	}
	if filter.Offset > 0 {
		params["offset"] = filter.Offset
		offset = ` OFFSET :offset`
	}

	query := "SELECT composition_id, genre, tags, listen_count, like_count FROM composition_metadatas WHERE delted_at is null "
	query = query + queryFilter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compositions []models.CompositionMetadata
	for rows.Next() {
		var metadata models.CompositionMetadata
		err := rows.Scan(&metadata.CompositionID, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, metadata)
	}

	return &compositions, nil
}

func (repo *CompositionRepository) GetCompositionTrending(filter *models.CompositionMetadata) (*[]models.CompositionMetadata, error) {
	rows, err := repo.Db.Query("SELECT   composition_id, genre, tags, listen_count, like_count  from composition_meta_datas where like_count = (SELECT MAX(like_count) FROM composition_meta_datas)")
	if err != nil {
		return nil, err
	}
	var compositions []models.CompositionMetadata
	for rows.Next() {
		var metadata models.CompositionMetadata
		err := rows.Scan(&metadata.CompositionID, &metadata.Genre, pq.Array(&metadata.Tags), &metadata.ListenCount, &metadata.LikeCount)
		if err != nil {
			return nil, err
		}
		compositions = append(compositions, metadata)
	}

	return &compositions, nil
}
func (repo *CompositionRepository) DeleteCompositionLike(compositionId int) (interface{}, error) {
	return repo.Db.Exec("update composition_meta_datas set like_count=like_count-1 where composition_id=$1 ", compositionId)
}
func (repo *CompositionRepository) CreateCompositionLike(compositionId int) (interface{}, error) {
	return repo.Db.Exec("update composition_meta_datas set like_count=like_count+1 where composition_id=$1 ", compositionId)
}
