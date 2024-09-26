// Package utils is for shared utils
package utils

import (
	"context"
	"os"
	"time"

	"github.com/wingocard/serum"
	"github.com/wingocard/serum/secretprovider"
	"github.com/wingocard/serum/secretprovider/gsmanager"
)

const (
	serumTimeout    = 5 * time.Second
	localModeEnvKey = "LOCAL_MODE"
)

// PrimeEnv populates the environment with the values
// from file specified in envFilePath. It also decrypts any
// secrets using serum.
func PrimeEnv(envFilePath string, localMode bool) error {
	if localMode {
		if err := os.Setenv(localModeEnvKey, "true"); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), serumTimeout)
	defer cancel()

	ij, err := serum.NewInjector(
		serum.FromFile(envFilePath),
		serum.WithSecretProviderFunc(func() (secretprovider.SecretProvider, error) {
			if localMode {
				return nil, nil
			}

			return gsmanager.New(ctx)
		}),
	)
	if err != nil {
		return err
	}

	defer ij.Close()
	return ij.Inject(ctx)
}
