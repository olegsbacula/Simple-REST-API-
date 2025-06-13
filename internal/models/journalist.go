package models

type Journalist struct {
   	ID        int    `json:"journalist_id"`
    Journalist_name      string `json:"journalist_name"`
	Journalist_last_name string `json:"journalist_last_name"`
	Email_journalist string `json:"email_journalist"`
    Created_articles_Id  string `json:"created_articles_Id"`
}
