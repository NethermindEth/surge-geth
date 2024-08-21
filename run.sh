./build/bin/geth \
  --gcmode archive \
  --syncmode full \
  --datadir ~/taiko-helder-geth \
  --metrics \
  --metrics.expensive \
  --metrics.addr "0.0.0.0" \
  --networkid "167010" \
  --http \
  --http.addr "0.0.0.0" \
  --http.vhosts "*" \
  --http.corsdomain "*" \
  --http.api eth,net,web3,txpool,miner,taiko,admin \
  --ws \
  --ws.addr "0.0.0.0" \
  --ws.origins "*" \
  --authrpc.addr "0.0.0.0" \
  --authrpc.port "8551" \
  --authrpc.vhosts "*" \
  --authrpc.jwtsecret ~/taiko-helder-geth/geth/jwtsecret \
  --allow-insecure-unlock \
  --ws.api eth,net,web3,txpool,miner,taiko \
  --taiko \
  --maxpeers 100 \
  console