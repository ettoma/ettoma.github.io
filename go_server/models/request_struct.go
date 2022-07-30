package models

type Create_account_holder_request struct {
	Account_holder_code string `json:"accountHolderCode"`
	Country             string `json:"countryAccount"`
	Doing_business_as   string `json:"doingBusinessAs"`
	Legal_business_name string `json:"legalBusinessName"`
	Registration_number string `json:"registrationNumber"`
	Shareholder_type    string `json:"shareholderType"`
	Job_title           string `json:"jobTitle"`
	First_name          string `json:"firstName"`
	Gender              string `json:"gender"`
	Last_name           string `json:"lastName"`
	Shareholder_email   string `json:"shareholderEmail"`
	Email               string `json:"email"`
	Address_country     string `json:"addressCountry"`
	Web_address         string `json:"webAddress"`
}

type Field_Type struct {
	Field      string `json:"field"`
	Field_name string `json:"fieldName"`
}

type InvalidFields struct {
	Error_code        int        `json:"errorCode"`
	Error_description string     `json:"errorDescription"`
	FieldType         Field_Type `json:"fieldType"`
}

type Adyen_account_holder_response_invalid struct {
	Invalid_fields []InvalidFields `json:"invalidFields"`
	Psp_reference  string          `json:"pspReference"`
}

type Adyen_account_holder_response_valid struct {
	Invalid_fields []map[string]interface{} `json:"invalidFields"`
	Psp_reference  string                   `json:"pspReference"`
}

type Create_store_request struct {
	Store_code string `json:"storeCode"`
}
