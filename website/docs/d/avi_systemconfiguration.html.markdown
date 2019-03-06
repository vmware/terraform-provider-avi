
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
page_title: "AVI: avi_systemconfiguration"
sidebar_current: "docs-avi-datasource-systemconfiguration"
description: |-
  Get information of Avi SystemConfiguration.
---

# avi_systemconfiguration

This data source is used to to get avi_systemconfiguration objects.

## Example Usage

```hcl
data "SystemConfiguration" "foo_SystemConfiguration" {
    uuid = "SystemConfiguration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SystemConfiguration by name.
* `uuid` - (Optional) Search SystemConfiguration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `admin_auth_configuration` - General description.
* `default_license_tier` - Specifies the default license tier which would be used by new clouds.
* `dns_configuration` - General description.
* `dns_virtualservice_refs` - Dns virtualservices hosting fqdn records for applications across avi vantage.
* `docker_mode` - General description.
* `email_configuration` - General description.
* `global_tenant_config` - General description.
* `linux_configuration` - General description.
* `mgmt_ip_access_control` - Configure ip access control for controller to restrict open access.
* `ntp_configuration` - General description.
* `portal_configuration` - General description.
* `proxy_configuration` - General description.
* `secure_channel_configuration` - Configure secure channel properties.
* `snmp_configuration` - General description.
* `ssh_ciphers` - Allowed ciphers list for ssh to the management interface on the controller and service engines.
* `ssh_hmacs` - Allowed hmac list for ssh to the management interface on the controller and service engines.
* `uuid` - General description.

