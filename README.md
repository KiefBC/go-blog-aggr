# Blog Aggregator CLI

A command-line RSS feed aggregator built in Go with PostgreSQL. Manage RSS feeds, follow feeds from other users, and aggregate content from your favorite blogs.

## Prerequisites

- Go 1.19+
- PostgreSQL database
- Basic command-line knowledge

## Quick Start

### 1. Clone and Build

```bash
git clone https://github.com/KiefBC/blog-aggr
cd blog-aggr
go build -o blog-aggr .
```

### 2. Database Setup

Set your PostgreSQL connection string in `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost/dbname?sslmode=disable",
  "current_user_name": ""
}
```

### 3. Run Migrations

```bash
goose -dir sql/schema postgres "your_connection_string" up
```

### 4. Get Started

```bash
# Create an account
./blog-aggr register yourname

# Login
./blog-aggr login yourname

# Add your first RSS feed
./blog-aggr addfeed "Hacker News" "https://hnrss.org/newest"

# Follow an existing feed
./blog-aggr follow "https://hnrss.org/newest"

# See what you're following
./blog-aggr following
```

## Key Commands

### User Management

- `register <username>` - Create a new account
- `login <username>` - Login to your account
- `users` - List all users

### Feed Management

- `addfeed <name> <url>` - Add a new RSS feed (requires login)
- `feeds` - List all available feeds
- `follow <url>` - Follow an existing feed (requires login)
- `following` - Show feeds you're following (requires login)

### Other

- `agg <url>` - Fetch and display RSS content
- `help [command]` - Show help
- `reset` - Clear all data (development only)

## Features

- **User Authentication** - Register and login system with session persistence
- **RSS Feed Management** - Add feeds and follow feeds created by other users
- **Type-Safe Queries** - Generated Go code from SQL using SQLC
- **CLI Interface** - Simple command-line interface with helpful error messages
- **PostgreSQL Integration** - Reliable data storage with proper relationships

## Development

This project uses:

- **SQLC** for type-safe database queries
- **Goose** for database migrations
- **PostgreSQL** for data persistence

### Regenerate Database Code

```bash
sqlc generate
```

### Run Tests

```bash
go test ./...
```

## License

This project is for anyone to use or modify.
