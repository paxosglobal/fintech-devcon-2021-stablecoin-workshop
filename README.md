# fintech-devcon-2021-stablecoin-workshop
Fintech DevCon 2021 Stablecoin Workshop

# Slides

https://github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/raw/master/slides.pdf

# Setup

- install golang https://golang.org/doc/install
- install docker https://docs.docker.com/get-docker/
- install node https://nodejs.org/en/download/

# This is Exercise 4: Use MetaMask To Receive Tokens

- Install metamask https://metamask.io/
- Create a wallet and a first address to send new tokens to
- Mint tokens to your metamask address using the UI http://localhost:3000/
  - Make sure to follow `Run it` below to get the app started
- Add the token to metamask to see the tokens you recieved
  - The contract address to add should be `0xc4680463046E64b10Da390d9049D24b8EC43AaAB`

# Run it

In separate terminals
- make start-local
- make run-backend
- make run-frontend

Navigate to http://localhost:3000/ to use the app!
