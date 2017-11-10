package data

type Database interface {
	Create(string, map[string]interface{}) (string, error)
	Get(string, string) (map[string]string, error)
	List(string, string) (map[string]string, error)
	Update(string, string, map[string]interface{}) error
	Delete(string, string) error
	Flush(string) error
}