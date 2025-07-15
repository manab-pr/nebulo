package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/device/data/datasource"
	"github.com/manab-pr/nebulo/modules/device/data/sql/model"
	"github.com/manab-pr/nebulo/modules/device/domain/entities"
	repoiface "github.com/manab-pr/nebulo/modules/device/domain/repository"
)

// Ensure interface compliance
var _ repoiface.DeviceRepository = (*DeviceRepositoryGorm)(nil)

// DeviceRepositoryGorm implements the DeviceRepository interface using GORM
type DeviceRepositoryGorm struct {
	ds *datasource.DeviceGormDatasource
}

// NewDeviceRepositoryGorm creates a new instance of DeviceRepositoryGorm
func NewDeviceRepositoryGorm(ds *datasource.DeviceGormDatasource) *DeviceRepositoryGorm {
	return &DeviceRepositoryGorm{ds: ds}
}

// RegisterDevice inserts a new device using the datasource
func (r *DeviceRepositoryGorm) RegisterDevice(ctx context.Context, device *entities.Devices) error {
	modelDevice := model.FromEntity(device)
	return r.ds.Insert(ctx, modelDevice)
}

// GetDeviceByID fetches a device by ID
func (r *DeviceRepositoryGorm) GetDeviceByID(ctx context.Context, id uint64) (*entities.Devices, error) {
	modelDevice, err := r.ds.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if modelDevice == nil {
		return nil, nil
	}
	return modelDevice.ToEntity(), nil
}

// GetAllDevices fetches all devices matching filters
func (r *DeviceRepositoryGorm) GetAllDevices(ctx context.Context, filters map[string]interface{}) ([]entities.Devices, error) {
	modelDevices, err := r.ds.FindAll(ctx, filters)
	if err != nil {
		return nil, err
	}
	devices := make([]entities.Devices, len(modelDevices))
	for i, m := range modelDevices {
		devices[i] = *m.ToEntity()
	}
	return devices, nil
}

// UpdateDevice updates a device by ID
func (r *DeviceRepositoryGorm) UpdateDevice(ctx context.Context, id uint64, updates map[string]interface{}) error {
	return r.ds.Update(ctx, id, updates)
}

// DeleteDevice permanently deletes a device by ID
func (r *DeviceRepositoryGorm) DeleteDevice(ctx context.Context, id uint64) error {
	return r.ds.Delete(ctx, id)
}
