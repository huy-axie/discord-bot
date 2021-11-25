## Usage

- How to run ? 
    + Env:
        - CF_API_TOKEN: is Cloudflare API token with DNS zone edit permission.
        - TOKEN: is Discord bot token.
    + Build your own image:
            ```docker build -t discord-bot .```
    + Docker command:
            ```docker run -e CF_API_TOKEN="" -e TOKEN="" discord-bot```
