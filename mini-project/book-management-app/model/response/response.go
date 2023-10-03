package response

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
