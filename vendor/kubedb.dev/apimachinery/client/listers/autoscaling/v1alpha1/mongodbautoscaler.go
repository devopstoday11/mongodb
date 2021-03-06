/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubedb.dev/apimachinery/apis/autoscaling/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MongoDBAutoscalerLister helps list MongoDBAutoscalers.
type MongoDBAutoscalerLister interface {
	// List lists all MongoDBAutoscalers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBAutoscaler, err error)
	// MongoDBAutoscalers returns an object that can list and get MongoDBAutoscalers.
	MongoDBAutoscalers(namespace string) MongoDBAutoscalerNamespaceLister
	MongoDBAutoscalerListerExpansion
}

// mongoDBAutoscalerLister implements the MongoDBAutoscalerLister interface.
type mongoDBAutoscalerLister struct {
	indexer cache.Indexer
}

// NewMongoDBAutoscalerLister returns a new MongoDBAutoscalerLister.
func NewMongoDBAutoscalerLister(indexer cache.Indexer) MongoDBAutoscalerLister {
	return &mongoDBAutoscalerLister{indexer: indexer}
}

// List lists all MongoDBAutoscalers in the indexer.
func (s *mongoDBAutoscalerLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBAutoscaler))
	})
	return ret, err
}

// MongoDBAutoscalers returns an object that can list and get MongoDBAutoscalers.
func (s *mongoDBAutoscalerLister) MongoDBAutoscalers(namespace string) MongoDBAutoscalerNamespaceLister {
	return mongoDBAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MongoDBAutoscalerNamespaceLister helps list and get MongoDBAutoscalers.
type MongoDBAutoscalerNamespaceLister interface {
	// List lists all MongoDBAutoscalers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBAutoscaler, err error)
	// Get retrieves the MongoDBAutoscaler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MongoDBAutoscaler, error)
	MongoDBAutoscalerNamespaceListerExpansion
}

// mongoDBAutoscalerNamespaceLister implements the MongoDBAutoscalerNamespaceLister
// interface.
type mongoDBAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MongoDBAutoscalers in the indexer for a given namespace.
func (s mongoDBAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBAutoscaler))
	})
	return ret, err
}

// Get retrieves the MongoDBAutoscaler from the indexer for a given namespace and name.
func (s mongoDBAutoscalerNamespaceLister) Get(name string) (*v1alpha1.MongoDBAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mongodbautoscaler"), name)
	}
	return obj.(*v1alpha1.MongoDBAutoscaler), nil
}
