/*
Copyright The KubeVirt Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeCDIConfigs implements CDIConfigInterface
type FakeCDIConfigs struct {
	Fake *FakeCdiV1beta1
}

var cdiconfigsResource = v1beta1.SchemeGroupVersion.WithResource("cdiconfigs")

var cdiconfigsKind = v1beta1.SchemeGroupVersion.WithKind("CDIConfig")

// Get takes name of the cDIConfig, and returns the corresponding cDIConfig object, and an error if there is any.
func (c *FakeCDIConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CDIConfig, err error) {
	emptyResult := &v1beta1.CDIConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(cdiconfigsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CDIConfig), err
}

// List takes label and field selectors, and returns the list of CDIConfigs that match those selectors.
func (c *FakeCDIConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CDIConfigList, err error) {
	emptyResult := &v1beta1.CDIConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(cdiconfigsResource, cdiconfigsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CDIConfigList{ListMeta: obj.(*v1beta1.CDIConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.CDIConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cDIConfigs.
func (c *FakeCDIConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(cdiconfigsResource, opts))
}

// Create takes the representation of a cDIConfig and creates it.  Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *FakeCDIConfigs) Create(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.CreateOptions) (result *v1beta1.CDIConfig, err error) {
	emptyResult := &v1beta1.CDIConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(cdiconfigsResource, cDIConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CDIConfig), err
}

// Update takes the representation of a cDIConfig and updates it. Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *FakeCDIConfigs) Update(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (result *v1beta1.CDIConfig, err error) {
	emptyResult := &v1beta1.CDIConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(cdiconfigsResource, cDIConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CDIConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCDIConfigs) UpdateStatus(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (result *v1beta1.CDIConfig, err error) {
	emptyResult := &v1beta1.CDIConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(cdiconfigsResource, "status", cDIConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CDIConfig), err
}

// Delete takes name of the cDIConfig and deletes it. Returns an error if one occurs.
func (c *FakeCDIConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(cdiconfigsResource, name, opts), &v1beta1.CDIConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCDIConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(cdiconfigsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CDIConfigList{})
	return err
}

// Patch applies the patch and returns the patched cDIConfig.
func (c *FakeCDIConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CDIConfig, err error) {
	emptyResult := &v1beta1.CDIConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(cdiconfigsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CDIConfig), err
}