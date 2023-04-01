package adapters

type IConnection interface {
	Query(query string, args []([]interface{})) error
}
