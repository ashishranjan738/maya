package v1alpha2

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type artifact struct {
	object *unstructured.Unstructured
	yaml   string
}

type list struct {
	items []*artifact
}

type builder struct {
	artifact *artifact
}

func (l *builder) Build() *artifact {
	// ...
}

func (l *builder) Unstructured() (*unstructured.Unstructured, error) {
	// ...
}

type listBuilder struct {
	list *list
}

func (l *listBuilder) Build() *list {
	// ...
}

func (l *listBuilder) Unstructured() ([]*unstructured.Unstructured, error) {
	// ...
}
