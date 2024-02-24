package response

type GetProductResponse struct {
	Id string `json:"id"`
}

type GetServicesResponse struct {
	Services []Service `json:"services"`
}

type ComponentCost struct {
	Id                     int64                   `json:"id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	ComponentCostLineItems []ComponentCostLineItem `json:"component_cost_line_items"`
}

type GetComponentCostLineItemDropdown struct {
	ComponentCostLineItems []ComponentCostLineItem `json:"component_cost_line_items"`
}

type ComponentCostLineItem struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int64   `json:"qty"`
	Notes string  `json:"notes"`
}

type Service struct {
	Id                      int64           `json:"id"`
	Name                    string          `json:"name"`
	Description             string          `json:"description"`
	CostOfOrigins           []ComponentCost `json:"cost_of_origins"`
	CostOfFinalDestinations []ComponentCost `json:"cost_of_final_destinations"`
	FreightCosts            []ComponentCost `json:"freight_costs"`
}

type GetPermissionsResponse struct {
	Permissions []Permission `json:"permission"`
}

type Permission struct {
	Id              int64  `json:"id"`
	ContentTypeId   int64  `json:"content_type_id"`
	Name            string `json:"name"`
	ContentTypeName string `json:"content_type_name"`
}

type GetComponentCostsResponse struct {
	ComponentCosts []ComponentCost `json:"component_costs"`
}

type GetGoodsResponse struct {
	Goods []Good `json:"goods"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserContentTypesResponse struct {
	UserId       int64         `json:"user_id"`
	UserName     string        `json:"user_name"`
	UserType     string        `json:"user_type"`
	ContentTypes []ContentType `json:"content_types"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Type     string `json:"type"`
	ClientId int64  `json:"clientId"`
	Roles    []Role `json:"roles"`
}

type GetRolesResponse struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	CreatedAt   string       `json:"created_at"`
	Permissions []Permission `json:"permission"`
}

type GetContentTypesResponse struct {
	ContentTypes []ContentType `json:"content_types"`
}

type ContentType struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	CreatedAt   string       `json:"created_at"`
	Permissions []Permission `json:"permissions"`
}

type GetDocumentsResponse struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	Id            int64  `json:"id"`
	Category      string `json:"category"`
	CategoryLabel string `json:"category_label"`
	CreatedAt     string `json:"created_at"`
}

type GetDocumentDetailResponse struct {
	Id            int64  `json:"id"`
	Category      string `json:"category"`
	CategoryLabel string `json:"category_label"`
	CreatedAt     string `json:"created_at"`

	ReceiverName  string `json:"receiver_name"`
	CompanyName   string `json:"company_name"`
	NoPE          string `json:"no_pe"`
	StuffingDate  string `json:"stuffing_date"`
	Destination   string `json:"destination"`
	VehicleType   string `json:"vehicle_type"`
	TicketNumber  string `json:"ticket_number"`
	PlateNumber   string `json:"plate_number"`
	SignatureName string `json:"signature_name"`

	DocumentNumber string `json:"document_number"`
	ProvinceName   string `json:"province_name"`
	Date           string `json:"date"`
	Attention      string `json:"attention"`
	ShipperName    string `json:"shipper_name"`
	ConsigneeName  string `json:"consignee_name"`
	NotifyParty    string `json:"notify_party"`
	POL            string `json:"pol"`
	POD            string `json:"pod"`

	Subject         string `json:"subject"`
	ContainerNumber string `json:"container_number"`
	TravelNumber    string `json:"travel_number"`
	Dooring         string `json:"dooring"`
	Detail          string `json:"detail"`

	SenderName string `json:"sender_name"`
	SeatNumber string `json:"seat_number"`

	ShipperAddress             string `json:"shipper_address"`
	ShipperDestination         string `json:"shipper_destination"`
	ShipperPhoneNumber         string `json:"shipper_phone_number"`
	ShipperMobilePhoneNumber   string `json:"shipper_mobile_phone_number"`
	ShipperEmailAddress        string `json:"shipper_email_address"`
	ShipperNpwp                string `json:"shipper_npwp"`
	ConsigneeAddress           string `json:"consignee_address"`
	ConsigneeDestination       string `json:"consignee_destination"`
	ConsigneePhoneNumber       string `json:"consignee_phone_number"`
	ConsigneeMobilePhoneNumber string `json:"consignee_mobile_phone_number"`
	ConsigneeEmailAddress      string `json:"consignee_email_address"`
	ConsigneeNpwp              string `json:"consignee_npwp"`
	NotifyPartyAddress         string `json:"notify_party_address"`
	NotifyPartyPhoneNumber     string `json:"notify_party_phone_number"`
	ShippingTerm               string `json:"shipping_term"`
	VesselName                 string `json:"vessel_name"`
	VoyageNumber               string `json:"voyage_number"`
	OceanFreightPrepaid        string `json:"ocean_freight_prepaid"`
	OceanFreightCollect        string `json:"ocean_freight_collect"`
	POLPrepaid                 string `json:"pol_prepaid"`
	POLCollect                 string `json:"pol_collect"`
	PODPrepaid                 string `json:"pod_prepaid"`
	PODCollect                 string `json:"pod_collect"`
	DangerousClass             string `json:"dangerous_class"`
	DangerousUnNumber          string `json:"dangerous_un_number"`
	ReeferTemperature          string `json:"reefer_temperature"`
	ReeferNote                 string `json:"reefer_note"`
	TranshipmentImport1        string `json:"transhipment_import_1"`
	TranshipmentImport2        string `json:"transhipment_import_2"`
	TranshipmentImport3        string `json:"transhipment_import_3"`
	TranshipmentExport1        string `json:"transhipment_export_1"`
	TranshipmentExport2        string `json:"transhipment_export_2"`
	TranshipmentExport3        string `json:"transhipment_export_3"`
	MadeBy                     string `json:"made_by"`
	Position                   string `json:"position"`

	Voy                    string `json:"voy"`
	Etd                    string `json:"etd"`
	Eta                    string `json:"eta"`
	ConnectingVesselName   string `json:"connecting_vessel_name"`
	ConnectingVesselEtd    string `json:"connecting_vessel_etd"`
	ConnectingVesselEta    string `json:"connecting_vessel_eta"`
	ConnectingVesselVoy    string `json:"connecting_vessel_voy"`
	SpkNumber              string `json:"spk_number"`
	SpkDalNumber           string `json:"spk_dal_number"`
	PoNumber               string `json:"po_number"`
	TermsOfShipment        string `json:"terms_of_shipment"`
	FreightType            string `json:"freight_type"`
	Movement               string `json:"movement"`
	NotifyNumber           string `json:"notify_number"`
	ConsigneeContactPerson string `json:"consignee_contact_person"`

	ConsigneeId        int64  `json:"consignee_id"`
	ConsigneeAddressId int64  `json:"consignee_address_id"`
	CurrentYear        string `json:"current_year"`

	DocumentLineItems          []DocumentLineItem `json:"document_line_items"`
	SecondaryDocumentLineItems []DocumentLineItem `json:"secondary_document_line_items"`
}

type DocumentLineItem struct {
	Id              int64  `json:"id"`
	Category        string `json:"category"`
	Order           int64  `json:"order"`
	ItemCode        string `json:"item_code"`
	Qty             string `json:"qty"`
	ItemName        string `json:"item_name"`
	Weight          string `json:"weight"`
	Description     string `json:"description"`
	ContainerNumber string `json:"container_number"`
	SealNumber      string `json:"seal_number"`
	PlateNumber     string `json:"plate_number"`
	StuffingDate    string `json:"stuffing_date"`
	Name            string `json:"name"`
}

type Good struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	Description       string `json:"description"`
}

type GetTotalResponse struct {
	Total int64 `json:"total"`
}

type SuccessMessageResponse struct {
	SuccessMessage string `json:"success_message"`
}

type GetVendorTrucksResponse struct {
	VendorTrucks []VendorTruck `json:"vendor_trucks"`
}

type GetVendorTrucksDropdownByVendorResponse struct {
	VendorTruckId int64   `json:"vendor_truck_id"`
	Trucks        []Truck `json:"trucks"`
}

type VendorTruck struct {
	Id         int64   `json:"id"`
	VendorId   int64   `json:"vendor_id"`
	VendorName string  `json:"vendor_name"`
	Trucks     []Truck `json:"trucks"`
}

type Truck struct {
	Id           int64   `json:"id"`
	ProvinceId   int64   `json:"province_id"`
	ProvinceName string  `json:"province_name"`
	Price        float64 `json:"price"`
	Name         string  `json:"name"`
}

type GetServiceDropdownComponentCostsResponse struct {
}

type GetQuotationsResponse struct {
	Quotations []Quotation `json:"quotations"`
}

type Quotation struct {
	Id                  int64  `json:"id"`
	QuotationNumber     string `json:"quotation_number"`
	ClientId            int64  `json:"client_id"`
	ClientName          string `json:"client_name"`
	ConsigneeId         int64  `json:"consignee_id"`
	ConsigneeName       string `json:"consignee_name"`
	PortOfLoadingId     int64  `json:"port_of_loading_id"`
	PortOfDestinationId int64  `json:"port_of_destination_id"`
	Date                string `json:"date"`
}

type GetShippingsResponse struct {
	Shippings []Shipping `json:"shippings"`
}

type Shipping struct {
	Id            int64  `json:"id"`
	DONumber      string `json:"do_number"`
	ClientId      int64  `json:"client_id"`
	ClientName    string `json:"client_name"`
	ConsigneeId   int64  `json:"consignee_id"`
	ConsigneeName string `json:"consignee_name"`
	Progress      string `json:"progress"`
	Deadline      string `json:"deadline"`
	DateCreated   string `json:"date_created"`
}

type GetShippingDetailResponse struct {
	Id                       int64                    `json:"id"`
	DONumber                 string                   `json:"do_number"`
	QuotationId              int64                    `json:"quotation_id"`
	ClientId                 int64                    `json:"client_id"`
	ClientName               string                   `json:"client_name"`
	ConsigneeId              int64                    `json:"consignee_id"`
	ConsigneeName            string                   `json:"consignee_name"`
	PortOfLoadingId          int64                    `json:"port_of_loading_id"`
	PortOfDestinationId      int64                    `json:"port_of_destination_id"`
	SpkNumber                string                   `json:"spk_number"`
	Deadline                 string                   `json:"deadline"`
	Status                   string                   `json:"status"`
	DateCreated              string                   `json:"date_created"`
	MotherVessel             string                   `json:"mother_vessel"`
	ShippingTrackingStatuses []ShippingTrackingStatus `json:"shipping_tracking_statuses"`
	Trackings                []Tracking               `json:"trackings"`
	Cost                     ShippingCost             `json:"cost"`
	ShippingTasks            []ShippingTask           `json:"shipping_tasks"`
}

type ShippingTask struct {
	TaskId         int64  `json:"id"`
	ShippingId     int64  `json:"shipping_id"`
	Status         string `json:"status"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	CompletionDate string `json:"td"`
	DueDate        string `json:"due_date"`
}

type ShippingCost struct {
	Id                      int64           `json:"id"`
	Status                  string          `json:"status"`
	CommissionAType         string          `json:"commission_a_type"`
	CommissionATypeValue    float64         `json:"commission_a_type_value"`
	CommissionAValue        float64         `json:"commission_a_value"`
	CommissionANote         string          `json:"commission_a_note"`
	CommissionBType         string          `json:"commission_b_type"`
	CommissionBTypeValue    float64         `json:"commission_b_type_value"`
	CommissionBValue        float64         `json:"commission_b_value"`
	CommissionBNote         string          `json:"commission_b_note"`
	MarginType              string          `json:"margin_type"`
	MarginTypeValue         float64         `json:"margin_type_value"`
	MarginValue             float64         `json:"margin_value"`
	MarginRoundValue        float64         `json:"margin_round_value"`
	MarginNote              string          `json:"margin_note"`
	FinalPriceValue         float64         `json:"final_price_value"`
	FinalPriceNote          string          `json:"final_price_note"`
	CostOfOrigins           []QuotationInfo `json:"cost_of_origins"`
	FreightCosts            []QuotationInfo `json:"freight_costs"`
	CostOfFinalDestinations []QuotationInfo `json:"cost_of_final_destinations"`
}

type Tracking struct {
	Id                   int64              `json:"id"`
	ContainerId          int64              `json:"container_id"`
	TravelDocumentNumber string             `json:"travel_document_number"`
	DriverNumber         string             `json:"driver_number"`
	PlateNumber          string             `json:"plate_number"`
	ContainerNumber      string             `json:"container_number"`
	ShippingRoute        string             `json:"shipping_route"`
	SealNumber           string             `json:"seal_number"`
	MotherVessel         string             `json:"mother_vessel"`
	FeNumber             string             `json:"fe_number"`
	Voyage               string             `json:"voyage"`
	StuffingDate         string             `json:"stuffing_date"`
	Note                 string             `json:"note"`
	TrackingProgresses   []TrackingProgress `json:"tracking_progresses"`
}

type TrackingProgress struct {
	Id                    int64  `json:"id"`
	Order                 int64  `json:"order"`
	Category              string `json:"category"`
	VesselName            string `json:"vessel_name"`
	VoyageName            string `json:"voyage_name"`
	EstimatedTimeDelivery string `json:"estimated_time_delivery"`
	TimeDelivery          string `json:"time_delivery"`
	EstimatedTimeArrival  string `json:"estimated_time_arrival"`
	TimeArrival           string `json:"time_arrival"`
	Berthing              string `json:"berthing"`
	Unloading             string `json:"unloading"`
	Note                  string `json:"note"`
}

type GetShippingTrackingsResponse struct {
	ShippingTrackings []ShippingTracking `json:"trackings"`
}

type ShippingTracking struct {
	ShippingId           int64  `json:"shipping_id"`
	QuotationId          int64  `json:"quotation_id"`
	DoNumber             string `json:"do_number"`
	TrackingId           int64  `json:"tracking_id"`
	ContainerId          int64  `json:"container_id"`
	TravelDocumentNumber string `json:"travel_document_number"`
	DriverNumber         string `json:"driver_number"`
	PlateNumber          string `json:"plate_number"`
	ContainerNumber      string `json:"container_number"`
	ShippingRoute        string `json:"shipping_route"`
	SealNumber           string `json:"seal_number"`
	MotherVessel         string `json:"mother_vessel"`
	FeNumber             string `json:"fe_number"`
	Voyage               string `json:"voyage"`
	StuffingDate         string `json:"stuffing_date"`
	Note                 string `json:"note"`
}

type GetQuotationCollectionsResponse struct {
	QuotationCollections []QuotationCollection `json:"quotation_collections"`
}

type GetQuotationsOnCollectionsDropdownResponse struct {
	Quotations []QuotationOnCollectionDropdown `json:"quotations"`
}

type QuotationOnCollectionDropdown struct {
	Id                  int64  `json:"id"`
	QuotationNumber     string `json:"quotation_number"`
	ClientId            int64  `json:"client_id"`
	ConsigneeId         int64  `json:"consignee_id"`
	PortOfLoadingId     int64  `json:"port_of_loading_id"`
	PortOfDestinationId int64  `json:"port_of_destination_id"`
}

type QuotationCollection struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	ClientName  string `json:"client_name"`
}

