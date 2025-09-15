package analyzer

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func TestS3List(bucketName string) bool {
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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		bodyStr := string(body)

		// Check if we can list objects
		if resp.StatusCode == 200 && (strings.Contains(bodyStr, "<ListBucketResult") ||
			strings.Contains(bodyStr, "<Contents>") ||
			strings.Contains(bodyStr, "<?xml")) {
			return true
		}
	}

	return false
}

func TestS3Upload(bucketName string) bool {
	testFile := "test-upload-" + fmt.Sprintf("%d", time.Now().Unix()) + ".txt"
	testContent := "This is a test file for S3 misconfiguration detection"

	urls := []string{
		fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, testFile),
		fmt.Sprintf("https://s3.amazonaws.com/%s/%s", bucketName, testFile),
	}

	for _, url := range urls {
		req, err := http.NewRequest("PUT", url, strings.NewReader(testContent))
		if err != nil {
			continue
		}

		req.Header.Set("Content-Type", "text/plain")

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Check if upload was successful
		if resp.StatusCode == 200 || resp.StatusCode == 201 {
			// Try to clean up the test file
			deleteReq, _ := http.NewRequest("DELETE", url, nil)
			client.Do(deleteReq)
			return true
		}
	}

	return false
}

func TestS3Delete(bucketName string) bool {
	// Create a test file first
	testFile := "test-delete-" + fmt.Sprintf("%d", time.Now().Unix()) + ".txt"

	urls := []string{
		fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, testFile),
		fmt.Sprintf("https://s3.amazonaws.com/%s/%s", bucketName, testFile),
	}

	for _, url := range urls {
		// Try to delete (even if file doesn't exist, we check permissions)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			continue
		}

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// If we get 204 (No Content) or 200, delete permissions exist
		// If we get 404, the object doesn't exist but we might have delete permissions
		if resp.StatusCode == 204 || resp.StatusCode == 200 || resp.StatusCode == 404 {
			// Check the response for permission indicators
			body, _ := io.ReadAll(resp.Body)
			bodyStr := string(body)

			// If we don't get an AccessDenied error, we might have delete permissions
			if !strings.Contains(bodyStr, "AccessDenied") && !strings.Contains(bodyStr, "Forbidden") {
				return true
			}
		}
	}

	return false
}
