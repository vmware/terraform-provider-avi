package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// WafPSMLocation waf p s m location
// swagger:model WafPSMLocation
type WafPSMLocation struct {

	// Freetext comment about this location. Field introduced in 19.1.1.
	Description *string `json:"description,omitempty"`

	// Location index, this is used to determine the order of the locations. Field introduced in 19.1.1.
	// Required: true
	Index *int32 `json:"index"`

	// Apply these rules only if the request is matching this description. Field introduced in 19.1.1.
	Match *WafPSMLocationMatch `json:"match,omitempty"`

	// User defined name for this location, if not set by the user, it is filled with data from the match. Field introduced in 19.1.1.
	Name *string `json:"name,omitempty"`

	// A list of rules which should be applied on this location. Field introduced in 19.1.1.
	Rules []*WafPSMRule `json:"rules,omitempty"`
}