type GetQuotationsByIdsResponse struct {
	Quotations []GetQuotation `json:"quotations"`
}

type GetQuotation struct {
	Id                         int64           `json:"id"`
	QuotationNumber            string          `json:"quotation_number"`
	Date                       string          `json:"date"`
	ServiceId                  int64           `json:"service_id"`
	ServiceName                string          `json:"service_name"`
	Note                       string          `json:"note"`
	PeNumber                   string          `json:"pe_number"`
	ScheduleId                 int64           `json:"schedule_id"`
	ScheduleETD                string          `json:"schedule_etd"`
	ScheduleETA                string          `json:"schedule_eta"`
	ScheduleDaysUntilArrived   int64           `json:"schedule_days_till_arrived"`
	ScheduleFinalETA           string          `json:"schedule_final_eta"`
	ScheduleNote               string          `json:"schedule_note"`
	ScheduleClosingTime        string          `json:"schedule_closing_time"`
	ClientId                   int64           `json:"client_id"`
	ClientName                 string          `json:"client_name"`
	ClientWarehouseId          int64           `json:"client_warehouse_id"`
	ClientWarehouseAddress     string          `json:"client_warehouse_address"`
	ConsigneeId                int64           `json:"consignee_id"`
	ConsigneeName              string          `json:"consignee_name"`
	ConsigneeWarehouseId       int64           `json:"consignee_warehouse_id"`
	ConsigneeWarehouseAddress  string          `json:"consignee_warehouse_address"`
	PortOfLoadingId            int64           `json:"port_of_loading_id"`
	PortOfDestinationId        int64           `json:"port_of_destination_id"`
	GoodId                     int64           `json:"good_id"`
	GoodName                   string          `json:"good_name"`
	GoodUnitOfMeasurement      string          `json:"good_unit_of_measurement"`
	GoodQty                    int64           `json:"good_qty"`
	GoodWidth                  float64         `json:"good_width"`
	GoodLength                 float64         `json:"good_length"`
	GoodDepth                  float64         `json:"good_depth"`
	GoodWeight                 string          `json:"good_weight"`
	ContainerId                int64           `json:"container_id"`
	ContainerUnitOfMeasurement string          `json:"container_unit_of_measurement"`
	ContainerNote              string          `json:"container_note"`
	CommissionAType            string          `json:"commission_a_type"`
	CommissionATypeValue       float64         `json:"commission_a_type_value"`
	CommissionAValue           float64         `json:"commission_a_value"`
	CommissionANote            string          `json:"commission_a_note"`
	CommissionBType            string          `json:"commission_b_type"`
	CommissionBTypeValue       float64         `json:"commission_b_type_value"`
	CommissionBValue           float64         `json:"commission_b_value"`
	CommissionBNote            string          `json:"commission_b_note"`
	MarginType                 string          `json:"margin_type"`
	MarginTypeValue            float64         `json:"margin_type_value"`
	MarginValue                float64         `json:"margin_value"`
	MarginRoundValue           float64         `json:"margin_round_value"`
	MarginNote                 string          `json:"margin_note"`
	FinalPriceValue            float64         `json:"final_price_value"`
	FinalPriceNote             string          `json:"final_price_note"`
	CostOfOrigins              []QuotationInfo `json:"cost_of_origins"`
	FreightCosts               []QuotationInfo `json:"freight_costs"`
	CostOfFinalDestinations    []QuotationInfo `json:"cost_of_final_destinations"`
}

