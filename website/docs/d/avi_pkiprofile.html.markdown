
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
page_title: "AVI: avi_pkiprofile"
sidebar_current: "docs-avi-datasource-pkiprofile"
description: |-
  Get information of Avi PKIProfile.
---

# avi_pkiprofile

This data source is used to to get avi_pkiprofile objects.

## Example Usage

```hcl
data "PKIProfile" "foo_PKIProfile" {
    uuid = "PKIProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PKIProfile by name.
* `uuid` - (Optional) Search PKIProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ca_certs` - List of certificate authorities (root and intermediate) trusted that is used for certificate validation.
* `created_by` - Creator name.
* `crl_check` - When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked.
* `crls` - Certificate revocation lists.
* `ignore_peer_chain` - When enabled, avi will not trust intermediate and root certs presented by a client.
* `is_federated` - This field describes the object's replication scope.
* `name` - Name of the pki profile.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.
* `validate_only_leaf_crl` - When enabled, avi will only validate the revocation status of the leaf certificate using crl.

