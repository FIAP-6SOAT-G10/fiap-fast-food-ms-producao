package dto

import "fiap-fast-food-ms-producao/domain/models"

type ProductionOrderDTO struct {
	ID              string `json:"id"`
	Status          string `json:"status" binding:"required"`
	ExternalOrderId string `json:"externalOrderId"`
}

func ToProductionOrderDTO(productionOrderModel *models.ProductionOrder) *ProductionOrderDTO {
	poDTO := ProductionOrderDTO{
		ID:              productionOrderModel.ID.Hex(),
		Status:          models.StatusFromInt(productionOrderModel.Status),
		ExternalOrderId: productionOrderModel.ExternalId,
	}
	return &poDTO
}
