// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	clientset "k8s.io/metrics/pkg/client/clientset_generated/clientset"
)

func AnyPtrToClientsetClientset() *clientset.Clientset {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*clientset.Clientset))(nil)).Elem()))
	var nullValue *clientset.Clientset
	return nullValue
}

func EqPtrToClientsetClientset(value *clientset.Clientset) *clientset.Clientset {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *clientset.Clientset
	return nullValue
}
