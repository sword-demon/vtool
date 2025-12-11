package mapx

import "github.com/sword-demon/vtool/internal/maps"

// Get 获取map中指定key的值
func Get[K comparable, V any](src map[K]V, key K) (V, bool) {
	val, exists := maps.Get(src, key)
	return val, exists
}

// Set 设置map中key的值为value
func Set[K comparable, V any](src map[K]V, key K, value V) {
	maps.Set(src, key, value)
}

// Delete 删除map中指定的key
func Delete[K comparable, V any](src map[K]V, key K) {
	maps.Delete(src, key)
}

// Has 检查map中是否包含指定的key
func Has[K comparable, V any](src map[K]V, key K) bool {
	return maps.Has(src, key)
}

// Keys 返回map中所有key的切片
func Keys[K comparable, V any](src map[K]V) []K {
	return maps.Keys(src)
}

// Values 返回map中所有value的切片
func Values[K comparable, V any](src map[K]V) []V {
	return maps.Values(src)
}
