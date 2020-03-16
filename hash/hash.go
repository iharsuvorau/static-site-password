package main

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("compareHexStrings", js.FuncOf(CompareHexStringsJS))
	js.Global().Set("calculateHexString", js.FuncOf(CalculateHexStringJS))

	<-c
}

func CompareHexStringsJS(this js.Value, args []js.Value) interface{} {
	s := args[0].String()
	outHex := args[1].String()

	calculated, err := CalculateHexString(s)
	if err != nil || calculated != outHex {
		return false
	}
	return true
}

func CalculateHexStringJS(this js.Value, args []js.Value) interface{} {
	s := args[0].String()

	h := sha512.New512_256()
	_, err := io.WriteString(h, s)
	if err != nil {
		return ""
	}

	calculated := hex.EncodeToString(h.Sum(nil))
	return calculated
}

func CompareHexStrings(s, outHex string) bool {
	calculated, err := CalculateHexString(s)
	if err != nil || calculated != outHex {
		return false
	}
	return true
}

func CalculateHexString(s string) (string, error) {
	h := sha512.New512_256()
	_, err := io.WriteString(h, s)
	if err != nil {
		return "", err
	}

	calculated := hex.EncodeToString(h.Sum(nil))
	return calculated, nil
}
