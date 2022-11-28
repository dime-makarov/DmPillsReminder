package dataaccess

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

type Prescription struct {
	Id                     uint      `json:"id"`
	MedicineName           string    `json:"medicineName"`
	IsActive               bool      `json:"isActive"`
	TimesInPeriod          uint      `json:"timesInPeriod"`
	PeriodLengthInMinutes  uint      `json:"periodLengthInMinutes"`
	TotalDurationInMinutes uint      `json:"totalDurationInMinutes"`
	StartDate              time.Time `json:"startDate"`
	CountTaken             uint      `json:"countTaken"`
	CountLeft              uint      `json:"countLeft"`
}

func InitializeDB(dataSourceName string) error {
    var err error

	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		return err
	}

	return db.Ping()
}

func GetPrescriptions() ([]Prescription, error) {
	var prescriptions []Prescription

	rows, err := db.Query("select * from prescriptions")

	if err != nil {
		return nil, fmt.Errorf("GetPrescriptions: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var prescr Prescription
		var isActiveDb int
		var timeDb sql.NullTime

		if err := rows.Scan(
			&prescr.Id,
			&prescr.MedicineName,
			&isActiveDb,
			&prescr.TimesInPeriod,
			&prescr.PeriodLengthInMinutes,
			&prescr.TotalDurationInMinutes,
			&timeDb,
			&prescr.CountTaken,
			&prescr.CountLeft); err != nil {

			return nil, fmt.Errorf("GetPrescriptions: %v", err)
		}

		if isActiveDb == 1 {
			prescr.IsActive = true
		}

		if timeDb.Valid {
			prescr.StartDate = timeDb.Time
		}

		prescriptions = append(prescriptions, prescr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetPrescriptions: %v", err)
	}

	return prescriptions, nil
}

func GetPrescriptionById(id uint) (Prescription, error) {
	var prescr Prescription
	var isActiveDb int
	var timeDb sql.NullTime

	row := db.QueryRow("select * from prescriptions where Id = ?", id)

	if err := row.Scan(
		&prescr.Id,
		&prescr.MedicineName,
		&isActiveDb,
		&prescr.TimesInPeriod,
		&prescr.PeriodLengthInMinutes,
		&prescr.TotalDurationInMinutes,
		&timeDb,
		&prescr.CountTaken,
		&prescr.CountLeft); err != nil {

		if err == sql.ErrNoRows {
			return prescr, fmt.Errorf("GetPrescriptionById %d: No such prescription", id)
		}

		return prescr, fmt.Errorf("GetPrescriptionById: %v", err)
	}

	if isActiveDb == 1 {
		prescr.IsActive = true
	}

	if timeDb.Valid {
		prescr.StartDate = timeDb.Time
	}

	return prescr, nil
}

func AddPrescription(
	medicineName string,
	timesInPeriod uint,
	periodLengthInMinutes uint,
	totalDurationInMinutes uint) (uint, error) {

	countLeft := totalDurationInMinutes / periodLengthInMinutes * timesInPeriod

	result, err := db.Exec(`INSERT INTO prescriptions (
        Id,
        MedicineName,
        IsActive,
        TimesInPeriod,
        PeriodLengthInMinutes,
        TotalDurationInMinutes,
        StartDate,
        CountTaken,
        CountLeft)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		nil,
		medicineName,
		0,
		timesInPeriod,
		periodLengthInMinutes,
		totalDurationInMinutes,
		nil,
		0,
		countLeft)

	if err != nil {
		return 0, fmt.Errorf("AddPrescription: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("AddPrescription: %v", err)
	}

	return uint(id), nil
}
