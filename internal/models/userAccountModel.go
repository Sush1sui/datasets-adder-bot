package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// NutritionalData is a map of nutrient name to value
type NutritionalData map[string]float64

type Meal struct {
    Name    string  `bson:"name" json:"name"`
    Calorie float64 `bson:"calorie" json:"calorie"`
}

type DietHistory struct {
    Date            time.Time         `bson:"date,omitempty" json:"date,omitempty"`
    NutritionalData []NutritionalData `bson:"nutritionalData,omitempty" json:"nutritionalData,omitempty"`
    Breakfast       []Meal            `bson:"breakfast,omitempty" json:"breakfast,omitempty"`
    Lunch           []Meal            `bson:"lunch,omitempty" json:"lunch,omitempty"`
    Dinner          []Meal            `bson:"dinner,omitempty" json:"dinner,omitempty"`
    OtherMealTime   []Meal            `bson:"otherMealTime,omitempty" json:"otherMealTime,omitempty"`
}

type LoggedWeight struct {
    Value float64 `bson:"value" json:"value"`
    Label string  `bson:"label" json:"label"`
}

type UserAccount struct {
    ID                bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    GmailID           *string            `bson:"gmailId,omitempty" json:"gmailId,omitempty"`
    ProfileLink       *string            `bson:"profileLink,omitempty" json:"profileLink,omitempty"`
    ProfilePublicID   *string            `bson:"profilePublicId,omitempty" json:"profilePublicId,omitempty"`
    Gender            *string            `bson:"gender,omitempty" json:"gender,omitempty"`
    BirthDate         *time.Time         `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
    HeightFeet        *float64           `bson:"heightFeet,omitempty" json:"heightFeet,omitempty"`           // in feet
    HeightInches      *float64           `bson:"heightInches,omitempty" json:"heightInches,omitempty"`       // in inches
    Weight            *float64           `bson:"weight,omitempty" json:"weight,omitempty"`           // in kg
    WeightGoal        *string            `bson:"weightGoal,omitempty" json:"weightGoal,omitempty"`   // e.g., "lose", "maintain", "gain"
    TargetWeight      *float64           `bson:"targetWeight,omitempty" json:"targetWeight,omitempty"` // in kg
    BMI               *float64           `bson:"bmi,omitempty" json:"bmi,omitempty"`
    Allergens         []string           `bson:"allergens,omitempty" json:"allergens,omitempty"`
    MedicalConditions []string           `bson:"medicalConditions,omitempty" json:"medicalConditions,omitempty"`
    DietHistory       []DietHistory      `bson:"dietHistory,omitempty" json:"dietHistory,omitempty"`
    LoggedWeights     []LoggedWeight     `bson:"loggedWeights,omitempty" json:"loggedWeights,omitempty"`
    Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
    FirstName         *string            `bson:"firstName,omitempty" json:"firstName,omitempty"`
    LastName          *string            `bson:"lastName,omitempty" json:"lastName,omitempty"`
    Email             *string            `bson:"email,omitempty" json:"email,omitempty"`
    Password          *string            `bson:"password,omitempty" json:"password,omitempty"`
    OTP               *string            `bson:"otp,omitempty" json:"otp,omitempty"`
    OTPExpires        *time.Time         `bson:"otpExpires,omitempty" json:"otpExpires,omitempty"`
    IsVerified        bool               `bson:"isVerified" json:"isVerified"`
}