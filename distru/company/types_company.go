package company

import "time"

type CompanyResponse struct {
	Company  Company         `json:"company"`
	MetaData CompanyMetaData `json:"meta_data"`
}

type CompanyMetaData struct {
	CategoryOptions []string `json:"category_options"`
}

type QbAuthToken struct {
	ID int `json:"id"`
}

type APIKeys struct {
	ID         int       `json:"id"`
	InsertedAt time.Time `json:"inserted_at"`
	MetrcKey   string    `json:"metrc_key"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StableMeasurement struct {
	MinDuration             int    `json:"min_duration"`
	SensitivityQuantity     string `json:"sensitivity_quantity"`
	SensitivityUnitTypeName string `json:"sensitivity_unit_type_name"`
}

type TimeMeasurement struct {
	MinDuration int `json:"min_duration"`
}

type SettingsAutoWeighTriggerConfig struct {
	StableMeasurement     StableMeasurement `json:"stable_measurement"`
	ThresholdQuantity     string            `json:"threshold_quantity"`
	ThresholdUnitTypeName string            `json:"threshold_unit_type_name"`
	TimeMeasurement       TimeMeasurement   `json:"time_measurement"`
	Timeout               int               `json:"timeout"`
}

type PaymentTerms struct {
	CompanyID int    `json:"company_id"`
	Days      int    `json:"days"`
	ID        int    `json:"id"`
	Locked    bool   `json:"locked"`
	Name      string `json:"name"`
	TimeOfDay string `json:"time_of_day"`
}

type License struct {
	Active         bool      `json:"active"`
	APIKeyID       int       `json:"api_key_id"`
	CompanyID      int       `json:"company_id"`
	ComplianceType string    `json:"compliance_type"`
	CreatorID      int       `json:"creator_id"`
	ExpiryDate     time.Time `json:"expiry_date"`
	FileAttachment any       `json:"file_attachment"`
	ID             int       `json:"id"`
	IssueDate      time.Time `json:"issue_date"`
	LicenseNumber  string    `json:"license_number"`
	LicenseType    string    `json:"license_type"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type GeometryLocation struct{}

type Ha struct {
	Hi float64 `json:"hi"`
	Lo float64 `json:"lo"`
}

type Va struct {
	Hi float64 `json:"hi"`
	Lo float64 `json:"lo"`
}

type Viewport struct {
	Ha Ha `json:"ha"`
	Va Va `json:"va"`
}

type Geometry struct {
	Location GeometryLocation `json:"location"`
	Viewport Viewport         `json:"viewport"`
}

type Gmaps struct {
	AddressComponents   []AddressComponents `json:"address_components"`
	AdrAddress          string              `json:"adr_address"`
	FormattedAddress    string              `json:"formatted_address"`
	Geometry            Geometry            `json:"geometry"`
	HTMLAttributions    []any               `json:"html_attributions"`
	Icon                string              `json:"icon"`
	IconBackgroundColor string              `json:"icon_background_color"`
	IconMaskBaseURI     string              `json:"icon_mask_base_uri"`
	Name                string              `json:"name"`
	PlaceID             string              `json:"place_id"`
	Reference           string              `json:"reference"`
	Types               []string            `json:"types"`
	URL                 string              `json:"url"`
	UtcOffset           int                 `json:"utc_offset"`
	UtcOffsetMinutes    int                 `json:"utc_offset_minutes"`
	Vicinity            string              `json:"vicinity"`
}

type LocationLatLon struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type MatchedSubstrings struct {
	Length int `json:"length"`
	Offset int `json:"offset"`
}

type LocationMetaData struct {
	Description       string            `json:"description"`
	Gmaps             Gmaps             `json:"gmaps"`
	IsFixture         bool              `json:"is_fixture"`
	Label             string            `json:"label"`
	Location          LocationLatLon    `json:"location"`
	MatchedSubstrings MatchedSubstrings `json:"matched_substrings"`
	PlaceID           string            `json:"place_id"`
}

type Location struct {
	Address        string           `json:"address"`
	Apt            any              `json:"apt"`
	City           string           `json:"city"`
	CompanyID      int              `json:"company_id"`
	Country        string           `json:"country"`
	ID             int              `json:"id"`
	Latitude       float64          `json:"latitude"`
	License        License          `json:"license"`
	LicenseID      int              `json:"license_id"`
	Longitude      float64          `json:"longitude"`
	Manual         bool             `json:"manual"`
	MetaData       LocationMetaData `json:"meta_data"`
	Name           string           `json:"name"`
	NumberOfLights any              `json:"number_of_lights"`
	SquareFeet     any              `json:"square_feet"`
	State          string           `json:"state"`
	StreetAddress  string           `json:"street_address"`
	Wattage        any              `json:"wattage"`
	Zip            string           `json:"zip"`
}

type Settings struct {
	EmailSettingsPromptInvoiceDialog        bool   `json:"email_settings_prompt_invoice_dialog"`
	EmailSettingsPromptPackingSlipDialog    bool   `json:"email_settings_prompt_packing_slip_dialog"`
	EmailSettingsPromptPurchaseOrderDialog  bool   `json:"email_settings_prompt_purchase_order_dialog"`
	InvoiceSettingsShowVendorInInvoicePdf   bool   `json:"invoice_settings_show_vendor_in_invoice_pdf"`
	OrderSettingsCreateNotificationEmails   string `json:"order_settings_create_notification_emails"`
	OrderSettingsDefaultWholesalePrice      bool   `json:"order_settings_default_wholesale_price"`
	OrderSettingsShowVendorInSalesOrderPdf  bool   `json:"order_settings_show_vendor_in_sales_order_pdf"`
	RequestSettingsCreateNotificationEmails string `json:"request_settings_create_notification_emails"`
	RequestSettingsDefaultRequestSlipEmails string `json:"request_settings_default_request_slip_emails"`
	ShowOnboardingCard                      bool   `json:"show_onboarding_card"`
}

type QuickbooksItems struct {
	Nodes []any `json:"nodes"`
}

type DefaultOrderPaymentTerm struct {
	CompanyID int    `json:"company_id"`
	Days      int    `json:"days"`
	ID        int    `json:"id"`
	Locked    bool   `json:"locked"`
	Name      string `json:"name"`
	TimeOfDay string `json:"time_of_day"`
}

type Company struct {
	Timezone                                               string                         `json:"timezone"`
	SettingsOrderNumberNext                                int                            `json:"settings_order_number_next"`
	SettingsInvoiceNumberPrefix                            string                         `json:"settings_invoice_number_prefix"`
	QbAuthToken                                            QbAuthToken                    `json:"qb_auth_token"`
	QbIncomeAccountID                                      any                            `json:"qb_income_account_id"`
	SettingsDefaultInvoiceColor                            string                         `json:"settings_default_invoice_color"`
	Email                                                  string                         `json:"email"`
	QbSettingsSyncPurchases                                bool                           `json:"qb_settings_sync_purchases"`
	AllowNegativeAvailableQuantity                         bool                           `json:"allow_negative_available_quantity"`
	DefaultOrderShipmentShippingManifestShipperID          any                            `json:"default_order_shipment_shipping_manifest_shipper_id"`
	SettingsOrderShipmentNumberPrefix                      string                         `json:"settings_order_shipment_number_prefix"`
	SettingsDefaultSalesOrderColor                         string                         `json:"settings_default_sales_order_color"`
	OrganizationName                                       string                         `json:"organization_name"`
	SettingsInvoiceDateBasedOffOrderDeliveryDate           bool                           `json:"settings_invoice_date_based_off_order_delivery_date"`
	ShowPaymentDelinquentWarning                           bool                           `json:"show_payment_delinquent_warning"`
	APIKeys                                                []APIKeys                      `json:"api_keys"`
	InternalUse                                            bool                           `json:"internal_use"`
	SettingsAutoWeighTriggerConfig                         SettingsAutoWeighTriggerConfig `json:"settings_auto_weigh_trigger_config"`
	SettingsInvoiceNumberNext                              int                            `json:"settings_invoice_number_next"`
	DefaultPurchaseShippingManifestShipper                 any                            `json:"default_purchase_shipping_manifest_shipper"`
	SettingsPurchaseNumberPrefix                           string                         `json:"settings_purchase_number_prefix"`
	MetrcBackend                                           string                         `json:"metrc_backend"`
	SettingsBreakdownNumberPrefix                          string                         `json:"settings_breakdown_number_prefix"`
	ID                                                     int                            `json:"id"`
	QbSettingsPoSyncStatuses                               []string                       `json:"qb_settings_po_sync_statuses"`
	EnabledModules                                         []string                       `json:"enabled_modules"`
	SettingsPlantWeighingMethod                            string                         `json:"settings_plant_weighing_method"`
	QbPurchaseDepositToAccountID                           any                            `json:"qb_purchase_deposit_to_account_id"`
	MigrationInProcess                                     bool                           `json:"migration_in_process"`
	DefaultPurchaseShippingManifestShipperID               any                            `json:"default_purchase_shipping_manifest_shipper_id"`
	Name                                                   string                         `json:"name"`
	SettingsReturnNumberPrefix                             string                         `json:"settings_return_number_prefix"`
	SettingsDefaultSalesOrderStatusOnCreation              any                            `json:"settings_default_sales_order_status_on_creation"`
	Phone                                                  string                         `json:"phone"`
	PaymentTerms                                           []PaymentTerms                 `json:"payment_terms"`
	SettingsDefaultPurchaseOrderColor                      string                         `json:"settings_default_purchase_order_color"`
	SettingsMetrcOutgoingTransferTransporterLicenseNumbers []string                       `json:"settings_metrc_outgoing_transfer_transporter_license_numbers"`
	SettingsPaymentNumberPrefix                            string                         `json:"settings_payment_number_prefix"`
	Description                                            any                            `json:"description"`
	DefaultStockTransferShippingManifestDistributor        any                            `json:"default_stock_transfer_shipping_manifest_distributor"`
	DefaultOrderExternalNotes                              any                            `json:"default_order_external_notes"`
	SettingsHideVendorOnPackingSlipPdf                     bool                           `json:"settings_hide_vendor_on_packing_slip_pdf"`
	SettingsDefaultOrderShipmentColor                      string                         `json:"settings_default_order_shipment_color"`
	StorageRemainingInBytes                                int                            `json:"storage_remaining_in_bytes"`
	AllowAPIAccess                                         any                            `json:"allow_api_access"`
	OrganizationID                                         string                         `json:"organization_id"`
	SettingsCalendarDefaultAssemblyDate                    string                         `json:"settings_calendar_default_assembly_date"`
	DefaultPurchaseShippingManifestDistributor             any                            `json:"default_purchase_shipping_manifest_distributor"`
	Deactivated                                            bool                           `json:"deactivated"`
	SettingsAllowPartialMatchingOfTransfersToSalesOrders   bool                           `json:"settings_allow_partial_matching_of_transfers_to_sales_orders"`
	DefaultPurchasePaymentTerm                             any                            `json:"default_purchase_payment_term"`
	Locations                                              []Location                     `json:"locations"`
	ComplianceType                                         string                         `json:"compliance_type"`
	InvoiceSettingsHideCompanyBranding                     bool                           `json:"invoice_settings_hide_company_branding"`
	Settings                                               Settings                       `json:"settings"`
	QbExpenseAccountID                                     any                            `json:"qb_expense_account_id"`
	DefaultOrderPaymentTermID                              int                            `json:"default_order_payment_term_id"`
	Demo                                                   bool                           `json:"demo"`
	SettingsOrderShipmentNumberNext                        int                            `json:"settings_order_shipment_number_next"`
	QuickbooksItems                                        QuickbooksItems                `json:"quickbooks_items"`
	DefaultPurchaseDescription                             any                            `json:"default_purchase_description"`
	SettingsLeaflinkSyncPayments                           bool                           `json:"settings_leaflink_sync_payments"`
	PaymentDelinquentWarningVisibility                     string                         `json:"payment_delinquent_warning_visibility"`
	Licenses                                               []License                      `json:"licenses"`
	HubspotCompanyID                                       any                            `json:"hubspot_company_id"`
	DefaultOrderShipmentShippingManifestDistributor        any                            `json:"default_order_shipment_shipping_manifest_distributor"`
	SettingsStockTransferNumberPrefix                      string                         `json:"settings_stock_transfer_number_prefix"`
	SettingsBreakdownNumberNext                            int                            `json:"settings_breakdown_number_next"`
	Images                                                 []any                          `json:"images"`
	QbSettingsSyncReturns                                  bool                           `json:"qb_settings_sync_returns"`
	DefaultOrderShipmentShippingManifestDistributorID      any                            `json:"default_order_shipment_shipping_manifest_distributor_id"`
	QbSettingsSyncZeroDollarPurchases                      bool                           `json:"qb_settings_sync_zero_dollar_purchases"`
	SettingsCalendarDefaultInvoiceDate                     string                         `json:"settings_calendar_default_invoice_date"`
	SettingsOrderNumberPrefix                              string                         `json:"settings_order_number_prefix"`
	SettingsCalendarDefaultTaskDate                        string                         `json:"settings_calendar_default_task_date"`
	SettingsCalendarDefaultOrderShipmentDate               string                         `json:"settings_calendar_default_order_shipment_date"`
	QbSalesItemID                                          any                            `json:"qb_sales_item_id"`
	QbSettingsSyncPurchasePayments                         bool                           `json:"qb_settings_sync_purchase_payments"`
	SettingsLeaflinkSyncProductPrices                      bool                           `json:"settings_leaflink_sync_product_prices"`
	UUID                                                   string                         `json:"uuid"`
	LeaflinkEnabled                                        bool                           `json:"leaflink_enabled"`
	QbSettingsSyncInvoicePayments                          bool                           `json:"qb_settings_sync_invoice_payments"`
	SettingsTableauSiteID                                  any                            `json:"settings_tableau_site_id"`
	APITokenExpiryDate                                     any                            `json:"api_token_expiry_date"`
	SettingsDefaultAssemblyColor                           string                         `json:"settings_default_assembly_color"`
	SettingsReturnNumberNext                               int                            `json:"settings_return_number_next"`
	OrdersAlwaysRequireInventory                           bool                           `json:"orders_always_require_inventory"`
	URL                                                    string                         `json:"url"`
	SettingsDefaultInvoiceTermsAndServices                 any                            `json:"settings_default_invoice_terms_and_services"`
	SettingsAssemblyNumberNext                             int                            `json:"settings_assembly_number_next"`
	DefaultStockTransferShippingManifestDistributorID      any                            `json:"default_stock_transfer_shipping_manifest_distributor_id"`
	TrymEnabled                                            bool                           `json:"trym_enabled"`
	Category                                               string                         `json:"category"`
	SettingsPaymentNumberNext                              int                            `json:"settings_payment_number_next"`
	DefaultOrderPaymentTerm                                DefaultOrderPaymentTerm        `json:"default_order_payment_term"`
	DefaultStockTransferShippingManifestShipper            any                            `json:"default_stock_transfer_shipping_manifest_shipper"`
	DefaultStockTransferShippingManifestShipperID          any                            `json:"default_stock_transfer_shipping_manifest_shipper_id"`
	DefaultOrderShipmentShippingManifestShipper            any                            `json:"default_order_shipment_shipping_manifest_shipper"`
	DueDateStrategy                                        string                         `json:"due_date_strategy"`
	SubscriptionTier                                       int                            `json:"subscription_tier"`
	SettingsCalendarDefaultSalesOrderDate                  string                         `json:"settings_calendar_default_sales_order_date"`
	UsState                                                string                         `json:"us_state"`
	SettingsDefaultTaskColor                               string                         `json:"settings_default_task_color"`
	SettingsCalendarDefaultPurchaseOrderDate               string                         `json:"settings_calendar_default_purchase_order_date"`
	SettingsAssemblyNumberPrefix                           string                         `json:"settings_assembly_number_prefix"`
	DefaultPurchaseShippingManifestDistributorID           any                            `json:"default_purchase_shipping_manifest_distributor_id"`
	APIToken                                               any                            `json:"api_token"`
	SettingsPurchaseNumberNext                             int                            `json:"settings_purchase_number_next"`
	SettingsStockTransferNumberNext                        int                            `json:"settings_stock_transfer_number_next"`
	DefaultOrderInventoryLocationID                        int                            `json:"default_order_inventory_location_id"`
	DefaultPurchasePaymentTermID                           any                            `json:"default_purchase_payment_term_id"`
	SettingsWeighingUnitTypeID                             any                            `json:"settings_weighing_unit_type_id"`
	SettingsAutoWeighTriggerCriteria                       string                         `json:"settings_auto_weigh_trigger_criteria"`
	SettingsTableauEmail                                   any                            `json:"settings_tableau_email"`
	LegalBusinessName                                      string                         `json:"legal_business_name"`
	SettingsGrafanaUser                                    any                            `json:"settings_grafana_user"`
	QbDepositToAccountID                                   any                            `json:"qb_deposit_to_account_id"`
	SettingsOrderStatusOnTransferMatching                  any                            `json:"settings_order_status_on_transfer_matching"`
}

type CompanyRelationship struct {
	AccountNumber              any            `json:"account_number"`
	Contacts                   []any          `json:"contacts"`
	CustomData                 CustomData     `json:"custom_data"`
	DefaultPaymentTermID       any            `json:"default_payment_term_id"`
	DocumentAddressDisplayName string         `json:"document_address_display_name"`
	GroupID                    any            `json:"group_id"`
	ID                         int            `json:"id"`
	InvoiceEmail               any            `json:"invoice_email"`
	LocationID                 any            `json:"location_id"`
	OrderEmail                 any            `json:"order_email"`
	OrderShipmentEmail         any            `json:"order_shipment_email"`
	OwnerID                    any            `json:"owner_id"`
	PriceTierID                any            `json:"price_tier_id"`
	PurchaseEmail              any            `json:"purchase_email"`
	RelatedCompany             RelatedCompany `json:"related_company"`
	RelationshipTypeID         int            `json:"relationship_type_id"`
	UserID                     any            `json:"user_id"`
}

type CustomData map[string]any

type RelatedCompany struct {
	Weedmaps                  any        `json:"weedmaps"`
	LegalBusinessName         string     `json:"legal_business_name"`
	Facebook                  any        `json:"facebook"`
	DefaultOrderExternalNotes any        `json:"default_order_external_notes"`
	Leafly                    any        `json:"leafly"`
	Name                      string     `json:"name"`
	Phone                     string     `json:"phone"`
	Licenses                  []License  `json:"licenses"`
	Instagram                 any        `json:"instagram"`
	Twitter                   any        `json:"twitter"`
	Images                    []any      `json:"images"`
	Website                   any        `json:"website"`
	ID                        int        `json:"id"`
	Email                     any        `json:"email"`
	Description               any        `json:"description"`
	Category                  any        `json:"category"`
	Locations                 []Location `json:"locations"`
}
