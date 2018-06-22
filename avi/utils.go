/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"log"
	"reflect"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func SchemaToAviData(d interface{}, s map[string]*schema.Schema) (interface{}, error) {
	switch d.(type) {
	default:
		log.Printf("[INFO] SchemaToAviData: resource d: %v(%v)", d, reflect.TypeOf(d))
		//return d, nil
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToAviData(v, nil); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in parsing k: %v v: %v", err, k, v)
			}
		}
		return m, nil
	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToAviData(varray[i], nil)
			if err == nil && obj != nil {
				objList = append(objList, obj)
			}
		}
		if len(objList) == 0 {
			return nil, nil
		}
		return objList, nil

	case *schema.Set:
		if len(d.(*schema.Set).List()) == 0 {
			return nil, nil
		}
		obj, err := SchemaToAviData(d.(*schema.Set).List()[0], nil)
		return obj, err

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s {
			if obj, err := SchemaToAviData(r.Get(k), nil); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in converting k: %v v: %v", err, k, v)
			}
		}
		return m, nil
	}
	/** Return the same object as there is nothing special about **/
	return d, nil
}

func CommonHash(v interface{}) int {
	return hashcode.String("avi")
}

func SetDefaultsInAPIRes(api_res interface{}, d_local interface{}, s map[string]*schema.Schema) (interface{}, error) {
	switch d_local.(type) {
	default:
	case map[string]interface{}:
		for k, v := range d_local.(map[string]interface{}) {
			switch v.(type) {
			//Getting key, value for given d_local
			default:
				if _, ok := api_res.(map[string]interface{})[k]; !ok {
					//Cheking if field is present in schema
					if dval, ok := s[k]; ok {
						//Getting default values from schema
						default_val, err := dval.DefaultValue()
						if err != nil {
							log.Printf("[ERROR] SetDefaultsInAPIRes %v", err)
						} else {
							if default_val != nil {
								api_res.(map[string]interface{})[k] = default_val
								log.Printf("[INFO] SetDefaultsInAPIRes setting default for field: %v\t val: %v", k, default_val)
							}
						}

					}
				}
			//d_local nested dictionary.
			case map[string]interface{}:
				s2, err := s[k]
				//As err returned is boolean value
				if err {
					log.Printf("[INFO] SetDefaultsInAPIRes %v", err)
				}
				switch s2.Elem.(type) {
				default:
				case *schema.Resource:
					api_res1, err := SetDefaultsInAPIRes(api_res.(map[string]interface{})[k], v, s2.Elem.(*schema.Resource).Schema)
					if err != nil {
						log.Printf("[INFO] SetDefaultsInAPIRes %v", err)

					} else {
						api_res.(map[string]interface{})[k] = api_res1
					}
				}
			//d_local is array of dictionaries.
			case []interface{}:
				var objList []interface{}
				varray2 := api_res.(map[string]interface{})[k].([]interface{})
				//getting schema for nested object.
				s2, err := s[k]
				//As err returned is boolean value
				if err {
					log.Printf("[ERROR] SetDefaultsInAPIRes %v", err)
				}
				for x, y := range v.([]interface{}) {
					switch s2.Elem.(type) {
					default:
					case *schema.Resource:
						obj, err := SetDefaultsInAPIRes(varray2[x], y, s2.Elem.(*schema.Resource).Schema)
						if err != nil {
							log.Printf("[INFO] SetDefaultsInAPIRes %v", err)
						} else {
							objList = append(objList, obj)
						}
					case *schema.Schema:
						objList = append(objList, y)
					}
				}
				api_res.(map[string]interface{})[k] = objList
			}
		}
	}
	return api_res, nil

}

