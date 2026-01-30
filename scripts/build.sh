#!/usr/bin/env bash
set -e

IMAGE_NAME=monopoly-go-bank-heist
TAG=latest

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "Project root: ${ROOT_DIR}"
echo "Building Docker image..."

docker build \
  -t ${IMAGE_NAME}:${TAG} \
  -f "${ROOT_DIR}/scripts/Dockerfile" \
  "${ROOT_DIR}"

echo "Build complete"
echo
echo "To run:"
echo "docker run -p 8080:8080 ${IMAGE_NAME}:${TAG}"
