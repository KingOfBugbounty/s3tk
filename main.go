package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"s3scanJAAAH/pkg/analyzer"
	"s3scanJAAAH/pkg/banner"
	"s3scanJAAAH/pkg/report"
)

func main() {
	banner.PrintBanner()

	fmt.Printf("%s[INFO]%s Reading S3 bucket URLs from stdin...\n", "\033[94m", banner.Reset)
	fmt.Printf("%s[INFO]%s Supported formats: bucket-name, https://bucket.s3.amazonaws.com, s3://bucket-name\n\n", "\033[94m", banner.Reset)

	var buckets []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			bucketName := analyzer.ExtractBucketName(line)
			if bucketName != "" {
				buckets = append(buckets, bucketName)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%s[ERROR]%s Error reading from stdin: %v\n", "\033[91m", banner.Reset, err)
		os.Exit(1)
	}

	if len(buckets) == 0 {
		fmt.Printf("%s[ERROR]%s No valid S3 bucket URLs found in input\n", "\033[91m", banner.Reset)
		os.Exit(1)
	}

	fmt.Printf("%s[INFO]%s Found %d bucket(s) to scan\n\n", "\033[94m", banner.Reset, len(buckets))

	var results []analyzer.S3Test

	for _, bucket := range buckets {
		result := analyzer.ScanS3Bucket(bucket)
		results = append(results, result)
		fmt.Println()
	}

	report.PrintResults(results)
}
