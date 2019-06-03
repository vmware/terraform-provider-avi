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

func ResourceServiceEngineSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"availability_zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"container_mode": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"container_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CONTAINER_TYPE_HOST",
		},
		"controller_created": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"controller_ip": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"data_vnics": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcevNICSchema(),
		},
		"enable_state": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SE_STATE_ENABLED",
		},
		"flavor": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hypervisor": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mgmt_vnic": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcevNICSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VM name unknown",
		},
		"resources": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSeResourcesSchema(),
		},
		"se_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviServiceEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEngineCreate,
		Read:   ResourceAviServiceEngineRead,
		Update: resourceAviServiceEngineUpdate,
		Delete: resourceAviServiceEngineDelete,
		Schema: ResourceServiceEngineSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceServiceEngineImporter,
		},
	}
}

func ResourceServiceEngineImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServiceEngineSchema()
	return ResourceImporter(d, m, "serviceengine", s)
}

func ResourceAviServiceEngineRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	err := ApiRead(d, meta, "serviceengine", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServiceEngineCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	err := ApiCreateOrUpdate(d, meta, "serviceengine", s)
	if err == nil {
		err = ResourceAviServiceEngineRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "serviceengine", s)
	if err == nil {
		err = ResourceAviServiceEngineRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceengine"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviServiceEngineDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
