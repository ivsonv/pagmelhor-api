package enums

type BenefitUsageStatus string

const (
	BenefitUsageStatusUsed      BenefitUsageStatus = "used"
	BenefitUsageStatusReserved  BenefitUsageStatus = "reserved"
	BenefitUsageStatusCancelled BenefitUsageStatus = "cancelled"
)
