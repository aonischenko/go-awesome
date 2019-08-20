package model

type Operation struct {
	Name string `json:"name"`
}

type UnaryOp struct {
	Operation
	Operand int `json:"x" mapstructure:"x"`
}

type BinaryOp struct {
	Operation
	Left  int `json:"x" mapstructure:"x"`
	Right int `json:"y" mapstructure:"y"`
}

type OpResult struct {
	Operation interface{} `json:"operation"`
	Success   bool        `json:"success"`
	Result    interface{} `json:"result"`
}
