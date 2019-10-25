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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConditions implements ConditionInterface
type FakeConditions struct {
	Fake *FakeTektonV1alpha1
	ns   string
}

var conditionsResource = schema.GroupVersionResource{Group: "tekton.dev", Version: "v1alpha1", Resource: "conditions"}

var conditionsKind = schema.GroupVersionKind{Group: "tekton.dev", Version: "v1alpha1", Kind: "Condition"}

// Get takes name of the condition, and returns the corresponding condition object, and an error if there is any.
func (c *FakeConditions) Get(name string, options v1.GetOptions) (result *v1alpha1.Condition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(conditionsResource, c.ns, name), &v1alpha1.Condition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Condition), err
}

// List takes label and field selectors, and returns the list of Conditions that match those selectors.
func (c *FakeConditions) List(opts v1.ListOptions) (result *v1alpha1.ConditionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(conditionsResource, conditionsKind, c.ns, opts), &v1alpha1.ConditionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConditionList{ListMeta: obj.(*v1alpha1.ConditionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConditionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested conditions.
func (c *FakeConditions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(conditionsResource, c.ns, opts))

}

// Create takes the representation of a condition and creates it.  Returns the server's representation of the condition, and an error, if there is any.
func (c *FakeConditions) Create(condition *v1alpha1.Condition) (result *v1alpha1.Condition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(conditionsResource, c.ns, condition), &v1alpha1.Condition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Condition), err
}

// Update takes the representation of a condition and updates it. Returns the server's representation of the condition, and an error, if there is any.
func (c *FakeConditions) Update(condition *v1alpha1.Condition) (result *v1alpha1.Condition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(conditionsResource, c.ns, condition), &v1alpha1.Condition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Condition), err
}

// Delete takes name of the condition and deletes it. Returns an error if one occurs.
func (c *FakeConditions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(conditionsResource, c.ns, name), &v1alpha1.Condition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConditions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(conditionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConditionList{})
	return err
}

// Patch applies the patch and returns the patched condition.
func (c *FakeConditions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Condition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(conditionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Condition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Condition), err
}
