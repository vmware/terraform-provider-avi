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

func ResourceFileObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"checksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"compressed": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"created": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"expires_at": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"read_only": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"restrict_download": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviFileObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviFileObjectCreate,
		Read:   ResourceAviFileObjectRead,
		Update: resourceAviFileObjectUpdate,
		Delete: resourceAviFileObjectDelete,
		Schema: ResourceFileObjectSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFileObjectImporter,
		},
	}
}

func ResourceFileObjectImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFileObjectSchema()
	return ResourceImporter(d, m, "fileobject", s)
}

func ResourceAviFileObjectRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	err := ApiRead(d, meta, "fileobject", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviFileObjectCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	err := ApiCreateOrUpdate(d, meta, "fileobject", s)
	if err == nil {
		err = ResourceAviFileObjectRead(d, meta)
	}
	return err
}

func resourceAviFileObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileObjectSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "fileobject", s)
	if err == nil {
		err = ResourceAviFileObjectRead(d, meta)
	}
	return err
}

func resourceAviFileObjectDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "fileobject"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviFileObjectDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
