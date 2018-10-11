package daos

type userDAO struct{}

func (u userDAO) IsUserExist(walletAddress string) bool {
	return false
}
