package membuat_context

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
}
