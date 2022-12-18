package outpatientSession

import (
	"HMS-16-BE/model"
	"database/sql"
)

type ConditionRepository interface {
	Create(condition model.Conditions) error
	GetAllByPatient(patientId string) ([]model.Conditions, error)
	GetById(id uint) (model.Conditions, error)
	GetBySchedule(scheduleId uint) (model.Conditions, error)
}

type conditionRepository struct {
	db *sql.DB
}

func NewConditionRepository(db *sql.DB) *conditionRepository {
	return &conditionRepository{db}
}

func (c *conditionRepository) Create(condition model.Conditions) error {
	query := `INSERT INTO conditions VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := c.db.Exec(query, condition.ID, condition.CreatedAt, condition.UpdatedAt, condition.DeletedAt,
		condition.NurseId, condition.Height, condition.Weight, condition.BloodPressure, condition.SugarAnalysis,
		condition.BodyTemperature, condition.HeartRate, condition.BreathRate, condition.Cholesterol, condition.Note,
		condition.ScheduleId, condition.Status)
	if err != nil {
		return err
	}
	return nil
}

func (c *conditionRepository) GetAllByPatient(patientId string) ([]model.Conditions, error) {
	query := `SELECT c.* FROM conditions c INNER JOIN schedules s ON s.id == c.schedule_id
           INNER JOIN patients p ON p.id == s.patient_id WHERE p.id = ?`
	row, err := c.db.Query(query, patientId)
	if err != nil {
		return nil, err
	}
	var conditions []model.Conditions
	defer row.Close()
	for row.Next() {
		var condition model.Conditions
		err = row.Scan(&condition.ID, &condition.CreatedAt, &condition.UpdatedAt, &condition.DeletedAt,
			&condition.Height, &condition.Weight, &condition.BloodPressure, &condition.SugarAnalysis,
			&condition.BodyTemperature, &condition.HeartRate, &condition.BreathRate, &condition.Cholesterol,
			&condition.Note, &condition.ScheduleId, &condition.Status)
		if err != nil {
			return nil, err
		}
		conditions = append(conditions, condition)
	}
	return conditions, nil
}

func (c *conditionRepository) GetById(id uint) (model.Conditions, error) {
	query := `SELECT * FROM conditions WHERE id = ?`
	row, err := c.db.Query(query, id)
	if err != nil {
		return model.Conditions{}, err
	}
	var condition model.Conditions
	defer row.Close()
	for row.Next() {
		err = row.Scan(&condition.ID, &condition.CreatedAt, &condition.UpdatedAt, &condition.DeletedAt,
			&condition.Height, &condition.Weight, &condition.BloodPressure, &condition.SugarAnalysis,
			&condition.BodyTemperature, &condition.HeartRate, &condition.BreathRate, &condition.Cholesterol,
			&condition.Note, &condition.ScheduleId, &condition.Status)
		if err != nil {
			return model.Conditions{}, err
		}
	}
	return condition, err
}
func (c *conditionRepository) GetBySchedule(scheduleId uint) (model.Conditions, error) {
	query := `SELECT * FROM conditions WHERE schedule_id = ?`
	row, err := c.db.Query(query, scheduleId)
	if err != nil {
		return model.Conditions{}, err
	}
	var condition model.Conditions
	defer row.Close()
	for row.Next() {
		err = row.Scan(&condition.ID, &condition.CreatedAt, &condition.UpdatedAt, &condition.DeletedAt,
			&condition.Height, &condition.Weight, &condition.BloodPressure, &condition.SugarAnalysis,
			&condition.BodyTemperature, &condition.HeartRate, &condition.BreathRate, &condition.Cholesterol,
			&condition.Note, &condition.ScheduleId, &condition.Status)
		if err != nil {
			return model.Conditions{}, err
		}
	}
	return condition, err
}
