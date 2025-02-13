/*
Copyright 2022 The OpenFunction Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	corev1alpha2 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1alpha2"
	corev1beta1 "github.com/openfunction/pkg/client/clientset/versioned/typed/core/v1beta1"
	eventsv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/events/v1alpha1"
	networkingv1alpha1 "github.com/openfunction/pkg/client/clientset/versioned/typed/networking/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CoreV1alpha2() corev1alpha2.CoreV1alpha2Interface
	CoreV1beta1() corev1beta1.CoreV1beta1Interface
	EventsV1alpha1() eventsv1alpha1.EventsV1alpha1Interface
	NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	coreV1alpha2       *corev1alpha2.CoreV1alpha2Client
	coreV1beta1        *corev1beta1.CoreV1beta1Client
	eventsV1alpha1     *eventsv1alpha1.EventsV1alpha1Client
	networkingV1alpha1 *networkingv1alpha1.NetworkingV1alpha1Client
}

// CoreV1alpha2 retrieves the CoreV1alpha2Client
func (c *Clientset) CoreV1alpha2() corev1alpha2.CoreV1alpha2Interface {
	return c.coreV1alpha2
}

// CoreV1beta1 retrieves the CoreV1beta1Client
func (c *Clientset) CoreV1beta1() corev1beta1.CoreV1beta1Interface {
	return c.coreV1beta1
}

// EventsV1alpha1 retrieves the EventsV1alpha1Client
func (c *Clientset) EventsV1alpha1() eventsv1alpha1.EventsV1alpha1Interface {
	return c.eventsV1alpha1
}

// NetworkingV1alpha1 retrieves the NetworkingV1alpha1Client
func (c *Clientset) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return c.networkingV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.coreV1alpha2, err = corev1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1beta1, err = corev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventsV1alpha1, err = eventsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1alpha1, err = networkingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.coreV1alpha2 = corev1alpha2.NewForConfigOrDie(c)
	cs.coreV1beta1 = corev1beta1.NewForConfigOrDie(c)
	cs.eventsV1alpha1 = eventsv1alpha1.NewForConfigOrDie(c)
	cs.networkingV1alpha1 = networkingv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.coreV1alpha2 = corev1alpha2.New(c)
	cs.coreV1beta1 = corev1beta1.New(c)
	cs.eventsV1alpha1 = eventsv1alpha1.New(c)
	cs.networkingV1alpha1 = networkingv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
