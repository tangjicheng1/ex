FROM golang:1.24-bookworm

RUN apt-get update && apt-get install -y \
    git \
    curl \
    vim \
    protobuf-compiler
