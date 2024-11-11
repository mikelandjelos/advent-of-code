package main

import "testing"

func TestIsNice(t *testing.T) {
	testedString := "ugknbfddgicrmopn"
	if isNice := IsNice(testedString); !isNice {
		t.Error(testedString, isNice)
	}

	testedString = "aaa"
	if isNice := IsNice(testedString); !isNice {
		t.Error(testedString, isNice)
	}

	testedString = "jchzalrnumimnmhp"
	if isNice := IsNice(testedString); isNice {
		t.Error(testedString, isNice)
	}

	testedString = "haegwjzuvuyypxyu"
	if isNice := IsNice(testedString); isNice {
		t.Error(testedString, isNice)
	}

	testedString = "dvszwmarrgswjxmb"
	if isNice := IsNice(testedString); isNice {
		t.Error(testedString, isNice)
	}
}

func TestIsNiceNewRules(t *testing.T) {
	testedString := "qjhvhtzxzqqjkmpb"
	if isNice := IsNiceNewRules(testedString); !isNice {
		t.Error(testedString, isNice)
	}

	testedString = "xxyxx"
	if isNice := IsNiceNewRules(testedString); !isNice {
		t.Error(testedString, isNice)
	}

	testedString = "uurcxstgmygtbstg"
	if isNice := IsNiceNewRules(testedString); isNice {
		t.Error(testedString, isNice)
	}

	testedString = "ieodomkazucvgmuy"
	if isNice := IsNiceNewRules(testedString); isNice {
		t.Error(testedString, isNice)
	}
}
