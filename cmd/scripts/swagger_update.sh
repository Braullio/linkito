#!/bin/bash
# https://github.com/swaggo/swag

cd ../..

echo "# RUN: pwd"
echo "# OUT: $(pwd)"
echo ""
echo "# RUN: swag init ./cmd/server/main.go"
echo "# OUT: "
echo "$(swag init ./cmd/server/main.go)"
echo ""
