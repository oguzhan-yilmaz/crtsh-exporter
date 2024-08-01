package collector

import (
	"fmt"
	"testing"
)

var (
	domainName = "example.com"
	hosts      = []string{
		"x",
		"x.y",
		"x.y.z",
	}
	tests = func(h []string) map[string]string {
		x := make(map[string]string, len(h))
		for _, val := range h {
			key := fmt.Sprintf("%s.%s", val, domainName)
			x[key] = val
		}
		return x
	}(hosts)
)

func TestDomain(t *testing.T) {
	d, err := NewDomain(domainName)
	if err != nil {
		t.Fatal("Expected to create Domain successfully")
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := d.Hostname(name)
			if err != nil {
				t.Fatal("Expected valid fully-qualified name")
			}

			if got != want {
				t.Errorf("got: %s; want: %s", got, want)
			}
		})
	}
}
