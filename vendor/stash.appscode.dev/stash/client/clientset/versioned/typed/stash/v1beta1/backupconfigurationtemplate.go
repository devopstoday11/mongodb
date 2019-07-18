/*
Copyright 2019 The Stash Authors.

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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "stash.appscode.dev/stash/apis/stash/v1beta1"
	scheme "stash.appscode.dev/stash/client/clientset/versioned/scheme"
)

// BackupConfigurationTemplatesGetter has a method to return a BackupConfigurationTemplateInterface.
// A group's client should implement this interface.
type BackupConfigurationTemplatesGetter interface {
	BackupConfigurationTemplates() BackupConfigurationTemplateInterface
}

// BackupConfigurationTemplateInterface has methods to work with BackupConfigurationTemplate resources.
type BackupConfigurationTemplateInterface interface {
	Create(*v1beta1.BackupConfigurationTemplate) (*v1beta1.BackupConfigurationTemplate, error)
	Update(*v1beta1.BackupConfigurationTemplate) (*v1beta1.BackupConfigurationTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.BackupConfigurationTemplate, error)
	List(opts v1.ListOptions) (*v1beta1.BackupConfigurationTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupConfigurationTemplate, err error)
	BackupConfigurationTemplateExpansion
}

// backupConfigurationTemplates implements BackupConfigurationTemplateInterface
type backupConfigurationTemplates struct {
	client rest.Interface
}

// newBackupConfigurationTemplates returns a BackupConfigurationTemplates
func newBackupConfigurationTemplates(c *StashV1beta1Client) *backupConfigurationTemplates {
	return &backupConfigurationTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the backupConfigurationTemplate, and returns the corresponding backupConfigurationTemplate object, and an error if there is any.
func (c *backupConfigurationTemplates) Get(name string, options v1.GetOptions) (result *v1beta1.BackupConfigurationTemplate, err error) {
	result = &v1beta1.BackupConfigurationTemplate{}
	err = c.client.Get().
		Resource("backupconfigurationtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupConfigurationTemplates that match those selectors.
func (c *backupConfigurationTemplates) List(opts v1.ListOptions) (result *v1beta1.BackupConfigurationTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackupConfigurationTemplateList{}
	err = c.client.Get().
		Resource("backupconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupConfigurationTemplates.
func (c *backupConfigurationTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backupconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a backupConfigurationTemplate and creates it.  Returns the server's representation of the backupConfigurationTemplate, and an error, if there is any.
func (c *backupConfigurationTemplates) Create(backupConfigurationTemplate *v1beta1.BackupConfigurationTemplate) (result *v1beta1.BackupConfigurationTemplate, err error) {
	result = &v1beta1.BackupConfigurationTemplate{}
	err = c.client.Post().
		Resource("backupconfigurationtemplates").
		Body(backupConfigurationTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupConfigurationTemplate and updates it. Returns the server's representation of the backupConfigurationTemplate, and an error, if there is any.
func (c *backupConfigurationTemplates) Update(backupConfigurationTemplate *v1beta1.BackupConfigurationTemplate) (result *v1beta1.BackupConfigurationTemplate, err error) {
	result = &v1beta1.BackupConfigurationTemplate{}
	err = c.client.Put().
		Resource("backupconfigurationtemplates").
		Name(backupConfigurationTemplate.Name).
		Body(backupConfigurationTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupConfigurationTemplate and deletes it. Returns an error if one occurs.
func (c *backupConfigurationTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("backupconfigurationtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupConfigurationTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("backupconfigurationtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupConfigurationTemplate.
func (c *backupConfigurationTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupConfigurationTemplate, err error) {
	result = &v1beta1.BackupConfigurationTemplate{}
	err = c.client.Patch(pt).
		Resource("backupconfigurationtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}