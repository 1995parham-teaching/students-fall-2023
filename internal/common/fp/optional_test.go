package fp_test

import (
	"fmt"
	"testing"

	"github.com/1995parham-teaching/students-fall-2023/internal/common/fp"
)

func TestOptional(t *testing.T) {
	t.Parallel()

	cases := []any{
		"Parham",
		1373,
		"Elahe",
		1378,
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("convert %v to pinter", c), func(t *testing.T) {
			t.Parallel()

			if fp.Optional(c) == nil {
				t.Fatalf("failed to convert %v to pointer", c)
			}
		})
	}
}
