package ratelimit

type Bucket struct {
	Limit     int
	Remaining int
	Reset     int64
}

// NewBucket creates a new bucket with the provided id and item count.
func NewBucket(id string, limit int, remaining int, reset int64) Bucket {
	b := Bucket{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}

	Buckets[id] = b
	return b
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
