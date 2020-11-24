package common

import (
	"encoding/gob"
)

func init() {
	gob.Register(&Address{})
	gob.Register(&Hash{})
	gob.Register(&[]*Hash{})
}
