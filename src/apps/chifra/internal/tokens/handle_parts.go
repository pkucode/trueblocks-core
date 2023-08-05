package tokensPkg

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpcClient"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/token"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/ethereum/go-ethereum"
)

func (opts *TokensOptions) HandleParts() error {
	chain := opts.Globals.Chain
	testMode := opts.Globals.TestMode

	// TODO: Why does this have to dirty the caller?
	settings := rpcClient.ConnectionSettings{
		Chain: chain,
		Opts:  opts,
	}
	opts.Conn = settings.DefaultRpcOptions()

	ctx, cancel := context.WithCancel(context.Background())
	fetchData := func(modelChan chan types.Modeler[types.RawToken], errorChan chan error) {
		for _, address := range opts.Addrs {
			addr := base.HexToAddress(address)
			for _, br := range opts.BlockIds {
				blockNums, err := br.ResolveBlocks(chain)
				if err != nil {
					errorChan <- err
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					cancel()
					return
				}

				for _, bn := range blockNums {
					if state, err := token.GetTokenState(chain, addr, fmt.Sprintf("0x%x", bn)); err != nil {
						errorChan <- err
					} else {
						s := &types.SimpleToken{
							Address:     state.Address,
							BlockNumber: bn,
							TotalSupply: state.TotalSupply,
							Decimals:    uint64(state.Decimals),
						}
						modelChan <- s
					}
				}
			}
		}
	}

	nameTypes := names.Custom | names.Prefund | names.Regular
	namesMap, err := names.LoadNamesMap(chain, nameTypes, nil)
	if err != nil {
		return err
	}

	extra := map[string]interface{}{
		"testMode": testMode,
		"namesMap": namesMap,
		"parts":    opts.Parts,
	}

	return output.StreamMany(ctx, fetchData, opts.Globals.OutputOptsWithExtra(extra))
}
