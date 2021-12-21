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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/azuredisk/v1alpha1"
)

// FakeAzVolumeAttachments implements AzVolumeAttachmentInterface
type FakeAzVolumeAttachments struct {
	Fake *FakeDiskV1alpha1
	ns   string
}

var azvolumeattachmentsResource = schema.GroupVersionResource{Group: "disk.csi.azure.com", Version: "v1alpha1", Resource: "azvolumeattachments"}

var azvolumeattachmentsKind = schema.GroupVersionKind{Group: "disk.csi.azure.com", Version: "v1alpha1", Kind: "AzVolumeAttachment"}

// Get takes name of the azVolumeAttachment, and returns the corresponding azVolumeAttachment object, and an error if there is any.
func (c *FakeAzVolumeAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azvolumeattachmentsResource, c.ns, name), &v1alpha1.AzVolumeAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzVolumeAttachment), err
}

// List takes label and field selectors, and returns the list of AzVolumeAttachments that match those selectors.
func (c *FakeAzVolumeAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzVolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azvolumeattachmentsResource, azvolumeattachmentsKind, c.ns, opts), &v1alpha1.AzVolumeAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzVolumeAttachmentList{ListMeta: obj.(*v1alpha1.AzVolumeAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzVolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azVolumeAttachments.
func (c *FakeAzVolumeAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azvolumeattachmentsResource, c.ns, opts))

}

// Create takes the representation of a azVolumeAttachment and creates it.  Returns the server's representation of the azVolumeAttachment, and an error, if there is any.
func (c *FakeAzVolumeAttachments) Create(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.CreateOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azvolumeattachmentsResource, c.ns, azVolumeAttachment), &v1alpha1.AzVolumeAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzVolumeAttachment), err
}

// Update takes the representation of a azVolumeAttachment and updates it. Returns the server's representation of the azVolumeAttachment, and an error, if there is any.
func (c *FakeAzVolumeAttachments) Update(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azvolumeattachmentsResource, c.ns, azVolumeAttachment), &v1alpha1.AzVolumeAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzVolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzVolumeAttachments) UpdateStatus(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (*v1alpha1.AzVolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azvolumeattachmentsResource, "status", c.ns, azVolumeAttachment), &v1alpha1.AzVolumeAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzVolumeAttachment), err
}

// Delete takes name of the azVolumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAzVolumeAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(azvolumeattachmentsResource, c.ns, name, opts), &v1alpha1.AzVolumeAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzVolumeAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azvolumeattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzVolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched azVolumeAttachment.
func (c *FakeAzVolumeAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azvolumeattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzVolumeAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzVolumeAttachment), err
}