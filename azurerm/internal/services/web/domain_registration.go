package web

import (
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2019-08-01/web"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func domainRegistrationContactSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"mailing_address": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"address_1": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							"address_2": {
								Type:         schema.TypeString,
								Optional:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							"city": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							"country": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							"postal_code": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							"state": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},
					},
				},
				"email": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty, // TODO - a not horrible approximation at email validation?
				},

				"first_name": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringLenBetween(1, 255),
				},

				"last_name": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringLenBetween(1, 255),
				},

				"phone_number": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"middle_name": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringLenBetween(1, 255),
				},

				"organisation": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringLenBetween(1, 255),
				},

				"job_title": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"fax": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func flattenDomainRegistrationContact(input *web.Contact) map[string]interface{} {
	contact := make(map[string]interface{})
	if input == nil {
		return contact
	}

	if input.Email != nil {
		contact["email"] = input.Email
	}

	if input.NameFirst != nil {
		contact["first_name"] = input.NameFirst
	}

	if input.NameLast != nil {
		contact["last_name"] = input.NameLast
	}

	if input.NameFirst != nil {
		contact["middle_name"] = input.NameMiddle
	}

	if input.Phone != nil {
		contact["phone_number"] = input.Phone
	}

	if input.AddressMailing != nil {
		contact["mailing_address"] = flattenDomainRegistrationContactMailingAddress(input.AddressMailing)
	}

	return contact
}

func flattenDomainRegistrationContactMailingAddress(input *web.Address) map[string]interface{} {
	address := make(map[string]interface{})
	if input == nil {
		return address
	}

	if input.Address1 != nil {
		address["address_1"] = input.Address1
	}

	if input.Address2 != nil {
		address["address_2"] = input.Address1
	}

	if input.City != nil {
		address["city"] = input.City
	}

	if input.State != nil {
		address["state"] = input.City
	}

	if input.Country != nil {
		address["country"] = input.City
	}

	if input.PostalCode != nil {
		address["postal_code"] = input.City
	}

	return address
}

func expandDomainRegistrationContact(input []interface{}) *web.Contact {
	contactRaw := input[0].(map[string]interface{})

	contact := web.Contact{
		AddressMailing: expandDomainRegistrationContactMailingAddress(contactRaw["mailing_address"].([]interface{})),
		Email:          utils.String(contactRaw["email"].(string)),
		NameFirst:      utils.String(contactRaw["first_name"].(string)),
		NameLast:       utils.String(contactRaw["last_name"].(string)),
		Phone:          utils.String(contactRaw["phone_number"].(string)),
	}

	if contactRaw["fax"] != nil {
		contact.Fax = utils.String(contactRaw["fax"].(string))
	}

	if contactRaw["job_title"] != nil {
		contact.JobTitle = utils.String(contactRaw["job_title"].(string))
	}

	if contactRaw["organisation"] != nil {
		contact.Organization = utils.String(contactRaw["organisation"].(string))
	}

	if contactRaw["middle_name"] != nil {
		contact.NameMiddle = utils.String(contactRaw["middle_name"].(string))
	}

	return &contact
}

func expandDomainRegistrationContactMailingAddress(input []interface{}) *web.Address {
	address := input[0].(map[string]interface{})

	contactAddress := &web.Address{
		Address1:   utils.String(address["address_1"].(string)),
		City:       utils.String(address["city"].(string)),
		Country:    utils.String(address["country"].(string)),
		PostalCode: utils.String(address["postal_code"].(string)),
		State:      utils.String(address["state"].(string)),
	}

	if address["address_2"] != nil {
		contactAddress.Address2 = utils.String(address["address_2"].(string))
	}

	return contactAddress
}