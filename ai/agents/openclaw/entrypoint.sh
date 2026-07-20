#!/usr/bin/env bash
set -e

if ! openclaw plugins inspect discord > /dev/null 2>&1; then
  NPM_CONFIG_CACHE=/tmp/.npm openclaw plugins install @openclaw/discord
fi

exec openclaw gateway run --bind lan
