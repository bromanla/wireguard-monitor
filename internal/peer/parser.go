package peer

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

const (
	// BytesMultiplier doubles traffic to account for both directions
	BytesMultiplier = 2
)

// ReadPeers runs `wg show all dump` and parses peers with non-zero traffic.
func ReadPeers() ([]Peer, error) {
	// cmd := exec.Command("cat", "./mock/dump.std") // Only dev mode
	cmd := exec.Command("wg", "show", "all", "dump")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("wg dump failed: %w", err)
	}

	peers := make([]Peer, 0)
	scanner := bufio.NewScanner(bytes.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 9 {
			continue
		}

		time, _ := strconv.ParseInt(parts[5], 10, 64)
		rx, _ := strconv.ParseUint(parts[6], 10, 64)
		tx, _ := strconv.ParseUint(parts[7], 10, 64)

		peer := Peer{
			Key:  parts[1],
			Time: time,
			// Multiply by BytesMultiplier to calculate total traffic relative to VDS
			// Swap RX and TX to show traffic from client's perspective, not server's
			Rx: tx * BytesMultiplier, // TX from server = RX for client
			Tx: rx * BytesMultiplier, // RX from server = TX for client
		}

		if peer.Rx > 0 || peer.Tx > 0 {
			peers = append(peers, peer)
		}
	}

	// Sort peers by descending Tx
	sort.Slice(peers, func(i, j int) bool {
		return peers[i].Tx > peers[j].Tx
	})

	return peers, nil
}
