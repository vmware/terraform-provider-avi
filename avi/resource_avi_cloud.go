/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceCloudSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apic_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAPICConfigurationSchema(),
		},
		"apic_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"autoscale_polling_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"aws_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAwsConfigurationSchema(),
		},
		"azure_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAzureConfigurationSchema(),
		},
		"cloudstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceCloudStackConfigurationSchema(),
		},
		"custom_tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"dhcp_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"docker_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDockerConfigurationSchema(),
		},
		"east_west_dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"east_west_ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_vip_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"gcp_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGCPConfigurationSchema(),
		},
		"ip6_autocfg_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"license_tier": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"license_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"linuxserver_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceLinuxServerConfigurationSchema(),
		},
		"mtu": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"nsx_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceNsxConfigurationSchema(),
		},
		"obj_name_prefix": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"openstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceOpenStackConfigurationSchema(),
		},
		"oshiftk8s_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceOShiftK8SConfigurationSchema(),
		},
		"prefer_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"rancher_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRancherConfigurationSchema(),
		},
		"state_based_dns_registration": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vca_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcevCloudAirConfigurationSchema(),
		},
		"vcenter_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcevCenterConfigurationSchema(),
		},
		"vtype": &schema.Schema{
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
	err := ApiRead(d, meta, "cloud", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCloudCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := ApiCreateOrUpdate(d, meta, "cloud", s)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "cloud", s)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloud"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
