// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package sdk

import (
	"io"
	"net/url"

	monitors "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/monitors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	outputHelpers "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output/helpers"
)

// Monitors provides an interface to the command line chifra monitors through the SDK.
func Monitors(rCtx *output.RenderCtx, w io.Writer, values url.Values) error {
	monitors.ResetOptions(sdkTestMode)
	opts := monitors.MonitorsFinishParseInternal(w, values)
	// EXISTING_CODE
	// EXISTING_CODE
	outputHelpers.InitJsonWriterApi("monitors", w, &opts.Globals)
	err := opts.MonitorsInternal(rCtx)
	outputHelpers.CloseJsonWriterIfNeededApi("monitors", err, &opts.Globals)

	return err
}

// EXISTING_CODE
// EXISTING_CODE
