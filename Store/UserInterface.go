package store
import(
	"Backend/project/Models"
)

type Userstore interface{
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUserInfoByID(userID string, user *models.User) error
	UpdateUserCredit(userID int, newCredit float64) error
	CheckDBConnection() error

}