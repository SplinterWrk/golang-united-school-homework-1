package solution

import (
	"os"
	"testing"

	"github.com/kyokomi/emoji"
)

var helloWorld string

func TestMain(m *testing.M) {
	helloWorld = emoji.Sprint("Hello :world_map:!")
	os.Exit(m.Run())
}
