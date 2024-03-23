package user_role_assoc_models

import (
	"log"

	"github.com/ljlimjk10/users-ms/models/role_models"
	"github.com/ljlimjk10/users-ms/models/user_models"

	"github.com/go-pg/pg/v10"
)

// 1 user -> 1 role only
type UserRoleAssociation struct {
	tableName  struct{}          `pg:"user_role_association"`
	UserRoleID int               `pg:"user_role_id, pk"`
	UserID     int               `pg:"user_id, notnull"`
	RoleID     int               `pg:"role_id, notnull"`
	User       *user_models.User `pg:"rel:has-one"`
	Role       *role_models.Role `pg:"rel:has-one"`
}

func CreateUserRoleAssoc(db *pg.DB, newUserRoleAssoc *UserRoleAssociation) error {
	_, err := db.Model(newUserRoleAssoc).Insert()
	return err
}

func GetUserRoleID(db *pg.DB, userID int) (int, error) {
	userRoleAssoc := new(UserRoleAssociation)
	err := db.Model(userRoleAssoc).Where("user_id = ?", userID).Select()

	if err != nil {
		log.Println(err)
		return -1, err
	}
	return userRoleAssoc.RoleID, nil
}
