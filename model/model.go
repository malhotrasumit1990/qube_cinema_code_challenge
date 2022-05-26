package model

type MyPartnerInterface interface {
}

type Partner_data struct {
	Theatre_ID   string `csv:"Theatre"`
	Content_Size string `csv:"Size Slab (in GB)"`
	Min_Cost     int    `csv:"Minimum cost"`
	Cost_PerGB   int    `csv:"Cost Per GB"`
	Partner_ID   string `csv:"Partner ID"`
}

type Delivery_Data struct {
	Delivery_ID  string
	Content_Size int
	Theatre_ID   string
}

type Delivery_Result struct {
	Delivery_ID string
	Possiblity  bool
	Partner_ID  string
	Cost        string
}
