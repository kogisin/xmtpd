package testutils

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fastjson"
	"github.com/xmtp/xmtpd/pkg/config"
)

const BLOCKCHAIN_RPC_URL = "http://localhost:7545"

// This is the private key that anvil has funded by default
// This is safe to commit
const TEST_PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

/*
*
In tests it's weirdly difficult to get the working directory of the project root.

Keep moving up the folder hierarchy until you find a go.mod
*
*/
func rootPath(t *testing.T) string {
	dir, err := os.Getwd()
	require.NoError(t, err)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir { // reached the root directory
			t.Fatal("Could not find the root directory")
		}
		dir = parent
	}
}

/*
*
Parse the JSON file at this location to get the deployed contract proxy address
*
*/
func getContractAddress(t *testing.T, fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read json: %v", err)
	}

	switch {
	case strings.Contains(fileName, "GroupMessages.json"):
		return fastjson.GetString(data, "addresses", "groupMessagesProxy")
	case strings.Contains(fileName, "IdentityUpdates.json"):
		return fastjson.GetString(data, "addresses", "identityUpdatesProxy")
	case strings.Contains(fileName, "XMTPNodeRegistry.json"):
		return fastjson.GetString(data, "addresses", "XMTPNodeRegistry")
	case strings.Contains(fileName, "RatesManager.json"):
		return fastjson.GetString(data, "addresses", "ratesManagerProxy")
	default:
		return ""
	}
}

func GetContractsOptions(t *testing.T) config.ContractsOptions {
	rootDir := rootPath(t)

	return config.ContractsOptions{
		RpcUrl: BLOCKCHAIN_RPC_URL,
		MessagesContractAddress: getContractAddress(
			t,
			path.Join(rootDir, "./contracts/config/anvil_localnet/GroupMessages.json"),
		),
		NodesContractAddress: getContractAddress(
			t,
			path.Join(rootDir, "./contracts/config/anvil_localnet/XMTPNodeRegistry.json"),
		),
		IdentityUpdatesContractAddress: getContractAddress(
			t,
			path.Join(rootDir, "./contracts/config/anvil_localnet/IdentityUpdates.json"),
		),
		RatesManagerContractAddress: getContractAddress(
			t,
			path.Join(rootDir, "./contracts/config/anvil_localnet/RatesManager.json"),
		),
		RegistryRefreshInterval: 100 * time.Millisecond,
		RatesRefreshInterval:    100 * time.Millisecond,
		ChainID:                 31337,
	}
}

func GetPayerOptions(t *testing.T) config.PayerOptions {
	return config.PayerOptions{
		PrivateKey: TEST_PRIVATE_KEY,
	}
}
