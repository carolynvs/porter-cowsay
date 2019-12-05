package cowsay

import (
	"fmt"

	"get.porter.sh/mixin/cowsay/pkg"
)

func (m *Mixin) PrintVersion() {
	fmt.Fprintf(m.Out, "cowsay mixin %s (%s)\n", pkg.Version, pkg.Commit)
}
