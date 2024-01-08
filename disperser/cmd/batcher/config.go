package main

import (
	"github.com/urfave/cli"
	"github.com/zero-gravity-labs/zgda/common/aws"
	"github.com/zero-gravity-labs/zgda/common/geth"
	"github.com/zero-gravity-labs/zgda/common/logging"
	"github.com/zero-gravity-labs/zgda/common/storage_node"
	"github.com/zero-gravity-labs/zgda/core/encoding"
	"github.com/zero-gravity-labs/zgda/disperser/batcher"
	"github.com/zero-gravity-labs/zgda/disperser/cmd/batcher/flags"
	"github.com/zero-gravity-labs/zgda/disperser/common/blobstore"
)

type Config struct {
	BatcherConfig     batcher.Config
	TimeoutConfig     batcher.TimeoutConfig
	BlobstoreConfig   blobstore.Config
	EthClientConfig   geth.EthClientConfig
	AwsClientConfig   aws.ClientConfig
	EncoderConfig     encoding.EncoderConfig
	LoggerConfig      logging.Config
	MetricsConfig     batcher.MetricsConfig
	StorageNodeConfig storage_node.ClientConfig
}

func NewConfig(ctx *cli.Context) Config {
	config := Config{
		BlobstoreConfig: blobstore.Config{
			BucketName: ctx.GlobalString(flags.S3BucketNameFlag.Name),
			TableName:  ctx.GlobalString(flags.DynamoDBTableNameFlag.Name),
		},
		EthClientConfig: geth.ReadEthClientConfig(ctx),
		AwsClientConfig: aws.ReadClientConfig(ctx, flags.FlagPrefix),
		EncoderConfig:   encoding.ReadCLIConfig(ctx),
		LoggerConfig:    logging.ReadCLIConfig(ctx, flags.FlagPrefix),
		BatcherConfig: batcher.Config{
			PullInterval:             ctx.GlobalDuration(flags.PullIntervalFlag.Name),
			FinalizerInterval:        ctx.GlobalDuration(flags.FinalizerIntervalFlag.Name),
			EncoderSocket:            ctx.GlobalString(flags.EncoderSocket.Name),
			NumConnections:           ctx.GlobalInt(flags.NumConnectionsFlag.Name),
			EncodingRequestQueueSize: ctx.GlobalInt(flags.EncodingRequestQueueSizeFlag.Name),
			BatchSizeMBLimit:         ctx.GlobalUint(flags.BatchSizeLimitFlag.Name),
			SRSOrder:                 ctx.GlobalInt(flags.SRSOrderFlag.Name),
			MaxNumRetriesPerBlob:     ctx.GlobalUint(flags.MaxNumRetriesPerBlobFlag.Name),
			TargetNumChunks:          ctx.GlobalUint(flags.TargetNumChunksFlag.Name),
		},
		TimeoutConfig: batcher.TimeoutConfig{
			EncodingTimeout:   ctx.GlobalDuration(flags.EncodingTimeoutFlag.Name),
			ChainReadTimeout:  ctx.GlobalDuration(flags.ChainReadTimeoutFlag.Name),
			ChainWriteTimeout: ctx.GlobalDuration(flags.ChainWriteTimeoutFlag.Name),
		},
		MetricsConfig: batcher.MetricsConfig{
			HTTPPort:      ctx.GlobalString(flags.MetricsHTTPPort.Name),
			EnableMetrics: ctx.GlobalBool(flags.EnableMetrics.Name),
		},
		StorageNodeConfig: storage_node.ReadClientConfig(ctx, flags.FlagPrefix),
	}
	return config
}
