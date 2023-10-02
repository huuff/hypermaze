package hal

type Link struct {
  Href string `json:"href"`
}


type Links struct {
  Self Link `json:"self"`
}
