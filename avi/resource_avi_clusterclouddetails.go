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

func ResourceClusterCloudDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_info": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAzureClusterInfoSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviClusterCloudDetails() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviClusterCloudDetailsCreate,
		Read:   ResourceAviClusterCloudDetailsRead,
		Update: resourceAviClusterCloudDetailsUpdate,
		Delete: resourceAviClusterCloudDetailsDelete,
		Schema: ResourceClusterCloudDetailsSchema(),
	}
}

func ResourceAviClusterCloudDetailsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/clusterclouddetails/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviClusterCloudDetailsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	err := ApiCreateOrUpdate(d, meta, "clusterclouddetails", s)
	if err == nil {
		err = ResourceAviClusterCloudDetailsRead(d, meta)
	}
	return err
}

func resourceAviClusterCloudDetailsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterCloudDetailsSchema()
	err := ApiCreateOrUpdate(d, meta, "clusterclouddetails", s)
	if err == nil {
		err = ResourceAviClusterCloudDetailsRead(d, meta)
	}
	return err
}

func resourceAviClusterCloudDetailsDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "clusterclouddetails"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviClusterCloudDetailsDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
