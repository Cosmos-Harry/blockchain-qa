package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const stateFileName = ".poll-state.json"

// PollState persists context between CLI commands so users don't need to
// manually copy-paste addresses and nonces between steps.
type PollState struct {
	FactoryAddress string `json:"factory_address,omitempty"`
	OracleAddress  string `json:"oracle_address,omitempty"`
	PollAddress    string `json:"poll_address,omitempty"`
	Choice         uint64 `json:"choice,omitempty"`
	Nonce          string `json:"nonce,omitempty"`
}

func stateFilePath() string {
	return filepath.Join(".", stateFileName)
}

func loadState() PollState {
	data, err := os.ReadFile(stateFilePath())
	if err != nil {
		return PollState{}
	}
	var s PollState
	if err := json.Unmarshal(data, &s); err != nil {
		return PollState{}
	}
	return s
}

func saveState(s PollState) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFilePath(), data, 0600)
}

// resolveFlag returns the flag value if set, otherwise falls back to the
// state field. Returns an error if neither is available.
func resolveFlag(flagValue, stateValue, name string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	if stateValue != "" {
		return stateValue, nil
	}
	return "", fmt.Errorf("%s not provided (use --%s flag or run the previous step first to auto-save it)", name, name)
}
