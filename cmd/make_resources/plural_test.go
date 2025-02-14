package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluralGoCondition(t *testing.T) {
	prData := `e = 0 and i != 0 and i % 1000000 = 0 and v = 0 or e != 0..5 @integer 1000000, 1c6, 2c6, 3c6, 4c6, 5c6, 6c6, … @decimal 1.0000001c6, 1.1c6, 2.0000001c6, 2.1c6, 3.0000001c6, 3.1c6, …`
	pr := &PluralRule{
		Count: "many",
		Rule:  prData,
	}

	const expected = "intEqualsAny(ops.C, 0) && !intEqualsAny(ops.I, 0) && intEqualsAny(ops.I % 1000000, 0) && intEqualsAny(ops.V, 0) ||\n!intInRange(ops.C, 0, 5)"
	assert.Equal(t, expected, pr.GoCondition())
}
