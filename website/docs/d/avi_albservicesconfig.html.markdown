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

* `app_signature_config` - Default values to be used for application signature sync. Field introduced in 20.1.4. Allowed in basic edition, essentials edition, enterprise edition.
* `asset_contact` - Information about the default contact for this controller cluster. Field introduced in 20.1.1.
* `case_config` - Default values to be used for pulse case management. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `feature_opt_in_status` - Information about the portal features opted in for controller. Field introduced in 20.1.1.
* `ip_reputation_config` - Default values to be used for ip reputation sync. Field introduced in 20.1.1.
* `mode` - Mode helps log collection and upload. Enum options - SALESFORCE, SYSTEST, MYVMWARE. Field introduced in 20.1.2. Allowed in basic(allowed values- salesforce,myvmware,systest) edition, essentials(allowed values- salesforce,myvmware,systest) edition, enterprise edition.
* `polling_interval` - Time interval in minutes. Allowed values are 5-60. Field introduced in 18.2.6.
* `portal_url` - The fqdn or ip address of the customer portal. Field introduced in 18.2.6.
* `split_proxy_configuration` - Split proxy configuration to connect external pulse services. Field introduced in 20.1.1.
* `use_split_proxy` - By default, pulse uses proxy added in system configuration. If pulse needs to use a seperate proxy, set this flag to true and configure split proxy configuration. Field introduced in 20.1.1.
* `use_tls` - Secure the controller to pulse communication over tls. Field introduced in 20.1.3. Allowed in basic edition, enterprise edition.
* `user_agent_db_config` - Default values to be used for user agent db service. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `uuid` - Field introduced in 18.2.6.
* `waf_config` - Default values to be used for pulse waf management. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.

