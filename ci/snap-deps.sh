#!/bin/bash

set -e

ROOT="$(dirname $( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd ))"

KUBECTL_VERSION=1.18.5


curl -sL "https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl" \
  > $ROOT/kubectl

echo $(readlink $ROOT/kubectl)
