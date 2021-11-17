package handlers

type Queue struct {
	elems []interface{}
	ttl   int64
}

func newQueue(ttl int64) *Queue {
	return &Queue{elems: make([]interface{}, 0), ttl: ttl}
}
