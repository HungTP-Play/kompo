package core

import (
	"errors"
	"strings"
)

type KompoOption struct {
	In         string // Input folder or file , Default = ./docker-compose.yaml
	Out        string // Output folder, Default = helm
	K8sVersion string // Kubernetes version, Default = 1.28
	Mode       string // Mode: apply, heml, kustomize
	Reserve    bool   // Default = false
}

func ValidateKompoMode(mode string) (bool, error) {
	modes := []string{
		"apply",
		"helm",
		"kustomize",
	}

	for i := 0; i < len(modes); i++ {
		if modes[i] == strings.ToLower(mode) {
			return true, nil
		}
	}

	return false, errors.New("[Kompo] -- Invalid Mode")
}

func ValidateKompoK8sSupportedVersion(k8sVersion string) (bool, error) {
	versions := []string{
		"1.24",
		"1.25",
		"1.26",
		"1.27",
		"1.28",
	}

	for i := 0; i < len(versions); i++ {
		if versions[i] == strings.ToLower(k8sVersion) {
			return true, nil
		}
	}

	return false, errors.New("[Kompo] -- Unsupported K8s Version")
}
