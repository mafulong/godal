package model

import "time"

//TestTb2 corresponds to test_tb2
type TestTb2 struct {
	TestId         int       `gorm:"Column:test_id" json:"test_id"`
	TestTitle      string    `gorm:"Column:test_title" json:"test_title"`
	TestAuthor     string    `gorm:"Column:test_author" json:"test_author"`
	SubmissionDate time.Time `gorm:"Column:submission_date" json:"submission_date"`
}

//TableName func
func (TestTb2) TableName() string {
	return "test_tb2"
}
