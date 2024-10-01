package db

// DB is an interface that have Connect(), Read() and Write() methods
type DB interface {
	Connect() error
	Read() string
	Write() error
}
