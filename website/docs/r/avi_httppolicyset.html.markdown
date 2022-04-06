<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_httppolicyset"
sidebar_current: "docs-avi-resource-httppolicyset"
description: |-
  Creates and manages Avi HTTPPolicySet.
---

# avi_httppolicyset

The HTTPPolicySet resource allows the creation and management of Avi HTTPPolicySet

## Example Usage

```hcl
resource "avi_httppolicyset" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the http policy set. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for pool. Internally set by cloud connector. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - (Optional) Creator name. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `geo_db_ref` - (Optional) Geo database. It is a reference to an object of type geodb. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `http_request_policy` - (Optional) Http request policy for the virtual service. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `http_response_policy` - (Optional) Http response policy for the virtual service. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `http_security_policy` - (Optional) Http security policy for the virtual service. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ip_reputation_db_ref` - (Optional) Ip reputation database. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `is_internal_policy` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the http policy set. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

