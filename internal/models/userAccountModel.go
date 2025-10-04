package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// NutritionItem for nutritionData.items
type NutritionItem struct {
    Name  string  `bson:"name,omitempty" json:"name,omitempty"`
    Value float64 `bson:"value,omitempty" json:"value,omitempty"`
    Unit  string  `bson:"unit,omitempty" json:"unit,omitempty"`
}

// NutritionData for ScanResultType.nutritionData
type NutritionData struct {
    Title string          `bson:"title,omitempty" json:"title,omitempty"`
    Items []NutritionItem `bson:"items,omitempty" json:"items,omitempty"`
}

// TriggeredAllergen for ScanResultType.triggeredAllergens
type TriggeredAllergen struct {
    Ingredient string `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
    Allergen   string `bson:"allergen,omitempty" json:"allergen,omitempty"`
}

// ScanResultType for meals in DietHistory
type ScanResultType struct {
    ID                 interface{}         `bson:"id,omitempty" json:"id,omitempty"`
    Name               string              `bson:"name,omitempty" json:"name,omitempty"`
    FoodName           string              `bson:"foodName,omitempty" json:"foodName,omitempty"`
    Brand              string              `bson:"brand,omitempty" json:"brand,omitempty"`
    ServingSize        string              `bson:"servingSize,omitempty" json:"servingSize,omitempty"`
    Ingredients        []string            `bson:"ingredients,omitempty" json:"ingredients,omitempty"`
    TriggeredAllergens []TriggeredAllergen `bson:"triggeredAllergens,omitempty" json:"triggeredAllergens,omitempty"`
    NutritionData      []NutritionData     `bson:"nutritionData,omitempty" json:"nutritionData,omitempty"`
    Source             string              `bson:"source,omitempty" json:"source,omitempty"`
    Quantity           float64             `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

// DietHistory struct
type DietHistory struct {
    Date          string           `bson:"date,omitempty" json:"date,omitempty"`
    Breakfast     []ScanResultType `bson:"breakfast,omitempty" json:"breakfast,omitempty"`
    Lunch         []ScanResultType `bson:"lunch,omitempty" json:"lunch,omitempty"`
    Dinner        []ScanResultType `bson:"dinner,omitempty" json:"dinner,omitempty"`
    OtherMealTime []ScanResultType `bson:"otherMealTime,omitempty" json:"otherMealTime,omitempty"`
}

// LoggedWeight struct
type LoggedWeight struct {
    Value float64 `bson:"value" json:"value"`
    Date  string  `bson:"date" json:"date"`
}

// DailyRecommendation struct
type DailyRecommendation struct {
    Calories float64 `bson:"calories" json:"calories"`
    Carbs    float64 `bson:"carbs" json:"carbs"`
    Protein  float64 `bson:"protein" json:"protein"`
    Fat      float64 `bson:"fat" json:"fat"`
}

// UserAccount struct
type UserAccount struct {
    ID                  bson.ObjectID         `bson:"_id,omitempty" json:"id,omitempty"`
    GmailID             *string               `bson:"gmailId,omitempty" json:"gmailId,omitempty"`
    ProfileLink         *string               `bson:"profileLink,omitempty" json:"profileLink,omitempty"`
    ProfilePublicID     *string               `bson:"profilePublicId,omitempty" json:"profilePublicId,omitempty"`
    Gender              *string               `bson:"gender,omitempty" json:"gender,omitempty"`
    BirthDate           *time.Time            `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
    HeightFeet          *float64              `bson:"heightFeet,omitempty" json:"heightFeet,omitempty"`
    HeightInches        *float64              `bson:"heightInches,omitempty" json:"heightInches,omitempty"`
    Weight              *float64              `bson:"weight,omitempty" json:"weight,omitempty"`
    WeightGoal          *string               `bson:"weightGoal,omitempty" json:"weightGoal,omitempty"`
    TargetWeight        *float64              `bson:"targetWeight,omitempty" json:"targetWeight,omitempty"`
    BMI                 *float64              `bson:"bmi,omitempty" json:"bmi,omitempty"`
    Allergens           []string              `bson:"allergens,omitempty" json:"allergens,omitempty"`
    MedicalConditions   []string              `bson:"medicalConditions,omitempty" json:"medicalConditions,omitempty"`
    DietHistory         []DietHistory         `bson:"dietHistory,omitempty" json:"dietHistory,omitempty"`
    LoggedWeights       []LoggedWeight        `bson:"loggedWeights,omitempty" json:"loggedWeights,omitempty"`
    Name                *string               `bson:"name,omitempty" json:"name,omitempty"`
    FirstName           *string               `bson:"firstName,omitempty" json:"firstName,omitempty"`
    LastName            *string               `bson:"lastName,omitempty" json:"lastName,omitempty"`
    Email               *string               `bson:"email,omitempty" json:"email,omitempty"`
    Password            *string               `bson:"password,omitempty" json:"password,omitempty"`
    OTP                 *string               `bson:"otp,omitempty" json:"otp,omitempty"`
    OTPExpires          *time.Time            `bson:"otpExpires,omitempty" json:"otpExpires,omitempty"`
    IsVerified          bool                  `bson:"isVerified" json:"isVerified"`
    LoginAttempts       int                   `bson:"loginAttempts,omitempty" json:"loginAttempts,omitempty"`
    LockUntil           *time.Time            `bson:"lockUntil,omitempty" json:"lockUntil,omitempty"`
    DietType            *string               `bson:"dietType,omitempty" json:"dietType,omitempty"`
    DailyRecommendation *DailyRecommendation  `bson:"dailyRecommendation,omitempty" json:"dailyRecommendation,omitempty"`
    ActivityLevel       *string               `bson:"activityLevel,omitempty" json:"activityLevel,omitempty"`
}