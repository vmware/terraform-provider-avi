---
layout: "avi"
page_title: "Avi: avi_sslkeyandcertificate"
sidebar_current: "docs-avi-resource-sslkeyandcertificate"
description: |-
  Creates and manages Avi SSLKeyAndCertificate.
---

# avi_sslkeyandcertificate

The SSLKeyAndCertificate resource allows the creation and management of Avi SSLKeyAndCertificate

## Example Usage

```hcl
resource "avi_sslkeyandcertificate" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `certificate` - (Required) Dict settings for sslkeyandcertificate.
* `name` - (Required) Name of the object.
* `ca_certs` - (Optional) Ca certificates in certificate chain.
* `certificate_base64` - (Optional) States if the certificate is base64 encoded.
* `certificate_management_profile_ref` - (Optional) It is a reference to an object of type certificatemanagementprofile.
* `created_by` - (Optional) Creator name.
* `dynamic_params` - (Optional) Dynamic parameters needed for certificate management profile.
* `enable_ocsp_stapling` - (Optional) Enables ocsp stapling.
* `enckey_base64` - (Optional) Encrypted private key corresponding to the private key (e.g.
* `enckey_name` - (Optional) Name of the encrypted private key (e.g.
* `format` - (Optional) Format of the key/certificate file.
* `hardwaresecuritymodulegroup_ref` - (Optional) It is a reference to an object of type hardwaresecuritymodulegroup.
* `key` - (Optional) Private key.
* `key_base64` - (Optional) States if the private key is base64 encoded.
* `key_params` - (Optional) Dict settings for sslkeyandcertificate.
* `key_passphrase` - (Optional) Passphrase used to encrypt the private key.
* `ocsp_config` - (Optional) Configuration related to ocsp.
* `ocsp_error_status` - (Optional) Error reported during ocsp status query.
* `ocsp_response_info` - (Optional) Information related to ocsp response.
* `status` - (Optional) Enum options - ssl_certificate_finished, ssl_certificate_pending.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `type` - (Optional) Enum options - ssl_certificate_type_virtualservice, ssl_certificate_type_system, ssl_certificate_type_ca.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

