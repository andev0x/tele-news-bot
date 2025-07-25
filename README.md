# Tele-News-Bot

A simple Telegram bot that fetches news from RSS feeds and sends new articles to a specified Telegram chat.

## Features

- Fetches news from multiple RSS feeds concurrently.
- Sends new articles to a specified Telegram chat.
- Remembers which articles have been sent to avoid duplicates.
- Configurable via environment variables.

## How it Works

The bot periodically (every 10 minutes) fetches items from the provided RSS feeds. It checks if an item has been published within the last 2 hours and if it has not been sent before. If both conditions are met, it sends a message to the configured Telegram chat with the article's title and link, and then stores the link in a local file to mark it as sent.

## Configuration

The bot is configured using a `.env` file. Create a `.env` file in the root of the project and add the following variables:

```
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
USER_CHAT_ID=your_telegram_chat_id
RSS_FEEDS=https://feed1.example.com,https://feed2.example.com
```

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token.
- `USER_CHAT_ID`: The ID of the Telegram chat where you want to receive news.
- `RSS_FEEDS`: A comma-separated list of RSS feed URLs.

An example `.env.example` file is provided.

## Installation and Usage

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/andev0x/tele-news-bot.git
    cd tele-news-bot
    ```

2.  **Create and configure your `.env` file:**
    ```bash
    cp .env.example .env
    # Edit .env with your details
    ```

3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run the bot:**
    ```bash
    go run main.go
    ```

The bot will start, and you will see log messages in your terminal.

## Dependencies

- [gopkg.in/telebot.v3](https://gopkg.in/telebot.v3) - A popular framework for building Telegram bots in Go.
- [github.com/mmcdole/gofeed](https://github.com/mmcdole/gofeed) - A robust RSS and Atom feed parser for Go.
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - A Go port of the Ruby dotenv library (loads environment variables from a `.env` file).
