package ssh

import (
	"os"
	"path/filepath"

	sshconfig "github.com/kevinburke/ssh_config"
)

// Host represents an SSH host configuration
type Host struct {
	Name    string
	Address string
	User    string
	Port    string
}

// ReadConfig reads the SSH config file and returns a list of hosts
func ReadConfig() ([]Host, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".ssh", "config")
	f, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Host{}, nil
		}
		return nil, err
	}
	defer f.Close()

	cfg, err := sshconfig.Decode(f)
	if err != nil {
		return nil, err
	}

	var hosts []Host
	for _, host := range cfg.Hosts {
		// Skip patterns with wildcards
		if containsWildcard(host.Patterns) {
			continue
		}

		for _, pattern := range host.Patterns {
			h := Host{
				Name:    pattern.String(),
				Address: getConfigValue(cfg, pattern.String(), "Hostname"),
				User:    getConfigValue(cfg, pattern.String(), "User"),
				Port:    getConfigValue(cfg, pattern.String(), "Port"),
			}

			// Only add if we have either a hostname or user specified
			if h.Address != "" || h.User != "" {
				hosts = append(hosts, h)
			}
		}
	}

	return hosts, nil
}

// getConfigValue safely gets a value from the SSH config
func getConfigValue(cfg *sshconfig.Config, host, key string) string {
	val, err := cfg.Get(host, key)
	if err != nil || val == "" {
		return ""
	}
	return val
}

// containsWildcard checks if any of the patterns contain wildcards
func containsWildcard(patterns []*sshconfig.Pattern) bool {
	for _, p := range patterns {
		if p.String() == "*" {
			return true
		}
	}
	return false
}
