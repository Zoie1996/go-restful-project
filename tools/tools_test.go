package tools

import (
	"fmt"
	"testing"
	"time"
)

func Test_Fmt(t *testing.T) {
	// a := fmt.Sprintf("%v", "test")
	fmt.Println("123")
	tt := time.Now()
	a := fmt.Sprintf("%s_%d", "admin1", tt.Unix())
	fmt.Println(a)
}
