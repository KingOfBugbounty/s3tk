package analyzer

import (
	"fmt"
	"s3scanJAAAH/pkg/banner"
)

type S3Test struct {
	BucketName   string
	CanList      bool
	CanUpload    bool
	CanDelete    bool
	CanTakeover  bool
	BucketExists bool
}

func ScanS3Bucket(bucketName string) S3Test {
	result := S3Test{BucketName: bucketName}

	fmt.Printf("[*] Scanning bucket: %s\n", bucketName)

	// Check if bucket exists first
	fmt.Print("  [+] Checking bucket existence... ")
	result.BucketExists = CheckBucketExistence(bucketName)
	if result.BucketExists {
		fmt.Printf("%s[EXISTS]%s\n", "\033[92m", banner.Reset)
	} else {
		fmt.Printf("%s[NOT FOUND]%s\n", "\033[93m", banner.Reset)
	}

	// Test for bucket takeover if bucket doesn't exist
	if !result.BucketExists {
		fmt.Print("  [+] Testing TAKEOVER potential... ")
		result.CanTakeover = TestBucketTakeover(bucketName)
		if result.CanTakeover {
			fmt.Printf("%s[VULNERABLE]%s\n", "\033[91m", banner.Reset)
		} else {
			fmt.Printf("%s[SECURE]%s\n", "\033[92m", banner.Reset)
		}
	}

	// Only test other permissions if bucket exists
	if result.BucketExists {
		// Test LIST permissions
		fmt.Print("  [+] Testing LIST permissions... ")
		result.CanList = TestS3List(bucketName)
		if result.CanList {
			fmt.Printf("%s[VULNERABLE]%s\n", "\033[91m", banner.Reset)
		} else {
			fmt.Printf("%s[SECURE]%s\n", "\033[92m", banner.Reset)
		}

		// Test UPLOAD permissions
		fmt.Print("  [+] Testing UPLOAD permissions... ")
		result.CanUpload = TestS3Upload(bucketName)
		if result.CanUpload {
			fmt.Printf("%s[VULNERABLE]%s\n", "\033[91m", banner.Reset)
		} else {
			fmt.Printf("%s[SECURE]%s\n", "\033[92m", banner.Reset)
		}

		// Test DELETE permissions
		fmt.Print("  [+] Testing DELETE permissions... ")
		result.CanDelete = TestS3Delete(bucketName)
		if result.CanDelete {
			fmt.Printf("%s[VULNERABLE]%s\n", "\033[91m", banner.Reset)
		} else {
			fmt.Printf("%s[SECURE]%s\n", "\033[92m", banner.Reset)
		}
	}

	return result
}
