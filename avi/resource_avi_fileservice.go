package avi

import (
	"log"
	"os"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFileServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"local_file": {
			Type:     schema.TypeString,
			Required: true,
		},
		//upload flag to state current local file will be uploaded to remote server.
		"upload": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
}

func resourceAviFileService() *schema.Resource {
	return &schema.Resource{
		Read:   ResourceAviFileServiceRead,
		Create: ResourceAviFileServiceCreate,
		Update: ResourceAviFileServiceUpdate,
		Delete: ResourceAviFileServiceDelete,
		Schema: ResourceFileServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFileServiceImporter,
		},
	}
}

func ResourceFileServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFileServiceSchema()
	return ResourceImporter(d, m, "fileservice", s)
}

func ResourceAviFileServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	var res interface{}
	switch upload := d.Get("upload").(bool); upload {
	case true:
		switch uri := d.Get("uri").(string); uri {
		case "license":
			path := "/api/license"
			err := client.AviSession.Get(path, &res)
			log.Printf("[DEBUG] ResourceAviFileServiceRead response: %v\n\n", res)
			if err != nil {
				log.Printf("[ERROR] ResourceAviFileServiceRead %v in GET of path %v\n", err, path)
				return err
			}
		default:
			uri := strings.Split(d.Get("uri").(string), "?")[0]
			path := "/api/fileservice?uri=controller://" + uri
			log.Printf("[DEBUG] ResourceAviFileServiceRead reading fileservice API status path %v\n", path)
			err := client.AviSession.Get(path, &res)
			log.Printf("[DEBUG] ResourceAviFileServiceRead response: %v\n\n", res)
			if err != nil {
				log.Printf("[ERROR] ResourceAviFileServiceRead %v in GET of path %v\n", err, path)
				return err
			}
		}
	case false:
		localFile := d.Get("local_file").(string)
		log.Printf("[DEBUG] ResourceAviFileServiceRead reading local file %v\n", localFile)
		if _, err := os.Stat(localFile); os.IsNotExist(err) {
			log.Printf("File does not exist")
			return err
		}
		log.Printf("File exists")
		return nil
	default:
		return nil
	}
	return nil
}

func ResourceAviFileServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFileServiceSchema()
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceCreate Error during upload/download %v\n", err)
		return err
	}
	return nil
}

func ResourceAviFileServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ResourceAviFileServiceUpdate")
	s := ResourceFileServiceSchema()
	err := MultipartUploadOrDownload(d, meta, s)
	if err != nil {
		log.Printf("[ERROR] ResourceAviFileServiceUpdate Error during upload/download %v\n", err)
		return err
	}
	err = ResourceAviFileServiceRead(d, meta)
	return err
}

func ResourceAviFileServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	localFile := d.Get("local_file").(string)
	switch upload := d.Get("upload").(bool); upload {
	case true:
		switch uri := d.Get("uri").(string); uri {
		case "license":
			path := "/api/" + uri + "/" + d.Id()
			err := client.AviSession.Delete(path)
			if err != nil {
				log.Printf("[ERROR] ResourceAviFileServiceDelete %v Deleting file of path %v\n", err, path)
			}
		default:
			uri := strings.Split(d.Get("uri").(string), "?")[0]
			path := "/api/fileservice?uri=controller://" + uri + "/" + d.Id()
			log.Printf("[DEBUG] ResourceAviFileServiceDelete deleting file using fileservice API status path %v\n", path)
			err := client.AviSession.Delete(path)
			if err != nil {
				log.Printf("[ERROR] ResourceAviFileServiceDelete %v Deleting file of path %v\n", err, path)
				return err
			}
		}
	case false:
		// delete file
		var err = os.Remove(localFile)
		if err != nil {
			log.Printf("[ERROR] ResourceAviFileServiceDelete Error for deleting file %v\n", localFile)
			return err
		}
		log.Printf("[INFO] ResourceAviFileServiceDelete file %v deleted\n", localFile)
	default:
		return nil
	}
	d.SetId("")
	return nil
}
