package testsupport

import "testing"

func AssertDifferentString(t *testing.T, notExpected string, actual string) {
	if notExpected == actual {
		t.Logf("The actual value [%s] equals the NOT expected [%s]", actual, notExpected)
		t.Fail()
	}
}

func AssertEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Logf("The actual value [%v] differes from the expected [%v]", actual, expected)
		t.Fail()
	}
}

func AssertEqualString(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Logf("The actual value [%v] differes from the expected [%v]", actual, expected)
		t.Fail()
	}
}
