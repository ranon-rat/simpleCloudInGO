package stuff

type Api struct {
	File []File `json:"Files"`
	Size int    `json:"size"`
}
type File struct {
	Id   int    `json:"id"`
	Name string `json:"filename"`
}
type HttpCodeError struct {
	Err  string
	Code int
}
