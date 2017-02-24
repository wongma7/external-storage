/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeExtensions
	ns   string
}

var networkpoliciesResource = schema.GroupVersionResource{Group: "extensions", Version: "", Resource: "networkpolicies"}

func (c *FakeNetworkPolicies) Create(networkPolicy *extensions.NetworkPolicy) (result *extensions.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(core.NewCreateAction(networkpoliciesResource, c.ns, networkPolicy), &extensions.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) Update(networkPolicy *extensions.NetworkPolicy) (result *extensions.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(core.NewUpdateAction(networkpoliciesResource, c.ns, networkPolicy), &extensions.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewDeleteAction(networkpoliciesResource, c.ns, name), &extensions.NetworkPolicy{})

	return err
}

func (c *FakeNetworkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := core.NewDeleteCollectionAction(networkpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &extensions.NetworkPolicyList{})
	return err
}

func (c *FakeNetworkPolicies) Get(name string, options v1.GetOptions) (result *extensions.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetAction(networkpoliciesResource, c.ns, name), &extensions.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) List(opts v1.ListOptions) (result *extensions.NetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewListAction(networkpoliciesResource, c.ns, opts), &extensions.NetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensions.NetworkPolicyList{}
	for _, item := range obj.(*extensions.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewWatchAction(networkpoliciesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched networkPolicy.
func (c *FakeNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(core.NewPatchSubresourceAction(networkpoliciesResource, c.ns, name, data, subresources...), &extensions.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), err
}
