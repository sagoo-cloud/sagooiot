package jsinterpreter

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"strings"
)

// RunScript 运行js脚本
func RunScript(jsonStr string, jsCode string) (string, error) {
	vm := otto.New()
	if _, err := vm.Run(jsCode); err != nil {
		return "", fmt.Errorf("failed to run JavaScript code: %v", err)
	}
	value, err := vm.Call("parse", nil, jsonStr)
	if err != nil {
		return "", fmt.Errorf("failed to call: %v", err)
	}
	data, err := value.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("failed to MarshalJSON: %v", err)
	}
	return strings.TrimPrefix(strings.TrimSuffix(string(data), "\""), "\""), nil
}
