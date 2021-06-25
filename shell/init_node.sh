#!/bin/sh

NODE_HOME="n2"
NODE_NAME="node2"
NODE_USER="user2"

rm -rf ./$NODE_HOME

artchaind init $NODE_NAME --home $NODE_HOME --chain-id artchain

artchaind keys add $NODE_USER --home $NODE_HOME

artchaind add-genesis-account $(artchaind keys show $NODE_USER -a --home $NODE_HOME) 100000000stake,1credit --home $NODE_HOME

artchaind gentx $NODE_USER 100000000stake --chain-id artchain --home $NODE_HOME

echo "Collecting genesis txs..."
artchaind collect-gentxs --home $NODE_HOME

echo "Validating genesis file..."
artchaind validate-genesis --home $NODE_HOME
