package models

type SearchResults struct {
	Search []Movie `json:"Search"`
}
type Movie struct {
	Metascore  string   `json:"Metascore"`
	BoxOffice  string   `json:"BoxOffice"`
	Website    string   `json:"Website"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	Ratings    []Rating `json:"Ratings"`
	Runtime    string   `json:"Runtime"`
	Language   string   `json:"Language"`
	Rated      string   `json:"Rated"`
	Production string   `json:"Production"`
	Released   string   `json:"Released"`
	ImdbID     string   `json:"imdbID"`
	Plot       string   `json:"Plot"`
	Director   string   `json:"Director"`
	Title      string   `json:"Title"`
	Actors     string   `json:"Actors"`
	Response   string   `json:"Response"`
	Type       string   `json:"Type"`
	Awards     string   `json:"Awards"`
	DVD        string   `json:"DVD"`
	Year       string   `json:"Year"`
	Poster     string   `json:"Poster"`
	Country    string   `json:"Country"`
	Genre      string   `json:"Genre"`
	Writer     string   `json:"Writer"`
}
type Rating struct {
	Value  string `json:"Value"`
	Source string `json:"Source"`
}
