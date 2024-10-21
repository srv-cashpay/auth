package seeder

import (
	"github.com/srv-cashpay/auth/configs"
	"github.com/srv-cashpay/auth/entity"
)

func Role() {
	db := configs.InitDB()

	var limits []entity.Role

	var limit = entity.Role{
		ID:   "956f2014-f8ab-41e2-88c1-0c3871524665",
		Role: "Superadmin",
	}

	limits = append(limits, limit)

	var limit2 = entity.Role{
		ID:   "956f2014-f8ab-41e2-88c1-0c3871524665",
		Role: "admin",
	}

	limits = append(limits, limit2)

	var limit3 = entity.Role{
		ID:   "956f2014-f8ab-41e2-88c1-0c3871524665",
		Role: "kasir",
	}

	limits = append(limits, limit3)

	var limit4 = entity.Role{
		ID:   "956f2014-f8ab-41e2-88c1-0e3871524665",
		Role: "gudang",
	}

	limits = append(limits, limit4)

	var limit5 = entity.Role{
		ID:   "956f2014-f8ab-41e2-88c1-0e3871524665",
		Role: "gudang",
	}

	limits = append(limits, limit5)

	if err := db.Create(&limits).Error; err != nil {
		return
	}
}

func RunSeeder() {
	Role()
}
