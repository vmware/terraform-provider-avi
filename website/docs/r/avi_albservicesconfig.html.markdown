<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_albservicesconfig"
sidebar_current: "docs-avi-resource-albservicesconfig"
description: |-
  Creates and manages Avi ALBServicesConfig.
---

# avi_albservicesconfig

The ALBServicesConfig resource allows the creation and management of Avi ALBServicesConfig

## Example Usage

```hcl
resource "avi_albservicesconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `app_signature_config` - (Required) Default values for application signature sync. Field introduced in 20.1.4. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `case_config` - (Required) Default values for case management. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `feature_opt_in_status` - (Required) Features opt-in for pulse cloud services. Field introduced in 20.1.1.
* `ip_reputation_config` - (Required) Default values to be used for ip reputation sync. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `portal_url` - (Required) The fqdn or ip address of the pulse cloud services. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `saas_licensing_config` - (Required) Saas licensing configuration. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `user_agent_db_config` - (Required) Default values for user agent db service. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `waf_config` - (Required) Default values for waf management. Field introduced in 21.1.1. Allowed in essentials edition with any value, basic edition with any value, enterprise, enterprise with cloud services edition.
* `asset_contact` - (Optional) Default contact for this controller cluster. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `mode` - (Optional) Mode helps log collection and upload. Enum options - MODE_UNKNOWN, SALESFORCE, SYSTEST, MYVMWARE. Field introduced in 20.1.2. Allowed in enterprise edition with any value, essentials edition(allowed values- salesforce,myvmware,systest), basic edition(allowed values- salesforce,myvmware,systest), enterprise with cloud services edition.
* `name` - (Optional) Name of the albservicesconfig object. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `polling_interval` - (Optional) Time interval in minutes. Allowed values are 5-60. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `split_proxy_configuration` - (Optional) Split proxy configuration to connect external pulse cloud services. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `use_split_proxy` - (Optional) By default, pulse cloud services uses proxy added in system configuration. If it should use a separate proxy, set this flag to true and configure split proxy configuration. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `use_tls` - (Optional) Secure the controller to pulse cloud services communication over tls. Field introduced in 20.1.3. Allowed in enterprise edition with any value, basic edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

