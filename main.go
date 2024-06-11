package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Member struct {
	ID          uint
	CompanyName string
	Status      string
	IndustryID  uint
	CountryID   uint
}

type FilterMemberInput struct {
	CompanyName *string
	Status      *string
	IndustryID  *uint
	CountryID   *uint
}

func FilteredMembers(db *gorm.DB, input FilterMemberInput) ([]Member, error) {
	var filteredMembers []Member

	query := db.Model(&Member{})

	if input.CompanyName != nil {
		query = query.Where("company_name = ?", *input.CompanyName)
	}
	if input.Status != nil {
		query = query.Where("status = ?", *input.Status)
	}
	if input.IndustryID != nil {
		query = query.Where("industry_id = ?", *input.IndustryID)
	}
	if input.CountryID != nil {
		query = query.Where("country_id = ?", *input.CountryID)
	}

	result := query.Find(&filteredMembers)
	if result.Error != nil {
		return nil, result.Error
	}
	return filteredMembers, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&Member{})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var member Member
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if result := db.Create(&member); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(member)
	})

	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		var input FilterMemberInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		members, err := FilteredMembers(db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(members)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
