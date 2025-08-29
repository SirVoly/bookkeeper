package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func promptString(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) == 0 {
				return fmt.Errorf("input required")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptInt(label string) (int64, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid number")
			}
			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(result, 10, 64)
}

func promptNullableString(label string) sql.NullString {
	prompt := promptui.Prompt{Label: label}
	val, _ := prompt.Run()
	if strings.TrimSpace(val) == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: val, Valid: true}
}

func promptNullableInt(label string) sql.NullInt64 {
	prompt := promptui.Prompt{Label: label}
    val, _ := prompt.Run()
    if strings.TrimSpace(val) == "" {
        return sql.NullInt64{Valid: false}
    }
    num, err := strconv.ParseInt(val, 10, 64)
    if err != nil {
        return sql.NullInt64{Valid: false}
    }
    return sql.NullInt64{Int64: num, Valid: true}
}
