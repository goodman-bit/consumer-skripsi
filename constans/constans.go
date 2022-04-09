package constans

const (
	SUCCESS_CODE = "01"
	FAILED_CODE  = "02"
	PENDING_CODE = "03"
	VOID_CODE    = "04"

	// Error berhubungan dengan data
	DATA_NOT_FOUND_CODE = "201"
	VALIDATE_ERROR_CODE = "202"
	GetJKey             = "JWTMKPmobile123456@"

	SYSTEM_ERROR_CODE    = "501"
	UNDEFINED_ERROR_CODE = "502"

	EMPTY_VALUE     = ""
	REDIS_CLIENT    = "REDIS_CLIENT"
	EMPTY_VALUE_INT = 0
	NULL_LONG_VALUE = -99

	STATUS_SUCCESS_CODE     = "01"
	STATUS_PENDING_CODE     = "03"
	STATUS_FAILED_CODE      = "02"
	STATUS_VOID_CODE        = "04"
	STATUS_ON_PROGRESS_CODE = "10"
	STATUS_SUCCESS_DESC     = "SUCCESS"
	STATUS_PENDING_DESC     = "PENDING"
	STATUS_FAILED_DESC      = "FAILED"
	STATUS_ON_PROGRESS_DESC = "PROCESSED"
	DATATYPE_TRX_ITEM_BATCH = "TRX_ITEM_BATCH"
	DATATYPE_TRX            = "TRX"
	OTHERS_PAYMENT          = "OTHERS"
	SYNC                    = "SYNC-LOCAL"

	TEMP_LOG_PREPAID_COLLECTION            = "temp_log_prepaid"
	BATCH_CALLBACK_URL                     = "batch_callback_url"
	TRX_ITEMS_PREPAID_COLLECTION           = "trx_items_prepaid"
	WHITELIST_ROUTE                        = "WhitelistRoute"
	TRX_ITEMS_PREPAID_FAIL_OVER_COLLECTION = "trx_items_prepaid_fail_over"

	DOCUMENT_PARKING = "PP"
	YES              = "Y"
	NO               = "N"

	PRODUCT_COLLECTION              = "product_col"
	QUEUE_APPS2PAY_USF              = "A2P-USF"
	QUEUE_TRX_PREPIAD_CALLBACK      = "MPRK_TRX_PREPAID_CALLBACK"
	QUEUE_PG_SCHEDULING_PROGRESSIVE = "pg.scheduling-progressive-mpos"

	TABLE_POLICY_OU_PRODUCT            = "PolicyOUProduct"
	TABLE_DEVICE_OU                    = "DeviceOU"
	TRX_COLLECTION                     = "Trx"
	TRX_OUTSTANDING_COLLECTION         = "TrxOutstanding"
	TRX_INVOICE_ITEM_COLLECTION        = "TrxInvoiceItem"
	TRX_INVOICE_DETAIL_ITEM_COLLECTION = "TrxInvoiceDetailItem"
	MILLISECOND                        = 60000
	PROGRESSIVE_TYPE                   = "PROGRESSIVE"
	MAX_PROGRESSIVE_TYPE               = "MAX_PROGRESSIVE"
	CALLBACK_PAYMENT                   = "CALLBACK_PAYMENT"
	INQUIRY                            = "INQUIRY"
	CONFIRM                            = "CONFIRM"
	CARDTYPE_FLAZZ                     = "01"
	CARDTYPE_PREPAID_MANDIRI           = "02"
	CARDTYPE_TAPCASH                   = "03"
	CARDTYPE_BRIZZI                    = "04"
	PAYMENT_METHOD_QRIS                = "QRIS"
	PG_PARKING_CHECKIN                 = "PG_PARKING_CHECKIN"
	QUEUE_PROGRESSIVE                  = "SCHEDULING_PROJECT"

	MEMBER_COLLECTIONS = "Member"

	STATUSTRX_SUCCESS = 1
	STATUSTRX_PENDING = 2
	QTY_TRX           = 1
)
