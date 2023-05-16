package plan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayPlan_Run(t *testing.T) {
	dp := NewDayPlan()
	dp.SetNumsPerDay(20)
	dp.SelectGlossary("IELTS")
	assert.NoError(t, dp.Run())
	assert.Equal(t, 20, len(dp.Words))
	fmt.Println(dp.Words)
}
