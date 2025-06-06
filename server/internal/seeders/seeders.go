package seeders

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"server/internal/models"
	"server/pkg/utils"
)

func SeedUsers(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	adminUser := models.User{
		Email:    "admin@fitness.com",
		Password: string(password),
		Role:     "admin",
		Fullname: "Fitness Booking Admin",
		Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=admin",
	}

	customerUsers := []models.User{
		{
			Email:    "customer1@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Helena Mourise",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=helena",
			Gender:   "female",
		},
		{
			Email:    "customer2@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "David Van der googs",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=david",
			Gender:   "male",
		},
		{
			Email:    "customer3@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Alexandre Joshephine",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=joshephine",
			Gender:   "female",
		},
		{
			Email:    "elena.morris@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Elena Morris",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=EM",
			Gender:   "female",
		},
		{
			Email:    "brandon.tan@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Brandon Tan",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
			Gender:   "male",
		},
		{
			Email:    "yuki.sato@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Yuki Sato",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=YS",
			Gender:   "female",
		},
		{
			Email:    "elena.morrisaga@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Elena Morris",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=EM",
			Gender:   "female",
		},
		{
			Email:    "elvis.presley@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "Elvis Presley",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
			Gender:   "male",
		},
		{
			Email:    "david.jovovich@fitness.com",
			Password: string(password),
			Role:     "customer",
			Fullname: "David Jovovich",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=BT",
			Gender:   "male",
		},
	}

	instructorUsers := []models.User{
		{
			Email:    "instructor1@fitness.com",
			Password: string(password),
			Role:     "instructor",
			Fullname: "Nurmalasari Permata",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=AB",
			Gender:   "female",
		},
		{
			Email:    "instructor2@fitness.com",
			Password: string(password),
			Role:     "instructor",
			Fullname: "Eisenberg Josephine",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=ZA",
			Gender:   "female",
		},
		{
			Email:    "instructor3@fitness.com",
			Password: string(password),
			Role:     "instructor",
			Fullname: "David Lee",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=GH",
			Gender:   "male",
		},
		{
			Email:    "instructor4@fitness.com",
			Password: string(password),
			Role:     "instructor",
			Fullname: "Ellena Morris",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=GH",
			Gender:   "female",
		},
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Println("Failed to seed admin:", err)
	}
	if err := db.Create(&customerUsers).Error; err != nil {
		log.Println("Failed to seed customers:", err)
	}
	if err := db.Create(&instructorUsers).Error; err != nil {
		log.Println("Failed to seed instructors:", err)
	}

	allUsers := []models.User{adminUser}
	allUsers = append(allUsers, customerUsers...)
	allUsers = append(allUsers, instructorUsers...)

	for _, user := range allUsers {
		generateNotificationSettingsForUser(db, user)
	}

	log.Println("User seeding completed with notification settings!")
}
func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)

	if count > 0 {
		log.Println("Categories already seeded, skipping...")
		return
	}

	categories := []models.Category{
		{ID: uuid.New(), Name: "Yoga"},
		{ID: uuid.New(), Name: "Pilates"},
		{ID: uuid.New(), Name: "Cardio"},
		{ID: uuid.New(), Name: "Strength Training"},
		{ID: uuid.New(), Name: "Martial Arts"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Printf("failed seeding categories: %v", err)
	} else {
		log.Println("Categories seeding completed!")
	}
}

func SeedSubcategories(db *gorm.DB) {
	var count int64
	db.Model(&models.Subcategory{}).Count(&count)

	if count > 0 {
		log.Println("Subcategories already seeded, skipping...")
		return
	}

	subcategoryData := map[string][]string{
		"Yoga":         {"Hatha Yoga", "Vinyasa Yoga"},
		"Pilates":      {"Mat Pilates", "Reformer Pilates"},
		"Cardio":       {"HIIT", "Zumba", "Aerobic Dance"},
		"Martial Arts": {"Boxing", "Muay Thai", "Kickboxing"},
	}

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Printf("failed to fetch categories: %v", err)
		return
	}

	categoryMap := make(map[string]uuid.UUID)
	for _, c := range categories {
		categoryMap[c.Name] = c.ID
	}

	var subcategories []models.Subcategory
	for categoryName, subList := range subcategoryData {
		categoryID, exists := categoryMap[categoryName]
		if !exists {
			log.Printf("category %s not found, skipping subcategories...", categoryName)
			continue
		}
		for _, sub := range subList {
			subcategories = append(subcategories, models.Subcategory{
				ID:         uuid.New(),
				Name:       sub,
				CategoryID: categoryID,
			})
		}
	}

	if err := db.Create(&subcategories).Error; err != nil {
		log.Printf("failed seeding subcategories: %v", err)
	} else {
		log.Println("Subcategories seeding completed!")
	}
}

