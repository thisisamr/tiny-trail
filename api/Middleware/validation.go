package MiddleWare

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"time"
)

// ValidateURL checks if a URL is well-formed and not pointing to localhost or private IPs
func ValidateURL(input string) error {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	// Ensure the scheme is HTTP or HTTPS
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported URL scheme: %s", u.Scheme)
	}

	// Check if the host is not localhost or a private IP
	// Regex to match disallowed domains, IP addresses, and subdomains
	disallowedPattern := `(^|\.)localhost$|localhost\.|^127\.0\.0\.1$|^0\.0\.0\.0$`

	host := u.Hostname()

	// Check if the host is a disallowed domain, IP, or private IP
	if regexp.MustCompile(disallowedPattern).MatchString(host) || isPrivateIP(host) {
		return fmt.Errorf("URL points to a disallowed host: %s", host)
	}
	return nil
}

// isPrivateIP checks if the given hostname is a private IP
func isPrivateIP(host string) bool {
	ip := net.ParseIP(host)
	if ip == nil {
		return false
	}

	// Check for private IP ranges
	privateIPBlocks := []*net.IPNet{
		{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)},
		{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)},
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// ValidateExpiryTime checks if the expiry time is in the future and within a reasonable range
func ValidateExpiryTime(expiry time.Duration) error {
	maxexpire := time.Now().Add(24 * time.Hour)

	// Optional: Set a maximum expiry duration (e.g., 1 year)
	if time.Now().Add(expiry).After(maxexpire) {
		return fmt.Errorf("expiry time is too far in the future")
	}

	return nil
}
