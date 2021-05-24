// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceVrfContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"bfd_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBfdProfileSchema(),
		},
		"bgp_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBgpProfileSchema(),
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"debugvrfcontext": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDebugVrfContextSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gateway_mon": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGatewayMonitorSchema(),
		},
		"internal_gateway_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceInternalGatewayMonitorSchema(),
		},
		"lldp_enable": {
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
		"static_routes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceStaticRouteSchema(),
		},
		"system_default": {
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
	}
}

func resourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVrfContextCreate,
		Read:   ResourceAviVrfContextRead,
		Update: resourceAviVrfContextUpdate,
		Delete: resourceAviVrfContextDelete,
		Schema: ResourceVrfContextSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVrfContextImporter,
		},
	}
}

func ResourceVrfContextImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVrfContextSchema()
	return ResourceImporter(d, m, "vrfcontext", s)
}

func ResourceAviVrfContextRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := APIRead(d, meta, "vrfcontext", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVrfContextCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := APICreateOrUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vrfcontext"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviVrfContextDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
