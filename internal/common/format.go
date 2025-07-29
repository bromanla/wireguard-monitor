package common

import (
	"fmt"
	"time"
)

// FormatBytes returns a human-readable string for byte counts (B, KiB, MiB, GiB).
func FormatBytes(bytes uint64) string {
	const (
		KiB = 1 << 10
		MiB = 1 << 20
		GiB = 1 << 30
	)

	switch {
	case bytes >= GiB:
		return fmt.Sprintf("%.2f GiB", float64(bytes)/GiB)
	case bytes >= MiB:
		return fmt.Sprintf("%.2f MiB", float64(bytes)/MiB)
	case bytes >= KiB:
		return fmt.Sprintf("%.2f KiB", float64(bytes)/KiB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// FormatAgo returns a relative time description (e.g., "5 min ago").
func FormatAgo(ts int64) string {
	if ts <= 0 {
		return "N/A"
	}

	now := time.Now().Unix()
	diff := now - ts

	switch {
	case diff < 60:
		return fmt.Sprintf("%d sec ago", diff)
	case diff < 3600:
		return fmt.Sprintf("%d min ago", diff/60)
	case diff < 86400:
		return fmt.Sprintf("%d h ago", diff/3600)
	default:
		return fmt.Sprintf("%d d ago", diff/86400)
	}
}
