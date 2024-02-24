package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
)

func (r *perqaraRepo) GetRolePermissions(ctx context.Context, req *payload.GetRolePermissionsPayload) ([]*model.RolePermission, error) {
	var result []*model.RolePermission

	q := r.db.Model(&result).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		if len(req.RoleIds) > 0 {
			q.Where("role_id IN ?", req.RoleIds)
		}
		if len(req.RolePermissionIds) > 0 {
			q.Where("id IN ?", req.RolePermissionIds)
		}
	}
	err := q.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
