package locations

type Type string

const (
	BODY Type = "body"
	FORM Type = "formData"
	HEADER Type = "header"
	PATH Type = "path"
	QUERY Type = "query"
)
