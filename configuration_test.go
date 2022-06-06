package main

import (
	"testing"
)

func TestStepShouldTakeAllGoctorCodes(t *testing.T) {
	step := Step{
		Run: `
			echo "goctor::MY_ERROR_1"
			echo "goctor::MY_ERROR_2"
		`,
	}

	_, goctorcodes, err := step.execute()

	if err == nil {
		t.Fatal("I can't see a error")
	}

	if len(goctorcodes) == 0 {
		t.Fatal("Without goctorcodes")
	}
}

func TestStepShouldTakeNot0ExitCode(t *testing.T) {
	step := Step{
		Run: "exit 1",
	}

	exit, _, err := step.execute()

	if err == nil {
		t.Fatal("I can't see a error")
	}

	if exit != 1 {
		t.Fatal("Wrong exist code, or without one")
	}
}

func TestStepShouldTakeAllGoctorcodeAndExitCode(t *testing.T) {
	step := Step{
		Run: `
			echo "goctor::ONE"
			echo "goctor::TWO"
			exit 1
		`,
	}

	exit, goctorcodes, err := step.execute()

	if err == nil {
		t.Fatal("I can't see a error")
	}

	if exit != 1 {
		t.Fatal("Wrong exist code, or without one")
	}

	if len(goctorcodes) != 2 {
		t.Fatal("Need to have 2")
	}
}
