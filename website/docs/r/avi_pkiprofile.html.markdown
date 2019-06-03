---
layout: "avi"
page_title: "Avi: avi_pkiprofile"
sidebar_current: "docs-avi-resource-pkiprofile"
description: |-
  Creates and manages Avi PKIProfile.
---

# avi_pkiprofile

The PKIProfile resource allows the creation and management of Avi PKIProfile

## Example Usage

```hcl
resource "avi_pkiprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the pki profile.
* `ca_certs` - (Optional) List of certificate authorities (root and intermediate) trusted that is used for certificate validation.
* `created_by` - (Optional) Creator name.
* `crl_check` - (Optional) When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked.
* `crls` - (Optional) Certificate revocation lists.
* `ignore_peer_chain` - (Optional) When enabled, avi will not trust intermediate and root certs presented by a client.
* `is_federated` - (Optional) This field describes the object's replication scope.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `validate_only_leaf_crl` - (Optional) When enabled, avi will only validate the revocation status of the leaf certificate using crl.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

