package analyzer

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func TestBucketTakeover(bucketName string) bool {
	// First check if bucket exists
	if CheckBucketExistence(bucketName) {
		return false // Can't takeover existing bucket
	}

	// Test different regions where bucket might be claimed
	regions := []string{
		"us-east-1",
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"eu-central-1",
		"ap-southeast-1",
		"ap-northeast-1",
	}

	for _, region := range regions {
		url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucketName, region)

		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Check for specific error patterns that indicate takeover potential
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		bodyStr := string(body)

		// Look for error messages that indicate the bucket doesn't exist
		// and could potentially be claimed
		if strings.Contains(bodyStr, "NoSuchBucket") ||
			strings.Contains(bodyStr, "The specified bucket does not exist") ||
			resp.StatusCode == 404 {
			return true
		}
	}

	return false
}

func CheckBucketExistence(bucketName string) bool {
	urls := []string{
		fmt.Sprintf("https://%s.s3.amazonaws.com/", bucketName),
		fmt.Sprintf("https://s3.amazonaws.com/%s/", bucketName),
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// If bucket exists, we should get some response (200, 403, etc.)
		// If bucket doesn't exist, we typically get 404
		if resp.StatusCode != 404 {
			return true
		}

		// Check response body for specific error messages
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		bodyStr := string(body)
		// If we get NoSuchBucket error, the bucket definitely doesn't exist
		if strings.Contains(bodyStr, "NoSuchBucket") || strings.Contains(bodyStr, "The specified bucket does not exist") {
			return false
		}
	}

	return true // Default to assuming bucket exists if unclear
}
