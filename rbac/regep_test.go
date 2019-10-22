package rbac

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	regp := regexp.MustCompile("[\\^]*")
	fmt.Println(regp.MatchString("abc"), regp.MatchString(``))
}
