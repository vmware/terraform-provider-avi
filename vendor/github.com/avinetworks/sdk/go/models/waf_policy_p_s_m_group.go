package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// WafPolicyPSMGroup waf policy p s m group
// swagger:model WafPolicyPSMGroup
type WafPolicyPSMGroup struct {

	// UNIX time since epoch in microseconds. Units(MICROSECONDS).
	// Read Only: true
	LastModified *string `json:"_last_modified,omitempty"`

	// Freetext comment about this group. Field introduced in 19.1.1.
	Description *string `json:"description,omitempty"`

	// Enable or disable this WAF rule group. Field introduced in 19.1.1.
	Enable *bool `json:"enable,omitempty"`

	// If a rule in this group matches the match_value pattern, this action will be executed. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 19.1.1.
	HitAction *string `json:"hit_action,omitempty"`

	// Positive Security Model locations. These are used to partition the application name space. Field introduced in 19.1.1.
	Locations []*WafPSMLocation `json:"locations,omitempty"`

	// If a rule in this group does not match the match_value pattern, this action will be executed. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 19.1.1.
	MissAction *string `json:"miss_action,omitempty"`

	// User defined name of the group. Field introduced in 19.1.1.
	// Required: true
	Name *string `json:"name"`

	// Tenant that this object belongs to. It is a reference to an object of type Tenant. Field introduced in 19.1.1.
	TenantRef *string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL *string `json:"url,omitempty"`

	// UUID of this object. Field introduced in 19.1.1.
	UUID *string `json:"uuid,omitempty"`
}
