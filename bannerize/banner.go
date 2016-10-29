package bannerize

import (
	"os"

	"github.com/dimiro1/banner"
)

func init() {
	in, _ := os.Open("banner.txt")
	defer in.Close()
	banner.Init(os.Stdout, true, false, in)
}
