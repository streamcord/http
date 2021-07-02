package ratelimit

import (
	"net/http"
	"strconv"
)

func UpdateBucket(id string, res *http.Response) error {
	remaining, err := strconv.ParseInt(res.Header.Get("X-RateLimit-Remaining"), 10, 8)
	if err != nil {
		return err
	}

	reset, err := strconv.ParseInt(res.Header.Get("X-RateLimit-Reset"), 10, 64)
	if err != nil {
		return err
	}

	limit, err := strconv.ParseInt(res.Header.Get("X-RateLimit-Limit"), 10, 8)
	if err != nil {
		return err
	}

	bucket := GetBucket(id)
	if bucket != nil {
		NewBucket(id, int(remaining), int(limit), reset)
		return nil
	}

	bucket.Limit = int(limit)
	bucket.Remaining = int(remaining)
	bucket.Reset = reset

	return nil
}
