package dog

import (
	"strings"
)

func WhenDogGrowsUp(puppy_name string) string{
	return "when the puppy grows up, it says:" + strings.ToUpper(puppy_name)
}