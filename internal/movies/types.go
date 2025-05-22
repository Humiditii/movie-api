package movies

type Movie struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title"`
	Genre    string `json:"genre" bson:"genre"`
	Director string `json:"director" bson:"director"`
}
