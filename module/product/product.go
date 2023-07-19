package product

import "database/sql"

type Product struct {
	Id   sql.NullInt64  `db:"product.id" fieldtag:"view"`
	Name sql.NullString `db:"product.name" fieldtag:"view"`
}