func SeedTypes(db *gorm.DB) {
	var count int64
	db.Model(&models.Type{}).Count(&count)

	if count > 0 {
		log.Println("Types already seeded, skipping...")
		return
	}

	types := []models.Type{
		{ID: uuid.New(), Name: "Group Class"},
		{ID: uuid.New(), Name: "Private Class"},
		{ID: uuid.New(), Name: "Virtual Class"},
	}

	if err := db.Create(&types).Error; err != nil {
		log.Printf("failed seeding types: %v", err)
	} else {
		log.Println("Successfully seeded types!")
	}
}

func SeedLevels(db *gorm.DB) {
	var count int64
	db.Model(&models.Level{}).Count(&count)

	if count > 0 {
		log.Println("Levels already seeded, skipping...")
		return
	}

	levels := []models.Level{
		{ID: uuid.New(), Name: "Beginner"},
		{ID: uuid.New(), Name: "Advanced"},
		{ID: uuid.New(), Name: "Intermediate"},
	}

	if err := db.Create(&levels).Error; err != nil {
		log.Printf("failed seeding levels: %v", err)
	} else {
		log.Println("Successfully seeded levels!")
	}
}

func SeedLocations(db *gorm.DB) {
	var count int64
	db.Model(&models.Location{}).Count(&count)

	if count > 0 {
		log.Println("Locations already seeded, skipping...")
		return
	}

	locations := []models.Location{
		{
			Name:        "Sweat Up Studio A",
			Address:     "123 Fitness St, New York, NY",
			GeoLocation: "40.712776,-74.005974",
		},
		{
			Name:        "Sweat Up Studio B",
			Address:     "456 Gym Ave, Los Angeles, CA",
			GeoLocation: "34.052235,-118.243683",
		},
	}

	if err := db.Create(&locations).Error; err != nil {
		log.Printf("failed seeding locations: %v", err)
	} else {
		log.Println("Successfully seeded locations!")
	}
}

