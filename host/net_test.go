package host

import (
	"os/exec"
	"strings"
	"testing"
)

func TestPrimaryIP(t *testing.T) {
	want := ipByShell()
	if want == "" {
		t.Fatal()
	}
	got, err := PrimaryIP()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func ipByShell() string {
	raw, err := exec.Command("hostname", "-i").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(raw))
}
