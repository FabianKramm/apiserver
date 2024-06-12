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

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	networkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
)

// ServiceCIDRsGetter has a method to return a ServiceCIDRInterface.
// A group's client should implement this interface.
type ServiceCIDRsGetter interface {
	ServiceCIDRs() ServiceCIDRInterface
}

// ServiceCIDRInterface has methods to work with ServiceCIDR resources.
type ServiceCIDRInterface interface {
	Create(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.CreateOptions) (*v1alpha1.ServiceCIDR, error)
	Update(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (*v1alpha1.ServiceCIDR, error)
	UpdateStatus(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (*v1alpha1.ServiceCIDR, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceCIDR, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceCIDRList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCIDR, err error)
	Apply(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error)
	ApplyStatus(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error)
	ServiceCIDRExpansion
}

// serviceCIDRs implements ServiceCIDRInterface
type serviceCIDRs struct {
	client rest.Interface
}

// newServiceCIDRs returns a ServiceCIDRs
func newServiceCIDRs(c *NetworkingV1alpha1Client) *serviceCIDRs {
	return &serviceCIDRs{
		client: c.RESTClient(),
	}
}

// Get takes name of the serviceCIDR, and returns the corresponding serviceCIDR object, and an error if there is any.
func (c *serviceCIDRs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCIDR, err error) {
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Get().
		Resource("servicecidrs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceCIDRs that match those selectors.
func (c *serviceCIDRs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCIDRList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for servicecidrs", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of ServiceCIDRs that match those selectors.
func (c *serviceCIDRs) list(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCIDRList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceCIDRList{}
	err = c.client.Get().
		Resource("servicecidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceCIDRs.
func (c *serviceCIDRs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("servicecidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceCIDR and creates it.  Returns the server's representation of the serviceCIDR, and an error, if there is any.
func (c *serviceCIDRs) Create(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.CreateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Post().
		Resource("servicecidrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCIDR).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceCIDR and updates it. Returns the server's representation of the serviceCIDR, and an error, if there is any.
func (c *serviceCIDRs) Update(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Put().
		Resource("servicecidrs").
		Name(serviceCIDR.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCIDR).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceCIDRs) UpdateStatus(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Put().
		Resource("servicecidrs").
		Name(serviceCIDR.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCIDR).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceCIDR and deletes it. Returns an error if one occurs.
func (c *serviceCIDRs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("servicecidrs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceCIDRs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("servicecidrs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceCIDR.
func (c *serviceCIDRs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCIDR, err error) {
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Patch(pt).
		Resource("servicecidrs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied serviceCIDR.
func (c *serviceCIDRs) Apply(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error) {
	if serviceCIDR == nil {
		return nil, fmt.Errorf("serviceCIDR provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(serviceCIDR)
	if err != nil {
		return nil, err
	}
	name := serviceCIDR.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCIDR.Name must be provided to Apply")
	}
	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("servicecidrs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *serviceCIDRs) ApplyStatus(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error) {
	if serviceCIDR == nil {
		return nil, fmt.Errorf("serviceCIDR provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(serviceCIDR)
	if err != nil {
		return nil, err
	}

	name := serviceCIDR.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCIDR.Name must be provided to Apply")
	}

	result = &v1alpha1.ServiceCIDR{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("servicecidrs").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
