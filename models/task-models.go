package models

type TaskOneRequest struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
}

type TaskOneResponse struct {
	Result []int `json:"output"`
}

type TaskTwoRequest struct {
	Nums []int `json:"nums"`
}

type TaskTwoResponse struct {
	Result [][]int `json:"output"`
}

type TaskThreeRequest struct {
	S     string   `json:"s"`
	Words []string `json:"words"`
}

type TaskThreeResponse struct {
	Result []int `json:"output"`
}
