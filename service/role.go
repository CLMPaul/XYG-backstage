package service

import (
	"database/sql"
	"gorm.io/gorm/clause"
	"xueyigou_demo/db"
)

type roleService struct{}

// AuthenticatePermission 验证用户在项目下是否具有某权限
func (*roleService) AuthenticatePermission(userID string, action string) (granted bool, err error) {
	selectFrom := clause.Table{Name: "users", Alias: "u"}
	joins := []clause.Join{
		{
			Type:  clause.InnerJoin,
			Table: clause.Table{Name: "role_users", Alias: "ur"},
			ON: clause.Where{
				Exprs: []clause.Expression{
					clause.Eq{
						Column: clause.Column{Table: "ur", Name: "user_id"},
						Value:  clause.Column{Table: "u", Name: "id"},
					},
				},
			},
		},
		{
			Type:  clause.InnerJoin,
			Table: clause.Table{Name: "roles", Alias: "r"},
			ON: clause.Where{
				Exprs: []clause.Expression{
					clause.Eq{
						Column: clause.Column{Table: "r", Name: "id"},
						Value:  clause.Column{Table: "ur", Name: "role_id"},
					},
				},
			},
		},
		{
			Type:  clause.InnerJoin,
			Table: clause.Table{Name: "role_permissions", Alias: "rp"},
			ON: clause.Where{
				Exprs: []clause.Expression{
					clause.Eq{
						Column: clause.Column{Table: "rp", Name: "role_id"},
						Value:  clause.Column{Table: "ur", Name: "role_id"},
					},
				},
			},
		},
	}
	columns := []clause.Column{
		{Name: "1", Raw: true},
	}
	conditions := []clause.Expression{
		clause.Eq{
			Column: clause.Column{Table: "u", Name: "id"},
			Value:  userID,
		},
		clause.Eq{
			Column: clause.Column{Table: "rp", Name: "action"},
			Value:  action,
		},
	}
	subquery := db.DB.Table("?", selectFrom).
		Clauses(clause.From{Tables: []clause.Table{selectFrom}, Joins: joins}).
		Clauses(clause.Where{Exprs: conditions}).
		Clauses(clause.Select{Columns: columns})
	if db.DB.Dialector.Name() == "dm" {
		var innerGranted sql.NullBool
		err = db.DB.Raw("SELECT ? WHERE EXISTS (?)", true, subquery).Scan(&innerGranted).Error
		granted = err == nil && innerGranted.Valid && innerGranted.Bool
	} else {
		err = db.DB.Raw("SELECT EXISTS (?)", subquery).Scan(&granted).Error
	}
	return
}
