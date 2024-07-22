package dbcollection

import (
	"database/sql"
	"fmt"
	"log"
	"Backend/project/Models"
)

type UserStore struct {
	db *sql.DB
}

// GetUserByID implements store.Userstore.
func (us *UserStore) GetUserByID(userID string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUserCredit implements store.Userstore.
func (us *UserStore) UpdateUserCredit(userID int, newCredit float64) error {
	panic("unimplemented")
}

// UpdateUserInfoByID implements store.Userstore.
func (us *UserStore) UpdateUserInfoByID(userID string, user *models.User) error {
	panic("unimplemented")
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}
func (us *UserStore) CreateUser(user *models.User) error {
    query := "INSERT INTO users (name, email, password, credit, area, address) VALUES ($1, $2, $3, $4, $5, $6)"
    _, err := us.db.Exec(query, user.Name, user.Email, user.Password, user.Credit, user.Area, user.Address)
    if err != nil {
		log.Printf("Error creating user: %v", err)
         // Log the error
    }
    return err
}
//func (us *UserStore) GetUserByPhoneNumber(phoneNumber float64) (*models.User, error) {
	//user := new(models.User)
	//err := us.db.QueryRow("SELECT * FROM users WHERE phone_number = $1", phoneNumber).Scan(
		//&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Password, &user.Credit, &user.Area, &user.Address,
	//)
	//if err == sql.ErrNoRows {
		//return nil, nil
	//}
	//if err != nil {
		//return nil, err
	//}
	//return user, nil
//}

func (us *UserStore) GetUserByEmail(email string) (*models.User, error) {
    user := new(models.User)
    query := `SELECT id, name, email, phone_number, password, credit, area, address FROM users WHERE email=$1`
    row := us.db.QueryRow(query, email)
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Credit, &user.Area, &user.Address)
    if err!= nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        log.Printf("Error scanning row: %v", err)
        return nil, err
    }
    return user, nil
}

func (s *UserStore) CheckDBConnection() error {
	err := s.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	return nil
}