func SeedClasses(db *gorm.DB) {
	var count int64
	db.Model(&models.Class{}).Count(&count)
	if count > 0 {
		log.Println("Classes already seeded, skipping...")
		return
	}

	var categories []models.Category
	var subcategories []models.Subcategory
	var types []models.Type
	var levels []models.Level
	var locations []models.Location

	db.Find(&categories)
	db.Find(&subcategories)
	db.Find(&types)
	db.Find(&levels)
	db.Find(&locations)

	categoryMap := make(map[string]uuid.UUID)
	for _, c := range categories {
		categoryMap[c.Name] = c.ID
	}

	subcategoryMap := make(map[string]uuid.UUID)
	for _, s := range subcategories {
		subcategoryMap[s.Name] = s.ID
	}

	classes := []models.Class{
		{
			ID: uuid.New(), Title: "Hatha Yoga for Beginners",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879286/fitness_booking_app/fx8mm2yumhlpen2fckgf.webp", Duration: 60,
			Description:    "This class is designed for those who are new to yoga. With a focus on breathing techniques, gentle stretches, and foundational postures, participants are introduced to the calming and strengthening principles of yoga. The goal is to cultivate awareness of body alignment and mental stillness. The instructor will guide each movement at a steady pace, ensuring comfort and safety for all levels. This class is ideal for reducing stress, improving flexibility, and setting the stage for a long-term yoga practice.",
			AdditionalList: []string{"Hatha Yoga", "Yoga"},
			TypeID:         types[0].ID, LevelID: levels[0].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Yoga"], SubcategoryID: subcategoryMap["Hatha Yoga"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Vinyasa Flow Intermediate",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879490/fitness_booking_app/o5uyu6r3qtmmb2py0c6f.webp", Duration: 75,
			Description:    "Vinyasa Flow is a dynamic yoga style that links breath with movement. This intermediate-level class focuses on building strength and flexibility through a fluid sequence of poses. Expect smooth transitions, moderate intensity, and a creative flow that challenges both body and mind. Participants should have some yoga experience, as the class moves at a faster pace. It’s perfect for deepening your practice, enhancing balance, and developing endurance in a mindful way.",
			AdditionalList: []string{"Vinyasa Yoga", "Yoga"},
			TypeID:         types[0].ID, LevelID: levels[1].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Yoga"], SubcategoryID: subcategoryMap["Vinyasa Yoga"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Mat Pilates Core Challenge",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879424/fitness_booking_app/ckauvapfizz7fbyzrd9h.webp", Duration: 60,
			Description:    "This Mat Pilates class is centered on strengthening the core muscles using bodyweight exercises. With a series of precise and controlled movements, participants develop stability, coordination, and postural alignment, which are essential for everyday mobility and injury prevention. It’s suitable for all levels and serves as a great complement to other workouts. Over time, you'll notice improvements in core strength, spinal alignment, and total-body awareness.",
			AdditionalList: []string{"Mat Pilates", "Pilates"},
			TypeID:         types[0].ID, LevelID: levels[2].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Pilates"], SubcategoryID: subcategoryMap["Mat Pilates"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Reformer Pilates Sculpt",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879388/fitness_booking_app/ew3uvgaxgecyaxz0xmgv.webp", Duration: 60,
			Description:    "Using the Reformer machine, this class adds resistance training to the Pilates experience. Participants will engage in slow, deliberate movements that sculpt lean muscles and improve flexibility without putting stress on the joints. Each session targets core, glutes, legs, and arms while enhancing balance and coordination. Reformer Pilates is ideal for anyone seeking low-impact yet highly effective full-body conditioning",
			AdditionalList: []string{"Reformer Pilates", "Pilates"},
			TypeID:         types[1].ID, LevelID: levels[1].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Pilates"], SubcategoryID: subcategoryMap["Reformer Pilates"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "HIIT Total Body Burn",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879477/fitness_booking_app/uudvfiyquve8cd4fdexh.webp", Duration: 45,
			Description:    "This High-Intensity Interval Training (HIIT) class is designed to burn maximum calories in minimum time. It combines cardio drills and bodyweight strength exercises in timed intervals, making it both efficient and effective. You’ll improve cardiovascular fitness, metabolism, and muscular endurance. Suitable for all fitness levels, this class offers modifications for beginners and challenges for advanced participants.",
			AdditionalList: []string{"HIIT", "Cardio"},
			TypeID:         types[0].ID, LevelID: levels[2].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Cardio"], SubcategoryID: subcategoryMap["HIIT"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Zumba Dance Energy",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879399/fitness_booking_app/b9xvyghagg2jcyvixntj.webp", Duration: 60,
			Description:    "Zumba Dance Energy is a fun and high-energy workout that blends Latin rhythms with easy-to-follow dance routines. It’s a full-body cardio session that feels more like a dance party than a workout.Perfect for all levels, this class improves endurance, burns calories, and boosts your mood. No dance experience is required—just come ready to move and enjoy the beat!",
			AdditionalList: []string{"Zumba", "Cardio"},
			TypeID:         types[0].ID, LevelID: levels[2].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Cardio"], SubcategoryID: subcategoryMap["Zumba"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Boxing Fundamentals",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879364/fitness_booking_app/cfawvclumu5hjtgviini.webp", Duration: 60,
			Description:    "Learn the basics of boxing in this empowering, non-contact class. Participants will be introduced to proper stance, punches, footwork, and defensive techniques while improving coordination and reaction time. Boxing a full-body workout that enhances strength, agility, and mental focus. This class is beginner-friendly and also a great stress reliever.",
			AdditionalList: []string{"Boxing", "Martial Arts"},
			TypeID:         types[0].ID, LevelID: levels[0].ID, LocationID: locations[0].ID,
			CategoryID: categoryMap["Martial Arts"], SubcategoryID: subcategoryMap["Boxing"], IsActive: true,
		},
		{
			ID: uuid.New(), Title: "Muay Thai Conditioning",
			Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879461/fitness_booking_app/lbt5otnjqiujqcokugbm.webp", Duration: 75,
			Description:    "Muay Thai is a powerful martial art that combines striking techniques with intense physical conditioning. In this class, participants will practice punches, kicks, elbows, and knees in controlled combinations, along with bodyweight drills for endurance. hether you're training for fitness or technique, Muay Thai offers an excellent way to improve power, discipline, and cardiovascular performance in one dynamic session.",
			AdditionalList: []string{"Muay Thai", "Martial Arts"},
			TypeID:         types[0].ID, LevelID: levels[1].ID, LocationID: locations[1].ID,
			CategoryID: categoryMap["Martial Arts"], SubcategoryID: subcategoryMap["Muay Thai"], IsActive: true,
		},
	}

	if err := db.Create(&classes).Error; err != nil {
		log.Printf("failed seeding classes: %v", err)
	} else {
		log.Println("Successfully seeded 8 professional fitness classes!")
	}
}

