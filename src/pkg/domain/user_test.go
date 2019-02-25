package domain

import (
	"testing"
)

func TestGetFullName(t *testing.T) {
	user := User{
		FirstName: "first",
		LastName:  "last",
	}
	res := user.GetFullName()

	if res != "first last" {
		t.Fatalf("Expected %v, but got %v", "first last", res)
	}
}

func TestIsEmpty(t *testing.T) {
	testcases := []struct {
		id        int64
		firstName string
		lastName  string
		expected  bool
	}{
		{1, "first", "last", false},
		{0, "first", "last", true},
		{1, "", "last", true},
		{1, "first", "", true},
	}

	for _, tc := range testcases {
		user := User{tc.id, tc.firstName, tc.lastName}
		if user.IsEmpty() != tc.expected {
			t.Fatalf("Expected %v, but got %v", tc.expected, user.IsEmpty())
		}
	}
}

func TestFindById(t *testing.T) {
	users := Users{
		User{1, "john", "brown"},
		User{2, "ben", "horowitz"},
	}

	testcases := []struct {
		id       int64
		expected User
	}{
		{1, User{1, "john", "brown"}},
		{2, User{2, "ben", "horowitz"}},
		{3, User{}},
	}

	for _, tc := range testcases {
		if users.FindById(tc.id) != tc.expected {
			t.Fatalf("Expected %v, but got %v", tc.expected, users.FindById(tc.id))
		}
	}

}
