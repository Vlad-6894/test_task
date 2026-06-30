package core_http_types

import "github.com/Vlad-6894/test_task/internal/core/domain"

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil
		return nil
	}

}
