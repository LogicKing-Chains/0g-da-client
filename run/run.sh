/bin/combined \
	--chain.rpc  \
	--chain.private-key  \
	--chain.receipt-wait-rounds 180 \
	--chain.receipt-wait-interval 1s \
	--chain.gas-limit 2000000 \
	--combined-server.use-memory-db \
	--combined-server.storage.kv-db-path /runtime/ \
	--combined-server.storage.time-to-expire 2592000 \
	--disperser-server.grpc-port 51001 \
	--batcher.da-entrance-contract  \
	--batcher.da-signers-contract 0x0000000000000000000000000000000000001000 \
	--batcher.finalizer-interval 20s \
	--batcher.confirmer-num 3 \
	--batcher.max-num-retries-for-sign 3 \
	--batcher.finalized-block-count 50 \
	--batcher.batch-size-limit 500 \
	--batcher.encoding-interval 3s \
	--batcher.encoding-request-queue-size 1 \
	--batcher.pull-interval 10s \
	--batcher.signing-interval 3s \
	--batcher.signed-pull-interval 20s \
	--batcher.expiration-poll-interval 3600 \
	--encoder-socket  \
	--encoding-timeout 600s \
	--signing-timeout 600s \
	--chain-read-timeout 12s \
	--chain-write-timeout 13s \
	--combined-server.log.level-file trace \
	--combined-server.log.level-std  trace \
    --disperser-server.retriever-address  \
	--combined-server.log.path /runtime/run.log
