package descriptor

import (
	"fmt"

	maia "github.com/grpc-custom/maia/proto"
)

func (r *Registry) Component(typ maia.ComponentType) (*Component, error) {
	comp, ok := r.components[typ]
	if !ok {
		return nil, fmt.Errorf("no found %s", typ)
	}
	return comp, nil
}
