package intro

import "testing"

// POSITIVE TEST
func TestCube_Volume_Success(t *testing.T) {
	// GIVE -> WHEN -> THEN
	cubeMockup := Cube{Side: 2}
	exptected := float64(8)
	actual, err := cubeMockup.Volume()
	if exptected != actual {
		t.Errorf("The result should be %2.f", actual)
	}
	if err != nil {
		t.Errorf("Volume result error")
	}
}

// NEGATIVE TEST
func TestCube_Volume_Failed(t *testing.T) {
	cubeMockup := Cube{Side: -1}
	expected := float64(0)
	actual, err := cubeMockup.Volume()

	if err == nil {
		t.Error("Should return error when side number negative")
	}

	if actual != expected {
		t.Error("Should return 0 float")
	}
}

func TestCube_Area_Success(t *testing.T) {
	cubeMockup := Cube{2}
	expected := float64(24)
	actual, err := cubeMockup.Area()
	if actual != expected {
		t.Errorf("The result should be %2.f", actual)
	}
	if err != nil {
		t.Error("should return nil error")
	}
}

func TestCube_Area_Failed(t *testing.T) {
	cubeMockup := Cube{-1}
	expected := float64(0)
	actual, err := cubeMockup.Area()
	if actual != expected {
		t.Error("Should return error when side number negative")
	}
	if err == nil {
		t.Error("Should return 0 float")
	}
}

/**
method volume :
side : 2 (contoh nilai sidenya)
SUCCESS : megirimkan 2 value yaitu
1 => haril dari perpangkatan (2 * 2 * 2) = 8
2 => nil
FAILED: mengirimkan 2 value
1 => 0 float
2 => error
*/
