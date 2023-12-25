package tunnel

import (
	"tunnel"
)

func QueryMode() string {
	return tunnel.Mode().String()
}
