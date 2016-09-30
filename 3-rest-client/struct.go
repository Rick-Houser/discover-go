// package main
//
// type Movie struct {
// 	Title      string  `json:"Title"`
// 	Year       string  `json:"Year"`
// 	Rated      string  `json:"Rated"`
// 	Released   string  `json:"Released"`
// 	Runtime    string  `json:"Runtime"`
// 	Genre      string  `json:"Genre"`
// 	Director   string  `json:"Director"`
// 	Writer     string  `json:"Writer"`
// 	Actors     string  `json:"Actors"`
// 	Plot       string  `json:"Plot"`
// 	Language   string  `json:"Language"`
// 	Country    string  `json:"Country"`
// 	Awards     string  `json:"Awards"`
// 	Poster     string  `json:"Poster"`
// 	Metascore  string  `json:"Metascore"`
// 	ImdbRating float64 `json:"imdbRating,string"`
// 	ImdbVotes  string  `json:"imdbVotes"`
// 	ImdbID     string  `json:"imdbID"`
// 	Type       string  `json:"Type"`
// 	Response   string  `json:"Response"`
// }

package main

type response struct {
	Title      string
	Year       int `json:"Year,string"`
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  int     `json:"Metascore,string"`
	ImdbRating float64 `json:"imdbRating,string"`
	ImdbVotes  string
	ImdbID     string
	Type       string
	Response   bool
}
