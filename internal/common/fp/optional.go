package fp

func Optional[T any](v T) *T {
	return &v
}
