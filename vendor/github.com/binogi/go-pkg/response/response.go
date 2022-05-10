// Package response defines the structs for common responses that are to be used when outputting
// RESTful data, so that they always have a predictable format, such as paginated lists or errors
package response

type Error struct {
	Message string `json:"message"`
}

type PaginatedList struct {
	Results    []interface{} `json:"results"`
	TotalCount uint          `json:"totalCount"`
}
