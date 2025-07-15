package datasource

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/device/data/sql/model"
	"gorm.io/gorm"
)

type DeviceGormDatasource struct {
	db *gorm.DB
}

func NewDeviceGormDatasource(db *gorm.DB) *DeviceGormDatasource {
	return &DeviceGormDatasource{db: db}
}

//  Insert
func (ds *DeviceGormDatasource) Insert(ctx context.Context, device *model.Device) error {
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()
	device.LastSeen = time.Now()
	return ds.db.WithContext(ctx).Create(device).Error
}

//  Get by ID
func (ds *DeviceGormDatasource) FindByID(ctx context.Context, id uint64) (*model.Device, error) {
	var device model.Device
	err := ds.db.WithContext(ctx).Where("id = ? AND status != ?", id, "deleted").First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

//Find All with Filters
func (ds *DeviceGormDatasource) FindAll(ctx context.Context, filters map[string]interface{}) ([]model.Device, error) {
	var devices []model.Device
	query := ds.db.WithContext(ctx).Model(&model.Device{})

	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	err := query.Where("status != ?", "deleted").Find(&devices).Error
	return devices, err
}

// Generic Update by ID
func (ds *DeviceGormDatasource) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return ds.db.WithContext(ctx).
		Model(&model.Device{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (ds *DeviceGormDatasource) Delete(ctx context.Context, id uint64) error {
	return ds.db.WithContext(ctx).
		Delete(&model.Device{}, id).Error
}