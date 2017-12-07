package data

type Database interface {
	Create(string, ...interface{}) (string, error) // params: table, data (implementation dependent); returns primary key, error
	Get(string, string) (map[string]interface{}, error)  // params: table, key; returns data, error
	List(string, ...interface{}) ([]map[string]interface{}, error) // params: table, filters (implementation dependent); returns a list of data, error
	Update(string, string, ...interface{}) error // params: table, primary key, data (implementation dependent); returns error
	Delete(string, string) error // params: table, key; returns error 
	Flush(string) error // params: table; returns error
}

type SecretStore interface {
	Create(string, map[string]interface{}) error
	Get(string) error
	List() error
	Update(string, map[string]interface{}) error
	Delete(string) error
}

type LogCollector interface {
	GetLogs(string) error
}

type StorageDriver interface {
	Allocate() error 
	Deallocate() error
}