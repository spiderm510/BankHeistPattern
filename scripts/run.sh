#!/usr/bin/env bash
docker run \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  monopoly-go-bank-heist:latest
