package walletstatus

import "github.com/RyzhAlexWork/go-intern/Task1/pkg/models"

// WalletStatus ...
type WalletStatus interface {
	Get() models.Status
	Change(newText models.Status)
}

type walletStatus struct {
	text models.Status
}

// Get return walletStatus text
func (w *walletStatus) Get() models.Status {
	return w.text
}

// Change change the walletStatus text
func (w *walletStatus) Change(newText models.Status) {
	w.text = newText
}

// NewWalletStatus create walletStatus implementation for interface WalletStatus
func NewWalletStatus() WalletStatus {
	return &walletStatus{}
}
