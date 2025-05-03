# 🏔️ Piste

**Piste** is the beginning of a project aimed at building a **copy trading Telegram bot** for **Solana memecoins** launched via [Pump.fun](https://pump.fun). This initial version connects to the Pump.fun WebSocket API to stream real-time trade data from specified wallets.

## 🚀 Actual Features

* Connects to `wss://pumpportal.fun/api/data`
* Subscribes to `subscribeAccountTrade` events for one or more wallet addresses
* Parses and prints real-time trade events

## 🧪 Getting Started

### Prerequisites

* Go 1.18+
* Web access to `wss://pumpportal.fun`

### Installation

```bash
git clone https://github.com/yourusername/piste.git
cd piste
go build -o piste
```

### Usage

```bash
./piste <WALLET_ADDRESS_1> <WALLET_ADDRESS_2> ...
```

Example:

```bash
./piste 3Hg...abc E9x...xyz
```

## 🛠️ Roadmap

* [x] Initial WebSocket connection and subscription
* [ ] Build Telegram bot interface
* [ ] Mirror trades via on-chain transactions (copy trading logic)
* [ ] Deploy to a VPS for real-time monitoring

## 📜 License

MIT
