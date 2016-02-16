package far

import "testing"

var fixturePath = "../fixtures/cat.rb"

func TestFileExists(t *testing.T) {
	fileExists := FileExists(fixturePath)

	if !fileExists {
		t.Error("Could not find", fixturePath)
	}
}

func TestFindAndReplace(t *testing.T) {
	timesReplaced, err := FindAndReplace(fixturePath, "Cat", "Dog")

	if err != nil {
		t.Error("Did not expect error, value:", err)
	}

	if timesReplaced != 6 {
		t.Errorf("Expected Cat to be replaced 6 times, was replaced %v times", timesReplaced)
	}

	resetFixture()
}

func resetFixture() {
	FindAndReplace(fixturePath, "Dog", "Cat")
}
