############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_wafpolicy"
sidebar_current: "docs-avi-resource-wafpolicy"
description: |-
  Creates and manages Avi WafPolicy.
---

# avi_wafpolicy

The WafPolicy resource allows the creation and management of Avi WafPolicy

## Example Usage

```hcl
resource "avi_wafpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Field introduced in 17.2.1.
* `allow_mode_delegation` - (Optional) Allow rules to overwrite the policy mode.
* `created_by` - (Optional) Creator name.
* `crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization.
* `description` - (Optional) Field introduced in 17.2.1.
* `enable_app_learning` - (Optional) Enable application learning for this waf policy.
* `failure_mode` - (Optional) Waf policy failure mode.
* `mode` - (Optional) Waf policy mode.
* `paranoia_level` - (Optional) Waf ruleset paranoia  mode.
* `positive_security_model` - (Optional) The positive security model.
* `post_crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization.
* `pre_crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `waf_crs_ref` - (Optional) Waf core ruleset used for the crs part of this policy.
* `waf_profile_ref` - (Optional) Waf profile for waf policy.
* `whitelist` - (Optional) A set of rules which describe conditions under which the request will bypass the waf.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.1.

