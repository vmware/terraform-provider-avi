/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"routing_service": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRoutingServiceSchema(),
		},
		"se_group_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"service_type": {
			Type:     schema.TypeString,
			Required: true,
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
		"vrf_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviNetworkService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkServiceCreate,
		Read:   ResourceAviNetworkServiceRead,
		Update: resourceAviNetworkServiceUpdate,
		Delete: resourceAviNetworkServiceDelete,
		Schema: ResourceNetworkServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkServiceImporter,
		},
	}
}

func ResourceNetworkServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkServiceSchema()
	return ResourceImporter(d, m, "networkservice", s)
}

func ResourceAviNetworkServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	err := APIRead(d, meta, "networkservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	err := APICreateOrUpdate(d, meta, "networkservice", s)
	if err == nil {
		err = ResourceAviNetworkServiceRead(d, meta)
	}
	return err
}

func resourceAviNetworkServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkServiceSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "networkservice", s)
	if err == nil {
		err = ResourceAviNetworkServiceRead(d, meta)
	}
	return err
}

func resourceAviNetworkServiceDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networkservice"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkServiceDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
