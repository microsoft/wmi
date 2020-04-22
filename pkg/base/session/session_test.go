package session

import "testing"

func TestGetSession(t *testing.T) {
	defer StopWMI()

	_, err := GetSession("temp", "localhost", "", "", "")
	if err != nil {
		t.Fatal(err)
	}
}
