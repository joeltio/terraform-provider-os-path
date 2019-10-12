package ospath

import (
	"path"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceJoin() *schema.Resource {
	return &schema.Resource{
		Create: createJoin,
		Read:   schema.Noop,
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func createJoin(d *schema.ResourceData, m interface{}) error {
	pathArr := d.Get("path").([]string)
	joinedPath := path.Join(pathArr...)
	d.Set("result", joinedPath)
	d.SetId(joinedPath)
	return nil
}
