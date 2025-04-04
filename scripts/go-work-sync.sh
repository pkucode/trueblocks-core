#!/usr/bin/env bash

#------------------------------------------------
# Name:    go-work-sync.sh
# Purpose: Find go.mod files, save the directories
#          to a go.work file, and runs go work sync

#------------------------------------------------
# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Navigate to the root of the repo
cd "$SCRIPT_DIR/.."

#------------------------------------------------
# Function to check if git submodules are initialized
check_submodules() {
    echo "===> Checking Git submodules..."
    if git submodule status | grep -q '^-'; then
        echo "Error: Some submodules are not initialized!"
        echo "Run: git submodule update --init --recursive"
        exit 1
    fi
}

# Run submodule check
check_submodules

#------------------------------------------------
# Check if the go.work file already exists
GO_WORK_FILE="go.work"
rm -f "$GO_WORK_FILE"
if [ ! -f "$GO_WORK_FILE" ]; then
    echo "// Go Version" > "$GO_WORK_FILE"
    echo "go 1.23.1" >> "$GO_WORK_FILE"
fi

#------------------------------------------------
# Find all go.mod files in the src directory and
# add their directories to go.work.
echo "===> Getting latest sdk and chifra..."
find . -type f -name 'go.mod' | while read -r modfile; do
    moddir=$(dirname "$modfile")
    go work use "$moddir"

    echo "    Updated $moddir/go.mod"
    cd "$moddir" || exit

    isSdk="[ "$moddir" == "./sdk" ]"
    isChifra="[ "$moddir" == "./src/apps/chifra" ]"
    isGoMaker="[ "$moddir" == "./src/dev_tools/goMaker" ]"
    isNode="[ "$moddir" == "./node" ]"
    isFourbyte="[ "$moddir" == "./examples/four_bytes" ]"
    isKeystore="[ "$moddir" == "./examples/keystore" ]"
    isIndexMan="[ "$moddir" == "./src/dev_tools/indexManager" ]"
    isSimple="[ "$moddir" == "./examples/simple" ]"

    if $isGoMaker || $isNode || $isFourbyte; then
        go get github.com/btcsuite/btcd 2> /dev/null
    fi

    if ! $isSdk && ! $isChifra && ! $isGoMaker && ! $isFourbyte && ! $isIndexMan && ! $isKeystore; then
        go get github.com/TrueBlocks/trueblocks-sdk/v4@latest
    fi

    if ! $isChifra; then
        go get github.com/TrueBlocks/trueblocks-core/src/apps/chifra@latest
    fi

    go mod tidy
    cd - > /dev/null
done

#------------------------------------------------
echo
echo "===> Creating a fresh go.work file..."
go work sync
cat go.work

#------------------------------------------------
echo
echo "===> Tidying go.mod files..."
"$SCRIPT_DIR/go-mod-tidy.sh"

echo
echo "===> The repo imports these TrueBlocks modules..."
"$SCRIPT_DIR/go-show-trueblocks.sh"

cd - 2>&1 > /dev/null
echo
echo "===> Done..."

exit 0
