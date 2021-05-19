package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVICertificateManagementProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVICertificateManagementProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVICertificateManagementProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICertificateManagementProfileExists("avi_certificatemanagementprofile.testCertificateManagementProfile"),
					resource.TestCheckResourceAttr(
						"avi_certificatemanagementprofile.testCertificateManagementProfile", "name", "test-cert-test-abc"),
				),
			},
			{
				Config: testAccAVICertificateManagementProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICertificateManagementProfileExists("avi_certificatemanagementprofile.testCertificateManagementProfile"),
					resource.TestCheckResourceAttr(
						"avi_certificatemanagementprofile.testCertificateManagementProfile", "name", "test-cert-updated"),
				),
			},
			{
				ResourceName:      "avi_certificatemanagementprofile.testCertificateManagementProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVICertificateManagementProfileConfig,
			},
		},
	})

}

func testAccCheckAVICertificateManagementProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI CertificateManagementProfile ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVICertificateManagementProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_certificatemanagementprofile" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI CertificateManagementProfile still exists")
		}
	}
	return nil
}

const testAccAVICertificateManagementProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_alertscriptconfig" "testAlertScriptConfig" {
	name = "test-alertscriptconfig"
	tenant_ref = data.avi_tenant.default_tenant.id
	action_script = <<EOF
#!/usr/bin/python3
import os
import json
import sys
import time
from avi.sdk.avi_api import ApiSession
import urllib3
import requests

if hasattr(requests.packages.urllib3, 'disable_warnings'):
    requests.packages.urllib3.disable_warnings()

if hasattr(urllib3, 'disable_warnings'):
    urllib3.disable_warnings()


def ParseAviParams(argv):
    if len(argv) != 2:
        return
    alert_params = json.loads(argv[1])
    print(str(alert_params))
    return alert_params


def get_api_token():
    return os.environ.get('API_TOKEN')


def cleanup_old_spec_se(se_grp_uuid, session):
    if not se_grp_uuid:
        print('%s Empty SE Group UUID received' % time.ctime())
        return

    print('%s ===== CLEANUP OLD SPEC SE BEGIN =====' % time.ctime())

    rsp = session.get('serviceenginegroup/%s' % se_grp_uuid)
    if rsp.status_code != 200:
        print('%s Service Engine Group %s not found' % (time.ctime(), se_grp_uuid))
    se_grp_data = json.loads(rsp.content)
    print(se_grp_data)
    if 'instance_flavor' not in se_grp_data:
        print('%s No flavor information found in the SE Group data, '
              'nothing to do.' % (time.ctime()))
        return

    # Get all SE referring to this SE Group
    rsp = session.get('serviceengine?page_size=1000'
                      '&refers_to=serviceenginegroup:%s' % se_grp_uuid)
    if rsp.status_code != 200:
        print('%s Service Engines could not be retrieved, status: %s' %
              (time.ctime(), rsp.status_code))
    se_dict = json.loads(rsp.content)['results']
    old_spec_se = list()
    for se_conf in se_dict:
        if 'flavor' not in se_conf:
            continue
        if se_grp_data['instance_flavor'] != se_conf['flavor']:
            old_spec_se.append(se_conf)

    # Disable for placement all SE's using the older spec SE's
    disabled_placement_old_spec_se = list()
    for se_conf in old_spec_se:
        se_conf['enable_state'] = 'SE_STATE_DISABLED_FOR_PLACEMENT'
        rsp = session.put('serviceengine/%s/' % se_conf['uuid'], se_conf)
        if rsp.status_code != 200:
            print('%s Failed to disable for placement SE %s, status:%s, error:%s' %
                  (time.ctime(), se_conf['name'], rsp.status_code, rsp.content))
            continue
        disabled_placement_old_spec_se.append(se_conf)

    # Then, Disable all SE's using the older spec SE's
    disabled_old_spec_se = list()
    for se_conf in disabled_placement_old_spec_se:
        se_conf['enable_state'] = 'SE_STATE_DISABLED'
        rsp = session.put('serviceengine/%s/' % se_conf['uuid'], se_conf)
        if rsp.status_code != 200:
            print('%s Failed to disable for placement SE %s, status:%s, error:%s' %
                  (time.ctime(), se_conf['name'], rsp.status_code, rsp.content))
            continue
        disabled_old_spec_se.append(se_conf)

    # Wait for all SE to finish disabling
    for se_conf in disabled_old_spec_se:
        retry_count = 12
        sleep_interval = 10
        tries = 0
        while True:
            rsp = session.get('serviceengine/%s/runtime' % se_conf['uuid'])
            if rsp.status_code != 200:
                if rsp.status_code == 404:
                    print('%s SE %s no longer exists' % (time.ctime(), se_conf['name']))
                    break
                print('%s Failed to fetch runtime for SE %s' %
                      (time.ctime(), se_conf['name']))
            else:
                se_runtime = json.loads(rsp.content)
                if 'oper_status' in se_runtime and 'state' in se_runtime['oper_status']:
                    state = se_runtime['oper_status']['state']
                    if state != 'OPER_DISABLED':
                        print('%s SE %s not disabled yet, Oper state %s' %
                              (time.ctime(), se_conf['name'], state))
                    else:
                        print('%s SE %s has been disabled' %
                              (time.ctime(), se_conf['name']))
                        break
                else:
                    print('%s Oper state not found for SE %s' %
                          (time.ctime(), se_conf['name']))
            time.sleep(sleep_interval)
            tries += 1
            if tries > retry_count:
                print('%s Giving up disabling SE %s after %d seconds, trying next SE' %
                      (time.ctime(), se_conf['name'], retry_count * sleep_interval))
                break

    print('%s ===== CLEANUP OLD SPEC SE DONE =====' % time.ctime())


