// Copyright 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were generated with makeClass --run. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package sdk

import (
	// EXISTING_CODE
	"io"
	"net/url"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	when "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/sdk"
	// EXISTING_CODE
)

type WhenOptions struct {
	BlockIds   []string
	List       bool
	Timestamps bool
	Count      bool
	Truncate   base.Blknum
	Repair     bool
	Check      bool
	Update     bool
	Deep       bool
	Globals

	// EXISTING_CODE
	// EXISTING_CODE
}

// When implements the chifra when command for the SDK.
func (opts *WhenOptions) When(w io.Writer) error {
	values := make(url.Values)

	// EXISTING_CODE
	//   blocks - one or more dates, block numbers, hashes, or special named blocks (see notes)
	//   -l, --list         export a list of the 'special' blocks
	//   -t, --timestamps   display or process timestamps
	//   -U, --count        with --timestamps only, returns the number of timestamps in the cache
	//   -r, --repair       with --timestamps only, repairs block(s) in the block range by re-querying from the chain
	//   -c, --check        with --timestamps only, checks the validity of the timestamp data
	//   -u, --update       with --timestamps only, bring the timestamp database forward to the latest block
	//   -d, --deep         with --timestamps --check only, verifies timestamps from on chain (slow)
	for _, blockId := range opts.BlockIds {
		values.Add("blocks", blockId)
	}
	if opts.List {
		values.Set("list", "true")
	}
	if opts.Timestamps {
		values.Set("timestamps", "true")
	}
	if opts.Count {
		values.Set("count", "true")
	}
	if opts.Repair {
		values.Set("repair", "true")
	}
	if opts.Check {
		values.Set("check", "true")
	}
	if opts.Update {
		values.Set("update", "true")
	}
	if opts.Deep {
		values.Set("deep", "true")
	}
	// EXISTING_CODE
	opts.Globals.mapGlobals(values)

	return when.When(w, values)
}

// no enums

// EXISTING_CODE
// EXISTING_CODE

