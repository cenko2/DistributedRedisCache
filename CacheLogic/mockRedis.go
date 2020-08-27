package CacheLogic

type MockRedis struct {
	InstanceID string
}

func (m MockRedis) Get(key string) string {
	return m.InstanceID + key
}
func (m MockRedis) Insert(key string, val string, ttl int) {

}
func (m MockRedis) Keyexist(key string) bool {
	return false
}
