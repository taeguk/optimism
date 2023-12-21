package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var DevnetPresetId = 901

func DevnetPreset() (*Preset, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	root, err := findDevnetAllocs(cwd)
	if err != nil {
		return nil, err
	}

	devnetFilepath := filepath.Join(root, ".devnet", "addresses.json")
	if _, err := os.Stat(devnetFilepath); errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf(".devnet/addresses.json not found. `make devnet-allocs` in monorepo root: %w", err)
	}

	content, err := os.ReadFile(devnetFilepath)
	if err != nil {
		return nil, fmt.Errorf("unable to read .devnet/addressees.json")
	}

	var l1Contracts L1Contracts
	if err := json.Unmarshal(content, &l1Contracts); err != nil {
		return nil, err
	}

	return &Preset{
		Name:        "Local Devnet",
		ChainConfig: ChainConfig{Preset: DevnetPresetId, L1Contracts: l1Contracts},
	}, nil
}

// findDevnetAllocs will recursively search upwards for a .devnet directory.
// This depends on the structure of the monorepo generating the .devnet
// folder on devnet-allocs
func findDevnetAllocs(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	for {
		modulePath := filepath.Join(dir, ".devnet")
		if _, err := os.Stat(modulePath); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)

		// Check if we reached the filesystem root
		if parentDir == dir {
			break
		}

		dir = parentDir
	}

	return "", fmt.Errorf(".devnet dir not found. `make devnet-allocs in monorepo root")
}
