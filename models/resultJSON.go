package models

type Result struct {
	Status string      `json:"status"`
	Code   int         `json: "code"`
	Data   interface{} `json: "data"`
	Meta   struct {
		Page       int   `json: "page"`
		Size       int   `json: "size"`
		Count      int   `json: "count"`
		TotalPages int64 `json: "totalPages`
		TotalData  int   `json: "totalData"`
	} `json :"meta"`
}
