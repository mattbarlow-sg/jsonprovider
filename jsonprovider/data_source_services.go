package jsonprovider

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServicesRead,
		Schema: map[string]*schema.Schema{
			"servlets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"servletname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	content, err := ioutil.ReadFile("/Users/matt.barlow/example.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	d.Set("servlets", payload["web-app"].(map[string]interface{})["servlet"])
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	var diags diag.Diagnostics

	return diags
}
