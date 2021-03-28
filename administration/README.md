# Administrative notes

## Build

### Requirements

Go 1.16

If 1.16 is not available in your repo (as is with ubuntu), you may manually install it

https://www.kalilinux.in/2020/06/how-to-install-golang-in-kali-linux-new.html

```
git clone https://github.com/t-900-a/gemini-fortune-bot.git
cd gemini-fortune-bot
go build .
```
## Gemini Server
Bot was tested using diamant a simple gemini server

[See installation and setup section](https://git.umaneti.net/diamant/about/)

## index.gmi
You need to create your own index.gmi or use the included examples

The index.gmi within the gemlog folder will be updated automatically for you

## Parameters
You'll need fortunes db that contains the potential fortunes that the bot will use
```bigquery
git clone https://github.com/bmc/fortunes.git
```
See the babysitter script for an example

fortuneFile - specific file location for the fortune file

websiteUri - your website where this bot is hosted (used to generate rss feed)

txHash - this is passed for you by monero-wallet-rpc, represented as %s
pmtUri - Monero address uri "monero:47afuhgbauhg"
pmtViewKey - your Monero view key

Specify your view key so you can prove that the bot is reacting to on-chain transactions.

You can also share your bot's generated RSS feed on Gemmit to compare popularity to other Gemini gemlogs that accept Monero

https://github.com/t-900-a/gemmit

## Cron

```bigquery
*/5 * * * * /home/<USER>/fortune_babysitter.sh