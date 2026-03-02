package utils

import (
	"fmt"
	"time"
)

func Timestamp() string {
	ts := fmt.Sprintf("%d", time.Now().UnixNano())
	// Крошечная пауза, чтобы следующий вызов UnixNano гарантированно дал другое число
	time.Sleep(1 * time.Microsecond)
	return ts
}
