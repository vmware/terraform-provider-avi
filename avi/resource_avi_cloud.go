// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
	"github.com/vmware/alb-sdk/go/models"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ResourceCloudSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"autoscale_polling_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
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
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"custom_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"dhcp_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"dns_provider_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_resolution_on_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"dns_resolvers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsResolverSchema(),
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
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_vip_static_routes": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"gcp_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGCPConfigurationSchema(),
		},
		"ip6_autocfg_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
		"maintenance_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"metrics_polling_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"mtu": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1500",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nsxt_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceNsxtConfigurationSchema(),
		},
		"ntp_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceNTPConfigurationSchema(),
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
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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
		"vmc_deployment": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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

// Wait until cloud is ready for the placement
func waitForCloudState(cloudUUID string, expectedCloudState string, client *clients.AviClient, maxRetry int) error {
	var robj interface{}
	var err error
	var cloudState string
	path := "api/cloud-inventory?uuid=" + cloudUUID
	i := 0
	for ; i < maxRetry; i++ {
		if err = client.AviSession.Get(path, &robj); err == nil {
			if objCount := robj.(map[string]interface{})["count"].(float64); objCount == float64(1) {
				var resp *models.CloudInventory
				jsonString, marshalErr := json.Marshal(robj.(map[string]interface{})["results"].([]interface{})[0])
				if marshalErr != nil {
					log.Printf("[Error] Got error while marshaling the response from the cloud-inventory api. %s", marshalErr.Error())
				} else {
					unmarshalErr := json.Unmarshal(jsonString, &resp)
					if unmarshalErr != nil {
						log.Printf("[Error] Got error while unmarshaling the response from the cloud-inventory api. %s", unmarshalErr.Error())
					} else {
						if cloudState = *((*resp).Status.State); cloudState == expectedCloudState {
							log.Printf("Got expected cloud state %s", cloudState)
							break
						} else {
							log.Printf("Didn't get expected cloud state. Current cloud state is %s", cloudState)
						}
					}
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

func setupVcenterMgmtNetwork(d *schema.ResourceData, meta interface{}) error {
	//Setup management network for vcenter cloud
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
	if err := APIUpdate(d, meta, "cloud", s); err != nil {
		log.Printf("[Error] Got error for cloud create/update. Error: %s", err.Error())
		return err
	}
	uuid := d.Get("uuid").(string)
	if ok := strings.Contains(mgmtNetwork, "api/"); !ok {
		mgmtNetwork = "vimgrruntime?name=" + mgmtNetwork
		vcenterConfig.(*schema.Set).List()[0].(map[string]interface{})["management_network"] = mgmtNetwork
		if err := d.Set("vcenter_configuration", vcenterConfig); err != nil {
			return err
		}
		if err := waitForCloudState(uuid, "CLOUD_STATE_FAILED", client, maxRetry); err != nil {
			return err
		}
		if err := APIUpdate(d, meta, "cloud", s); err != nil {
			log.Printf("[Error] Got error for cloud create/update. Error: %s", err.Error())
			return err
		}
	}
	if err := waitForCloudState(uuid, "CLOUD_STATE_PLACEMENT_READY", client, maxRetry); err != nil {
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
	} else if err = APICreate(d, meta, "cloud", s); err == nil {
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
	} else if err = APIUpdate(d, meta, "cloud", s); err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "cloud")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
