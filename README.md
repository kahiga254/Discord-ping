# Discord Ping Bot

A simple Discord bot built with [discordgo](https://github.com/bwmarrin/discordgo) that responds to "ping" messages with "pong".

## 🚀 Features
- Listens for "ping" messages and responds with "pong".
- Uses environment variables to securely store the bot token.

## 🛠 Installation
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/your-username/discord-ping.git
cd discord-ping
```

### **2️⃣ Set Up Your Environment**
1. **Create a `.env` file** or set an environment variable:

   **For Linux/macOS:**
   ```sh
   export DISCORD_TOKEN="your-bot-token"
   ```
   **For Windows (PowerShell):**
   ```powershell
   $env:DISCORD_TOKEN="your-bot-token"
   ```

2. **Add `.env` to `.gitignore`** (to prevent committing secrets):
   ```sh
   echo ".env" >> .gitignore
   ```

### **3️⃣ Install Dependencies**
```sh
go mod tidy
```

### **4️⃣ Run the Bot**
```sh
go run main.go
```

## 📜 Project Structure
```
📂 discord-ping/
├── 📂 bot/
│   ├── bot.go       # Handles bot logic
├── 📂 config/
│   ├── config.go    # Loads configuration
├── go.mod           # Go module dependencies
├── go.sum           # Dependency checksums
├── main.go          # Entry point for the bot
└── .gitignore       # Files to ignore
```

## 🔧 Configuration
Instead of storing the **bot token** in a JSON file (which is unsafe), this project loads it from an **environment variable**.

Modify `bot/bot.go` to retrieve the token securely:
```go
import (
    "os"
)

goBot, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
```

## ❌ Handling Push Protection Errors
If GitHub blocks your push due to a leaked token, follow these steps:

1. **Remove the token from the file:**
   ```sh
   nano config.json  # Delete the token
   ```
2. **Rewrite commit history to erase the leaked token:**
   ```sh
   git filter-branch --force --index-filter \
     "git rm --cached --ignore-unmatch config.json" \
     --prune-empty --tag-name-filter cat -- --all
   ```
3. **Force push the cleaned history:**
   ```sh
   git push origin master --force
   ```
4. **Regenerate your bot token in the [Discord Developer Portal](https://discord.com/developers/applications).**

## 🎯 Future Enhancements
- Add more commands.
- Implement slash commands.
- Add error handling and logging.

## 🤝 Contributing
Feel free to fork and submit pull requests! 🚀

## 📜 License
This project is licensed under the MIT License. See the `LICENSE` file for details.

