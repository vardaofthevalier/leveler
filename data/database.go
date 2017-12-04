package data

type Database interface {
	Create(string, map[string]interface{}, string) (string, error)
	Get(string, string) (string, error)
	List(string, string) ([]string, error)
	Update(string, string, string) error
	Delete(string, string) error
	Flush(string) error
}

type SecretStore interface {
	Create(string, map[string]interface{}) error
	Get(string) error
	List() error
	Update(string, map[string]interface{}) error
	Delete(string) error
}