func SeedClassGalleries(db *gorm.DB) {
	var count int64
	db.Model(&models.ClassGallery{}).Count(&count)
	if count > 0 {
		log.Println("ClassGalleries already seeded, skipping...")
		return
	}

	var classes []models.Class
	if err := db.Find(&classes).Error; err != nil {
		log.Println("Failed to fetch classes:", err)
		return
	}

	if len(classes) == 0 {
		log.Println("No classes found for gallery seeding.")
		return
	}

	// Mapping: class title → image URLs
	galleryMap := map[string][]string{
		"Hatha Yoga for Beginners": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880275/fitness_booking_app/zfuaa53iljyvzl9ux3wc.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880274/fitness_booking_app/uluaivkhpndaymslbkht.webp",
		},
		"Vinyasa Flow Intermediate": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879960/fitness_booking_app/ez14mm0svmysibda2rzj.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746879959/fitness_booking_app/ozzdnzx5zvncibeknfqb.webp",
		},
		"HIIT Total Body Burn": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880255/fitness_booking_app/zaar3efvfdctw3qed79m.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880254/fitness_booking_app/di0zpt8a9l3rjcykpwah.webp",
		},
		"Zumba Dance Energy": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880169/fitness_booking_app/vnnrrr1nv6kt3yoq56nz.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880169/fitness_booking_app/ocmjzmyyhrkvudcqbq2q.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880168/fitness_booking_app/tolouxu0lecscg832dwi.webp",
		},
		"Mat Pilates Core Challenge": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880141/fitness_booking_app/uwlxbomqugaliffncscs.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880140/fitness_booking_app/tozp2hvl25kllrmq1amp.webp",
		},
		"Reformer Pilates Sculpt": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880139/fitness_booking_app/si4hk7k5txauk4rju517.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880141/fitness_booking_app/uwlxbomqugaliffncscs.webp",
		},
		"Boxing Fundamentals": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880115/fitness_booking_app/p7e0pl96xn8s7tnwponl.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880114/fitness_booking_app/dqdjuest25qrjqzhplup.webp",
		},
		"Muay Thai Conditioning": {
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880228/fitness_booking_app/ckskvayptwsjdytj1gus.webp",
			"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1746880227/fitness_booking_app/v78tjjzijecqoswyrmxf.webp",
		},
	}

	var galleries []models.ClassGallery
	for _, class := range classes {
		images := galleryMap[class.Title]
		for _, url := range images {
			galleries = append(galleries, models.ClassGallery{
				ClassID: class.ID,
				URL:     url,
			})
		}
	}

	if err := db.Create(&galleries).Error; err != nil {
		log.Printf("failed seeding class galleries: %v", err)
	} else {
		log.Println("ClassGalleries seeding completed with Cloudinary URLs!")
	}
}

