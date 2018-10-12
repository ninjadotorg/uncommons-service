package daos

import (
	"github.com/ninjadotorg/uncommons-service/models"
)

// UserDAO ...
type UserDAO struct{}

// IsUserExist ...
func (u UserDAO) IsUserExist(walletAddress string) (models.User, error) {
	user := models.User{}
	err := models.Database().Where("wallet_address = ?", walletAddress).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
