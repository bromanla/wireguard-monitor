package main

import (
	"fmt"
	"log"
	"strings"
	"wireguard-monitor/internal/common"
	"wireguard-monitor/internal/config"
	"wireguard-monitor/internal/notifier"
	"wireguard-monitor/internal/peer"
)

func main() {
	configPath := common.GetEnv("WG_CONFIG_PATH", "/etc/wireguard/wg0.conf")

	config, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Error when reading the config: %v", err)
	}

	peers, err := peer.ReadPeers()
	if err != nil {
		log.Fatalf("Error when reading the peers: %v", err)
	}

	var total uint64 = 0
	var body strings.Builder

	for _, p := range peers {
		total += p.Rx + p.Tx
		name := config.Get(p.Key)

		if body.Len() > 0 {
			body.WriteString("\n\n")
		}

		body.WriteString(p.Format(name))
	}

	header := fmt.Sprintf("Total traffic: %s", common.FormatBytes(total))
	message := header + "\n\n" + body.String()

	notifier.SendTelegram(message)
}