func SeedPackages(db *gorm.DB) {
	rand.Seed(time.Now().UTC().UnixNano())

	var count int64
	db.Model(&models.Package{}).Count(&count)
	if count > 0 {
		log.Println("Packages already seeded, skipping...")
		return
	}

	var classes []models.Class
	if err := db.Preload("Category").Find(&classes).Error; err != nil {
		log.Fatalf("Failed fetching classes: %v", err)
	}

	categoryClasses := map[string][]models.Class{}
	for _, class := range classes {
		categoryClasses[class.Category.Name] = append(categoryClasses[class.Category.Name], class)
	}

	packages := []models.Package{
		{
			Name:           "Yoga Wellness Trial",
			Description:    "Try 1 Yoga class to relieve stress and boost flexibility.",
			Price:          120000,
			Credit:         1,
			Discount:       20,
			Expired:        14,
			AdditionalList: []string{"Valid for 14 days after first booking."},
			Image:          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748043621/yoga-wellness_ueqi65.webp",
			IsActive:       true,
		},
		{
			Name:           "Pilates Core Pack (5x)",
			Description:    "Enjoy 5 Mat/Reformer Pilates sessions to build core strength and posture.",
			Price:          600000,
			Credit:         5,
			Discount:       10,
			Expired:        60,
			AdditionalList: []string{"Valid for 2 months", "Non-refundable"},
			Image:          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748043619/pilates-core_reeqdu.webp",
			IsActive:       true,
		},
		{
			Name:           "Cardio Burnout Pass (10x)",
			Description:    "Get your heart pumping with 10 sessions of HIIT, Zumba, and Aerobic workouts.",
			Price:          1000000,
			Credit:         10,
			Discount:       15,
			Expired:        120,
			AdditionalList: []string{"Valid for 4 months", "Non-refundable"},
			Image:          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748043620/cardio-burner_fi0nhe.webp",
			IsActive:       true,
		},
		{
			Name:           "Combat Starter (1x)",
			Description:    "Experience our Martial Arts class in one exciting session.",
			Price:          150000,
			Credit:         1,
			Discount:       0,
			Expired:        14,
			AdditionalList: []string{"Valid for 14 days after booking."},
			Image:          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748043620/combat-starter_wdrrxk.webp",
			IsActive:       true,
		},
		{
			Name:           "Warrior Pack (5x)",
			Description:    "Boost your skills with 5 sessions of Boxing, Muay Thai, or Kickboxing.",
			Price:          650000,
			Credit:         5,
			Discount:       5,
			Expired:        60,
			AdditionalList: []string{"Valid for 2 months", "No refunds after activation."},
			Image:          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748043619/warior-pack_dmpnoa.webp",
			IsActive:       true,
		},
	}

	if err := db.Create(&packages).Error; err != nil {
		log.Printf("Failed seeding packages: %v", err)
		return
	}
	log.Println("Successfully seeded packages!")

	packageClassMap := map[string]string{
		"Yoga Wellness Trial":       "Yoga",
		"Pilates Core Pack (5x)":    "Pilates",
		"Cardio Burnout Pass (10x)": "Cardio",
		"Combat Starter (1x)":       "Martial Arts",
		"Warrior Pack (5x)":         "Martial Arts",
	}

	for i, pkg := range packages {
		categoryName := packageClassMap[pkg.Name]
		classList := categoryClasses[categoryName]

		if len(classList) == 0 {
			log.Printf("⚠️ No class found for category: %s", categoryName)
			continue
		}

		n := min(2+rand.Intn(3), len(classList))
		rand.Shuffle(len(classList), func(i, j int) { classList[i], classList[j] = classList[j], classList[i] })
		selected := classList[:n]

		if err := db.Model(&packages[i]).Association("Classes").Replace(&selected); err != nil {
			log.Printf(" Failed associating package %s: %v", pkg.Name, err)
		} else {
			log.Printf("Associated %d classes to package %s", len(selected), pkg.Name)
		}
	}

	log.Println("Package-Class association per category completed!")
}

func SeedInstructors(db *gorm.DB) {
	var count int64
	db.Model(&models.Instructor{}).Count(&count)

	if count > 0 {
		log.Println("Instructors already seeded, skipping...")
		return
	}
	var instructorsUser []models.User
	if err := db.Where("role = ?", "instructor").Find(&instructorsUser).Error; err != nil {
		log.Println("Failed to fetch instructor users:", err)
		return
	}

	if len(instructorsUser) == 0 {
		log.Println("No instructor users found, skipping instructor seeding...")
		return
	}

	var instructors []models.Instructor
	for _, user := range instructorsUser {
		instructors = append(instructors, models.Instructor{
			ID:             uuid.New(),
			UserID:         user.ID,
			Experience:     3,
			Specialties:    "Yoga, Reformer Pilates",
			Rating:         5.0,
			TotalClass:     0,
			Certifications: "Certified Yoga Teacher, Certified Reformer Pilates Instructor",
		})
	}

	if err := db.Create(&instructors).Error; err != nil {
		log.Printf("failed seeding instructors: %v", err)
	} else {
		log.Println("Successfully seeded instructors!")
	}
}

