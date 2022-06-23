module github.com/jack139/artchain

go 1.15

require (
	github.com/Ferluci/fast-realip v1.0.0
	github.com/cosmos/cosmos-sdk v0.45.5
	github.com/cosmos/go-bip39 v1.0.0
	github.com/fasthttp/router v1.3.11
	github.com/go-redis/redis/v8 v8.10.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ipfs/go-ipfs-api v0.2.0
	github.com/irisnet/irismod v1.5.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.6
	github.com/valyala/fasthttp v1.23.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.45.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
