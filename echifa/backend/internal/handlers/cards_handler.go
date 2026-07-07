package handlers

import "github.com/gofiber/fiber/v2"

type ChifaCard struct {
	CardNumber      string `json:"cardNumber"`
	NIN             string `json:"nin"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Gender          string `json:"gender"`
	InsuranceType   string `json:"insuranceType"`
	BeneficiaryType string `json:"beneficiaryType"`
	Status          string `json:"status"`
	Agency          string `json:"agency"`
	ExpirationDate  string `json:"expirationDate"`
	PhotoURL        string `json:"photoUrl"`
}

func Cards(c *fiber.Ctx) error {
	cards := []ChifaCard{
		{
			CardNumber:      "CHIFA-000000001",
			NIN:             "198765432109876",
			FirstName:       "Ahmed",
			LastName:        "Benali",
			DateOfBirth:     "1990-06-15",
			Gender:          "Male",
			InsuranceType:   "Employee",
			BeneficiaryType: "Principal",
			Status:          "Active",
			Agency:          "CNAS Alger Centre",
			ExpirationDate:  "2030-12-31",
			PhotoURL:        "https://placehold.co/120x150?text=Ahmed",
		},
		{
			CardNumber:      "CHIFA-000000002",
			NIN:             "299876543210987",
			FirstName:       "Sara",
			LastName:        "Khelifi",
			DateOfBirth:     "1995-02-10",
			Gender:          "Female",
			InsuranceType:   "Employee",
			BeneficiaryType: "Principal",
			Status:          "Active",
			Agency:          "CNAS Oran",
			ExpirationDate:  "2031-05-20",
			PhotoURL:        "https://placehold.co/120x150?text=Sara",
		},
		{
			CardNumber:      "CHIFA-000000003",
			NIN:             "187654321098765",
			FirstName:       "Mohamed",
			LastName:        "Bouzid",
			DateOfBirth:     "1987-11-08",
			Gender:          "Male",
			InsuranceType:   "Retired",
			BeneficiaryType: "Principal",
			Status:          "Active",
			Agency:          "CNAS Constantine",
			ExpirationDate:  "2029-09-15",
			PhotoURL:        "https://placehold.co/120x150?text=Mohamed",
		},
		{
			CardNumber:      "CHIFA-000000004",
			NIN:             "296543210987654",
			FirstName:       "Nadia",
			LastName:        "Mansouri",
			DateOfBirth:     "1996-08-19",
			Gender:          "Female",
			InsuranceType:   "Student",
			BeneficiaryType: "Dependent",
			Status:          "Suspended",
			Agency:          "CNAS Annaba",
			ExpirationDate:  "2028-12-31",
			PhotoURL:        "https://placehold.co/120x150?text=Nadia",
		},
		{
			CardNumber:      "CHIFA-000000005",
			NIN:             "181234567890123",
			FirstName:       "Yacine",
			LastName:        "Zebda",
			DateOfBirth:     "1998-03-22",
			Gender:          "Male",
			InsuranceType:   "Employee",
			BeneficiaryType: "Principal",
			Status:          "Active",
			Agency:          "CNAS Laghouat",
			ExpirationDate:  "2032-01-10",
			PhotoURL:        "https://placehold.co/120x150?text=Yacine",
		},
	}

	return c.JSON(fiber.Map{
		"success": true,
		"count":   len(cards),
		"data":    cards,
	})
}
