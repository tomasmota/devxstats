package sources

type source interface {
	Fetch() error
}