type QuotationInfo struct {
	Id                          int64                        `json:"id"`
	VendorId                    int64                        `json:"vendor_id"`
	ComponentCostCategory       string                       `json:"component_cost_category"`
	Note                        string                       `json:"notes"`
	QuotationTruckInfos         []QuotationTruckInfo         `json:"trucks"`
	QuotationAddCostInfos       []QuotationAddCostInfo       `json:"additional_costs"`
	QuotationComponentCostInfos []QuotationComponentCostInfo `json:"component_costs"`
}

type QuotationTruckInfo struct {
	Id            int64   `json:"id"`
	VendorTruckId int64   `json:"vendor_truck_id"`
	TruckId       int64   `json:"truck_id"`
	Price         float64 `json:"price"`
	Qty           int64   `json:"qty"`
	Note          string  `json:"note"`
}

type QuotationAddCostInfo struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int64   `json:"qty"`
}

type QuotationComponentCostInfo struct {
	Id                     int64                   `json:"id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	ComponentCostLineItems []ComponentCostLineItem `json:"component_cost_line_items"`
}

type GetTasksResponse struct {
	Tasks []Task `json:"shipping_tasks"`
}

type Task struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type GetExternalShippingTrackingStatusesResponse struct {
	SpkNumber                string                   `json:"spk_number"`
	DONumber                 string                   `json:"do_number"`
	Deadline                 string                   `json:"deadline"`
	ShippingTrackingStatuses []ShippingTrackingStatus `json:"shipping_tracking_statuses"`
}

type ShippingTrackingStatus struct {
	Id                    int64  `json:"id"`
	ContainerName         string `json:"container_name"`
	StuffingDate          string `json:"stuffing_date"`
	EstimatedTimeDelivery string `json:"estimated_time_delivery"`
	TimeDelivery          string `json:"time_delivery"`
	EstimatedTimeArrival  string `json:"estimated_time_arrival"`
	Berthing              string `json:"berthing"`
	Unloading             string `json:"unloading"`
	PlanDelivery          string `json:"plan_delivery"`
	Delivery              string `json:"delivery"`
	Status                string `json:"status"`
}

type GetInvoicesResponse struct {
	Invoices []Invoice `json:"invoices"`
}

type GetInvoiceDetailResponse struct {
	Id                int64                      `json:"id"`
	ClientId          int64                      `json:"client_id"`
	InvoiceNumber     string                     `json:"invoice_number"`
	Status            string                     `json:"status"`
	Date              string                     `json:"date"`
	ShipperId         int64                      `json:"shipper_id"`
	ShipperName       string                     `json:"shipper_name"`
	ConsigneeId       int64                      `json:"consignee_id"`
	ConsigneeName     string                     `json:"consignee_name"`
	NotifyParty       string                     `json:"notify_party"`
	GoodId            int64                      `json:"good_id"`
	GoodDescription   string                     `json:"good_description"`
	PolId             int64                      `json:"pol_id"`
	PolName           string                     `json:"pol_name"`
	PodId             int64                      `json:"pod_id"`
	PodName           string                     `json:"pod_name"`
	Vessel            string                     `json:"vessel"`
	Etd               string                     `json:"etd"`
	TotalPrice        float64                    `json:"total_price"`
	Quotations        []InvoiceQuotationRelation `json:"quotations"`
	InvoiceContainers []InvoiceContainer         `json:"invoice_containers"`
	InvoiceAddCosts   []InvoiceAddCost           `json:"invoice_add_costs"`
}

type Invoice struct {
	Id                int64                      `json:"id"`
	ClientId          int64                      `json:"client_id"`
	InvoiceNumber     string                     `json:"invoice_number"`
	Status            string                     `json:"status"`
	Date              string                     `json:"date"`
	ShipperId         int64                      `json:"shipper_id"`
	ShipperName       string                     `json:"shipper_name"`
	ConsigneeId       int64                      `json:"consignee_id"`
	ConsigneeName     string                     `json:"consignee_name"`
	NotifyParty       string                     `json:"notify_party"`
	GoodId            int64                      `json:"good_id"`
	GoodDescription   string                     `json:"good_description"`
	PolId             int64                      `json:"pol_id"`
	PolName           string                     `json:"pol_name"`
	PodId             int64                      `json:"pod_id"`
	PodName           string                     `json:"pod_name"`
	Vessel            string                     `json:"vessel"`
	Etd               string                     `json:"etd"`
	TotalPrice        float64                    `json:"total_price"`
	Quotations        []InvoiceQuotationRelation `json:"quotations"`
	InvoiceContainers []InvoiceContainer         `json:"invoice_containers"`
	InvoiceAddCosts   []InvoiceAddCost           `json:"invoice_add_costs"`
}

type InvoiceQuotationRelation struct {
	Id              int64  `json:"id"`
	QuotationId     int64  `json:"quotation_id"`
	QuotationNumber string `json:"quotation_number"`
}

type InvoiceContainer struct {
	Id                   int64   `json:"id"`
	ContainerId          int64   `json:"container_id"`
	ContainerNumber      string  `json:"container_number"`
	DoNumber             string  `json:"do_number"`
	Seal                 string  `json:"seal"`
	PlateNumber          string  `json:"plate_number"`
	TravelDocumentNumber string  `json:"travel_document_number"`
	Qty                  int64   `json:"qty"`
	TosServiceName       string  `json:"tos_service_name"`
	TosContainerName     string  `json:"tos_container_name"`
	FinalPriceValue      float64 `json:"final_price_value"`
}

type InvoiceAddCost struct {
	Id          int64   `json:"id"`
	Description string  `json:"description"`
	Qty         int64   `json:"qty"`
	Cost        float64 `json:"cost"`
}

type GetInvoiceDropdownByClientResponse struct {
	Invoices []InvoiceByClient `json:"invoices"`
}

type InvoiceByClient struct {
	Id            int64   `json:"id"`
	InvoiceNumber string  `json:"invoice_number"`
	TotalPrice    float64 `json:"total_price"`
	PaymentAmount float64 `json:"payment_amount"`
}

type GetInvoicePaymentsResponse struct {
	InvoicePayments []InvoicePayment `json:"invoice_payments"`
}

type InvoicePayment struct {
	Id            int64   `json:"id"`
	InvoiceNumber string  `json:"invoice_number"`
	Date          string  `json:"payment_date"`
	ClientId      int64   `json:"client_id"`
	ClientName    string  `json:"client_name"`
	PaymentMethod string  `json:"payment_method"`
	PaymentAmount float64 `json:"payment_amount"`
	Status        string  `json:"status"`
	Note          string  `json:"note"`
}

type GetInvoicePaymentDetailResponse struct {
	Id                int64   `json:"id"`
	Date              string  `json:"date"`
	ClientId          int64   `json:"client_id"`
	ClientName        string  `json:"client_name"`
	Status            string  `json:"status"`
	InvoiceId         int64   `json:"invoice_id"`
	InvoiceNumber     string  `json:"invoice_number"`
	InvoiceTotalPrice float64 `json:"invoice_total_price"`
	PaymentAmount     float64 `json:"payment_amount"`
	PaymentMethod     string  `json:"payment_method"`
	RestOfTheBill     float64 `json:"rest_of_the_bill"`
	BankAccount       string  `json:"bank_account"`
	Note              string  `json:"note"`
}

type GetInvoicePaymentHistoryResponse struct {
	InvoiceTotalPrice  float64          `json:"invoice_total_price"`
	TotalPaymentAmount float64          `json:"total_payment"`
	PaymentStatus      string           `json:"payment_status"`
	ProofOfCutPph      string           `json:"proof_of_cut_pph"`
	PaymentHistories   []PaymentHistory `json:"payment_histories"`
}

type PaymentHistory struct {
	Date          string  `json:"payment_date"`
	PaymentAmount float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	BankAccount   string  `json:"bank_account"`
	Note          string  `json:"note"`
}

type GetExpensesResponse struct {
	Expenses []Expense `json:"expenses"`
}

type Expense struct {
	Id               int64             `json:"id"`
	ExpenseNumber    string            `json:"expense_number"`
	Status           string            `json:"status"`
	TotalAmount      float64           `json:"total_amount"`
	Date             string            `json:"date"`
	PaymentMethod    string            `json:"payment_method"`
	BankAccountId    int64             `json:"bank_account_id"`
	BankAccountName  string            `json:"bank_account_name"`
	Description      string            `json:"description"`
	ExpenseLineItems []ExpenseLineItem `json:"expense_line_items"`
}

type ExpenseLineItem struct {
	Id          int64   `json:"id"`
	CoaId       int64   `json:"coa_id"`
	CoaName     string  `json:"coa_name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Note        string  `json:"note"`
}

