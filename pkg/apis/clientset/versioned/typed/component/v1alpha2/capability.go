/*
Copyright The Kubernetes Authors.

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

package v1alpha2

import (
	"time"

	scheme "github.com/snowdrop/component-api/pkg/apis/clientset/versioned/scheme"
	v1alpha2 "github.com/snowdrop/component-api/pkg/apis/component/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CapabilitiesGetter has a method to return a CapabilityInterface.
// A group's client should implement this interface.
type CapabilitiesGetter interface {
	Capabilities(namespace string) CapabilityInterface
}

// CapabilityInterface has methods to work with Capability resources.
type CapabilityInterface interface {
	Create(*v1alpha2.Capability) (*v1alpha2.Capability, error)
	Update(*v1alpha2.Capability) (*v1alpha2.Capability, error)
	UpdateStatus(*v1alpha2.Capability) (*v1alpha2.Capability, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.Capability, error)
	List(opts v1.ListOptions) (*v1alpha2.CapabilityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Capability, err error)
	CapabilityExpansion
}

// capabilities implements CapabilityInterface
type capabilities struct {
	client rest.Interface
	ns     string
}

// newCapabilities returns a Capabilities
func newCapabilities(c *DevexpV1alpha2Client, namespace string) *capabilities {
	return &capabilities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the capability, and returns the corresponding capability object, and an error if there is any.
func (c *capabilities) Get(name string, options v1.GetOptions) (result *v1alpha2.Capability, err error) {
	result = &v1alpha2.Capability{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("capabilities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Capabilities that match those selectors.
func (c *capabilities) List(opts v1.ListOptions) (result *v1alpha2.CapabilityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.CapabilityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("capabilities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested capabilities.
func (c *capabilities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("capabilities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a capability and creates it.  Returns the server's representation of the capability, and an error, if there is any.
func (c *capabilities) Create(capability *v1alpha2.Capability) (result *v1alpha2.Capability, err error) {
	result = &v1alpha2.Capability{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("capabilities").
		Body(capability).
		Do().
		Into(result)
	return
}

// Update takes the representation of a capability and updates it. Returns the server's representation of the capability, and an error, if there is any.
func (c *capabilities) Update(capability *v1alpha2.Capability) (result *v1alpha2.Capability, err error) {
	result = &v1alpha2.Capability{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("capabilities").
		Name(capability.Name).
		Body(capability).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *capabilities) UpdateStatus(capability *v1alpha2.Capability) (result *v1alpha2.Capability, err error) {
	result = &v1alpha2.Capability{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("capabilities").
		Name(capability.Name).
		SubResource("status").
		Body(capability).
		Do().
		Into(result)
	return
}

// Delete takes name of the capability and deletes it. Returns an error if one occurs.
func (c *capabilities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("capabilities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *capabilities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("capabilities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched capability.
func (c *capabilities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Capability, err error) {
	result = &v1alpha2.Capability{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("capabilities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
