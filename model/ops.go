package model

type Operation struct {
	Name string `json:"name"`
}

type UnaryOp struct {
	Operation
	Operand int `json:"x" form:"x" binding:"exists"`
}

type BinaryOp struct {
	Operation
	Left  int `json:"x" form:"x" binding:"exists"`
	Right int `json:"y" form:"y" binding:"exists"`
}

type OpResult struct {
	Operation interface{} `json:"operation"`
	Success   bool        `json:"success"`
	Result    interface{} `json:"result"`
}
