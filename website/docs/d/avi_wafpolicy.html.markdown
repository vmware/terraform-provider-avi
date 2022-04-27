<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafpolicy"
sidebar_current: "docs-avi-datasource-wafpolicy"
description: |-
  Get information of Avi WafPolicy.
---

# avi_wafpolicy

This data source is used to to get avi_wafpolicy objects.

## Example Usage

```hcl
data "avi_wafpolicy" "foo_wafpolicy" {
    uuid = "wafpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafPolicy by name.
* `uuid` - (Optional) Search WafPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_mode_delegation` - Allow rules to overwrite the policy mode. This must be set if the policy mode is set to enforcement. Field introduced in 18.1.5, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allowlist` - A set of rules which describe conditions under which the request will bypass the waf. This will be processed in the request header phase before any other waf related code. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `application_signatures` - Application specific signatures. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `bypass_static_extensions` - Enable the functionality to bypass waf for static file extensions. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `confidence_override` - Configure thresholds for confidence labels. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - Creator name. Field introduced in 17.2.4. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crs_overrides` - Override attributes for crs rules. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `description` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_app_learning` - Enable application learning for this waf policy. Field introduced in 18.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_auto_rule_updates` - Enable application learning based rule updates on the waf profile. Rules will be programmed in dedicated waf learning group. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_regex_learning` - Enable dynamic regex generation for positive security model rules. This is an experimental feature and shouldn't be used in production. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `failure_mode` - Waf policy failure mode. This can be 'open' or 'closed'. Enum options - WAF_FAILURE_MODE_OPEN, WAF_FAILURE_MODE_CLOSED. Field introduced in 18.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `geo_db_ref` - Geo location mapping database used by this wafpolicy. It is a reference to an object of type geodb. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `learning_params` - Parameters for tuning application learning. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `min_confidence` - Minimum confidence label required for auto rule updates. Enum options - CONFIDENCE_VERY_HIGH, CONFIDENCE_HIGH, CONFIDENCE_PROBABLE, CONFIDENCE_LOW, CONFIDENCE_NONE. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `mode` - Waf policy mode. This can be detection or enforcement. It can be overwritten by rules if allow_mode_delegation is set. Enum options - WAF_MODE_DETECTION_ONLY, WAF_MODE_ENFORCEMENT. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `paranoia_level` - Waf ruleset paranoia  mode. This is used to select rules based on the paranoia-level tag. Enum options - WAF_PARANOIA_LEVEL_LOW, WAF_PARANOIA_LEVEL_MEDIUM, WAF_PARANOIA_LEVEL_HIGH, WAF_PARANOIA_LEVEL_EXTREME. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `positive_security_model` - The positive security model. This is used to describe how the request or parts of the request should look like. It is executed in the request body phase of avi waf. Field introduced in 18.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `post_crs_groups` - Waf rules are categorized in to groups based on their characterization. These groups are created by the user and will be enforced after the crs groups. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pre_crs_groups` - Waf rules are categorized in to groups based on their characterization. These groups are created by the user and will be  enforced before the crs groups. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `waf_crs_ref` - Waf core ruleset used for the crs part of this policy. It is a reference to an object of type wafcrs. Field introduced in 18.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `waf_profile_ref` - Waf profile for waf policy. It is a reference to an object of type wafprofile. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

