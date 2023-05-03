<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_albservicesconfig"
sidebar_current: "docs-avi-datasource-albservicesconfig"
description: |-
  Get information of Avi ALBServicesConfig.
---

# avi_albservicesconfig

This data source is used to to get avi_albservicesconfig objects.

## Example Usage

```hcl
data "avi_albservicesconfig" "foo_albservicesconfig" {
    uuid = "albservicesconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ALBServicesConfig by name.
* `uuid` - (Optional) Search ALBServicesConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_signature_config` - Default values for application signature sync. Field introduced in 20.1.4. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `asset_contact` - Default contact for this controller cluster. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `case_config` - Default values for case management. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `feature_opt_in_status` - Features opt-in for pulse cloud services. Field introduced in 20.1.1.
* `ip_reputation_config` - Default values to be used for ip reputation sync. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `mode` - Mode helps log collection and upload. Enum options - MODE_UNKNOWN, SALESFORCE, SYSTEST, MYVMWARE. Field introduced in 20.1.2. Allowed in enterprise edition with any value, essentials edition(allowed values- salesforce,myvmware,systest), basic edition(allowed values- salesforce,myvmware,systest), enterprise with cloud services edition.
* `name` - Name of the albservicesconfig object. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `operations_config` - Operations configuration. Field introduced in 22.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `polling_interval` - Time interval in minutes. Allowed values are 5-60. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `portal_url` - The fqdn or ip address of the pulse cloud services. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `saas_licensing_config` - Saas licensing configuration. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `split_proxy_configuration` - Split proxy configuration to connect external pulse cloud services. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `use_split_proxy` - By default, pulse cloud services uses proxy added in system configuration. If it should use a separate proxy, set this flag to true and configure split proxy configuration. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `use_tls` - Secure the controller to pulse cloud services communication over tls. Field introduced in 20.1.3. Allowed in enterprise edition with any value, basic edition with any value, enterprise with cloud services edition.
* `user_agent_db_config` - Default values for user agent db service. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `uuid` - Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `waf_config` - Default values for waf management. Field introduced in 21.1.1. Allowed in essentials edition with any value, basic edition with any value, enterprise, enterprise with cloud services edition.

