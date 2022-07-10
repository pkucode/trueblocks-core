// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * This file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

package tracesPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/identifiers"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

// TracesOptions provides all command options for the chifra traces command.
type TracesOptions struct {
	// a space-separated list of one or more transaction identifiers
	Transactions []string `json:"transactions,omitempty"`
	// transaction identifiers
	TransactionIds []identifiers.Identifier `json:"transactionIds,omitempty"`
	// articulate the retrieved data if ABIs can be found
	Articulate bool `json:"articulate,omitempty"`
	// call the node's trace_filter routine with bang-separated filter
	Filter string `json:"filter,omitempty"`
	// export state diff traces (not implemented)
	Statediff bool `json:"statediff,omitempty"`
	// show the number of traces for the transaction only (fast)
	Count bool `json:"count,omitempty"`
	// skip over the 2016 ddos during export ('on' by default)
	SkipDdos bool `json:"skipDdos,omitempty"`
	// if --skip_ddos is on, this many traces defines what a ddos transaction is
	Max uint64 `json:"max,omitempty"`
	// the global options
	Globals globals.GlobalOptions `json:"globals,omitempty"`
	// an error flag if needed
	BadFlag error `json:"badFlag,omitempty"`
}

var tracesCmdLineOptions TracesOptions

// testLog is used only during testing to export the options for this test case.
func (opts *TracesOptions) testLog() {
	logger.TestLog(len(opts.Transactions) > 0, "Transactions: ", opts.Transactions)
	logger.TestLog(opts.Articulate, "Articulate: ", opts.Articulate)
	logger.TestLog(len(opts.Filter) > 0, "Filter: ", opts.Filter)
	logger.TestLog(opts.Statediff, "Statediff: ", opts.Statediff)
	logger.TestLog(opts.Count, "Count: ", opts.Count)
	logger.TestLog(opts.SkipDdos, "SkipDdos: ", opts.SkipDdos)
	logger.TestLog(opts.Max != 250, "Max: ", opts.Max)
	opts.Globals.TestLog()
}

// String implements the Stringer interface
func (opts *TracesOptions) String() string {
	b, _ := json.MarshalIndent(opts, "", "\t")
	return string(b)
}

// getEnvStr allows for custom environment strings when calling to the system (helps debugging).
func (opts *TracesOptions) getEnvStr() []string {
	envStr := []string{}
	// EXISTING_CODE
	// EXISTING_CODE
	return envStr
}

// toCmdLine converts the option to a command line for calling out to the system.
func (opts *TracesOptions) toCmdLine() string {
	options := ""
	if opts.Articulate {
		options += " --articulate"
	}
	if len(opts.Filter) > 0 {
		options += " --filter " + opts.Filter
	}
	if opts.Statediff {
		options += " --statediff"
	}
	if opts.Count {
		options += " --count"
	}
	if opts.SkipDdos {
		options += " --skip_ddos"
	}
	if opts.Max != 250 {
		options += (" --max " + fmt.Sprintf("%d", opts.Max))
	}
	options += " " + strings.Join(opts.Transactions, " ")
	// EXISTING_CODE
	// EXISTING_CODE
	options += fmt.Sprintf("%s", "") // silence go compiler for auto gen
	return options
}

// tracesFinishParseApi finishes the parsing for server invocations. Returns a new TracesOptions.
func tracesFinishParseApi(w http.ResponseWriter, r *http.Request) *TracesOptions {
	opts := &TracesOptions{}
	opts.Max = 250
	for key, value := range r.URL.Query() {
		switch key {
		case "transactions":
			for _, val := range value {
				s := strings.Split(val, " ") // may contain space separated items
				opts.Transactions = append(opts.Transactions, s...)
			}
		case "articulate":
			opts.Articulate = true
		case "filter":
			opts.Filter = value[0]
		case "statediff":
			opts.Statediff = true
		case "count":
			opts.Count = true
		case "skipDdos":
			opts.SkipDdos = true
		case "max":
			opts.Max = globals.ToUint64(value[0])
		default:
			if !globals.IsGlobalOption(key) {
				opts.BadFlag = validate.Usage("Invalid key ({0}) in {1} route.", key, "traces")
				return opts
			}
		}
	}
	opts.Globals = *globals.GlobalsFinishParseApi(w, r)
	// EXISTING_CODE
	// EXISTING_CODE

	return opts
}

// tracesFinishParse finishes the parsing for command line invocations. Returns a new TracesOptions.
func tracesFinishParse(args []string) *TracesOptions {
	opts := GetOptions()
	opts.Globals.FinishParse(args)
	defFmt := "txt"
	// EXISTING_CODE
	opts.Transactions = args
	// EXISTING_CODE
	if len(opts.Globals.Format) == 0 || opts.Globals.Format == "none" {
		opts.Globals.Format = defFmt
	}
	return opts
}

func GetOptions() *TracesOptions {
	// EXISTING_CODE
	// EXISTING_CODE
	return &tracesCmdLineOptions
}
