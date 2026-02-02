package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/dto"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/entity"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/port/in"
	"github.com/TranXuanPhong25/ecom/services/voucher-service/internal/core/port/out"
)

// VoucherService implements the VoucherServicePort interface
type VoucherService struct {
	repository out.VoucherRepositoryPort
}

// NewVoucherService creates a new instance of VoucherService
func NewVoucherService(repository out.VoucherRepositoryPort) in.VoucherServicePort {
	return &VoucherService{
		repository: repository,
	}
}

// CreateVoucher creates a new voucher with validation
func (s *VoucherService) CreateVoucher(ctx context.Context, request dto.CreateVoucherRequest) (*entity.Voucher, error) {
	// Validate voucher code format (alphanumeric, hyphen, underscore)
	if !isValidVoucherCode(request.Code) {
		return nil, errors.New("voucher code must contain only alphanumeric characters, hyphens, or underscores")
	}

	// Validate discount value based on type
	if request.DiscountType == entity.DiscountTypePercentage {
		if request.DiscountValue <= 0 || request.DiscountValue > 100 {
			return nil, errors.New("percentage discount must be between 0 and 100")
		}
	} else if request.DiscountType == entity.DiscountTypeFixedAmount {
		if request.DiscountValue <= 0 {
			return nil, errors.New("fixed amount discount must be greater than 0")
		}
	}

	// Validate expiration date (must be in the future)
	if request.ExpiresAt != nil && request.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("expiration date must be in the future")
	}

	// Check if voucher code already exists
	existing, err := s.repository.FindByCode(ctx, strings.ToUpper(request.Code))
	if err == nil && existing != nil {
		return nil, errors.New("voucher code already exists")
	}

	voucher := &entity.Voucher{
		Code:          strings.ToUpper(request.Code),
		DiscountType:  request.DiscountType,
		DiscountValue: request.DiscountValue,
		MinOrderValue: request.MinOrderValue,
		MaxUsage:      request.MaxUsage,
		UsedCount:     0,
		ExpiresAt:     request.ExpiresAt,
		IsActive:      true,
		Description:   request.Description,
	}

	if err := s.repository.Create(ctx, voucher); err != nil {
		return nil, fmt.Errorf("failed to create voucher: %w", err)
	}

	return voucher, nil
}

// GetVoucherByCode retrieves a voucher by its code
func (s *VoucherService) GetVoucherByCode(ctx context.Context, code string) (*entity.Voucher, error) {
	voucher, err := s.repository.FindByCode(ctx, strings.ToUpper(code))
	if err != nil {
		return nil, fmt.Errorf("voucher not found: %w", err)
	}
	return voucher, nil
}

// GetVoucherByID retrieves a voucher by its ID
func (s *VoucherService) GetVoucherByID(ctx context.Context, id uint) (*entity.Voucher, error) {
	voucher, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("voucher not found: %w", err)
	}
	return voucher, nil
}

// ListVouchers retrieves a list of vouchers with pagination
func (s *VoucherService) ListVouchers(ctx context.Context, request dto.ListVouchersRequest) ([]entity.Voucher, int64, error) {
	// Set default values
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Limit <= 0 {
		request.Limit = 20
	}

	offset := (request.Page - 1) * request.Limit

	vouchers, total, err := s.repository.FindAll(ctx, request.Limit, offset, request.Active)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list vouchers: %w", err)
	}

	return vouchers, total, nil
}

// UpdateVoucher updates an existing voucher
func (s *VoucherService) UpdateVoucher(ctx context.Context, id uint, request dto.UpdateVoucherRequest) (*entity.Voucher, error) {
	// Get existing voucher
	voucher, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("voucher not found: %w", err)
	}

	// Update fields if provided
	if request.DiscountValue != nil {
		// Validate discount value based on type
		if voucher.DiscountType == entity.DiscountTypePercentage {
			if *request.DiscountValue <= 0 || *request.DiscountValue > 100 {
				return nil, errors.New("percentage discount must be between 0 and 100")
			}
		} else if voucher.DiscountType == entity.DiscountTypeFixedAmount {
			if *request.DiscountValue <= 0 {
				return nil, errors.New("fixed amount discount must be greater than 0")
			}
		}
		voucher.DiscountValue = *request.DiscountValue
	}

	if request.MinOrderValue != nil {
		voucher.MinOrderValue = *request.MinOrderValue
	}

	if request.MaxUsage != nil {
		voucher.MaxUsage = *request.MaxUsage
	}

	if request.ExpiresAt != nil {
		if request.ExpiresAt.Before(time.Now()) {
			return nil, errors.New("expiration date must be in the future")
		}
		voucher.ExpiresAt = request.ExpiresAt
	}

	if request.IsActive != nil {
		voucher.IsActive = *request.IsActive
	}

	if request.Description != nil {
		voucher.Description = *request.Description
	}

	if err := s.repository.Update(ctx, voucher); err != nil {
		return nil, fmt.Errorf("failed to update voucher: %w", err)
	}

	return voucher, nil
}

// DeleteVoucher deletes a voucher by its ID
func (s *VoucherService) DeleteVoucher(ctx context.Context, id uint) error {
	// Check if voucher exists
	_, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("voucher not found: %w", err)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete voucher: %w", err)
	}

	return nil
}

// isValidVoucherCode checks if the voucher code contains only valid characters
func isValidVoucherCode(code string) bool {
	if len(code) == 0 {
		return false
	}
	for _, c := range code {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}
