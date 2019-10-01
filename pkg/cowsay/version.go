package cowsay

import (
	"fmt"

	"github.com/deislabs/porter-cowsay/pkg"
)

func (m *Mixin) PrintVersion() {
	fmt.Fprintf(m.Out, "cowsay mixin %s (%s)\n", pkg.Version, pkg.Commit)
}
