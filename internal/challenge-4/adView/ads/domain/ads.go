package Ad

import (
	"github.com/javier-tw/learning-go/internal/challenge-4/shared/domain/collection"
)

type Ads []Ad

func (ads Ads) ToPrimitives() []Primitives {
	return collection.ToPrimitives[Primitives](ads)
}

func CollectionFromPrimitives(primitives []Primitives) Ads {
	return collection.FromPrimitives(primitives, FromPrimitives)
}
