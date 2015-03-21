package even //<1>

import "testing" //<2>

func TestEven(t *testing.T) { //<3>
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}
