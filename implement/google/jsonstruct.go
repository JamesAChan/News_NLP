package google

// define json structure for unmarshal
type GoogleJS struct {
	Status    string	`json:"status"`
	Total     int64		`json:"totalResults"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	Sources 	Source   `json:"source"`
	Author      string	 `json:"author"`
	Title		string	 `json:"title"`
	Description string	 `json:"description"`
	URL     	string	 `json:"url"`
	URL2image   string	 `json:"urlToImage"`
	Time 		string	 `json:"PublishedAt"`
	Content   	string	 `json:"content"`
}

type MyArticle struct {
	Sources 	Source
	Author      string
	Title		string
	Description string
	URL     	string
	URL2image   string
	Time 		string
	Content   	string
	Category    string
}



type Source struct {
	ID		string		`json:"id"`
	Name  	string		`json:"name"`
}