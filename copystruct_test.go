package rztool

import (
	"fmt"
	"testing"
	"time"
)

type A struct {
	Title   string
	Content string
	Create  string
}

type B struct {
	Title  string
	Create string
}

func Test_copy(t *testing.T) {
	a := new(A)
	a.Title = "titlte"
	a.Content = "content"
	a.Create = time.Now().Format("2005-01-02 15:04:05")

	b := new(B)

	CopyStruct(a, b)

	fmt.Println(*b)

}
