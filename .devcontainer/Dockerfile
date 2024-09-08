FROM buildpack-deps:jammy-curl

LABEL dev.containers.features="common"

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && apt-get install -y \
    curl \
    git \
    sudo \
    vim \
    && rm -rf /var/lib/apt/lists/*

