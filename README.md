# **BotBot LeetCode Daily Challenge Notifier**
BotBot LeetCode Daily Challenge Notifier is a simple automated bot designed to fetch daily coding challenges from LeetCode and send notifications with the challenge details. Currently, it supports notifications via Telegram.

## **Features**
- Fetches the daily coding challenge from LeetCode.
- Formats messages appropriately for different platforms `(now is only for Telegram)`.
- Sends formatted challenge details directly as a personal message.

## **Getting Started**

### **Prerequisites**

- Go 1.18 or later.
- An active LeetCode API endpoint that provides daily challenges.
- A Telegram bot token and chat ID for sending messages through Telegram.

### **Installation**

1. **Clone the Repository**
    ```bash
    git clone https://github.com/phuonganhniie/botbot-leetcode.git
    ```
    ```bash
    cd botbot-leetcode
    ```
    
2. **Configure the Application**

    Copy the sample configuration file and modify it according to your needs. The config file must be named `config.env`.
    
    Edit **`config/config.env`** with your details:
    
    ```yaml
    LEETCODE_DAILY_URL: "https://example.com/api/leetcode_challenge"

    TELEGRAM_BOT_TOKEN: "12xxx:ABC-DExxxxhIkl-zyx57W2xxxxx3ew11"
    ```

 * About the LeetCode URL: I used the [alfa-leetcode-api](https://github.com/alfaArghya/alfa-leetcode-api) for getting the well-structured LeetCode's data. Special thanks to `alfaarghya` for contributing helpful APIs.

### **Usage**
Run the application to start fetching and sending the daily LeetCode challenge:

```bash
go run main.go
```

The application is set up with a cron job that triggers at 10:00 AM daily to perform these tasks automatically by `AWS Lambda` & `AWS EventBridge`.


## **License**
This project is licensed under the MIT License - see the [LICENSE.md](https://chat.openai.com/g/g-n7Rs0IK86-grimoire/c/LICENSE.md) file for details.