func SeedPayments(db *gorm.DB) {
	var count int64
	db.Model(&models.Payment{}).Count(&count)
	if count > 0 {
		log.Println("Payments already seeded, skipping...")
		return
	}

	// Get sample users
	var customer1, customer2 models.User
	if err := db.Preload("Profile").Where("email = ?", "customer1@fitness.com").First(&customer1).Error; err != nil {
		log.Println("Failed to find customer1@fitness.com:", err)
		return
	}
	if err := db.Preload("Profile").Where("email = ?", "customer2@fitness.com").First(&customer2).Error; err != nil {
		log.Println("Failed to find customer2@fitness.com:", err)
		return
	}

	// Get package
	var pkg models.Package
	if err := db.First(&pkg).Error; err != nil {
		log.Println("Failed to find package:", err)
		return
	}

	now := time.Now().UTC()
	taxRate := utils.GetTaxRate()
	base := pkg.Price * (1 - pkg.Discount/100)
	tax := base * taxRate
	total := base + tax

	// Seed payments
	payments := []models.Payment{
		createPayment(customer1, pkg, base, tax, total, now.AddDate(0, 0, -3), "success"),
		createPayment(customer2, pkg, base, tax, total, now, "success"),
		createPayment(customer2, pkg, base, tax, total, now, "success"),
		createPayment(customer2, pkg, base, tax, total, now.AddDate(0, 0, -1), "success"),
		createPayment(customer2, pkg, base, tax, total, now.AddDate(0, 0, -2), "success"),
		createPayment(customer2, pkg, base, tax, total, now.AddDate(0, 0, -2), "failed"),
	}

	if err := db.Create(&payments).Error; err != nil {
		log.Printf("Failed seeding payments: %v", err)
	} else {
		log.Println("Seeded 6 payments: 1 for customer01, 5 for customer02")
	}
}

func createPayment(user models.User, pkg models.Package, base, tax, total float64, paidAt time.Time, status string) models.Payment {
	id := uuid.New()
	return models.Payment{
		ID:            id,
		UserID:        user.ID,
		Email:         user.Email,
		Fullname:      user.Fullname,
		PackageID:     pkg.ID,
		PackageName:   pkg.Name,
		PaymentMethod: "bank_transfer",
		BasePrice:     base,
		Tax:           tax,
		Total:         total,
		Status:        status,
		PaidAt:        paidAt,
		InvoiceNumber: utils.GenerateInvoiceNumber(id),
	}
}

func SeedUserPackages(db *gorm.DB) {
	var count int64
	db.Model(&models.UserPackage{}).Count(&count)
	if count > 0 {
		log.Println("UserPackages already seeded, skipping...")
		return
	}

	var customer1, customer2 models.User
	if err := db.Where("email = ?", "customer1@fitness.com").First(&customer1).Error; err != nil {
		log.Println("Failed to find customer1@fitness.com")
		return
	}
	if err := db.Where("email = ?", "customer2@fitness.com").First(&customer2).Error; err != nil {
		log.Println("Failed to find customer2@fitness.com")
		return
	}

	var pkg models.Package
	if err := db.First(&pkg).Error; err != nil {
		log.Println("Failed to find package")
		return
	}

	now := time.Now().UTC()
	threeDaysAgo := now.AddDate(0, 0, -3)
	expired := threeDaysAgo.AddDate(0, 0, 30)

	var userPackages []models.UserPackage

	userPackages = append(userPackages, models.UserPackage{
		ID:              uuid.New(),
		UserID:          customer1.ID,
		PackageID:       pkg.ID,
		PackageName:     pkg.Name,
		RemainingCredit: pkg.Credit,
		PurchasedAt:     threeDaysAgo,
		ExpiredAt:       &expired,
	})

	userPackages = append(userPackages, models.UserPackage{
		ID:              uuid.New(),
		UserID:          customer2.ID,
		PackageID:       pkg.ID,
		PackageName:     pkg.Name,
		RemainingCredit: pkg.Credit,
		PurchasedAt:     threeDaysAgo,
		ExpiredAt:       &expired,
	})

	if err := db.Create(&userPackages).Error; err != nil {
		log.Printf("Failed seeding user packages: %v", err)
	} else {
		log.Println("Seeded user packages successfully")
	}
}

