package trxRepository

import (
	"consumerskripsi/repositories"
)

const coloumnTagihan = `no_header, device_id, member_code, member_name, member_type, 
card_number_uuid, card_number, type_card, beginning_balance, grand_total, 
product_id, product_code, product_name, ou_id, check_in_time, 
check_out_time, duration_time, price, base_time, progressive_time, 
progressive_price, is_pct, progressive_pct, max_price, is_24h,
overnight_time, overnight_price, grace_period, vehicle_number_in, vehicle_number_out, 
log_trans, qr_text, merchant_key, check_in_datetime, check_out_datetime,
created_at`

type trxRepository struct {
	RepoDB repositories.Repository
}

func NewTrxRepository(repoDB repositories.Repository) trxRepository {
	return trxRepository{
		RepoDB: repoDB,
	}
}

//
//func (ctx trxRepository) FindTagihanByNoHeader(noHeader, kodeProduk string) (models.Tagihan, error) {
//	var result models.Tagihan
//
//	query := `SELECT` + coloumnTagihan + `FROM tagihan
//	WHERE no_header = $1 AND product_code = $2`
//
//	err := ctx.RepoDB.DB.QueryRow(query, noHeader, kodeProduk).Scan(&result.DeviceId, &result.MemberCode, &result.MemberName, &result.MemberType,
//		&result.CardNumberUUID, &result.CardNumber, &result.TypeCard, &result.BeginningBalance, &result.GracePeriod,
//		&result.ProductId, &result.ProductCode, &result.ProductName, &result.OuId, &result.CheckInTime,
//		&result.CheckOutTime, &result.DurationTime, &result.Price, &result.BaseTime, &result.ProgressiveTime,
//		&result.ProgressivePrice, &result.IsPct, &result.ProgressivePct, &result.MaxPrice, &result.Is24H,
//		&result.OvernightTime, &result.OvernightPrice, &result.GracePeriod, &result.VehicleNumberIn, &result.VehicleNumberOut,
//		&result.LogTrans, &result.QrText, &result.MerchantKey, &result.CheckInDatetime, &result.CheckOutDatetime,
//		&result.Created_at)
//	if err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
