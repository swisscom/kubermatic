// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AllowedRegistryLister helps list AllowedRegistries.
// All objects returned here must be treated as read-only.
type AllowedRegistryLister interface {
	// List lists all AllowedRegistries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error)
	// Get retrieves the AllowedRegistry from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AllowedRegistry, error)
	AllowedRegistryListerExpansion
}

// allowedRegistryLister implements the AllowedRegistryLister interface.
type allowedRegistryLister struct {
	indexer cache.Indexer
}

// NewAllowedRegistryLister returns a new AllowedRegistryLister.
func NewAllowedRegistryLister(indexer cache.Indexer) AllowedRegistryLister {
	return &allowedRegistryLister{indexer: indexer}
}

// List lists all AllowedRegistries in the indexer.
func (s *allowedRegistryLister) List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AllowedRegistry))
	})
	return ret, err
}

// Get retrieves the AllowedRegistry from the index for a given name.
func (s *allowedRegistryLister) Get(name string) (*v1.AllowedRegistry, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("allowedregistry"), name)
	}
	return obj.(*v1.AllowedRegistry), nil
}
