# Frosty Go Starter

🧊 A minimal Go app that uses the [Frosty AI](https://gofrosty.ai) REST API to send a prompt and route it based on cost or performance.

## ⚙️ Setup

1. **Clone this repo**

2. **Create a `.env` file**  
   Copy `.env.example` and fill in your Router ID and Key from the Frosty console:

   ```env
   ROUTER_ID=your_router_id
   ROUTER_KEY=your_router_key
    ```
3. Install dependencies

```bash
    go mod tidy
```
Run the app
```bash
    go run main.go
```

## 🧠 What It Does
Calls Frosty’s /chat endpoint with a prompt

Authenticates with router_id and router_key

Supports routing rules like cost, performance, or none

## 💡 Example Output
```json
{
  "trace_id": "a37a5ae7-62ec-481e-89b7-2952b721b023",
  "total_tokens": 40,
  "prompt_type": "chat",
  "prompt_tokens": 11,
  "response_tokens": 29,
  "model": "claude-3-sonnet-20240229",
  "provider": "Anthropic",
  "total_time": 878.59,
  "prompt": [
    {
      "role": "user",
      "content": "Tell me a joke"
    }
  ],
  "cost": "- -",
  "rule": "cost",
  "response": "Here's a silly joke for you: Why can't a bicycle stand up by itself? It's two-tired!",
  "success": "True"
}
```
## 📎 Helpful Links
- 🔐 [Frosty Console](https://console.gofrosty.ai)
- 📚 [Frosty API Docs](https://docs.gofrosty.ai/frosty-ai-docs/api-documentation)
