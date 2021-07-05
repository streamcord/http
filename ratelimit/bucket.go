package ratelimit

type Bucket struct {
	Limit     int
	Remaining int
	Reset     float64
}

var (
	Buckets map[string]Bucket
)

// Gets an existing bucket or returns nil if it doesn't exist
func GetBucket(id string) *Bucket {
	b, ok := Buckets[id]
	if !ok {
		return nil
	}

	return &b
}
