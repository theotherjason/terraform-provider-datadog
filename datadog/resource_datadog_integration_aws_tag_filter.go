package datadog

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/validators"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDatadogIntegrationAwsTagFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a Datadog AWS tag filter resource. This can be used to create and manage Datadog AWS tag filters - US site’s endpoint only",
		Create:      resourceDatadogIntegrationAwsTagFilterCreate,
		Update:      resourceDatadogIntegrationAwsTagFilterUpdate,
		Read:        resourceDatadogIntegrationAwsTagFilterRead,
		Delete:      resourceDatadogIntegrationAwsTagFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Your AWS Account ID without dashes.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description:  "The namespace associated with the tag filter entry. Allowed enum values: 'elb', 'application_elb', 'sqs', 'rds', 'custom', 'network_elb,lambda'",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validators.ValidateEnumValue(datadogV1.NewAWSNamespaceFromValue),
			},
			"tag_filter_str": {
				Description: "The tag filter string.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func buildDatadogIntegrationAwsTagFilter(d *schema.ResourceData) *datadogV1.AWSTagFilterCreateRequest {
	filterRequest := datadogV1.NewAWSTagFilterCreateRequestWithDefaults()
	if v, ok := d.GetOk("account_id"); ok {
		filterRequest.SetAccountId(v.(string))
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace := datadogV1.AWSNamespace(v.(string))
		filterRequest.SetNamespace(namespace)
	}
	if v, ok := d.GetOk("tag_filter_str"); ok {
		filterRequest.SetTagFilterStr(v.(string))
	}

	return filterRequest
}

func resourceDatadogIntegrationAwsTagFilterCreate(d *schema.ResourceData, meta interface{}) error {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1

	req := buildDatadogIntegrationAwsTagFilter(d)
	if _, _, err := datadogClientV1.AWSIntegrationApi.CreateAWSTagFilter(authV1, *req); err != nil {
		return utils.TranslateClientError(err, "error creating aws tag filter")
	}

	d.SetId(fmt.Sprintf("%s:%s", req.GetAccountId(), req.GetNamespace()))
	return resourceDatadogIntegrationAwsTagFilterRead(d, meta)
}

func resourceDatadogIntegrationAwsTagFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1

	req := buildDatadogIntegrationAwsTagFilter(d)
	if _, _, err := datadogClientV1.AWSIntegrationApi.CreateAWSTagFilter(authV1, *req); err != nil {
		return utils.TranslateClientError(err, "error updating aws tag filter")
	}

	return resourceDatadogIntegrationAwsTagFilterRead(d, meta)
}

func resourceDatadogIntegrationAwsTagFilterRead(d *schema.ResourceData, meta interface{}) error {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1

	accountID, tfNamespace, err := utils.AccountAndNamespaceFromID(d.Id())
	if err != nil {
		return err
	}
	namespace := datadogV1.AWSNamespace(tfNamespace)

	resp, _, err := datadogClientV1.AWSIntegrationApi.ListAWSTagFilters(authV1, accountID)
	if err != nil {
		return utils.TranslateClientError(err, "error listing aws tag filter")
	}

	for _, ns := range resp.GetFilters() {
		if ns.GetNamespace() == namespace {
			d.Set("account_id", accountID)
			d.Set("namespace", ns.GetNamespace())
			d.Set("tag_filter_str", ns.GetTagFilterStr())
			return nil
		}
	}

	// Set ID to an empty string if namespace is not found.
	// This allows Terraform to destroy the resource in state.
	d.SetId("")
	return nil
}

func resourceDatadogIntegrationAwsTagFilterDelete(d *schema.ResourceData, meta interface{}) error {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1

	accountID, tfNamespace, err := utils.AccountAndNamespaceFromID(d.Id())
	if err != nil {
		return err
	}
	namespace := datadogV1.AWSNamespace(tfNamespace)
	deleteRequest := datadogV1.AWSTagFilterDeleteRequest{
		AccountId: &accountID,
		Namespace: &namespace,
	}

	if _, _, err := datadogClientV1.AWSIntegrationApi.DeleteAWSTagFilter(authV1, deleteRequest); err != nil {
		return utils.TranslateClientError(err, "error deleting aws tag filter")
	}

	return nil
}
