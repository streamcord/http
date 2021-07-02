package ratelimit

type Bucket struct {
	ID        string
	Limit     int
	Remaining int
	Reset     int64
}

// NewBucket creates a new bucket with the provided id and item count.
func NewBucket(id string, limit int, remaining int, reset int64) *Bucket {
	b := Bucket{
		ID:        id,
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}
	Buckets = append(Buckets, &b)

	return &b
}

var (
	Buckets []*Bucket
)

// Gets an existing bucket or returns nil if it doesn't exist
func GetBucket(id string) *Bucket {
	for _, b := range Buckets {
		if b.ID == id {
			return b
		}
	}

	return nil
}
