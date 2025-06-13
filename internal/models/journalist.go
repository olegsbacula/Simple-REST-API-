package models

type Journalist struct {
   	ID        int    `db:"id" json:"id"`
    Name      string `db:"name" json:"name"`
	Last_name string `db:"last_name" json:"last_name"`
	Email_journalist string `db:"email_journalist" json:"email_journalist"`
    Created_articles_Id  string `db:"created_articles_Id" json:"created_articles_Id"`
}
