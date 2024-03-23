package role_models

type Role struct {
	tableName   struct{} `pg:"roles"`
	RoleID      int      `pg:"role_id, pk"`
	RoleName    string   `pg:"role_name, unique, notnull"`
	Description string   `pg:"description, notnull"`
}

func RoleIDToName(roleID int) string {
	switch roleID {
	case 1:
		return "User"
	case 2:
		return "Admin"
	default:
		return "Unknown"
	}
}
