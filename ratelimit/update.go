package ratelimit

import (
	"net/http"
	"strconv"
)

func UpdateBucket(id string, res *http.Response) error {
	remaining, err := strconv.Atoi(res.Header.Get("X-RateLimit-Remaining"))
	if err != nil {
		return err
	}

	reset, err := strconv.ParseFloat(res.Header.Get("X-RateLimit-Reset"), 64)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(res.Header.Get("X-RateLimit-Limit"))
	if err != nil {
		return err
	}

	Buckets[id] = Bucket{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}

	return nil
}
