package seeder

import (
	"github.com/srv-cashpay/auth/configs"
	"github.com/srv-cashpay/auth/entity"
)

func Role() {
	db := configs.InitDB()

	var limits []entity.Role

	var limit = entity.Role{
		ID:   "8gHwINv71XDy",
		Role: "God Cashpay",
	}

	limits = append(limits, limit)

	var limit2 = entity.Role{
		ID:   "e9Wl2JyVeBM_",
		Role: "Admin",
	}

	limits = append(limits, limit2)

	var limit3 = entity.Role{
		ID:   "JQn-Y=l=NvJ7",
		Role: "Kasir",
	}

	limits = append(limits, limit3)

	var limit4 = entity.Role{
		ID:   "vC8h4YlOcCHf",
		Role: "Gudang",
	}

	limits = append(limits, limit4)

	if err := db.Create(&limits).Error; err != nil {
		return
	}
}

func RunSeeder() {
	Role()
}
