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

func ResourceNsxtSegmentRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dhcp6_ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"dhcp_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"dhcp_ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nw_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nw_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"opaque_network_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segment_gw": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segment_gw6": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segment_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segname": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"subnet": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"subnet6": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tier1_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNsxtSegmentRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNsxtSegmentRuntimeCreate,
		Read:   ResourceAviNsxtSegmentRuntimeRead,
		Update: resourceAviNsxtSegmentRuntimeUpdate,
		Delete: resourceAviNsxtSegmentRuntimeDelete,
		Schema: ResourceNsxtSegmentRuntimeSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNsxtSegmentRuntimeImporter,
		},
	}
}

func ResourceNsxtSegmentRuntimeImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNsxtSegmentRuntimeSchema()
	return ResourceImporter(d, m, "nsxtsegmentruntime", s)
}

func ResourceAviNsxtSegmentRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	err := ApiRead(d, meta, "nsxtsegmentruntime", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	err := ApiCreateOrUpdate(d, meta, "nsxtsegmentruntime", s)
	if err == nil {
		err = ResourceAviNsxtSegmentRuntimeRead(d, meta)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "nsxtsegmentruntime", s)
	if err == nil {
		err = ResourceAviNsxtSegmentRuntimeRead(d, meta)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "nsxtsegmentruntime"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNsxtSegmentRuntimeDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
