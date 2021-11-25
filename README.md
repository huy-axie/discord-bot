## Usage
- Create a bot and add to channel :ref https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf
- How to run ? 
    + Env:
        - CF_API_TOKEN: is Cloudflare API token with DNS zone edit permission.
        - TOKEN: is Discord bot token.
    + Build your own image:
            ```docker build -t discord-bot .```
    + Docker command:
            ```docker run -e CF_API_TOKEN="" -e TOKEN="" discord-bot```

- Intereact with bot:
    + Go to channel and type: `ping` to verify bot, always response with `Pong!`.
    + Type `flarectl -h`. 