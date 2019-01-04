package main

import (
	"fmt"
)

// ManifestColumnItem Manifest Column Item
type ManifestColumnItem struct {
	ID int64
}

// NewManifestColumnItem Manifest Column Item from Hash
func NewManifestColumnItem(hash int64) (item *ManifestColumnItem, error error) {
	isValidHash := hash&(1<<(32-1)) != 0
	if isValidHash {
		return &ManifestColumnItem{hash - (1 << 32)}, nil
	}
	return nil, fmt.Errorf("unable to convert hash %v to manifest column item", hash)
}
