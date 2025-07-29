#!/bin/bash

set -e
cd "$(dirname "$0")"

if [ -f ".env" ]; then
  set -a
  source .env
  set +a
else
  echo "⚠️ File.env not found"
  exit 1
fi

exec "./wireguard-monitor" "$@"
