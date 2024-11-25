#!/bin/bash

rm -rf ~/.grogu

# config chain id
grogu config chain-id bandchain

# add validator to grogu config
grogu config validator $(bandd keys show validator -a --bech val --keyring-backend test)

# setup execution endpoint
grogu config bothan "$BOTHAN_URL"

# setup broadcast-timeout to grogu config
grogu config broadcast-timeout "1m"

# setup rpc-poll-interval to grogu config
grogu config rpc-poll-interval "1s"

# setup nodes to grogu config
grogu config nodes "http://localhost:26657"

# setup log-level to grogu config
grogu config log-level "info"

# setup max-try to grogu config
grogu config max-try 5

echo "y" | bandd tx oracle activate --from validator --gas-prices 0.0025uband --keyring-backend test --chain-id bandchain

# wait for activation transaction success
sleep 3

for i in $(eval echo {1..4})
do
  # add feeder key
  grogu keys add feeder$i
done

sleep 3

# send band tokens to feeders
echo "y" | bandd tx bank multi-send validator $(grogu keys list -a) 1000000uband --gas-prices 0.0025uband --keyring-backend test --chain-id bandchain
# wait for sending band tokens transaction success
sleep 3

# add feeder to bandchain
echo "y" | bandd tx feeds add-feeders $(grogu keys list -a) --from validator --gas-prices 0.0025uband --keyring-backend test --chain-id bandchain

# wait for addding feeder transaction success
sleep 3

# run grogu
grogu run