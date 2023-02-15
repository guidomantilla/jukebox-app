package encodingjson

type MarshalFunc func(v any) ([]byte, error)
type UnmarshalFunc func(data []byte, v any) error
