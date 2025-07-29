package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PeersMap maps public keys to human-friendly names.
type PeersMap map[string]string

// Get returns the name for a given key or "UNKNOWN" if not found.
func (p PeersMap) Get(key string) string {
	if name, ok := p[key]; ok {
		return name
	}
	return "UNKNOWN"
}

// ReadConfig parses the WireGuard config file at path and extracts peer comments.
func ReadConfig(path string) (PeersMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config %q: %w", path, err)
	}
	defer file.Close()

	peers := make(PeersMap)
	lastComment := "unknown"
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "PublicKey") {
			eqIdx := strings.Index(line, "=")
			if eqIdx != -1 {
				key := strings.TrimSpace(line[eqIdx+1:])
				peers[key] = lastComment[2:]
				lastComment = "unknown"
			}
		}

		if strings.HasPrefix(line, "#") {
			lastComment = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return peers, nil
}
