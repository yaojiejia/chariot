package Reader

// Reader is an interface that have Read() and Connect() methods
type Reader interface {
	Read() string
	ReadAndCache() error
	Get() (string, error)
}
