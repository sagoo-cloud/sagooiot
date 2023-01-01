package model

// Command 命令
type Command struct {
	Name       string      `json:"name"`
	Argument   bool        `json:"argument"`
	Directives []Directive `json:"directives"`
}

// Directive 指令
type Directive struct {
	Point string  `json:"point"`
	Value float64 `json:"value"`
	Delay int64   `json:"delay"`

	//使用表达式
	Expression string `json:"expression,omitempty"`
}
