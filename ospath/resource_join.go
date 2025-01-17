package ospath

import (
	"path"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceJoin() *schema.Resource {
	return &schema.Resource{
		Create: createJoin,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func toStringSlice(interfaceSlice []interface{}) []string {
	stringSlice := []string{}
	for i := range interfaceSlice {
		stringSlice = append(stringSlice, interfaceSlice[i].(string))
	}

	return stringSlice
}

func createJoin(d *schema.ResourceData, m interface{}) error {
	pathArr := toStringSlice(d.Get("path").([]interface{}))
	joinedPath := path.Join(pathArr...)
	d.Set("result", joinedPath)
	d.SetId(joinedPath)
	return nil
}
