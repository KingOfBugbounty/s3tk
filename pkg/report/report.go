package report

import (
	"fmt"
	"s3scanJAAAH/pkg/analyzer"
	"s3scanJAAAH/pkg/banner"
)

func PrintResults(results []analyzer.S3Test) {
	fmt.Printf("\n%s╔══════════════════════════════════════════════════════════════╗%s\n", "\033[93m", banner.Reset)
	fmt.Printf("%s║                        SCAN RESULTS                          ║%s\n", "\033[93m", banner.Reset)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════╝%s\n\n", "\033[93m", banner.Reset)

	vulnerableBuckets := 0
	takeoverBuckets := 0

	for _, result := range results {
		if result.CanList || result.CanUpload || result.CanDelete || result.CanTakeover {
			vulnerableBuckets++

			if result.CanTakeover {
				takeoverBuckets++
				fmt.Printf("%s[TAKEOVER POSSIBLE]%s %s\n", "\033[95m", banner.Reset, result.BucketName)
				fmt.Printf("  └─ %s[TAKEOVER]%s Bucket doesn't exist - can be claimed for subdomain takeover\n", "\033[35m", banner.Reset)
				fmt.Println()
			} else {
				fmt.Printf("%s[MISCONFIGURED]%s %s\n", "\033[91m", banner.Reset, result.BucketName)

				if result.CanList {
					fmt.Printf("  └─ %s[LIST]%s Public read access - can enumerate bucket contents\n", "\033[31m", banner.Reset)
				}
				if result.CanUpload {
					fmt.Printf("  └─ %s[UPLOAD]%s Public write access - can upload malicious files\n", "\033[31m", banner.Reset)
				}
				if result.CanDelete {
					fmt.Printf("  └─ %s[DELETE]%s Public delete access - can remove objects\n", "\033[31m", banner.Reset)
				}
				fmt.Println()
			}
		} else {
			if result.BucketExists {
				fmt.Printf("%s[SECURE]%s %s\n", "\033[92m", banner.Reset, result.BucketName)
			} else {
				fmt.Printf("%s[NOT FOUND]%s %s (bucket doesn't exist but not exploitable)\n", "\033[93m", banner.Reset, result.BucketName)
			}
		}
	}

	fmt.Printf("\n%s╔══════════════════════════════════════════════════════════════╗%s\n", "\033[96m", banner.Reset)
	fmt.Printf("%s║                          SUMMARY                             ║%s\n", "\033[96m", banner.Reset)
	fmt.Printf("%s╚══════════════════════════════════════════════════════════════╝%s\n", "\033[96m", banner.Reset)
	fmt.Printf("Total buckets scanned: %d\n", len(results))
	fmt.Printf("Vulnerable buckets found: %s%d%s\n", "\033[91m", vulnerableBuckets, banner.Reset)
	fmt.Printf("Takeover opportunities: %s%d%s\n", "\033[95m", takeoverBuckets, banner.Reset)
	fmt.Printf("Secure buckets: %s%d%s\n", "\033[92m", len(results)-vulnerableBuckets, banner.Reset)
}