func ApiDataToSchema(adata interface{}, d *schema.ResourceData, t map[string]*schema.Schema) (interface{}, error) {
	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			var m map[string]interface{}
			m = map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] ApiDataToSchema %v in converting k: %v v: %v", err, k, v)
				}
			}
			//var s schema.Set
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			//s.Add(m)
			return s, nil
		} else {
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
						err := d.Set(k, obj)
						if err != nil {
							log.Printf("[ERROR] ApiDataToSchema %v in setting %v", err, obj)
						}
					}
				}
			}
			return d, nil
		}
	case []interface{}:
		var objList []interface{}
		varray := adata.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := ApiDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch obj.(type) {
				default:
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			} else {
				log.Printf("[ERROR] ApiDataToSchema %v", err)
			}
		}
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func ApiCreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	var err error
	client := meta.(*clients.AviClient)
	var robj interface{}
	obj := d

	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/" + objType
		if objType == "cluster" {
			path = path + "?include_name=true&skip_default=true"
			err = client.AviSession.Put(path, data, &robj)
		} else if uuid, ok := d.GetOk("uuid"); ok {
			path = path + "/" + uuid.(string) + "?include_name=true&skip_default=true"
			err = client.AviSession.Put(path, data, &robj)
		} else {
			if name, ok := d.GetOk("name"); ok {
				var existing_obj interface{}
				if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
					cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
					// strip the # if it exists
					cloudUUID = strings.Split(cloudUUID, "#")[0]
					log.Printf("[INFO] ApiCreateOrUpdate: using cloud %v for obj %v name %s \n",
						cloudUUID, objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existing_obj), session.SetCloudUUID(cloudUUID),
						session.SetIncludeName(true))
				} else {
					log.Printf("[INFO] ApiCreateOrUpdate: reading obj %v name %s \n",
						objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existing_obj), session.SetIncludeName(true))
				}
				if err != nil {
					// object not found
					log.Printf("[INFO] ApiCreateOrUpdate: Creating obj type %v schema %v data %v\n", objType, d, data)
					err = client.AviSession.Post(path, data, &robj)
					if err != nil {
						log.Printf("[ERROR] ApiCreateOrUpdate creation failed %v object with name %v\n", err, name)
					}
				} else {
					// found existing object.
					uuid = existing_obj.(map[string]interface{})["uuid"].(string)
					d.Set("uuid", uuid)
					d.SetId(uuid.(string))
					path = path + "/" + uuid.(string) + "?include_name=true&skip_default=true"
					err = client.AviSession.Put(path, data, &robj)
				}
			} else {
				log.Printf("[INFO] ApiCreateOrUpdate: Creating obj %v schema %v data %v\n", objType, d, data)
				err = client.AviSession.Post(path, data, &robj)
				if err != nil {
					log.Printf("[ERROR] ApiCreateOrUpdate creation failed %v\n", err)
				}
			}
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiCreateOrUpdate: Error %v path %v id %v\n", err, path, d.Id())
			return err
		}
		log.Printf("[DEBUG] ApiCreateOrUpdate: object %v\n", robj)
		uuid := robj.(map[string]interface{})["uuid"].(string)
		d.Set("uuid", uuid)
		if url, ok := robj.(map[string]interface{})["url"].(string); ok && url != "" {
			d.SetId(url)
		} else {
			d.SetId(uuid)
		}
		//url = strings.SplitN(url, "#", 2)[0]
	} else {
		log.Printf("[ERROR] ApiCreateOrUpdate: Error %v", err)
	}
	return err
}

func ApiRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var obj interface{}
	uuid := ""
	log.Printf("[DEBUG] ApiRead reading object with objType %v id %v\n",
		objType, d.Id())

	if d.Id() != "" {
		// extract the uuid from it.
		log.Printf("[DEBUG] ApiRead reading object with objType %v id %v \n", objType, d.Id())
		url_parts := strings.Split(d.Id(), "/")
		uuid = url_parts[len(url_parts)-1]
		// need to strip #xxx if present
		uuid = strings.Split(uuid, "#")[0]
	} else if u, ok := d.GetOk("uuid"); ok {
		uuid = u.(string)
		log.Printf("[DEBUG] ApiRead reading object with uuid %v \n", uuid)
	}
	if uuid != "" {
		path := "api/" + objType + "/" + uuid + "?include_name=true&skip_default=true"
		log.Printf("[DEBUG] ApiRead reading object with id %v path %v\n", uuid, path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with uuid %v not found err %v\n", uuid, err)
			return nil
		}
	} else if name, ok := d.GetOk("name"); ok {
		var err error
		if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
			cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
			cloudUUID = strings.Split(cloudUUID, "#")[0]
			log.Printf("[DEBUG] ApiRead using cloud %v obj %v name %v\n", cloudUUID,
				objType, name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetCloudUUID(cloudUUID),
				session.SetIncludeName(true), session.SetSkipDefault(true))
		} else {
			log.Printf("[DEBUG] ApiRead using name %v \n", name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetIncludeName(true),
				session.SetSkipDefault(true))
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with name %v:%v not found err %v\n", objType, name, err)
			return nil
		}
	} else {
		d.SetId("")
		log.Printf("[ERROR] ApiRead not found %v\n", d.Get("uuid"))
		return nil
	}
	if local_data, err := SchemaToAviData(d, s); err == nil {
		mod_api_res, err := SetDefaultsInAPIRes(obj, local_data, s)
		if err != nil {
			log.Printf("[ERROR] ApiRead in modifying api response object %v\n", err)
		}
		if _, err := ApiDataToSchema(mod_api_res, d, s); err == nil {
			uuid := mod_api_res.(map[string]interface{})["uuid"].(string)
			url := mod_api_res.(map[string]interface{})["url"].(string)
			//url = strings.SplitN(url, "#", 2)[0]
			if url != "" {
				d.SetId(url)
				log.Printf("[DEBUG] ApiRead read object with id %v\n", url)
			} else {
				d.SetId(uuid)
				log.Printf("[DEBUG] ApiRead read object with id %v\n", uuid)
			}
		} else {
			log.Printf("[ERROR] ApiRead in setting read object %v\n", err)
		}
	}

	return nil
}

func ResourceImporter(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) ([]*schema.ResourceData, error) {
	if d.Id() != "" {
		// return the ID based import
		return []*schema.ResourceData{d}, nil
	}
	var data interface{}
	client := meta.(*clients.AviClient)
	path := "api/" + objType + "/?skip_default=true&include_name=true"
	log.Printf("[DEBUG] ResourceImporter reading object with path %v\n", path)

	err := client.AviSession.Get(path, &data)
	if err != nil {
		log.Printf("[ERROR] ResourceImporter %v in GET of path %v\n", err, path)
		return nil, err
	}
	count := int((data.(map[string]interface{})["count"]).(float64))
	log.Printf("[DEBUG] ResourceImporter read data with path %v -> count %v\n", path, count)
	if count > 0 {
		results := make([]*schema.ResourceData, count)
		apiResults := (data.(map[string]interface{})["results"]).([]interface{})
		for index := 0; index < count; index++ {
			obj := apiResults[index].(map[string]interface{})
			log.Printf("[DEBUG] ResourceImporter processing obj %v results %v\n", obj, results[index])
			result := new(schema.ResourceData)
			if _, err := ApiDataToSchema(obj, result, s); err == nil {
				log.Printf("[DEBUG] ResourceImporter Processing obj %v\n", obj)
				url := obj["url"].(string)
				uuid := obj["uuid"].(string)
				//url = strings.SplitN(url, "#", 2)[0]
				result.SetId(url)
				result.Set("uuid", uuid)
				result.SetType("avi_" + objType)
				results[index] = result
			}
		}
		return results, nil
	}
	return nil, nil
}

func ApiDeleteSystemDefaultCheck(d *schema.ResourceData) bool {
	var systemDefault bool
	var sysName string
	if sysdef, ok := d.GetOk("system_default"); ok {
		systemDefault = sysdef.(bool)
	}
	if name, ok := d.GetOk("name"); ok {
		sysName = name.(string)
	}
	if systemDefault || strings.HasPrefix(sysName, "System-") || sysName == "Default-Group" {
		d.SetId("")
		return true
	}
	return false
}
