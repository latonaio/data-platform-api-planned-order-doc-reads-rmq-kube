package requests

type ItemDoc struct {
	PlannedOrder             int    `json:"PlannedOrder"`
	PlannedOrderItem         int    `json:"PlannedOrderItem"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
