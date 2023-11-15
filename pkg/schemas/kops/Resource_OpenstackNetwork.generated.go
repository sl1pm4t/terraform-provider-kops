package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/sl1pm4t/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceOpenstackNetwork() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone_hints": OptionalList(String()),
			"ipv6_support_disabled":   OptionalBool(),
			"public_network_names":    OptionalList(String()),
			"internal_network_names":  OptionalList(String()),
			"address_sort_order":      OptionalString(),
		},
	}

	return res
}

func ExpandResourceOpenstackNetwork(in map[string]interface{}) kops.OpenstackNetwork {
	if in == nil {
		panic("expand OpenstackNetwork failure, in is nil")
	}
	return kops.OpenstackNetwork{
		AvailabilityZoneHints: func(in interface{}) []*string {
			return func(in interface{}) []*string {
				if in == nil {
					return nil
				}
				var out []*string
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *string {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in string) *string {
							return &in
						}(string(ExpandString(in)))
					}(in))
				}
				return out
			}(in)
		}(in["availability_zone_hints"]),
		IPv6SupportDisabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["ipv6_support_disabled"]),
		PublicNetworkNames: func(in interface{}) []*string {
			return func(in interface{}) []*string {
				if in == nil {
					return nil
				}
				var out []*string
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *string {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in string) *string {
							return &in
						}(string(ExpandString(in)))
					}(in))
				}
				return out
			}(in)
		}(in["public_network_names"]),
		InternalNetworkNames: func(in interface{}) []*string {
			return func(in interface{}) []*string {
				if in == nil {
					return nil
				}
				var out []*string
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) *string {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in string) *string {
							return &in
						}(string(ExpandString(in)))
					}(in))
				}
				return out
			}(in)
		}(in["internal_network_names"]),
		AddressSortOrder: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["address_sort_order"]),
	}
}

func FlattenResourceOpenstackNetworkInto(in kops.OpenstackNetwork, out map[string]interface{}) {
	out["availability_zone_hints"] = func(in []*string) interface{} {
		return func(in []*string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *string) interface{} {
					if in == nil {
						return nil
					}
					return func(in string) interface{} {
						return FlattenString(string(in))
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.AvailabilityZoneHints)
	out["ipv6_support_disabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.IPv6SupportDisabled)
	out["public_network_names"] = func(in []*string) interface{} {
		return func(in []*string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *string) interface{} {
					if in == nil {
						return nil
					}
					return func(in string) interface{} {
						return FlattenString(string(in))
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.PublicNetworkNames)
	out["internal_network_names"] = func(in []*string) interface{} {
		return func(in []*string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in *string) interface{} {
					if in == nil {
						return nil
					}
					return func(in string) interface{} {
						return FlattenString(string(in))
					}(*in)
				}(in))
			}
			return out
		}(in)
	}(in.InternalNetworkNames)
	out["address_sort_order"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.AddressSortOrder)
}

func FlattenResourceOpenstackNetwork(in kops.OpenstackNetwork) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceOpenstackNetworkInto(in, out)
	return out
}
