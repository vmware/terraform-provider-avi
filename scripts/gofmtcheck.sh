#!/usr/bin/env bash
############################################################################
# ========================================================================
# Copyright (c) 2025 Broadcom Inc. and/or its subsidiaries. All Rights Reserved. Broadcom Confidential.
# ========================================================================
###

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(gofmt -l `find . -name '*.go' | grep -v vendor`)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0
