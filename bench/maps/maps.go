package maps

var (
	cache = map[string]string{}
	cachePtr = map[string]string{}
)

func compute(id string) string {
	return id
}

func computePtr(id string) string {
	return id
}

func NormalGet(id string) string {
	if cache[id] != "" {
		return cache[id]
	}

	cache[id] = compute(id)
	return cache[id]
}

func PointerGet(id string) string {
	val := cachePtr[id]
	if val != "" {
		return val
	}

	v := computePtr(id)
	cachePtr[id] = v
	return v
}

func Clear() {
	cache = map[string]string{}
	cachePtr = map[string]string{}
}