package company

import "time"

type TableDataResponse struct {
	TableColumns []TableColumn `json:"table_columns"`
	TableData    []TableData   `json:"table_data"`
}

type TableColumn struct {
	CustomFieldID any    `json:"custom_field_id"`
	FieldPath     string `json:"field_path"`
	ID            int    `json:"id"`
	ImagePath     any    `json:"image_path"`
	IsCopyable    bool   `json:"is_copyable"`
	IsDate        bool   `json:"is_date"`
	IsExternalURL bool   `json:"is_external_url"`
	Name          string `json:"name"`
	Official      bool   `json:"official"`
	Position      int    `json:"position"`
	URLPath       string `json:"url_path"`
	Visible       bool   `json:"visible"`
}

type TableData struct {
	Email          string     `json:"email"`
	BatchCount     int        `json:"batch_count"`
	CompanyID      int        `json:"company_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
	LeaflinkBrand  string     `json:"leaflink_brand"`
	Weedmaps       string     `json:"weedmaps"`
	RelatedCompany struct {
		Category                     any      `json:"category"`
		DefaultOrderExternalNotes    any      `json:"default_order_external_notes"`
		DefaultOrderPaymentTermID    any      `json:"default_order_payment_term_id"`
		DefaultPurchasePaymentTermID any      `json:"default_purchase_payment_term_id"`
		Description                  any      `json:"description"`
		Email                        any      `json:"email"`
		ID                           int      `json:"id"`
		Images                       []any    `json:"images"`
		LegalBusinessName            string   `json:"legal_business_name"`
		Name                         string   `json:"name"`
		Phone                        string   `json:"phone"`
		Settings                     Settings `json:"settings"`
		URL                          string   `json:"url"`
	} `json:"related_company"`
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Phone        string     `json:"phone"`
	Description  string     `json:"description"`
	OwnerID      int        `json:"owner_id"`
	ProductCount int        `json:"product_count"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Locations    []struct {
		Address   string  `json:"address"`
		Apt       string  `json:"apt"`
		City      string  `json:"city"`
		CompanyID int     `json:"company_id"`
		Country   string  `json:"country"`
		ID        int     `json:"id"`
		Latitude  float64 `json:"latitude"`
		License   string  `json:"license"`
		LicenseID any     `json:"license_id"`
		Longitude float64 `json:"longitude"`
		Manual    bool    `json:"manual"`
		MetaData  struct {
			Description string `json:"description"`
			Gmaps       struct {
				AddressComponents []struct {
					LongName  string   `json:"long_name"`
					ShortName string   `json:"short_name"`
					Types     []string `json:"types"`
				} `json:"address_components"`
				AdrAddress       string `json:"adr_address"`
				FormattedAddress string `json:"formatted_address"`
				Geometry         struct {
					Location struct {
					} `json:"location"`
					Viewport struct {
						Ha struct {
							Hi float64 `json:"hi"`
							Lo float64 `json:"lo"`
						} `json:"ha"`
						Va struct {
							Hi float64 `json:"hi"`
							Lo float64 `json:"lo"`
						} `json:"va"`
					} `json:"viewport"`
				} `json:"geometry"`
				HTMLAttributions    []any    `json:"html_attributions"`
				Icon                string   `json:"icon"`
				IconBackgroundColor string   `json:"icon_background_color"`
				IconMaskBaseURI     string   `json:"icon_mask_base_uri"`
				Name                string   `json:"name"`
				PlaceID             string   `json:"place_id"`
				Reference           string   `json:"reference"`
				Types               []string `json:"types"`
				URL                 string   `json:"url"`
				UtcOffset           int      `json:"utc_offset"`
				UtcOffsetMinutes    int      `json:"utc_offset_minutes"`
				Vicinity            string   `json:"vicinity"`
			} `json:"gmaps"`
			IsFixture bool   `json:"is_fixture"`
			Label     string `json:"label"`
			Location  struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			MatchedSubstrings struct {
				Length int `json:"length"`
				Offset int `json:"offset"`
			} `json:"matched_substrings"`
			PlaceID string `json:"place_id"`
		} `json:"meta_data"`
		Name           string `json:"name"`
		NumberOfLights any    `json:"number_of_lights"`
		SquareFeet     any    `json:"square_feet"`
		State          string `json:"state"`
		StreetAddress  string `json:"street_address"`
		Wattage        any    `json:"wattage"`
		Zip            string `json:"zip"`
	} `json:"locations"`
	RelatedCompanyID     int `json:"related_company_id"`
	DefaultPaymentTermID any `json:"default_payment_term_id"`
	Owner                struct {
		Avatar             any    `json:"avatar"`
		CompanyID          int    `json:"company_id"`
		DeactivatedAt      any    `json:"deactivated_at"`
		DefaultDashboardID any    `json:"default_dashboard_id"`
		DeletedAt          any    `json:"deleted_at"`
		Description        any    `json:"description"`
		Email              string `json:"email"`
		FirstName          string `json:"first_name"`
		ID                 int    `json:"id"`
		Initials           string `json:"initials"`
		LastName           string `json:"last_name"`
		Name               string `json:"name"`
		Phone              any    `json:"phone"`
		Profile            struct {
			Avatar      any    `json:"avatar"`
			Description any    `json:"description"`
			FirstName   string `json:"first_name"`
			ID          int    `json:"id"`
			Initials    string `json:"initials"`
			LastName    string `json:"last_name"`
			Name        string `json:"name"`
			Phone       any    `json:"phone"`
			QbInitials  string `json:"qb_initials"`
			Title       any    `json:"title"`
			WorkPhone   any    `json:"work_phone"`
		} `json:"profile"`
		ProfileID  int    `json:"profile_id"`
		QbInitials string `json:"qb_initials"`
		RoleID     int    `json:"role_id"`
		Settings   struct {
			ActiveCompanyRelationshipFilterID int    `json:"active_company_relationship_filter_id"`
			ActiveContactFilterID             int    `json:"active_contact_filter_id"`
			ActiveOrderFilterID               int    `json:"active_order_filter_id"`
			ActiveProductFilterID             int    `json:"active_product_filter_id"`
			EmailInclusionInvoicesEmail       any    `json:"email_inclusion_invoices_email"`
			EmailInclusionInvoicesLevel       string `json:"email_inclusion_invoices_level"`
			EmailInclusionOrdersEmail         any    `json:"email_inclusion_orders_email"`
			EmailInclusionOrdersLevel         string `json:"email_inclusion_orders_level"`
			EmailInclusionPackingSlipsEmail   any    `json:"email_inclusion_packing_slips_email"`
			EmailInclusionPackingSlipsLevel   string `json:"email_inclusion_packing_slips_level"`
			EmailInclusionPurchasesEmail      any    `json:"email_inclusion_purchases_email"`
			EmailInclusionPurchasesLevel      string `json:"email_inclusion_purchases_level"`
			EmailInclusionTransfersEmail      any    `json:"email_inclusion_transfers_email"`
			EmailInclusionTransfersLevel      string `json:"email_inclusion_transfers_level"`
			OrdersIndexDefaultView            string `json:"orders_index_default_view"`
			ShowAssemblyHistoryTour           bool   `json:"show_assembly_history_tour"`
			ShowBatchTour                     bool   `json:"show_batch_tour"`
			ShowDashboardTour                 bool   `json:"show_dashboard_tour"`
			ShowHolidayFourTwenty             bool   `json:"show_holiday_four_twenty"`
			ShowHolidayHalloween              bool   `json:"show_holiday_halloween"`
			ShowHolidayNewYears               bool   `json:"show_holiday_new_years"`
			ShowHolidayThanksgiving           bool   `json:"show_holiday_thanksgiving"`
			ShowMapTour                       bool   `json:"show_map_tour"`
			ShowProductTour                   bool   `json:"show_product_tour"`
			ShowReportTableTour               bool   `json:"show_report_table_tour"`
			ShowTableTour                     bool   `json:"show_table_tour"`
			ShowTransferTour                  bool   `json:"show_transfer_tour"`
		} `json:"settings"`
		Superadmin      bool   `json:"superadmin"`
		TeamMemberships []any  `json:"team_memberships"`
		Title           any    `json:"title"`
		URL             string `json:"url"`
		WorkPhone       any    `json:"work_phone"`
	} `json:"owner"`
	ContactCount       int              `json:"contact_count"`
	RelationshipTypeID any              `json:"relationship_type_id"`
	Twitter            any              `json:"twitter"`
	Website            any              `json:"website"`
	URL                string           `json:"url"`
	Instagram          any              `json:"instagram"`
	InsertedAt         time.Time        `json:"inserted_at"`
	Category           any              `json:"category"`
	RelationshipType   any              `json:"relationship_type"`
	LeaflinkCustomer   LeaflinkCustomer `json:"leaflink_customer"`
	PosType            any              `json:"pos_type"`
	Facebook           any              `json:"facebook"`
	LegalBusinessName  string           `json:"legal_business_name"`
	Model              string           `json:"model"`
	CustomData         map[string]any   `json:"custom_data"`
	Leafly             any              `json:"leafly"`
}

type LeaflinkCustomer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
