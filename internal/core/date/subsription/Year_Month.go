package core_subcription_date

import (
	"fmt"
	"time"

	core_errors "github.com/Vlad-6894/test_task/internal/core/errors"
)

type YearMonth struct {
	Year  int
	Month time.Month
}

func ParseStartDateFromJson(m map[string]any) (YearMonth, error) {
	date, ok := m["start_date"]
	if !ok {
		err := fmt.Errorf("Error get start_date: %w", core_errors.ErrNotFound)
		return YearMonth{}, err
	}

	yearMonth, err := parseDateFromJson(date)
	if err != nil {
		return YearMonth{}, fmt.Errorf("Error parce start date! %w", err)
	}
	return yearMonth, nil
}

func ParseFinishDateFromJson(m map[string]any) (*YearMonth, error) {
	date, ok := m["finish_date"]
	if !ok {
		return nil, nil
	}
	yearMonth, err := parseDateFromJson(date)
	if err != nil {
		return nil, fmt.Errorf("Error parce start date! %w", err)
	}
	return &yearMonth, nil
}

func parseDateFromJson(date any) (YearMonth, error) {

	dateString := fmt.Sprintf("%v", date)

	var (
		year  int
		month int
	)

	if _, err := fmt.Sscanf(dateString, "%d-%d", &year, &month); err != nil {
		return YearMonth{}, fmt.Errorf("Fail to parce start date: %w", core_errors.ErrInvalidArgument)
	}

	monthRight, err := parseMonth(month)
	if err != nil {
		return YearMonth{}, fmt.Errorf("Failed to parse dateStart from Json : %w", err)
	}

	yearMonth := YearMonth{
		Year:  year,
		Month: monthRight,
	}

	return yearMonth, nil
}

func parseMonth(month int) (time.Month, error) {
	if month < 1 || month > 12 {
		return time.January, fmt.Errorf("parse month error: %w", core_errors.ErrInvalidArgument)
	}

	return time.Month(month), nil
}

func GetStartDateFromModel(date time.Time) YearMonth {
	year := date.Year()
	yearMonht := date.Month()

	return YearMonth{
		Year:  year,
		Month: yearMonht,
	}
}

func GetFinishDateFromModel(date *time.Time) *YearMonth {
	year := date.Year()
	yearMonht := date.Month()

	return &YearMonth{
		Year:  year,
		Month: yearMonht,
	}
}
