package repositories

import (
	"context"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/models"
)

func (r repository) CreateUser(user models.User) (models.User, error) {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	_, err := r.db.NewInsert().Model(&user).ExcludeColumn("id").Returning("*").Exec(context.TODO())
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r repository) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.db.NewSelect().Model(&user).Where("email = ?", email).Scan(context.TODO())
	return user, err
}

func (r repository) GetUser(id string) (models.User, error) {
	user := models.User{}
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.TODO())
	return user, err
}

func (r repository) UpdateUser(id string, user models.User) (models.User, error) {
	updateUser := map[string]interface{}{
		"updated_at": time.Now().UTC(),
	}
	query := r.db.NewUpdate().Model(&updateUser).TableExpr("users")
	if user.Password != "" {
		updateUser["password"] = user.Password
	}
	_, err := query.Where("id = ?", id).Returning("*").Exec(context.TODO())
	return user, err
}
