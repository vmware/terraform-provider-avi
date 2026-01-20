## 30.2.7 (Unreleased)
## 30.2.6 (Dec 05, 2025)
## 30.2.5 (Oct 14, 2025)
## 30.2.4 (Jul 11, 2025)
## 30.2.3 (Apr 3, 2025)
## 30.2.2 (Sept 11, 2024)
## 30.2.1 (May 08, 2024)
ENHANCEMENTS:
- Added new resource `avi_controllerportalregistration` [#553](https://github.com/vmware/terraform-provider-avi/pull/553)
- Added new resource `avi_serviceauthprofile` [#553](https://github.com/vmware/terraform-provider-avi/pull/553)
- Added new resource `avi_taskjournal` [#590](https://github.com/vmware/terraform-provider-avi/pull/590)

BUG FIXES:
- AV-202599 Resource avi_useraccount fails with generic error instead of providing user useful feedback
- Failure to create avi_useraccount resource. https://github.com/vmware/terraform-provider-avi/issues/314
- AV-187301: Entering wrong credentials on terraform using AVI as provider, while using Plan command breaks statefile
- AV-186452: AVI-provider for Terraform version 22.1.2 and 22.1.4 creates any object in the Admin tenant.
- AV-186737: Changes to certain subresources are not detected as changes by Terraform
- AV-187199: AVI terraform apply the resource object even if it is already exist without showing Error.
- Terraform Provider Avi overwriting existing resources. https://github.com/vmware/terraform-provider-avi/issues/358
- AV-194160 Terraform Provider crash during avi_serviceenginegroup destroy. https://github.com/vmware/terraform-provider-avi/issues/376

## 22.1.6 (March 08, 2024)
NOTE:

From version 22.1.6, updating system default configurations now requires first importing the existing configuration into the Terraform state.

Please refer [README.MD](https://github.com/vmware/terraform-provider-avi/tree/22.1.7)

ENHANCEMENTS:
- Added new resource `avi_systemreport` [#576](https://github.com/vmware/terraform-provider-avi/pull/576/files)
- Updated resources with newly added fields [#576](https://github.com/vmware/terraform-provider-avi/pull/576/files)

BUG FIXES:
- AV-187199: AVI terraform apply the resource object even if it is already exist without showing Error.
- Terraform Provider Avi overwriting existing resources. https://github.com/vmware/terraform-provider-avi/issues/358

## 22.1.5 (October 30, 2023)
BUG FIXES:
- AV-187301: Entering wrong credentials on terraform using AVI as provider, while using Plan command breaks statefile
- AV-186452: AVI-provider for Terraform version 22.1.2 and 22.1.4 creates any object in the Admin tenant.
- AV-186737: Changes to certain subresources are not detected as changes by Terraform

## 0.2.3 (June 16, 2020)
Features:
- Updated Avi Go SDK to latest.
- Update Avi 18.2.9 release

Bugs:
- AV-85581 Added seproperties into post not allowed list
- Fix for issue #65

## 0.2.2 (May 12, 2020)
FEATURES:
- Lazy Authentication support. The provider uses lazy authentication so Avi controllers can be created in a plan and it allows password reset.
- Added examples for VMware Horizon and VMware vsphere

BUG FIXES:
- AV-79347: Terraform: Unable to reconcile state of resources modified outside Terraform
- AV-76349: Avi terraform enable=false is ignored.

## 0.2.1 (November 20, 2019)
FEATURES:
- Added new resource and datasource for upgradestatussummary object.
- Added new resource and datasource for testsedatastorelevel1 object.
- Added new resource and datasource for testsedatastorelevel2 object.
- Added new resource and datasource for testsedatastorelevel3 object.

IMPROVEMENTS:
- Updated terraform provider avi with latest gosdk (AV-72035)
- Added cluster status field to cluster datasource and resource to monitor the cluster status.

BUG FIXES
- Fixed cluster deletion erroring out due to delete not allowed on cluster.

## 0.2.0 (October 01, 2019)
FEATURES:
- Added new resource and datasource for controllerportalregistration object.
- Added new resource and datasource for customerportalinfo object.
- Added new resource and datasource for image object. This is used in managing software updates to the Avi Controller.
- Added new resource and datasource for portalfileupload object.
- Added new datasource for fetching upgrade status.


IMPROVEMENTS:
- Updated terraform provider avi with latest gosdk (AV-69259). It has updated options for session handling and fixes the Avi Session and CSRF Token being set during login.
- Added api_timeout property to avi provider definition (AV-68694)


BUG FIXES:
- Miscallaneous bug fixes regarding handling of resources that do not have UUID like useraccount, cluster. It also fixed the crash that happens when any API does not return a JSON value by design.
- Cleanup Terraform VSVIP examples. It is recommended to use proper vsvip resource instead of embedding them inside Virtualservice resource which was supported for legacy purposes. (AV-64687)
- Go FMT changes.


## 0.1.2 (July 12, 2019)
## 0.1.1 (June 21, 2017)

NOTES:

Bumping the provider version to get around provider caching issues - still same functionality

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)