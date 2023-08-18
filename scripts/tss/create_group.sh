#!/bin/bash

export WALLET_NAME=validator

BASEDIR=$(dirname "$0")

# Submit create_group proposal
bandd tx gov submit-proposal $BASEDIR/proposal_create_group.json \
  --from $WALLET_NAME \
  -y --keyring-backend test --gas-prices 0.0025uband -b sync

sleep 5

# Query last proposal's id
PROPOSAL_ID=$(bandd query gov proposals --count-total --limit 1 --output json | jq -r  '.pagination.total')

# Vote on that proposal
echo "...Voting to proposal $PROPOSAL_ID..."
bandd tx gov vote $PROPOSAL_ID yes \
  --from $WALLET_NAME \
  -y --keyring-backend test --gas-prices 0.0025uband -b sync
