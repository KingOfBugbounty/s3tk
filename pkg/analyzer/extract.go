package analyzer

import "strings"

func ExtractBucketName(s3url string) string {
	s3url = strings.TrimSpace(s3url)

	// Handle different S3 URL formats
	if strings.HasPrefix(s3url, "https://") {
		s3url = strings.TrimPrefix(s3url, "https://")
	} else if strings.HasPrefix(s3url, "http://") {
		s3url = strings.TrimPrefix(s3url, "http://")
	}

	// Extract bucket name from different formats
	if strings.Contains(s3url, ".s3.") {
		// Format: bucket-name.s3.region.amazonaws.com
		return strings.Split(s3url, ".")[0]
	} else if strings.HasPrefix(s3url, "s3.") {
		// Format: s3.region.amazonaws.com/bucket-name
		parts := strings.Split(s3url, "/")
		if len(parts) > 1 {
			return parts[1]
		}
	}

	// If it's just the bucket name
	if !strings.Contains(s3url, "/") && !strings.Contains(s3url, ".") {
		return s3url
	}

	return s3url
}
