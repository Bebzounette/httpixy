package httpixy

import "strings"

// GenerateURL should generate for potential vulnerability URLs
func GenerateURL(u string) []string {
	var url []string

	if !strings.HasSuffix(u, "/") {
		u += "/"
	}

	url = append(url, u)
	
	return url
}
