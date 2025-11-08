package helper

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Daci1/url-shortener-atad/internal/types"
)

func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

func IsPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	privateBlocks := []string{
		"10.0.0.0/8",
		"100.64.0.0/10", // shared address space
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16", // APIPA
		"::1/128",
		"fc00::/7",
		"fe80::/10",
	}

	for _, cidr := range privateBlocks {
		_, block, _ := net.ParseCIDR(cidr)
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func ExtractClientIP(req *http.Request) string {

	if xff := req.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		for _, p := range parts {
			ipStr := strings.TrimSpace(p)
			if ipStr == "" {
				continue
			}
			ip := net.ParseIP(ipStr)
			if ip == nil {
				continue
			}
			if !IsPrivateIP(ip) {
				return ip.String()
			}
			// if it's private, continue to see if there's a public one later in the list
		}
		// none public — fallthrough to other headers
	}

	// 2) X-Real-IP
	if xr := req.Header.Get("X-Real-Ip"); xr != "" {
		ip := net.ParseIP(strings.TrimSpace(xr))
		if ip != nil && !IsPrivateIP(ip) {
			return ip.String()
		}
	}

	// 3) RemoteAddr fallback (may include port)
	if ra := req.RemoteAddr; ra != "" {
		host, _, err := net.SplitHostPort(ra)
		if err == nil {
			ip := net.ParseIP(host)
			if ip != nil {
				return ip.String()
			}
		} else {
			// Sometimes RemoteAddr has no port
			ip := net.ParseIP(ra)
			if ip != nil {
				return ip.String()
			}
		}
	}

	return ""
}

func GetGeoData(ip string) (*types.GeoData, error) {
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)

	client := http.Client{
		Timeout: 5 * time.Second, // prevent long hangs
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch geolocation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response from ipapi: %s", resp.Status)
	}

	var data types.GeoData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode geolocation JSON: %w", err)
	}

	return &data, nil
}
