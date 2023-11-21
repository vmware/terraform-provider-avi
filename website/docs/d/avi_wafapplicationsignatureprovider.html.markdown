<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafapplicationsignatureprovider"
sidebar_current: "docs-avi-datasource-wafapplicationsignatureprovider"
description: |-
  Get information of Avi WafApplicationSignatureProvider.
---

# avi_wafapplicationsignatureprovider

This data source is used to to get avi_wafapplicationsignatureprovider objects.

## Example Usage

```hcl
data "avi_wafapplicationsignatureprovider" "foo_wafapplicationsignatureprovider" {
    uuid = "wafapplicationsignatureprovider-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafApplicationSignatureProvider by name.
* `uuid` - (Optional) Search WafApplicationSignatureProvider by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `available_applications` - Available application names and the ruleset version, when the rules for an application changed the last time. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `filter_rules_on_import` - If this is set to false, all provided rules are imported when updating this object. If this is set to true, only newer rules are considered for import. Newer rules are rules where the rule id is not in the range of 2,000,000 to 2,080,000 or where the rule has a tag with a cve from 2013 or newer. All other rules are ignored on rule import. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of application specific ruleset provider. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ruleset_version` - Version of signatures. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `service_status` - If this object is managed by the application signatures update service, this field contain the status of this syncronization. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `signatures` - The waf rules. Not visible in the api. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

