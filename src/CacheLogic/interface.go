package CacheLogic

type CacheHandler interface {
	Get(string) string
	Insert(string, string, int)
	Keyexist(string) bool
}
