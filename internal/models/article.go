package models

type Article struct {
    Article_ID        int    `json:"article_id"`
    Article_Name      string `json:"article_name"`
    Title 	  string `json:"title"`
    Description  string `json:"description"`
	Article_Created_at	string `json:"article_created_at"`
}
