package belajar_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
	t.Log("ssssssssssssssssssssssssssssssssssssssssssssssssss")
	t.Logf("uuuuuuuuuuuuuuuuuuuuuuuuu")
}

func TestContextsss(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
	t.Log("ssssssssssssssssssssssssssssssssssssssssssssssssss")
	t.Logf("uuuuuuuuuuuuuuuuuuuuuuuuu")
}
