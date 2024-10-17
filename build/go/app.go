package app

type App struct {
	BundleID    string `json:"bundleId,omitempty"`
	Environment string `json:"environment,omitempty"`
	Name        string `json:"name,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	Release     string `json:"release,omitempty"`
	Version     string `json:"version,omitempty"`
}
