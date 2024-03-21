<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafapplicationsignatureprovider"
sidebar_current: "docs-avi-resource-wafapplicationsignatureprovider"
description: |-
  Creates and manages Avi WafApplicationSignatureProvider.
---

# avi_wafapplicationsignatureprovider

The WafApplicationSignatureProvider resource allows the creation and management of Avi WafApplicationSignatureProvider

## Example Usage

```hcl
resource "avi_wafapplicationsignatureprovider" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `available_applications` - (Optional) Available application names and the ruleset version, when the rules for an application changed the last time. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `filter_rules_on_import` - (Optional) If this is set to false, all provided rules are imported when updating this object. If this is set to true, only newer rules are considered for import. Newer rules are rules where the rule id is not in the range of 2,000,000 to 2,080,000 or where the rule has a tag with a cve from 2013 or newer. All other rules are ignored on rule import. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Optional) Name of application specific ruleset provider. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ruleset_version` - (Optional) Version of signatures. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `service_status` - (Optional) If this object is managed by the application signatures update service, this field contain the status of this syncronization. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `signatures` - (Optional) The waf rules. Not visible in the api. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

