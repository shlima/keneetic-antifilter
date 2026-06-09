package convertion

import (
	"keneetic-antifilter/internal/pkg/cidr"
)

type Line struct {
	Address *cidr.Address
	Comment string
}
