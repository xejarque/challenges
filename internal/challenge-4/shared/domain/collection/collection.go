package collection

type Entity[P any] interface {
	ToPrimitives() P
}

func ToPrimitives[P any, T Entity[P]](entities []T) []P {
	primitives := make([]P, len(entities))
	for i, entity := range entities {
		primitives[i] = entity.ToPrimitives()
	}
	return primitives
}

func FromPrimitives[P any, T Entity[P]](primitives []P, fromPrimitives func(P) T) []T {
	entities := make([]T, len(primitives))
	for i, adPrimitives := range primitives {
		entities[i] = fromPrimitives(adPrimitives)
	}
	return entities
}
