package manifest

import (
	"fmt"
)

// ColumnItem Manifest Column Item
type ColumnItem struct {
	ID int64
}

// NewManifestColumnItem Manifest Column Item from Hash
func NewManifestColumnItem(hash int64) (item *ColumnItem, error error) {
	isValidHash := hash&(1<<(32-1)) != 0
	if isValidHash {
		return &ColumnItem{hash - (1 << 32)}, nil
	}
	return nil, fmt.Errorf("unable to convert hash %v to manifest column item", hash)
}
