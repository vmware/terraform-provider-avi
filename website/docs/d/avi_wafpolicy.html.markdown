
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

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
data "WafPolicy" "foo_WafPolicy" {
    uuid = "WafPolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafPolicy by name.
* `uuid` - (Optional) Search WafPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_mode_delegation` - Allow rules to overwrite the policy mode.
* `created_by` - Creator name.
* `crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `description` - Field introduced in 17.2.1.
* `failure_mode` - Waf policy failure mode.
* `learning` - Configure parameters for waf learning.
* `mode` - Waf policy mode.
* `name` - Field introduced in 17.2.1.
* `paranoia_level` - Waf ruleset paranoia  mode.
* `post_crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `pre_crs_groups` - Waf rules are categorized in to groups based on their characterization.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.1.
* `waf_crs_ref` - Waf core ruleset used for the crs part of this policy.
* `waf_profile_ref` - Waf profile for waf policy.

