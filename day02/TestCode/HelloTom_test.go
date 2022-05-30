package HelloTom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
	//if output != expectOutput {
	//	t.Errorf("Test Fail!\nExpected %s do not match actual %s", expectOutput, output)
	//}
}
