package models

type Article struct {
    ID        int    `db:"id" json:"id"`
    Name      string `db:"name" json:"name"`
    Title 	  string `db:"title" json:"title"`
    Description  string `db:"description" json:"description"`
	Created_at	string `db:"created_at" json:"created_at"`
}
