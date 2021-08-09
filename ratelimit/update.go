package ratelimit

import (
	"net/http"
	"strconv"
)

// UpdateBucket will grab the ratelimit information from the response headers and use that to update the existing ratelimit bucket.
func UpdateBucket(id string, res *http.Response) error {
	remaining, err := strconv.Atoi(res.Header.Get("X-RateLimit-Remaining"))
	if err != nil {
		return nil
	}

	reset, err := strconv.ParseFloat(res.Header.Get("X-RateLimit-Reset"), 64)
	if err != nil {
		return nil
	}

	limit, err := strconv.Atoi(res.Header.Get("X-RateLimit-Limit"))
	if err != nil {
		return nil
	}

	Buckets[id] = Bucket{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}

	return nil
}
