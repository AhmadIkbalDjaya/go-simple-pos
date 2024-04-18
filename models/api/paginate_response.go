package api

type MetaPaginate struct {
	Page      int `json:"page"`
	Perpage   int `json:"perpage"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}

type PaginateResponse struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    any          `json:"data"`
	Meta    MetaPaginate `json:"meta"`
}