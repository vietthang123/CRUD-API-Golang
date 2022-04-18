package book

type Book struct {
	tableName struct{} `pg:"book"`
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Category  string   `json:"category"`
	Content   string   `jonn:"content"`
	Author    string   `json:"author"`
}
