package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/upup/pkg/fi/cloudup/gce"
)

func ServiceAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: ServiceAccountRead,
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"role": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"member": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		// SchemaVersion:  res.SchemaVersion,
		// StateUpgraders: res.StateUpgraders,
	}
}

func ServiceAccountRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clusterName := d.Get("cluster_name").(string)
	role := d.Get("role").(string)
	projectID := d.Get("project_id").(string)

	accountID := gce.ServiceAccountName(role, clusterName)

	email := accountID + "@" + projectID + ".iam.gserviceaccount.com"

	d.Set("account_id", accountID)
	d.Set("email", email)
	d.Set("member", "serviceAccount:"+email)
	d.SetId(email)
	return nil
}
