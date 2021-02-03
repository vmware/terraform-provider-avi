/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCloudSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apic_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAPICConfigurationSchema(),
		},
		"apic_mode": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"autoscale_polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"aws_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAwsConfigurationSchema(),
		},
		"azure_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAzureConfigurationSchema(),
		},
		"cloudstack_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCloudStackConfigurationSchema(),
		},
		"custom_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"dhcp_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dns_provider_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_resolution_on_se": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"docker_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDockerConfigurationSchema(),
		},
		"east_west_dns_provider_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"east_west_ipam_provider_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_vip_on_all_interfaces": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_vip_static_routes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"gcp_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGCPConfigurationSchema(),
		},
		"ip6_autocfg_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ipam_provider_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"license_tier": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"license_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"linuxserver_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLinuxServerConfigurationSchema(),
		},
		"mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nsx_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceNsxConfigurationSchema(),
		},
		"nsxt_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceNsxtConfigurationSchema(),
		},
		"obj_name_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"openstack_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOpenStackConfigurationSchema(),
		},
		"prefer_static_routes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"proxy_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"rancher_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRancherConfigurationSchema(),
		},
		"se_group_template_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state_based_dns_registration": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vca_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcevCloudAirConfigurationSchema(),
		},
		"vcenter_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcevCenterConfigurationSchema(),
		},
		"vtype": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudCreate,
		Read:   ResourceAviCloudRead,
		Update: resourceAviCloudUpdate,
		Delete: resourceAviCloudDelete,
		Schema: ResourceCloudSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCloudImporter,
		},
	}
}

func ResourceCloudImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCloudSchema()
	return ResourceImporter(d, m, "cloud", s)
}

func ResourceAviCloudRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := APIRead(d, meta, "cloud", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

// Verify cloudState is ready to configure mgmt IP
func isCloudReady(cloudState string) bool {
	switch cloudState {
	case
		"VCENTER_DISCOVERY_COMPLETE_NO_MGMT_NW",
		"VCENTER_DISCOVERY_COMPLETE",
		"VCENTER_DISCOVERY_WAITING_DC",
		"VCENTER_DISCOVERY_ONGOING":
		return true
	}
	return false
}

//wait until vcenter inventory gets completed to configure mgmt IP
func waitForVcenterState(cloudUUID string, client *clients.AviClient, maxRetry int) error {
	var robj interface{}
	var err error
	var inventoryState string
	path := "api/vimgrvcenterruntime?cloud_uuid=" + cloudUUID
	i := 0
	for ; i < maxRetry; i++ {
		if err = client.AviSession.Get(path, &robj); err == nil {
			if objCount := robj.(map[string]interface{})["count"].(float64); objCount == float64(1) {
				if inventoryState = robj.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["inventory_state"].(string); isCloudReady(inventoryState) {
					log.Printf("Got expected inventory state %s", inventoryState)
					break
				} else {
					log.Printf("Didn't get expected inventory state. Current state is %s", inventoryState)
				}
			} else {
				log.Printf("Inventory is not completed")
			}
		} else {
			log.Printf("[Error] Got error while retrieving vimgrvcenterruntime %s", err.Error())
		}
		time.Sleep(10 * time.Second)
	}
	if i == maxRetry && err == nil {
		err = errors.New("didn't get expected vcenter(vimgrvcenterruntime) inventory state . Current State: " + inventoryState)
	}
	return err
}

// Wait until cloud is ready for the placement
func waitForCloudReady(cloudUUID string, client *clients.AviClient, maxRetry int) error {
	var robj interface{}
	var err error
	var cloudState string
	path := "api/cloud-inventory?uuid=" + cloudUUID
	i := 0
	for ; i < maxRetry; i++ {
		if err = client.AviSession.Get(path, &robj); err == nil {
			if objCount := robj.(map[string]interface{})["count"].(float64); objCount == float64(1) {
				if cloudState = robj.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["status"].(map[string]interface{})["state"].(string); cloudState == "CLOUD_STATE_PLACEMENT_READY" {
					break
				} else {
					log.Printf("Didn't get expected cloud state. Current cloud state is %s", cloudState)
				}
			} else {
				log.Printf("Didn't get inventory for cloud")
			}
		} else {
			log.Printf("[Error] Got error while retrieving cloud-inventory %s", err.Error())
		}
		time.Sleep(10 * time.Second)
	}
	if i == maxRetry && err == nil {
		err = errors.New("didn't get expected state CLOUD_STATE_PLACEMENT_READY in cloud-inventory. Current State: " + cloudState)
	}
	return err
}

//Setup management network for vcenter cloud
func setupVcenterMgmtNetwork(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	var maxRetry int
	if retryCount, isRetry := os.LookupEnv("cloud_state_max_retry"); !isRetry {
		maxRetry = 50
	} else {
		intCount, err := strconv.Atoi(retryCount)
		if err != nil {
			maxRetry = 50
		} else {
			maxRetry = intCount
		}
	}
	client := meta.(*clients.AviClient)
	vcenterConfig, _ := d.GetOk("vcenter_configuration")
	mgmtNetwork := vcenterConfig.(*schema.Set).List()[0].(map[string]interface{})["management_network"].(string)
	mgmtNetwork = "vimgrruntime?name=" + mgmtNetwork
	if err := APICreateOrUpdate(d, meta, "cloud", s); err != nil {
		log.Printf("[Error] Got error for cloud create/update. Error: %s", err.Error())
		return err
	}
	uuid := d.Get("uuid").(string)
	if err := waitForVcenterState(uuid, client, maxRetry); err != nil {
		return err
	}
	vcenterConfig.(*schema.Set).List()[0].(map[string]interface{})["management_network"] = mgmtNetwork
	if err := d.Set("vcenter_configuration", vcenterConfig); err != nil {
		return err
	}
	if err := APICreateOrUpdate(d, meta, "cloud", s); err != nil {
		return err
	}
	if err := waitForCloudReady(uuid, client, maxRetry); err != nil {
		return err
	}
	return nil
}

func resourceAviCloudCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	var err error
	cloudType := d.Get("vtype")
	_, isVcenterConfig := d.GetOk("vcenter_configuration")
	if cloudType == "CLOUD_VCENTER" && isVcenterConfig {
		err = setupVcenterMgmtNetwork(d, meta)
	} else if err = APICreateOrUpdate(d, meta, "cloud", s); err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	var err error
	cloudType := d.Get("vtype")
	_, isVcenterConfig := d.GetOk("vcenter_configuration")
	if cloudType == "CLOUD_VCENTER" && isVcenterConfig {
		err = setupVcenterMgmtNetwork(d, meta)
	} else if err = APICreateOrUpdate(d, meta, "cloud", s); err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloud"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCloudDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
