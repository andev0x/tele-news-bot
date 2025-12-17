# Project Context

## Purpose
Telegram bot for reading and aggregating news feeds.

## Entry point
- cmd/bot/main.go

## Main packages
- feed/        : fetch and parse RSS feeds
- notify/     : send messages to Telegram
- store/      : store sent items and subscribers
- config/     : load configuration

## Runtime
- Single binary
- Runs as a long-lived process

## Notes
- This is a personal, basic project
- No web server, no database
