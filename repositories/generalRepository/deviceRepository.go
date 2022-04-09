package generalRepository

import (
	"consumerskripsi/models"
	"consumerskripsi/repositories"
)

type deviceRepository struct {
	RepoDB repositories.Repository
}

func NewDeviceRepository(repoDB repositories.Repository) deviceRepository {
	return deviceRepository{
		RepoDB: repoDB,
	}
}

func (ctx deviceRepository) FindDeviceByDeviceId(deviceId string) (models.Device, error) {
	var result []models.Device
	query := `
	SELECT A.id, A.device_id, A.mid, A.tid, A.idcorporate, A.acquiringmid,
		A.acquiringtid, A.merchantkey, A.created_at, A.updated_at
	FROM device A
	WHERE  EXISTS(
		SELECT 1 
		FROM corporate B
		WHERE A.idcorporate = B.id AND A.device_id = $1
		AND A.deleted_at IS NULL AND B.deleted_at IS NULL)`

	data, err := ctx.RepoDB.DB.Query(query, deviceId)
	if err != nil {
		return result[0], err
	}
	defer data.Close()

	for data.Next() {
		var val models.Device
		err := data.Scan(&val.ID, &val.Device_ID, &val.MID, &val.TID, &val.IDCorporate,
			&val.AcquiringMID, &val.AcquiringTID, &val.MerchantKey, &val.Created_at, &val.Updated_at)
		if err != nil {
			return val, err
		}
		result = append(result, val)
	}
	return result[0], nil
}
