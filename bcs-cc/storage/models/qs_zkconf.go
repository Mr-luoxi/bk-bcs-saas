// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ZookeeperConfigQuerySet

// ZookeeperConfigQuerySet is an queryset type for ZookeeperConfig
type ZookeeperConfigQuerySet struct {
	db *gorm.DB
}

// NewZookeeperConfigQuerySet constructs new ZookeeperConfigQuerySet
func NewZookeeperConfigQuerySet(db *gorm.DB) ZookeeperConfigQuerySet {
	return ZookeeperConfigQuerySet{
		db: db.Model(&ZookeeperConfig{}),
	}
}

func (qs ZookeeperConfigQuerySet) w(db *gorm.DB) ZookeeperConfigQuerySet {
	return NewZookeeperConfigQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *ZookeeperConfig) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *ZookeeperConfig) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) All(ret *[]ZookeeperConfig) error {
	return qs.db.Find(ret).Error
}

// BCSZookeeperEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) BCSZookeeperEq(bCSZookeeper string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("bcs_zookeeper = ?", bCSZookeeper))
}

// BCSZookeeperIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) BCSZookeeperIn(bCSZookeeper ...string) ZookeeperConfigQuerySet {
	if len(bCSZookeeper) == 0 {
		qs.db.AddError(errors.New("must at least pass one bCSZookeeper in BCSZookeeperIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bcs_zookeeper IN (?)", bCSZookeeper))
}

// BCSZookeeperNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) BCSZookeeperNe(bCSZookeeper string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("bcs_zookeeper != ?", bCSZookeeper))
}

// BCSZookeeperNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) BCSZookeeperNotIn(bCSZookeeper ...string) ZookeeperConfigQuerySet {
	if len(bCSZookeeper) == 0 {
		qs.db.AddError(errors.New("must at least pass one bCSZookeeper in BCSZookeeperNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bcs_zookeeper NOT IN (?)", bCSZookeeper))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtEq(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtGt(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtGte(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtLt(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtLte(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatedAtNe(createdAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CreatorEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatorEq(creator string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("creator = ?", creator))
}

// CreatorIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatorIn(creator ...string) ZookeeperConfigQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator IN (?)", creator))
}

// CreatorNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatorNe(creator string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("creator != ?", creator))
}

// CreatorNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) CreatorNotIn(creator ...string) ZookeeperConfigQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator NOT IN (?)", creator))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) Delete() error {
	return qs.db.Delete(ZookeeperConfig{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(ZookeeperConfig{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(ZookeeperConfig{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtEq(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtGt(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtGte(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtIsNotNull() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtIsNull() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtLt(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtLte(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DeletedAtNe(deletedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DescEq(desc string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("desc = ?", desc))
}

// DescIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DescIn(desc ...string) ZookeeperConfigQuerySet {
	if len(desc) == 0 {
		qs.db.AddError(errors.New("must at least pass one desc in DescIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("desc IN (?)", desc))
}

// DescNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DescNe(desc string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("desc != ?", desc))
}

// DescNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) DescNotIn(desc ...string) ZookeeperConfigQuerySet {
	if len(desc) == 0 {
		qs.db.AddError(errors.New("must at least pass one desc in DescNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("desc NOT IN (?)", desc))
}

// EnvironmentEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) EnvironmentEq(environment string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("environment = ?", environment))
}

// EnvironmentIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) EnvironmentIn(environment ...string) ZookeeperConfigQuerySet {
	if len(environment) == 0 {
		qs.db.AddError(errors.New("must at least pass one environment in EnvironmentIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("environment IN (?)", environment))
}

// EnvironmentNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) EnvironmentNe(environment string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("environment != ?", environment))
}

// EnvironmentNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) EnvironmentNotIn(environment ...string) ZookeeperConfigQuerySet {
	if len(environment) == 0 {
		qs.db.AddError(errors.New("must at least pass one environment in EnvironmentNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("environment NOT IN (?)", environment))
}

// ExtraEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ExtraEq(extra string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("extra = ?", extra))
}

// ExtraIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ExtraIn(extra ...string) ZookeeperConfigQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra IN (?)", extra))
}

// ExtraNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ExtraNe(extra string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("extra != ?", extra))
}

// ExtraNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ExtraNotIn(extra ...string) ZookeeperConfigQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra NOT IN (?)", extra))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) GetUpdater() ZookeeperConfigUpdater {
	return NewZookeeperConfigUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDEq(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDGt(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDGte(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDIn(ID ...uint) ZookeeperConfigQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDLt(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDLte(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDNe(ID uint) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) IDNotIn(ID ...uint) ZookeeperConfigQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// KindEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindEq(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind = ?", kind))
}

// KindGt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindGt(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind > ?", kind))
}

// KindGte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindGte(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind >= ?", kind))
}

// KindIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindIn(kind ...uint8) ZookeeperConfigQuerySet {
	if len(kind) == 0 {
		qs.db.AddError(errors.New("must at least pass one kind in KindIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("kind IN (?)", kind))
}

// KindLt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindLt(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind < ?", kind))
}

// KindLte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindLte(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind <= ?", kind))
}

// KindNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindNe(kind uint8) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("kind != ?", kind))
}

// KindNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) KindNotIn(kind ...uint8) ZookeeperConfigQuerySet {
	if len(kind) == 0 {
		qs.db.AddError(errors.New("must at least pass one kind in KindNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("kind NOT IN (?)", kind))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) Limit(limit int) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) NameEq(name string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) NameIn(name ...string) ZookeeperConfigQuerySet {
	if len(name) == 0 {
		qs.db.AddError(errors.New("must at least pass one name in NameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("name IN (?)", name))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) NameNe(name string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) NameNotIn(name ...string) ZookeeperConfigQuerySet {
	if len(name) == 0 {
		qs.db.AddError(errors.New("must at least pass one name in NameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", name))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) Offset(offset int) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ZookeeperConfigQuerySet) One(ret *ZookeeperConfig) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderAscByCreatedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderAscByDeletedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderAscByID() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByKind is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderAscByKind() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("kind ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderAscByUpdatedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderDescByCreatedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderDescByDeletedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderDescByID() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByKind is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderDescByKind() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("kind DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) OrderDescByUpdatedAt() ZookeeperConfigQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtEq(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtGt(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtGte(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtLt(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtLte(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatedAtNe(updatedAt time.Time) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UpdatorEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatorEq(updator string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updator = ?", updator))
}

// UpdatorIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatorIn(updator ...string) ZookeeperConfigQuerySet {
	if len(updator) == 0 {
		qs.db.AddError(errors.New("must at least pass one updator in UpdatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updator IN (?)", updator))
}

// UpdatorNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatorNe(updator string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("updator != ?", updator))
}

// UpdatorNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) UpdatorNotIn(updator ...string) ZookeeperConfigQuerySet {
	if len(updator) == 0 {
		qs.db.AddError(errors.New("must at least pass one updator in UpdatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updator NOT IN (?)", updator))
}

// ZookeeperEq is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ZookeeperEq(zookeeper string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("zookeeper = ?", zookeeper))
}

// ZookeeperIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ZookeeperIn(zookeeper ...string) ZookeeperConfigQuerySet {
	if len(zookeeper) == 0 {
		qs.db.AddError(errors.New("must at least pass one zookeeper in ZookeeperIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("zookeeper IN (?)", zookeeper))
}

// ZookeeperNe is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ZookeeperNe(zookeeper string) ZookeeperConfigQuerySet {
	return qs.w(qs.db.Where("zookeeper != ?", zookeeper))
}

// ZookeeperNotIn is an autogenerated method
// nolint: dupl
func (qs ZookeeperConfigQuerySet) ZookeeperNotIn(zookeeper ...string) ZookeeperConfigQuerySet {
	if len(zookeeper) == 0 {
		qs.db.AddError(errors.New("must at least pass one zookeeper in ZookeeperNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("zookeeper NOT IN (?)", zookeeper))
}

// SetBCSZookeeper is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetBCSZookeeper(bCSZookeeper string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.BCSZookeeper)] = bCSZookeeper
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetCreatedAt(createdAt time.Time) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCreator is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetCreator(creator string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Creator)] = creator
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetDeletedAt(deletedAt *time.Time) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetDesc is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetDesc(desc string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Desc)] = desc
	return u
}

// SetEnvironment is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetEnvironment(environment string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Environment)] = environment
	return u
}

// SetExtra is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetExtra(extra string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Extra)] = extra
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetID(ID uint) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.ID)] = ID
	return u
}

// SetKind is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetKind(kind uint8) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Kind)] = kind
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetName(name string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetUpdatedAt(updatedAt time.Time) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUpdator is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetUpdator(updator string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Updator)] = updator
	return u
}

// SetZookeeper is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) SetZookeeper(zookeeper string) ZookeeperConfigUpdater {
	u.fields[string(ZookeeperConfigDBSchema.Zookeeper)] = zookeeper
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ZookeeperConfigUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ZookeeperConfigQuerySet

// ===== BEGIN of ZookeeperConfig modifiers

// ZookeeperConfigDBSchemaField describes database schema field. It requires for method 'Update'
type ZookeeperConfigDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ZookeeperConfigDBSchemaField) String() string {
	return string(f)
}

// ZookeeperConfigDBSchema stores db field names of ZookeeperConfig
var ZookeeperConfigDBSchema = struct {
	ID           ZookeeperConfigDBSchemaField
	CreatedAt    ZookeeperConfigDBSchemaField
	UpdatedAt    ZookeeperConfigDBSchemaField
	DeletedAt    ZookeeperConfigDBSchemaField
	Extra        ZookeeperConfigDBSchemaField
	Name         ZookeeperConfigDBSchemaField
	Desc         ZookeeperConfigDBSchemaField
	Zookeeper    ZookeeperConfigDBSchemaField
	BCSZookeeper ZookeeperConfigDBSchemaField
	Environment  ZookeeperConfigDBSchemaField
	Creator      ZookeeperConfigDBSchemaField
	Updator      ZookeeperConfigDBSchemaField
	Kind         ZookeeperConfigDBSchemaField
}{

	ID:           ZookeeperConfigDBSchemaField("id"),
	CreatedAt:    ZookeeperConfigDBSchemaField("created_at"),
	UpdatedAt:    ZookeeperConfigDBSchemaField("updated_at"),
	DeletedAt:    ZookeeperConfigDBSchemaField("deleted_at"),
	Extra:        ZookeeperConfigDBSchemaField("extra"),
	Name:         ZookeeperConfigDBSchemaField("name"),
	Desc:         ZookeeperConfigDBSchemaField("desc"),
	Zookeeper:    ZookeeperConfigDBSchemaField("zookeeper"),
	BCSZookeeper: ZookeeperConfigDBSchemaField("bcs_zookeeper"),
	Environment:  ZookeeperConfigDBSchemaField("environment"),
	Creator:      ZookeeperConfigDBSchemaField("creator"),
	Updator:      ZookeeperConfigDBSchemaField("updator"),
	Kind:         ZookeeperConfigDBSchemaField("kind"),
}

// Update updates ZookeeperConfig fields by primary key
// nolint: dupl
func (o *ZookeeperConfig) Update(db *gorm.DB, fields ...ZookeeperConfigDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":            o.ID,
		"created_at":    o.CreatedAt,
		"updated_at":    o.UpdatedAt,
		"deleted_at":    o.DeletedAt,
		"extra":         o.Extra,
		"name":          o.Name,
		"desc":          o.Desc,
		"zookeeper":     o.Zookeeper,
		"bcs_zookeeper": o.BCSZookeeper,
		"environment":   o.Environment,
		"creator":       o.Creator,
		"updator":       o.Updator,
		"kind":          o.Kind,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update ZookeeperConfig %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ZookeeperConfigUpdater is an ZookeeperConfig updates manager
type ZookeeperConfigUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewZookeeperConfigUpdater creates new ZookeeperConfig updater
// nolint: dupl
func NewZookeeperConfigUpdater(db *gorm.DB) ZookeeperConfigUpdater {
	return ZookeeperConfigUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&ZookeeperConfig{}),
	}
}

// ===== END of ZookeeperConfig modifiers

// ===== END of all query sets