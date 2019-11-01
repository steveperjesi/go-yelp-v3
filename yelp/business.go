package yelp

type Business struct {
  Rating       float64    `json:"rating"`
  Price        string     `json:"price"`
  Phone        string     `json:"phone"`
  Id           string     `json:"id"`
  Alias        string     `json:"alias"`
  IsClosed     bool       `json:"is_closed"`
  Categories   []Category `json:"categories"`
  ReviewCount  int        `json:"review_count"`
  Name         string     `json:"name"`
  Url          string     `json:"url"`
  Coordinates  Coordinate `json:"coordinates"`
  ImageUrl     string     `json:"image_url"`
  Location     Location   `json:"location"`
  Distance     float64    `json:"distance"`
  Transactions []string   `json:"transactions"`
}
