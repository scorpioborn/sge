#!/usr/bin/env bash

# How to run manually:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run --rm -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

echo "Formatting protobuf files"
find ./ -name "*.proto" -exec clang-format -i {} \;

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./sge -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    # this regex checks if a proto file has its go_package set to cosmossdk.io/api/...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*cosmossdk.io/api' "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..


# move proto files to the right places
# mkdir -p github.com/sge-network/sge/x/legacy
# mv github.com/sge-network/sge/x/bet github.com/sge-network/sge/x/legacy/bet
# mv github.com/sge-network/sge/x/house github.com/sge-network/sge/x/legacy/house
# mv github.com/sge-network/sge/x/market github.com/sge-network/sge/x/legacy/market
# mv github.com/sge-network/sge/x/orderbook github.com/sge-network/sge/x/legacy/orderbook
# mv github.com/sge-network/sge/x/ovm github.com/sge-network/sge/x/legacy/ovm
# mv github.com/sge-network/sge/x/reward github.com/sge-network/sge/x/legacy/reward
# mv github.com/sge-network/sge/x/subaccount github.com/sge-network/sge/x/legacy/subaccount


cp -r ./github.com/sge-network/sge/x/* x/
cp -r ./github.com/sge-network/sge/types/* types/

rm -rf ./github.com

# # go mod tidy

./scripts/protocgen-pulsar.sh