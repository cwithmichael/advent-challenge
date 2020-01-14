package main

import "testing"

func TestHasPairOfEqualAdjacentDigits(t *testing.T) {
	res := hasPairOfEqualAdjacentDigits("111111")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
	}

	res = hasPairOfEqualAdjacentDigits("123789")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = hasPairOfEqualAdjacentDigits("121111")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}
}

func TestIsAllDigitsIncreasingOrSame(t *testing.T) {
	res := isAllDigitsIncreasingOrSame("111111")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
	}

	res = isAllDigitsIncreasingOrSame("223450")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isAllDigitsIncreasingOrSame("842599")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isAllDigitsIncreasingOrSame("776677")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isAllDigitsIncreasingOrSame("775599")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isAllDigitsIncreasingOrSame("121250")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}
}

func TestIsWithinRange(t *testing.T) {
	res := isWithinRange("100")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isWithinRange("199999393993")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, false)
	}

	res = isWithinRange("502397")
	if res != true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
	}
}

func TestCheckPassword(t *testing.T) {
	res := CheckPassword("111111")
	if res == true {
		t.Errorf("Incorrect result, got: %t, want: %t.", res, true)
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
