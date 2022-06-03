package main

import (
	"testing"
)

func TestExamShouldTakeAllGoctorCodes(t *testing.T) {
	exam := Exam{
		Run: `
			echo "goctor::MY_ERROR_1"
			echo "goctor::MY_ERROR_2"
		`,
	}

	_, goctorcodes, err := exam.execute()

	if err == nil {
		t.Fatal("I can't see a error")
	}

	if len(goctorcodes) == 0 {
		t.Fatal("Without goctorcodes")
	}
}

func TestExamShouldTakeNot0ExitCode(t *testing.T) {
	exam := Exam{
		Run: "exit 1",
	}

	exit, _, err := exam.execute()

	if err == nil {
		t.Fatal("I can't see a error")
	}

	if exit != 1 {
		t.Fatal("Wrong exist code, or without one")
	}
}

func TestExamShouldTakeAllGoctorcodeAndExitCode(t *testing.T) {
	exam := Exam{
		Run: `
			echo "goctor::ONE"
			echo "goctor::TWO"
			exit 1
		`,
	}

	exit, goctorcodes, err := exam.execute()

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
