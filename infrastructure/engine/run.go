package engine

import (
	provider2 "github.com/simon/mengine/infrastructure/provider"
	"github.com/simon/mengine/infrastructure/support/chain"
)

type providerHandleFunc func(provider3 provider2.Provider) error

// Run first provider must be root provider
func Run(providers ...provider2.Provider) error {
	rootProvider := providers[0]
	if rootProvider.Name() != ProviderName {

	}
	var (
		providerFunc = func(fn providerHandleFunc) error {
			for i := range providers {
				if err := fn(providers[i]); err != nil {
					return err
				}
			}
			return nil
		}
	)

	return chain.BuildChain(func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
		if err := rootProvider.Prepare(); err != nil {
			return nil, err
		}
		err := providerFunc(func(provider3 provider2.Provider) error {
			return provider3.Prepare()
		})
		return next, err
	}, func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
		//if err := rootProvider.Run(); err != nil {
		//	return nil, err
		//}
		err := providerFunc(func(provider3 provider2.Provider) error {
			return provider3.Run()
		})
		if err != nil {
			return nil, err
		}
		return next, rootProvider.Run()
	}, func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
		err := providerFunc(func(provider3 provider2.Provider) error {
			return provider3.Shutdown()
		})
		if err != nil {
			return nil, err
		}
		return next, rootProvider.Shutdown()
	})
}
