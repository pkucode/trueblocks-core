// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package whenPkg

import (
	"net/http"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/specials"
)

func (opts *WhenOptions) ListInternal() error {

	result := specials.GetSpecials(opts.Globals.TestMode)
	if opts.Globals.ApiMode {
		opts.Globals.Respond(opts.Globals.Writer, http.StatusOK, result)

	} else {
		err := opts.Globals.Output(os.Stdout, opts.Globals.Format, result)
		if err != nil {
			logger.Log(logger.Error, err)
		}
	}

	return nil
}
