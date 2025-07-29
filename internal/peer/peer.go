package peer

import (
	"fmt"
	"strings"
	"time"
	"wireguard-monitor/internal/common"
)

const (
	OnlineTimeout = 5 * time.Minute
)

// Peer represents one WireGuard client with traffic stats.
type Peer struct {
	Key  string
	Time int64  // last handshake timestamp
	Rx   uint64 //bytes received by client
	Tx   uint64 // bytes sent by client
}

// IsOnline returns true if the last handshake is within the online timeout.
func (peer Peer) IsOnline() bool {
	if peer.Time <= 0 {
		return false
	}

	return time.Since(time.Unix(peer.Time, 0)) < OnlineTimeout
}

// Format renders the peer info as an HTML-formatted string
func (peer Peer) Format(name string) string {
	status := "🔴"
	if peer.IsOnline() {
		status = "🟢"
	}

	var b strings.Builder

	fmt.Fprintf(&b, "👤 <b>%s</b>  %s\n", name, status)
	fmt.Fprintf(&b, "<pre>\n")
	fmt.Fprintf(&b, "Time: %s\n", common.FormatAgo(peer.Time))
	fmt.Fprintf(&b, "RX: %s\n", common.FormatBytes(peer.Rx))
	fmt.Fprintf(&b, "TX: %s\n", common.FormatBytes(peer.Tx))
	fmt.Fprintf(&b, "</pre>")

	return b.String()
}
