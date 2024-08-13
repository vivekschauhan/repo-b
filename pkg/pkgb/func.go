package pkgb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vivekschauhan/repo-a/pkg/pkga"
)

func FuncB() {
	pkga.FuncA()
	fmt.Printf("func b: %s\n", uuid.NewString())
}
