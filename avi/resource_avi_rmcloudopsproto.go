/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceRmCloudOpsProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"last_queried_se_creation_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pending_se_creation_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"pending_vnic_op_count": {
			Type:     schema.TypeInt,
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

func resourceAviRmCloudOpsProto() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRmCloudOpsProtoCreate,
		Read:   ResourceAviRmCloudOpsProtoRead,
		Update: resourceAviRmCloudOpsProtoUpdate,
		Delete: resourceAviRmCloudOpsProtoDelete,
		Schema: ResourceRmCloudOpsProtoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRmCloudOpsProtoImporter,
		},
	}
}

func ResourceRmCloudOpsProtoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRmCloudOpsProtoSchema()
	return ResourceImporter(d, m, "rmcloudopsproto", s)
}

func ResourceAviRmCloudOpsProtoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	err := APIRead(d, meta, "rmcloudopsproto", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRmCloudOpsProtoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	err := APICreateOrUpdate(d, meta, "rmcloudopsproto", s)
	if err == nil {
		err = ResourceAviRmCloudOpsProtoRead(d, meta)
	}
	return err
}

func resourceAviRmCloudOpsProtoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRmCloudOpsProtoSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "rmcloudopsproto", s)
	if err == nil {
		err = ResourceAviRmCloudOpsProtoRead(d, meta)
	}
	return err
}

func resourceAviRmCloudOpsProtoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "rmcloudopsproto"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviRmCloudOpsProtoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
