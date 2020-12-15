// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	multicluster_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for MultiClusterRoleClient from Clientset
func MultiClusterRoleClientFromClientsetProvider(clients multicluster_solo_io_v1alpha1.Clientset) multicluster_solo_io_v1alpha1.MultiClusterRoleClient {
	return clients.MultiClusterRoles()
}

// Provider for MultiClusterRole Client from Client
func MultiClusterRoleClientProvider(client client.Client) multicluster_solo_io_v1alpha1.MultiClusterRoleClient {
	return multicluster_solo_io_v1alpha1.NewMultiClusterRoleClient(client)
}

type MultiClusterRoleClientFactory func(client client.Client) multicluster_solo_io_v1alpha1.MultiClusterRoleClient

func MultiClusterRoleClientFactoryProvider() MultiClusterRoleClientFactory {
	return MultiClusterRoleClientProvider
}

type MultiClusterRoleClientFromConfigFactory func(cfg *rest.Config) (multicluster_solo_io_v1alpha1.MultiClusterRoleClient, error)

func MultiClusterRoleClientFromConfigFactoryProvider() MultiClusterRoleClientFromConfigFactory {
	return func(cfg *rest.Config) (multicluster_solo_io_v1alpha1.MultiClusterRoleClient, error) {
		clients, err := multicluster_solo_io_v1alpha1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.MultiClusterRoles(), nil
	}
}