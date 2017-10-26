/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	//"github.com/hashicorp/terraform/helper/schema"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
	"strings"
)

func SchemaToAviData(d interface{}, s map[string]*schema.Schema) (interface{}, error) {
	switch d.(type) {
	default:
		log.Printf("SchemaToAviData: resource d: %v(%v)", d, reflect.TypeOf(d))
		//return d, nil
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToAviData(v, nil); err == nil {
				m[k] = obj
			} else {
				log.Printf("Error in parsing k: %v v: %v", k, v)
			}
		}
		log.Printf("Set elem %v obj %v", d, m)
		return m, nil
	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToAviData(varray[i], nil)
			if err == nil {
				objList = append(objList, obj)
			}
		}
		log.Printf("List elem %v objList %v", d, objList)
		return objList, nil

	case *schema.Set:
		obj, err := SchemaToAviData(d.(*schema.Set).List()[0], nil)
		log.Printf("Set elem %v obj %v", d.(*schema.Set).List()[0], obj)
		return obj, err

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s {
			if obj, err := SchemaToAviData(r.Get(k), nil); err == nil {
				m[k] = obj
			} else {
				log.Printf("Error in converting k: %v v: %v", k, v)
			}
		}
		log.Printf("Converted Top level resource data %v -> %v", d, m)
		return m, nil
	}
	/** Return the same object as there is nothing special about **/
	log.Printf("SchemaToAviData: Returning resource d: %v(%v)", d, reflect.TypeOf(d))
	return d, nil
}

func CommonHash(v interface{}) int {
	return hashcode.String("avi")
}

func ApiDataToSchema(adata interface{}, d *schema.ResourceData, t map[string]*schema.Schema) (interface{}, error) {
	//log.Printf("Start Processing resource %v(%v)", adata, reflect.TypeOf(adata))
	var allowedKeys = map[string]bool{
		"name":                true,
		"uuid":                true,
		"fail_action":         true,
		"health_monitor_refs": true,
		"servers":             true,
		"type":                true,
		"ip":                  true,
		"addr":                true,
		"port":                true,
		"url":                 true,
	}

	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			var m map[string]interface{}
			m = map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := allowedKeys[k]; ok {
					if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
						log.Printf("ApiDataToSchema: Processed resource k %v value: %v(%v)", k, obj)
						m[k] = obj
						//log.Printf("Updating map %v", m)
					} else {
						//log.Printf("Error in converting k: %v v: %v", k, v)
					}
				}
			}
			//var s schema.Set
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			//log.Printf("Updating set object obj %v with m: %v", s, m)
			//s.Add(m)
			log.Printf("ApiDataToSchema: Returning Set obj %v\n", s)
			return s, nil
		} else {
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
						log.Printf("ApiDataToSchema: Returning Set obj %v\n", d)
						err := d.Set(k, obj)
						if err != nil {
							log.Print("ApiDataToSchema err %v in setting %v", err, obj)
						}
					}
				}
			}
			log.Printf("ApiDataToSchema: Returning Set obj %v\n", d)
			return d, nil
		}
	case []interface{}:
		//log.Printf("Processing List d: %v(%v)", adata, reflect.TypeOf(adata))
		var objList []interface{}
		varray := adata.([]interface{})
		//log.Printf("data array List d: %v", varray)
		for i := 0; i < len(varray); i++ {
			obj, err := ApiDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch obj.(type) {
				default:
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			}
		}
		log.Printf("ApiDataToSchema: List elem objList %v", objList)
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func ApiCreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	var err error
	client := meta.(*clients.AviClient)
	var robj interface{}
	//	obj := tObj.(*schema.Set).List()[0]
	obj := d
	log.Printf("[INFO] TF data %v %v\n", obj, reflect.TypeOf(obj))
	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/" + objType
		log.Printf("[INFO] POST data %v id %v\n", data, d.Id())
		if uuid, ok := d.GetOk("uuid"); ok {
			path = path + "/" + uuid.(string)
			err = client.AviSession.Put(path, data, &robj)
		} else {
			err = client.AviSession.Post(path, data, &robj)
		}
		log.Printf("ApiCreateOrUpdate: using Path %v Id %v", path, d.Id())
		if err != nil {
			log.Printf("[ERROR] ApiCreateOrUpdate: Error %v path %v id %v\n", err, path, d.Id())
			return err
		}
		url := robj.(map[string]interface{})["url"].(string)
		uuid := robj.(map[string]interface{})["uuid"].(string)
		log.Printf("[ERROR] ApiCreateOrUpdate: url %v uuid %v", url, uuid)
		url = strings.SplitN(url, "#", 2)[0]
		d.SetId(url)
		d.Set("uuid", uuid)
		//d.SetId(robj.(map[string]interface{})["uuid"].(string))
		log.Printf("[INFO] resourceAviPoolCreate: Create success %v uuid %v\n", err, url)
	} else {
		log.Printf("Error %v", err)
	}
	return err
}