type GetExpenseDetailResponse struct {
	Id               int64             `json:"id"`
	ExpenseNumber    string            `json:"expense_number"`
	Status           string            `json:"status"`
	TotalAmount      float64           `json:"total_amount"`
	Date             string            `json:"date"`
	PaymentMethod    string            `json:"payment_method"`
	BankAccountId    int64             `json:"bank_account_id"`
	BankAccountName  string            `json:"bank_account_name"`
	Description      string            `json:"description"`
	ExpenseLineItems []ExpenseLineItem `json:"expense_line_items"`
}

type GetMstVendorTypesDropdownResponse struct {
	VendorTypes []VendorType `json:"data"`
}

type VendorType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetMstVendorsResponse struct {
	Vendors []Vendor `json:"vendors"`
}

type Vendor struct {
	Id             int64  `json:"id"`
	CompanyName    string `json:"company_name"`
	PicName        string `json:"pic_name"`
	PicPhoneNumber string `json:"pic_phone_number"`
	ProvinceName   string `json:"location"`
	CreatedAt      string `json:"created_at"`
}

type VendorInformation struct {
	Id         int64 `json:"id"`
	ProvinceId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"province_id"`
	CityId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"city_id"`
	DistrictId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"district_id"`
	Zipcode      int64 `json:"zipcode"`
	VendorTypeId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"vendor_type_id"`
	CompanyName   string        `json:"company_name"`
	Email         string        `json:"email"`
	NPWP          string        `json:"npwp"`
	StreetAddress string        `json:"street_address"`
	PhoneNumbers  []PhoneNumber `json:"phone_numbers"`
}

