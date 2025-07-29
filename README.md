# 🛡️ WireGuard Monitor

A simple Go tool to monitor your WireGuard traffic and send reports to Telegram. Made because my cloud provider limits server traffic

---

## 🚀 Features

- Fetches peer stats using `wg show all dump`
- Formats byte counts and last-seen times
- Sends a summary to your Telegram chat via bot


## 🔧 Before you start

You'll need:

- WireGuard installed (`wg` command works)
- Telegram bot token (create one with [@BotFather](https://t.me/BotFather))
- Your Telegram chat ID


## 🌐 Environment Variables

Create a `.env` file in your project folder with these:

| Variable           | Description                                     |
| ------------------ | ----------------------------------------------- |
| `WG_CONFIG_PATH`   | Path to WireGuard config file (optional)        |
| `TELEGRAM_TOKEN`   | Telegram Bot API token (required)               |
| `TELEGRAM_CHAT_ID` | Telegram chat ID to send messages to (required) |

> **Tip:** If `WG_CONFIG_PATH` is not set, defaults to `/etc/wireguard/wg0.conf`.


## 📦 Installation

1. Download the latest version:
   ```bash
   wget https://github.com/bromanla/wireguard-monitor/releases/latest/download/wg-monitor
   wget https://raw.githubusercontent.com/bromanla/wireguard-monitor/refs/heads/main/entrypoint.sh
   ```

2. Make files executable:
   ```bash
   chmod +x wireguard-monitor entrypoint.sh
   ```

3. Place files in your project folder (e.g. `/home/yourname/wireguard-monitor`)

## ⚙️ Usage

1. Place your .env file next to the scripts.

2. Run the entrypoint script:
   ```bash
   ./entrypoint.sh
   ```

## ⏱️ Running via Cron

Add this to your crontab (crontab -e):

```bash
# Run every 30 minutes
*/30 * * * * /home/yourname/wireguard-monitor/entrypoint.sh
```

> **Tip:** Replace /home/yourname/wireguard-monitor with your actual folder path

---

Enjoy monitoring!
