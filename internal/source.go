package sources

type source interface {
	Sync() error
}
