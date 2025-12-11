package maps

// Get 获取map中指定key的值，如果key不存在返回零值
// 返回值表示key是否存在于map中
func Get[K comparable, V any](src map[K]V, key K) (V, bool) {
	val, exists := src[key]
	return val, exists
}

// Set 设置map中key的值为value
func Set[K comparable, V any](src map[K]V, key K, value V) {
	src[key] = value
}

// Delete 删除map中指定的key
func Delete[K comparable, V any](src map[K]V, key K) {
	delete(src, key)
}

// Has 检查map中是否包含指定的key
func Has[K comparable, V any](src map[K]V, key K) bool {
	_, exists := src[key]
	return exists
}

// Keys 返回map中所有key的切片
func Keys[K comparable, V any](src map[K]V) []K {
	keys := make([]K, 0, len(src))
	for k := range src {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回map中所有value的切片
func Values[K comparable, V any](src map[K]V) []V {
	values := make([]V, 0, len(src))
	for _, v := range src {
		values = append(values, v)
	}
	return values
}