type PhoneNumber struct {
	OrderNumber int64  `json:"order_number"`
	PhoneNumber string `json:"phone_number"`
}

type VendorPic struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Position     string        `json:"position"`
	OrderNumber  int64         `json:"order_number"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type VendorContainerLaborEquipmentPrice struct {
	Id             int64          `json:"id"`
	ContainerId    VclepContainer `json:"container_id"`
	LaborPrice     int64          `json:"labor_price"`
	EquipmentPrice int64          `json:"equipment_price"`
	Notes          string         `json:"notes"`
}

type VclepContainer struct {
	Id   int64  `json:"id"`
	Size string `json:"size"`
}

type GetClientDetailResponse struct {
	ClientInformation ClientInformation `json:"client_information"`
	ClientPics        []ClientPic       `json:"client_pics"`
	ClientWarehouses  []ClientWarehouse `json:"client_warehouses"`
}

type ClientWarehousePic struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	OrderNumber  int64         `json:"order_number"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type ClientWarehouse struct {
	Id         int64 `json:"id"`
	ProvinceId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"province_id"`
	CityId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"city_id"`
	DistrictId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"district_id"`
	Zipcode             int64                `json:"zipcode"`
	Name                string               `json:"name"`
	StreetAddress       string               `json:"street_address"`
	OrderNumber         int64                `json:"order_number"`
	ClientWarehousePics []ClientWarehousePic `json:"client_warehouse_pics"`
}

type ClientPic struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	OrderNumber  int64         `json:"order_number"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	Email        string        `json:"email"`
}

