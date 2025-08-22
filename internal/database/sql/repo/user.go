package repo

import (
	"fmt"
	"gin-starter/internal/database/sql/model"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	ModelRepository[model.User, uuid.UUID]
	FindByEmail(email string) (model.User, error)
}

var userRepositoryObect UserRepository
var userRepoOnce sync.Once

type userRepositoryImpl struct {
	readableModelRepositoryImpl[model.User, uuid.UUID]
	writableModelRepositoryImpl[model.User, uuid.UUID]
	db *gorm.DB
}

func (u *userRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User

	queryString := fmt.Sprintf("email = %s", model.USER_EMAIL_FIELD)

	result := u.db.Where(queryString, email).First(&user)
	if err := result.Error; err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepoOnce.Do(func() {
		userRepositoryObect = &userRepositoryImpl{
			db: db,
		}
	})

	return userRepositoryObect
}
