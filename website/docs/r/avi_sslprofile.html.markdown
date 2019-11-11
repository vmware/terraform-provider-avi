---
layout: "avi"
page_title: "Avi: avi_sslprofile"
sidebar_current: "docs-avi-resource-sslprofile"
description: |-
  Creates and manages Avi SSLProfile.
---

# avi_sslprofile

The SSLProfile resource allows the creation and management of Avi SSLProfile

## Example Usage

```hcl
resource "avi_sslprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `accepted_ciphers` - (Optional) Ciphers suites represented as defined by u(http //www.openssl.org/docs/apps/ciphers.html).
* `accepted_versions` - (Optional) Set of versions accepted by the server.
* `cipher_enums` - (Optional) Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_rsa_with_rc4_128_sha, tls_aes_256_gcm_sha384, tls_chacha20_poly1305_sha256, tls_aes_128_gcm_sha256.
* `ciphersuites` - (Optional) Tls 1.3 ciphers suites represented as defined by u(https //www.openssl.org/docs/manmaster/man1/ciphers.html).
* `description` - (Optional) User defined description for the object.
* `dhparam` - (Optional) Dh parameters used in ssl.
* `enable_early_data` - (Optional) Enable early data processing for tls1.3 connections.
* `enable_ssl_session_reuse` - (Optional) Enable ssl session re-use.
* `prefer_client_cipher_ordering` - (Optional) Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile.
* `send_close_notify` - (Optional) Send 'close notify' alert message for a clean shutdown of the ssl connection.
* `ssl_rating` - (Optional) Dict settings for sslprofile.
* `ssl_session_timeout` - (Optional) The amount of time in seconds before an ssl session expires.
* `tags` - (Optional) List of list.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `type` - (Optional) Ssl profile type.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

