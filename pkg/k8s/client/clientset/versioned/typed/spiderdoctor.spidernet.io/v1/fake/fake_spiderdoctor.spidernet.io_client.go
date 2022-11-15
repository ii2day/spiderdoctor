// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/spidernet-io/spiderdoctor/pkg/k8s/client/clientset/versioned/typed/spiderdoctor.spidernet.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpiderdoctorV1 struct {
	*testing.Fake
}

func (c *FakeSpiderdoctorV1) Netdnses() v1.NetdnsInterface {
	return &FakeNetdnses{c}
}

func (c *FakeSpiderdoctorV1) Nethttps() v1.NethttpInterface {
	return &FakeNethttps{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpiderdoctorV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
