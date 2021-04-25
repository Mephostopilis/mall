package types

// SearchReq search request options
type SearchReq struct {
	Q    string
	Size int32
	From int32
	Term string
	Ids  []string
}
