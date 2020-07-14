// +kubebuilder:object:generate=true
// +groupName=bootes.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupVersion = schema.GroupVersion{Group: "bootes.io", Version: "v1"}

	ClusterKind  = "Cluster"
	ListenerKind = "Listener"
	RouteKind    = "Route"
	EndpointKind = "Endpoint"
)
