package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlow() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to access information about an existing Kestra Flow",

		ReadContext: dataSourceFlowRead,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Description: "The namespace.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"flow_id": {
				Description: "The flow id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"revision": {
				Description: "The flow revision.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"content": {
				Description: "The flow content as yaml.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"keep_original_source": {
				Description: "Use the content as source code, keeping comment and indentation.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func dataSourceFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	namespaceId := d.Get("namespace")
	flowId := d.Get("flow_id")

	r, reqErr := c.request("GET", fmt.Sprintf("/api/v1/flows/%s/%s", namespaceId, flowId), nil)
	if reqErr != nil {
		return diag.FromErr(reqErr.Err)
	}

	errs := flowApiToSchema(r.(map[string]interface{}), d)
	if errs != nil {
		return errs
	}

	return diags
}
