package entities

type Entity interface {
	Map() map[string]interface{}
	ParseMap(map[string]interface{}) error
}
