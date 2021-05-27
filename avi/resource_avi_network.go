// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configured_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSubnetSchema(),
		},
		"dhcp_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"exclude_discovered_subnets": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ip6_autocfg_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"synced_from_se": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		"vcenter_dvs": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkCreate,
		Read:   ResourceAviNetworkRead,
		Update: resourceAviNetworkUpdate,
		Delete: resourceAviNetworkDelete,
		Schema: ResourceNetworkSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkImporter,
		},
	}
}

func ResourceNetworkImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkSchema()
	return ResourceImporter(d, m, "network", s)
}

func ResourceAviNetworkRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := APIRead(d, meta, "network", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := APICreateOrUpdate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "network"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
