package postgres

import (
	strorage "discovery_servcie/help"
	"discovery_servcie/models"
	"time"
)

func (repo *CompositionRepository) CreateUserInteraction(compositionId string, interaction *models.UserInteraction) (interface{}, error) {
	return repo.Db.Exec("INSERT INTO user_interactions (user_id, composition_id, interaction_type, created_at) VALUES ($1, $2, $3, $4)",
		interaction.UserID, compositionId, interaction.InteractionType, time.Now())
}

//	func (repo *CompositionRepository) GetUserInteractionByID(id int) (*models.UserInteraction, error) {
//		rows, err := repo.Db.Query("SELECT id, user_id, composition_id, interaction_type, created_at FROM user_interactions WHERE id = $1", id)
//		if err != nil {
//			return nil, err
//		}
//		var interaction models.UserInteraction
//		for rows.Next() {
//			err := rows.Scan(&interaction.ID, &interaction.UserID, &interaction.CompositionID, &interaction.InteractionType, &interaction.CreatedAt)
//			if err != nil {
//				return nil, err
//			}
//			return &interaction, nil
//		}
//		return nil, err
//
// }
//
//	func (repo *CompositionRepository) UpdateUserInteraction(interaction *models.UserInteraction, id string) (interface{}, error) {
//		return repo.Db.Exec("UPDATE user_interactions SET user_id = $1, composition_id = $2, interaction_type = $3, updated_at = $4 WHERE id = $5",
//			interaction.UserID, interaction.CompositionID, interaction.InteractionType, time.Now(), id)
//	}
func (repo *CompositionRepository) DeleteUserInteraction(compositionId string) (interface{}, error) {
	return repo.Db.Exec("update  user_interactions set deleted_at=$1 WHERE id = $2", time.Now(), id)
}
func (repo *CompositionRepository) GetUserInteractions(filter *models.UserInteraction) (*[]models.UserInteraction, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	queryFilter := ""

	if filter.UserID > 0 {
		params["user_id"] = filter.UserID
		queryFilter += " AND user_id = :user_id"
	}
	if filter.CompositionID > 0 {
		params["composition_id"] = filter.CompositionID
		queryFilter += " AND composition_id = :composition_id"
	}
	if filter.InteractionType != "" {
		params["interaction_type"] = filter.InteractionType
		queryFilter += " AND interaction_type = :interaction_type"
	}

	if filter.Limit > 0 {
		params["limit"] = filter.Limit
		limit = ` LIMIT :limit`
	}
	if filter.Offset > 0 {
		params["offset"] = filter.Offset
		offset = ` OFFSET :offset`
	}

	query := "SELECT id, user_id, composition_id, interaction_type, created_at FROM user_interactions WHERE 1=1"
	query = query + queryFilter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interactions []models.UserInteraction
	for rows.Next() {
		var interaction models.UserInteraction
		err := rows.Scan(&interaction.ID, &interaction.UserID, &interaction.CompositionID, &interaction.InteractionType, &interaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		interactions = append(interactions, interaction)
	}

	return &interactions, nil
}