func SeedClassSchedules(db *gorm.DB) {
	var user models.User
	if err := db.Where("email = ?", "customer1@fitness.com").First(&user).Error; err != nil {
		log.Println("User customer1@fitness.com not found")
		return
	}

	var class models.Class
	var instructor models.Instructor
	if err := db.Preload("Location").First(&class).Error; err != nil {
		log.Println("No class found")
		return
	}
	if err := db.Preload("User.Profile").First(&instructor).Error; err != nil {
		log.Println("No instructor found")
		return
	}

	now := time.Now().UTC().Truncate(time.Minute)

	threeDaysAgo := now.AddDate(0, 0, -3)

	zoomLink := "https://zoom.us/j/92613838319?pwd=cTlscGI5cGlTU2IwZVN1b0FuR2d2QT09"
	verificationCode := "123456"

	schedulePast1 := models.ClassSchedule{
		ClassID:          class.ID,
		ClassName:        class.Title,
		ClassImage:       class.Image,
		Location:         class.Location.Name,
		InstructorID:     instructor.ID,
		InstructorName:   instructor.User.Fullname,
		Date:             threeDaysAgo,
		StartHour:        9,
		StartMinute:      0,
		IsOpened:         true,
		Duration:         class.Duration,
		Capacity:         10,
		Booked:           1,
		ZoomLink:         &zoomLink,
		VerificationCode: &verificationCode,
		Color:            "#f59e0b",
	}
	db.Create(&schedulePast1)

	schedulePast2 := models.ClassSchedule{
		ClassID:          class.ID,
		ClassName:        class.Title,
		ClassImage:       class.Image,
		Location:         class.Location.Name,
		InstructorID:     instructor.ID,
		InstructorName:   instructor.User.Fullname,
		Date:             threeDaysAgo,
		StartHour:        13,
		StartMinute:      0,
		IsOpened:         true,
		Duration:         class.Duration,
		Capacity:         10,
		Booked:           1,
		ZoomLink:         &zoomLink,
		VerificationCode: &verificationCode,
		Color:            "#f97316",
	}
	db.Create(&schedulePast2)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowLocal := time.Now().In(loc)
	startTime := nowLocal.Add(1 * time.Hour)

	dateOnly := time.Date(
		nowLocal.Year(), nowLocal.Month(), nowLocal.Day(),
		0, 0, 0, 0, time.UTC,
	)

	scheduleToday := models.ClassSchedule{
		ID:             uuid.New(),
		ClassID:        class.ID,
		ClassName:      class.Title,
		ClassImage:     class.Image,
		InstructorID:   instructor.ID,
		InstructorName: instructor.User.Fullname,
		Location:       class.Location.Name,
		Date:           dateOnly,
		StartHour:      startTime.Hour(),
		StartMinute:    startTime.Minute(),
		Duration:       class.Duration,
		Capacity:       10,
		Booked:         1,
		Color:          "#10b981",
	}

	db.Create(&scheduleToday)

	CreateBookingWithAttendance(db, user, schedulePast1, "attended", true, true)
	CreateBookingWithAttendance(db, user, schedulePast2, "entered", true, false)
	CreateBookingWithAttendance(db, user, scheduleToday, "not-join", false, false)

	log.Println("Seeded ClassSchedules + Bookings + Attendance (past & today) safely.")
}

func CreateBookingWithAttendance(db *gorm.DB, user models.User, schedule models.ClassSchedule, status string, attended bool, reviewed bool) {
	startTime := time.Date(schedule.Date.Year(), schedule.Date.Month(), schedule.Date.Day(),
		schedule.StartHour, schedule.StartMinute, 0, 0, time.UTC)
	endTime := startTime.Add(time.Duration(schedule.Duration) * time.Minute)

	booking := models.Booking{
		UserID:          user.ID,
		ClassScheduleID: schedule.ID,
		Status:          "booked",
		CreatedAt:       startTime,
	}

	if err := db.Create(&booking).Error; err != nil {
		log.Printf("Failed to create booking for schedule %s: %v", schedule.ID.String(), err)
		return
	}

	var checkedAt *time.Time
	var verifiedAt *time.Time

	if attended {
		checkedAt = &startTime
	}
	if reviewed {
		verifiedAt = &endTime
	}

	attendance := models.Attendance{
		BookingID:  booking.ID,
		Status:     status,
		CheckedIn:  attended,
		CheckedOut: reviewed,
		IsReviewed: reviewed,
		CheckedAt:  checkedAt,
		VerifiedAt: verifiedAt,
	}

	if err := db.Create(&attendance).Error; err != nil {
		log.Printf("Failed to create attendance for booking %s: %v", booking.ID.String(), err)
		return
	}

	log.Printf("Attendance created: %s", attendance.ID.String())
}

