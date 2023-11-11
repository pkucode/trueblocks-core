package config

import "github.com/ethereum/go-ethereum/crypto"

type unchainedGroup struct {
	Comment            string `toml:"comment"`
	PreferredPublisher string `toml:"preferredPublisher,omitempty"`
	SmartContract      string `toml:"smartContract,omitempty"`
}

func GetUnchained() unchainedGroup {
	return GetRootConfig().Unchained
}

func HeaderHash(version string) []byte {
	return crypto.Keccak256([]byte(version))
}

func ExpectedVersion() string {
	return headerVersion
}

func GetPublisher(value string) string {
	if value == "" {
		value = GetUnchained().PreferredPublisher
		if value == "" {
			value = "publisher.unchainedindex.eth"
		}
	}
	return value
}

var headerVersion = "trueblocks-core@v2.0.0-release" //"trueblocks-core@v0.40.0" // "trueblocks-core@v2.0.0-release"

var VersionTags = map[string]string{
	"0x81ae14ba68e372bc9bd4a295b844abd8e72b1de10fcd706e624647701d911da1": "trueblocks-core@v0.40.0",
	"0x6fc0c6dd027719f456c1e50a329f6157767325aa937411fa6e7be9359d9e0046": "trueblocks-core@v2.0.0-release",
}
