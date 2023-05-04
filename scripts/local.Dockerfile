# syntax=docker/dockerfile:experimental

# This Dockerfile is meant to be used with the build_local_dep_image.sh script
# in order to build an image using the local version of coreth

# Changes to the minimum golang version must also be replicated in
# scripts/build_pepecoin.sh
# scripts/local.Dockerfile (here)
# Dockerfile
# README.md
# go.mod
FROM golang:1.19.6-buster

RUN mkdir -p /go/src/github.com/memeticofficial

WORKDIR $GOPATH/src/github.com/memeticofficial
COPY pepecoingo pepecoingo

WORKDIR $GOPATH/src/github.com/memeticofficial/pepecoingo
RUN ./scripts/build_pepecoin.sh

RUN ln -sv $GOPATH/src/github.com/memeticofficial/pepecoin-byzantine/ /pepecoingo
