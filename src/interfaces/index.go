package interfaces

type Api struct {
	Files []File `json:"files"`
	Size  int    `json:"size"`
}

type File struct {
	Id   int    `json:"id"`
	Name string `json:"filename"`
}

type HttpCodeError struct {
	Err  string
	Code int
}
