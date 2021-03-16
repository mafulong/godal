package testdb

import "time"

type testTb1 struct {
	TestId         int       `gorm:"Column:test_id" json:"test_id"`
	TestTitle      string    `gorm:"Column:test_title" json:"test_title"`
	TestAuthor     string    `gorm:"Column:test_author" json:"test_author"`
	SubmissionDate time.Time `gorm:"Column:submission_date" json:"submission_date"`
}

func (testTb1) TableName() string {
	return "test_tb1"
}
