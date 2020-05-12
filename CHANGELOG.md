## 0.2.2 (Unreleased)
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
