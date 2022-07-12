#!/bin/bash

build=release
work=statsd-push-agent
binary=statsd-push-agent
oss=(darwin darwin linux linux windows windows)
arcs=(amd64 arm64 amd64 arm64 amd64 arm64)

mkdir -p $build/$work
build_cmd="go build -o $build/$work/statsd-push-agent -trimpath agent/agent.go"

for ((i=0; i<6; i++)) do
  GOOS=${oss[i]} GOARCH=${arcs[i]} $build_cmd
  cp install/* $build/$work
  cd $build
  zip statsd-push-agent-${oss[i]}-${arcs[i]}.zip -r $work
  rm $work/*
  cd - 
done