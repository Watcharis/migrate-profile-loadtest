package models

import "time"

type Profile struct {
	ID                                int     `json:"id,omitempty" redis:"id" gorm:"column:id"`
	Uuid                              string  `json:"uuid,omitempty" redis:"uuid" gorm:"column:uuid"`
	Cid                               string  `json:"cid,omitempty" redis:"cid" gorm:"column:cid"`
	Macid                             string  `json:"macid,omitempty" redis:"macid" gorm:"column:macid"`
	Firstname                         string  `json:"firstname,omitempty" redis:"firstname" gorm:"column:firstname"`
	Lastname                          string  `json:"lastname,omitempty" redis:"lastname" gorm:"column:lastname"`
	MobileNo                          string  `json:"mobile_no,omitempty" redis:"mobile_no" gorm:"column:mobile_no"`
	Email                             string  `json:"email,omitempty" redis:"email" gorm:"column:email"`
	Pin                               string  `json:"pin,omitempty" redis:"pin" gorm:"column:pin"`
	CasaAc                            *string `json:"casa_ac,omitempty" redis:"-" gorm:"column:casa_ac"`
	GwalletAc                         string  `json:"gwallet_ac,omitempty" redis:"gwallet_ac" gorm:"column:gwallet_ac"`
	PaoCifId                          string  `json:"pao_cif_id,omitempty" redis:"pao_cif_id" gorm:"column:pao_cif_id"`
	PaoPinAppId                       string  `json:"pao_pin_app_id,omitempty" redis:"pao_pin_app_id" gorm:"-"`
	PaoPinUuId                        string  `json:"pao_pin_uuid,omitempty" redis:"pao_pin_uuid" gorm:"-"`
	PaoReqConsent                     bool    `json:"pao_req_consent,omitempty" redis:"pao_req_consent" gorm:"column:pao_req_consent"`
	TungReqConsent                    bool    `json:"tung_req_consent,omitempty" redis:"tung_req_consent" gorm:"column:tung_req_consent"`
	NotificationFlg                   bool    `json:"notification_flg,omitempty" redis:"notification_flg"`
	AccessToken_                      string  `json:"access_token,omitempty" redis:"access_token" gorm:"-"`
	RefreshToken_                     string  `json:"refresh_token,omitempty" redis:"refresh_token" gorm:"-"`
	BioToken_                         string  `json:"bio_token,omitempty" redis:"bio_token" gorm:"-"`
	BlockFlgPin                       bool    `json:"block_flg_pin,omitempty" redis:"block_flg_pin" gorm:"column:block_flg_pin"`
	BlockFlgRetail                    bool    `json:"block_flg_retail,omitempty" redis:"block_flg_retail" gorm:"-"`
	BuyerFlg                          bool    `json:"buy_flg" redis:"buyer_flg" gorm:"column:buyer_flg"`
	RetailerFlg                       bool    `json:"retailer_flg" redis:"retailer_flg" gorm:"column:retailer_flg"`
	RetailerYrFlg                     bool    `json:"retailer_yr_flg" redis:"retailer_yr_flg" gorm:"column:retailer_yr_flg"`
	RetailerYrMerID                   string  `json:"retailer_yr_id,omitempty" redis:"retailer_yr_id" gorm:"column:retailer_yr_id"`
	RetailerYrMerchantOwnerName       string  `json:"retailer_yr_merchant_owner_name,omitempty" redis:"retailer_yr_merchant_owner_name" gorm:"column:retailer_yr_merchant_owner_name"`
	RetailerYrMerchantName            string  `json:"retailer_yr_merchant_name,omitempty" redis:"retailer_yr_merchant_name" gorm:"column:retailer_yr_merchant_name"`
	RetailerYrMerchantBranchId        string  `json:"retailer_yr_merchant_branch_id,omitempty" redis:"retailer_yr_merchant_branch_id" gorm:"column:retailer_yr_merchant_branch_id"`
	RetailerYrMainMobileNo            string  `json:"retailer_yr_main_mobile,omitempty" redis:"retailer_yr_main_mobile" gorm:"column:retailer_yr_main_mobile"`
	RetailerYrActiveStatus            bool    `json:"retailer_yr_active_status,omitempty" redis:"retailer_yr_active_status" gorm:"column:retailer_yr_active_status"`
	RetailerYrCasaAc                  string  `json:"retailer_yr_casa_ac,omitempty" redis:"retailer_yr_casa_ac" gorm:"column:retailer_yr_casa_ac"`
	RetailerYrLocationLati            string  `json:"retailer_yr_location_latitude,omitempty" redis:"retailer_yr_location_latitude" gorm:"column:retailer_yr_location_latitude"`
	RetailerYrLocationLong            string  `json:"retailer_yr_location_longitude,omitempty" redis:"retailer_yr_location_longitude" gorm:"column:retailer_yr_location_longitude"`
	RetailerYrMerchantCategory        string  `json:"retailer_yr_merchant_category" redis:"retailer_yr_merchant_category" gorm:"column:retailer_yr_merchant_category"`
	RetailerYrMerchantSubcat          string  `json:"retailer_yr_merchant_subcat" redis:"retailer_yr_merchant_subcat" gorm:"column:retailer_yr_merchant_subcat"`
	RetailerYrMerchantAddrNo          string  `json:"retailer_yr_merchant_addr_no,omitempty" redis:"RetailerYrMerchantAddrNo" gorm:"column:retailer_yr_merchant_addr_no"`
	RetailerYrMerchantAddrSubdistrict string  `json:"retailer_yr_merchant_addr_subdistrict,omitempty" redis:"RetailerYrMerchantAddrSubdistrict" gorm:"column:retailer_yr_merchant_addr_subdistrict"`
	RetailerYrMerchantAddrDistrict    string  `json:"retailer_yr_merchant_addr_district,omitempty" redis:"RetailerYrMerchantAddrDistrict" gorm:"column:retailer_yr_merchant_addr_district"`
	RetailerYrMerchantAddrProvince    string  `json:"retailer_yr_merchant_addr_province,omitempty" redis:"RetailerYrMerchantAddrProvince" gorm:"column:retailer_yr_merchant_addr_province"`
	RetailerYrMerchantAddrPostalcode  string  `json:"retailer_yr_merchant_addr_postalcode,omitempty" redis:"RetailerYrMerchantAddrPostalcode" gorm:"column:retailer_yr_merchant_addr_postalcode"`
	RetailerYrMerchantStatus          string  `json:"retailer_yr_merchant_status,omitempty" redis:"retailer_yr_merchant_status" gorm:"column:retailer_yr_merchant_status"`
	RetailerYrLastUpdate              string  `json:"retailer_yr_last_update_at,omitempty" redis:"retailer_yr_last_update_at" gorm:"-"`
	RetailerYrFcm                     string  `json:"retailer_yr_fcm_token,omitempty" redis:"retailer_yr_fcm_token" gorm:"column:retailer_yr_fcm_token"`
	RetailerYrCompCode                string  `json:"retailer_yr_company_code,omitempty" redis:"retailer_yr_company_code" gorm:"column:retailer_yr_company_code"`
	RetailerAllowVerifyFlg            bool    `json:"retailer_allow_verify_flg" redis:"retailer_allow_verify_flg" gorm:"-"`
	RetailerUpliftStatus              string  `json:"retailer_uplift_status" redis:"retailer_uplift_status" gorm:"-"`
	IsJuristicFlg                     bool    `json:"is_juristic" redis:"is_juristic" gorm:"column:is_juristic"`
	DisabilityType                    int64   `json:"disability_type" redis:"disability_type"`
	UserAutoClaimStatus               string  `json:"user_auto_claim_status" gorm:"-"`
	Tags                              *string `json:"tags,omitempty" redis:"-"`
	SubJuristicID                     uint64  `json:"sub_juristic_id" redis:"sub_juristic_id" gorm:"column:sub_juristic_id"`
	// mvp16 new column cdi token
	CdiToken string `json:"cdi_token" redis:"cdi_token" gorm:"column:cdi_token"`

	// Time should be last to prevent redis scan error
	UpdateAt time.Time `json:"update_at" redis:"-" gorm:"column:update_at"`
	CreateAt time.Time `json:"create_at" redis:"-" gorm:"column:create_at"`

	// --- Cannot Scan field because type not support in redis such as pointer or time
	CasaAcForRedis   string `json:"-" redis:"casa_ac" gorm:"-"`
	TagsForRedis     string `json:"-" redis:"Tags" gorm:"-"`
	CreateAtForRedis string `json:"-" redis:"update_at" gorm:"-"`
	UpdateAtForRedis string `json:"-" redis:"create_at" gorm:"-"`
	BlacklistFlg     bool   `json:"blacklist_flg,omitempty" redis:"blacklist_flg" gorm:"column:blacklist_flg"`
}