type ClientInformation struct {
	Id         int64 `json:"id"`
	ProvinceId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"province_id"`
	CityId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"city_id"`
	DistrictId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"district_id"`
	Zipcode       int64         `json:"zipcode"`
	CompanyName   string        `json:"company_name"`
	Email         string        `json:"email"`
	NPWP          string        `json:"npwp"`
	StreetAddress string        `json:"street_address"`
	PhoneNumbers  []PhoneNumber `json:"phone_numbers"`
}

type GetMstConsigneeDetailResponse struct {
	ConsigneeInformation ConsigneeInformation `json:"consignee_information"`
	ConsigneePics        []ConsigneePic       `json:"consignee_pics"`
	ConsigneeWarehouses  []ConsigneeWarehouse `json:"consignee_warehouses"`
}

type ConsigneeWarehousePic struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	OrderNumber  int64         `json:"order_number"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type ConsigneeWarehouse struct {
	Id         int64 `json:"id"`
	ProvinceId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"province_id"`
	CityId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"city_id"`
	DistrictId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"district_id"`
	Zipcode                int64                   `json:"zipcode"`
	Name                   string                  `json:"name"`
	StreetAddress          string                  `json:"street_address"`
	OrderNumber            int64                   `json:"order_number"`
	ConsigneeWarehousePics []ConsigneeWarehousePic `json:"consignee_warehouse_pics"`
}

type ConsigneePic struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	OrderNumber  int64         `json:"order_number"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	Email        string        `json:"email"`
}

type ConsigneeInformation struct {
	Id         int64 `json:"id"`
	ProvinceId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"province_id"`
	CityId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"city_id"`
	DistrictId struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"district_id"`
	Zipcode       int64         `json:"zipcode"`
	CompanyName   string        `json:"company_name"`
	Email         string        `json:"email"`
	NPWP          string        `json:"npwp"`
	StreetAddress string        `json:"street_address"`
	Notes         string        `json:"notes"`
	PhoneNumbers  []PhoneNumber `json:"phone_numbers"`
}

type GetMstVendorDetailResponse struct {
	VendorInformation                   VendorInformation                    `json:"vendor_information"`
	VendorPics                          []VendorPic                          `json:"vendor_pics"`
	VendorContainerLaborEquipmentPrices []VendorContainerLaborEquipmentPrice `json:"vendor_container_labor_equipment_prices"`
}

type GetMstVendorsDropdownResponse struct {
	Vendors []CompanyDropdown `json:"vendors"`
}

type GetClientsDropdownResponse struct {
	Clients []CompanyDropdown `json:"clients"`
}

type GetClientsResponse struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	Id          int64  `json:"id"`
	CompanyName string `json:"company_name"`
	CreatedAt   string `json:"created_at"`
}

type GetMstConsigneesResponse struct {
	Consignees []Consignee `json:"consignees"`
}

type Consignee struct {
	Id             int64  `json:"id"`
	CompanyName    string `json:"company_name"`
	PicName        string `json:"pic_name"`
	PicPhoneNumber string `json:"pic_phone_number"`
	Location       string `json:"location"`
	CreatedAt      string `json:"created_at"`
}

type GetMstConsigneesDropdownResponse struct {
	Consignees []CompanyDropdown `json:"consignees"`
}

type CompanyDropdown struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetClientWarehousesDropdownResponse struct {
	Warehouses []WarehouseDropdown `json:"warehouses"`
}

type GetMstConsigneeWarehousesDropdownResponse struct {
	Warehouses []WarehouseDropdown `json:"warehouses"`
}

type GetMstVendorWarehousesDropdownResponse struct {
	Warehouses []WarehouseDropdown `json:"warehouses"`
}

type WarehouseDropdown struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetMstContainersDropdownResponse struct {
	Containers []ContainerDropdown `json:"containers"`
}

type ContainerDropdown struct {
	Id   int64  `json:"id"`
	Size string `json:"size"`
}

type GetMstHarboursDropdownResponse struct {
	Harbours []HarbourDropdown `json:"harbours"`
}

type HarbourDropdown struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetMstProvincesDropdownResponse struct {
	Provinces []AreaDropdown `json:"provinces"`
}

type GetMstCitiesDropdownResponse struct {
	Cities []AreaDropdown `json:"cities"`
}

type GetMstDistrictsDropdownResponse struct {
	Districts []AreaDropdown `json:"districts"`
}

type GetMstZipcodesDropdownResponse struct {
	Zipcodes []AreaDropdown `json:"zipcodes"`
}

type AreaDropdown struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetMstContainersResponse struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	Id          int64  `json:"id"`
	Size        string `json:"size"`
	Description string `json:"description"`
}

type GetMstContainerDetailResponse struct {
	Id          int64  `json:"id"`
	Size        string `json:"string"`
	Description string `json:"description"`
}

type GetMstBanksDropdownResponse struct {
	Banks []Bank `json:"banks"`
}

type GetMstBanksResponse struct {
	Banks []Bank `json:"banks"`
}

type GetMstBankDetailResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AccountNo   string `json:"account_no"`
	Description string `json:"description"`
}

type Bank struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AccountNo   string `json:"account_no"`
	Description string `json:"description"`
}
