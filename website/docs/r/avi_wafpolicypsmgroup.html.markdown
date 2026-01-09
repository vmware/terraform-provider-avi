<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafpolicypsmgroup"
sidebar_current: "docs-avi-resource-wafpolicypsmgroup"
description: |-
  Creates and manages Avi WafPolicyPSMGroup.
---

# avi_wafpolicypsmgroup

The WafPolicyPSMGroup resource allows the creation and management of Avi WafPolicyPSMGroup

## Example Usage

```hcl
resource "avi_wafpolicypsmgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) User defined name of the group. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `completely_described_match_elements` - (Optional) A list of all match element collections which are completely described in the psm group. Every input value which matches one of the elements in this list but is not handled by a waf psm rule, will run the match_element miss_action. Allowed values are waf_variable_args. Enum options - WAF_VARIABLE_ARGS, WAF_VARIABLE_ARGS_GET, WAF_VARIABLE_ARGS_POST, WAF_VARIABLE_ARGS_NAMES, WAF_VARIABLE_REQUEST_COOKIES, WAF_VARIABLE_QUERY_STRING, WAF_VARIABLE_REQUEST_BASENAME, WAF_VARIABLE_REQUEST_URI, WAF_VARIABLE_PATH_INFO, WAF_VARIABLE_REQUEST_HEADERS. Field introduced in 31.2.1. Maximum of 1 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Free-text comment about this group. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enable` - (Optional) Enable or disable this waf rule group. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `hit_action` - (Optional) If a rule in this group matches the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_allow_parameter. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `is_learning_group` - (Optional) This field indicates that this group is used for learning. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `location_match_miss_action` - (Optional) If there is no location matching the request, this action will be executed. Allowed actions are waf_action_no_op and waf_action_block. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `locations` - (Optional) Positive security model locations. These are used to partition the application name space. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `miss_action` - (Optional) If a rule in this group does not match the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_block. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