func SeedReviews(db *gorm.DB) {
	var count int64
	db.Model(&models.Review{}).Count(&count)

	if count > 0 {
		log.Println("Reviews already seeded, skipping...")
		return
	}

	var user models.User
	if err := db.First(&user, "role = ?", "customer").Error; err != nil {
		log.Println("Failed to find customer user:", err)
		return
	}

	var classes []models.Class
	if err := db.Limit(3).Find(&classes).Error; err != nil {
		log.Println("Failed to fetch classes:", err)
		return
	}

	if len(classes) == 0 {
		log.Println("No classes found for review seeding.")
		return
	}

	var reviews []models.Review
	for i, class := range classes {
		review := models.Review{
			ID:      uuid.New(),
			UserID:  user.ID,
			ClassID: class.ID,
			Rating:  4 + (i % 2),
			Comment: "Great class experience!",
		}
		reviews = append(reviews, review)
	}

	if err := db.Create(&reviews).Error; err != nil {
		log.Printf("failed seeding reviews: %v", err)
	} else {
		log.Println("Reviews seeding completed!")
	}
}

func SeedNotificationTypes(db *gorm.DB) {
	defaultTypes := []models.NotificationType{
		{ID: uuid.New(), Code: "system_message", Title: "System Announcement", Category: "announcement", DefaultEnabled: false},
		{ID: uuid.New(), Code: "class_reminder", Title: "Class Reminder", Category: "reminder", DefaultEnabled: false},
		{ID: uuid.New(), Code: "promo_offer", Title: "New Promotion Available", Category: "promotion", DefaultEnabled: false},
	}
	for _, t := range defaultTypes {
		db.FirstOrCreate(&t, "code = ?", t.Code)
	}
}

func generateNotificationSettingsForUser(db *gorm.DB, user models.User) {
	var notifTypes []models.NotificationType
	if err := db.Find(&notifTypes).Error; err != nil {
		log.Printf("Failed to get notification types for user %s: %v", user.Email, err)
		return
	}

	for _, nt := range notifTypes {
		for _, channel := range []string{"email", "browser"} {
			setting := models.NotificationSetting{
				ID:                 uuid.New(),
				UserID:             user.ID,
				NotificationTypeID: nt.ID,
				Channel:            channel,
				Enabled:            nt.DefaultEnabled,
			}
			if err := db.Create(&setting).Error; err != nil {
				log.Printf("Failed to create notification setting for user %s: %v", user.Email, err)
			}
		}
	}
}

func SeedDummyNotifications(db *gorm.DB) {
	var user models.User
	if err := db.Where("email = ?", "customer1@fitness.com").First(&user).Error; err != nil {
		log.Println("customer1@fitness.com not found")
		return
	}

	notifications := []models.Notification{
		{
			ID:       uuid.New(),
			UserID:   user.ID,
			TypeCode: "class_reminder",
			Title:    "Upcoming Class Reminder",
			Message:  "Don't forget your class starts in 1 hour!",
			Channel:  "browser",
			IsRead:   false,
		},
		{
			ID:       uuid.New(),
			UserID:   user.ID,
			TypeCode: "promo_offer",
			Title:    "Special Promo Just for You",
			Message:  "Get 20% off your next class using code: FIT20",
			Channel:  "browser",
			IsRead:   false,
		},
	}

	if err := db.Create(&notifications).Error; err != nil {
		log.Printf("Failed to seed dummy notifications: %v", err)
	} else {
		log.Println("Dummy notifications for customer1@fitness.com seeded!")
	}
}
func SeedVouchers(db *gorm.DB) {
	var count int64
	db.Model(&models.Voucher{}).Count(&count)
	if count > 0 {
		log.Println("Vouchers already seeded, skipping...")
		return
	}

	now := time.Now().UTC()
	expired := now.AddDate(0, 1, 0)

	max1 := 30000.0
	max2 := 50000.0

	voucher1 := models.Voucher{
		ID:           uuid.New(),
		Code:         "FIT50",
		Description:  "Dapatkan diskon 50% hingga 30.000",
		DiscountType: "percentage",
		Discount:     50,
		MaxDiscount:  &max1,
		Quota:        10,
		IsReusable:   false,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	voucher2 := models.Voucher{
		ID:           uuid.New(),
		Code:         "HEALTHY100K",
		Description:  "Diskon langsung 100.000",
		DiscountType: "fixed",
		Discount:     100000,
		MaxDiscount:  &max2,
		Quota:        10,
		IsReusable:   true,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	if err := db.Create([]models.Voucher{voucher1, voucher2}).Error; err != nil {
		log.Printf("Failed to seed vouchers: %v", err)
		return
	}

	log.Println("Vouchers seeding completed!")

}
