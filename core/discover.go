package core

import (
	"fmt"
	"net"
	"strings"
)

// Discover the Session Endpoint of a domain
func Discover(domain string) (string, error) {
	// Format: https://${hostname}[:${port}]/.well-known/jmap
	_, srvs, err := net.LookupSRV("jmap", "tcp", domain)
	if err != nil {
		return "", err
	}
	if len(srvs) < 1 {
		return "", fmt.Errorf("no jmap srv found")
	}
	srv := srvs[0]
	endpoint := strings.Builder{}
	endpoint.WriteString("https://")
	endpoint.WriteString(strings.TrimSuffix(srv.Target, "."))
	if srv.Port > 0 {
		endpoint.WriteString(fmt.Sprintf(":%d", srv.Port))
	}
	endpoint.WriteString("/.well-known/jmap")
	return endpoint.String(), nil
}
