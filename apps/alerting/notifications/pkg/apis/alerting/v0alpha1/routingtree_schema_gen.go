//
// Code generated by grafana-app-sdk. DO NOT EDIT.
//

package v0alpha1

import (
	"github.com/grafana/grafana-app-sdk/resource"
)

// schema is unexported to prevent accidental overwrites
var (
	schemaRoutingTree = resource.NewSimpleSchema("notifications.alerting.grafana.app", "v0alpha1", &RoutingTree{}, &RoutingTreeList{}, resource.WithKind("RoutingTree"),
		resource.WithPlural("routingtrees"), resource.WithScope(resource.NamespacedScope))
	kindRoutingTree = resource.Kind{
		Schema: schemaRoutingTree,
		Codecs: map[resource.KindEncoding]resource.Codec{
			resource.KindEncodingJSON: &RoutingTreeJSONCodec{},
		},
	}
)

// Kind returns a resource.Kind for this Schema with a JSON codec
func RoutingTreeKind() resource.Kind {
	return kindRoutingTree
}

// Schema returns a resource.SimpleSchema representation of RoutingTree
func RoutingTreeSchema() *resource.SimpleSchema {
	return schemaRoutingTree
}

// Interface compliance checks
var _ resource.Schema = kindRoutingTree