if __name__ == "__main__":
    alert_params = ParseAviParams(sys.argv)
    se_grp_uuid = None
    tenant_uuid = 'admin'
    for event in alert_params.get('events', []):
        if event['event_id'] == 'CONFIG_SE_GRP_FLAVOR_UPDATE':
            se_grp_uuid = event['obj_uuid']
            tenant_uuid = event['event_details']['config_se_grp_flv_update_details']['tenant_uuid']
    token = get_api_token()
    print('token = %s\n' % token)
    with ApiSession('localhost', 'admin', token=token,
                    tenant_uuid=tenant_uuid) as session:
        cleanup_old_spec_se(se_grp_uuid, session)
EOF
}
resource "avi_certificatemanagementprofile" "testCertificateManagementProfile" {
	name = "test-cert-test-abc"
	run_script_ref = avi_alertscriptconfig.testAlertScriptConfig.id
	tenant_ref = data.avi_tenant.default_tenant.id
}
`

const testAccAVICertificateManagementProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_alertscriptconfig" "testAlertScriptConfig" {
	name = "test-alertscriptconfig-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	action_script = <<EOF
#!/usr/bin/python3
import os
import json
import sys
import time
from avi.sdk.avi_api import ApiSession
import urllib3
import requests

if hasattr(requests.packages.urllib3, 'disable_warnings'):
    requests.packages.urllib3.disable_warnings()

if hasattr(urllib3, 'disable_warnings'):
    urllib3.disable_warnings()


def ParseAviParams(argv):
    if len(argv) != 2:
        return
    alert_params = json.loads(argv[1])
    print(str(alert_params))
    return alert_params


def get_api_token():
    return os.environ.get('API_TOKEN')


def cleanup_old_spec_se(se_grp_uuid, session):
    if not se_grp_uuid:
        print('%s Empty SE Group UUID received' % time.ctime())
        return

    print('%s ===== CLEANUP OLD SPEC SE BEGIN =====' % time.ctime())

    rsp = session.get('serviceenginegroup/%s' % se_grp_uuid)
    if rsp.status_code != 200:
        print('%s Service Engine Group %s not found' % (time.ctime(), se_grp_uuid))
    se_grp_data = json.loads(rsp.content)
    print(se_grp_data)
    if 'instance_flavor' not in se_grp_data:
        print('%s No flavor information found in the SE Group data, '
              'nothing to do.' % (time.ctime()))
        return

    # Get all SE referring to this SE Group
    rsp = session.get('serviceengine?page_size=1000'
                      '&refers_to=serviceenginegroup:%s' % se_grp_uuid)
    if rsp.status_code != 200:
        print('%s Service Engines could not be retrieved, status: %s' %
              (time.ctime(), rsp.status_code))
    se_dict = json.loads(rsp.content)['results']
    old_spec_se = list()
    for se_conf in se_dict:
        if 'flavor' not in se_conf:
            continue
        if se_grp_data['instance_flavor'] != se_conf['flavor']:
            old_spec_se.append(se_conf)

    # Disable for placement all SE's using the older spec SE's
    disabled_placement_old_spec_se = list()
    for se_conf in old_spec_se:
        se_conf['enable_state'] = 'SE_STATE_DISABLED_FOR_PLACEMENT'
        rsp = session.put('serviceengine/%s/' % se_conf['uuid'], se_conf)
        if rsp.status_code != 200:
            print('%s Failed to disable for placement SE %s, status:%s, error:%s' %
                  (time.ctime(), se_conf['name'], rsp.status_code, rsp.content))
            continue
        disabled_placement_old_spec_se.append(se_conf)

    # Then, Disable all SE's using the older spec SE's
    disabled_old_spec_se = list()
    for se_conf in disabled_placement_old_spec_se:
        se_conf['enable_state'] = 'SE_STATE_DISABLED'
        rsp = session.put('serviceengine/%s/' % se_conf['uuid'], se_conf)
        if rsp.status_code != 200:
            print('%s Failed to disable for placement SE %s, status:%s, error:%s' %
                  (time.ctime(), se_conf['name'], rsp.status_code, rsp.content))
            continue
        disabled_old_spec_se.append(se_conf)

    # Wait for all SE to finish disabling
    for se_conf in disabled_old_spec_se:
        retry_count = 12
        sleep_interval = 10
        tries = 0
        while True:
            rsp = session.get('serviceengine/%s/runtime' % se_conf['uuid'])
            if rsp.status_code != 200:
                if rsp.status_code == 404:
                    print('%s SE %s no longer exists' % (time.ctime(), se_conf['name']))
                    break
                print('%s Failed to fetch runtime for SE %s' %
                      (time.ctime(), se_conf['name']))
            else:
                se_runtime = json.loads(rsp.content)
                if 'oper_status' in se_runtime and 'state' in se_runtime['oper_status']:
                    state = se_runtime['oper_status']['state']
                    if state != 'OPER_DISABLED':
                        print('%s SE %s not disabled yet, Oper state %s' %
                              (time.ctime(), se_conf['name'], state))
                    else:
                        print('%s SE %s has been disabled' %
                              (time.ctime(), se_conf['name']))
                        break
                else:
                    print('%s Oper state not found for SE %s' %
                          (time.ctime(), se_conf['name']))
            time.sleep(sleep_interval)
            tries += 1
            if tries > retry_count:
                print('%s Giving up disabling SE %s after %d seconds, trying next SE' %
                      (time.ctime(), se_conf['name'], retry_count * sleep_interval))
                break

    print('%s ===== CLEANUP OLD SPEC SE DONE =====' % time.ctime())


if __name__ == "__main__":
    alert_params = ParseAviParams(sys.argv)
    se_grp_uuid = None
    tenant_uuid = 'admin'
    for event in alert_params.get('events', []):
        if event['event_id'] == 'CONFIG_SE_GRP_FLAVOR_UPDATE':
            se_grp_uuid = event['obj_uuid']
            tenant_uuid = event['event_details']['config_se_grp_flv_update_details']['tenant_uuid']
    token = get_api_token()
    print('token = %s\n' % token)
    with ApiSession('localhost', 'admin', token=token,
                    tenant_uuid=tenant_uuid) as session:
        cleanup_old_spec_se(se_grp_uuid, session)
EOF
}
resource "avi_certificatemanagementprofile" "testCertificateManagementProfile" {
	name = "test-cert-updated"
	run_script_ref = avi_alertscriptconfig.testAlertScriptConfig.id
	tenant_ref = data.avi_tenant.default_tenant.id
}
`
