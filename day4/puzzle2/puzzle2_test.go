package puzzle2

import "testing"

func TestHasSinglePairOfEqualAdjacentDigits(t *testing.T) {
	res := hasSinglePairOfEqualAdjacentDigits("112233")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
	}

	res = hasSinglePairOfEqualAdjacentDigits("123444")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = hasSinglePairOfEqualAdjacentDigits("111122")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
	}

	res = hasSinglePairOfEqualAdjacentDigits("777778")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

}

func TestCheckPassword(t *testing.T) {
	res := CheckPassword("111111")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = CheckPassword("223450")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = CheckPassword("123789")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}
}
