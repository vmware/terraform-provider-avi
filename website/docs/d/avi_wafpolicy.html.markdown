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

* `allow_mode_delegation` - Allow rules to overwrite the policy mode.
* `application_signatures` - Application specific signatures.
* `confidence_override` - Configure thresholds for confidence labels.
* `created_by` - Creator name.
* `crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `description` - Field introduced in 17.2.1.
* `enable_app_learning` - Enable application learning for this waf policy.
* `enable_auto_rule_updates` - Enable application learning based rule updates on the waf profile.
* `enable_regex_learning` - Enable dynamic regex generation for positive security model rules.
* `failure_mode` - Waf policy failure mode.
* `learning_params` - Parameters for tuning application learning.
* `min_confidence` - Minimum confidence label required for auto rule updates.
* `mode` - Waf policy mode.
* `name` - Field introduced in 17.2.1.
* `paranoia_level` - Waf ruleset paranoia  mode.
* `positive_security_model` - The positive security model.
* `post_crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `pre_crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.1.
* `waf_crs_ref` - Waf core ruleset used for the crs part of this policy.
* `waf_profile_ref` - Waf profile for waf policy.
* `whitelist` - A set of rules which describe conditions under which the request will bypass the waf.

