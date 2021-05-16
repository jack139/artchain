#!/usr/bin/env bash

rm -rf ./n1

artchaind init node1 --home n1 --chain-id artchain

artchaind keys add user0 --home n1
artchaind keys add gt --home n1
artchaind keys add faucet --home n1

artchaind add-genesis-account $(artchaind keys show user0 -a --home n1) 100000000stake,1credit --home n1
artchaind add-genesis-account $(artchaind keys show gt -a --home n1) 1credit --home n1
artchaind add-genesis-account $(artchaind keys show faucet -a --home n1) 21000000credit --home n1

artchaind gentx user0 100000000stake --chain-id artchain --home n1

echo "Collecting genesis txs..."
artchaind collect-gentxs --home n1

echo "Validating genesis file..."
artchaind validate-genesis --home n1


artchaind start --log_level warn --home n1