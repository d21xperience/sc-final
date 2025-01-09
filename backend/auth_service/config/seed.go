package config

// SeedProfiles seeds the profiles table with dummy data
// func SeedProfiles(db *gorm.DB, count int) error {
// 	// rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < count; i++ {

// 		// fmt.Fprintln(faker.Date())
// 		profile := model.User{
// 			Username: faker.Username(),
// 			Email:    faker.Email(),
// 			Password: faker.Password(),
// 			Nama:     faker.Name(),
// 			JK:       faker.Gender(),
// 			Phone:    faker.Phonenumber(),
// 			TptLahir: faker.Century(),
// 			// TglLahir:       tglLahir,
// 			// AlamatJalan:    faker.GetRealAddress(),
// 			KotaKab: faker.Century(),
// 			Prov:    faker.Century(),
// 			// KodePos:        faker.da,
// 			NamaAyah: faker.Name(),
// 			NamaIbu:  faker.Name(),
// 			// CreatedAt:      faker.Name(),
// 			// UpdatedAt:      faker.Name(),
// 			RoleID:         1,
// 			TenantID:       2,
// 			ProfilePicture: faker.URL(),

// 			// Name:      faker.Name(),
// 			// Email:     faker.Email(),
// 			// Age:       rand.Intn(47) + 18, // Age between 18 and 65
// 			// Location:  faker.Address(),
// 			// Bio:       faker.Sentence(),
// 			// CreatedAt: time.Now(),
// 		}

// 		if err := db.Create(&profile).Error; err != nil {
// 			return fmt.Errorf("failed to seed profile %d: %w", i+1, err)
// 		}
// 	}
// 	return nil
// }
