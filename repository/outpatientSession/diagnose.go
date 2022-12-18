package outpatientSession

import (
	"HMS-16-BE/model"
	"database/sql"
)

type DiagnoseRepository interface {
	Create(diagnose model.Diagnoses) error
	GetAllByPatient(patientId string) ([]model.Diagnoses, error)
	GetById(id uint) (model.Diagnoses, error)
	GetBySchedule(scheduleId uint) (model.Diagnoses, error)
}

type diagnoseRepository struct {
	db *sql.DB
}

func NewDiagnoseRepository(db *sql.DB) *diagnoseRepository {
	return &diagnoseRepository{db}
}

func (d *diagnoseRepository) Create(diagnose model.Diagnoses) error {
	query := `INSERT INTO diagnoses VALUES (?,?,?,?,?,?,?,?,?)`
	_, err := d.db.Exec(query, diagnose.ID, diagnose.CreatedAt, diagnose.UpdatedAt, diagnose.DeletedAt,
		diagnose.DoctorId, diagnose.Diagnose, diagnose.Prescription, diagnose.ScheduleId, diagnose.Status)
	if err != nil {
		return err
	}
	return nil
}

func (d *diagnoseRepository) GetAllByPatient(patientId string) ([]model.Diagnoses, error) {
	query := `SELECT d.* FROM diagnoses cdINNER JOIN schedules s ON s.id == d.schedule_id
           INNER JOIN patients p ON p.id == s.patient_id WHERE p.id = ?`
	row, err := d.db.Query(query, patientId)
	if err != nil {
		return nil, err
	}
	var diagnoses []model.Diagnoses
	defer row.Close()
	for row.Next() {
		var diagnose model.Diagnoses
		err = row.Scan(&diagnose.ID, &diagnose.CreatedAt, &diagnose.UpdatedAt, &diagnose.DeletedAt,
			&diagnose.DoctorId, &diagnose.Diagnose, &diagnose.Prescription, &diagnose.ScheduleId, &diagnose.Status)
		if err != nil {
			return nil, err
		}
		diagnoses = append(diagnoses, diagnose)
	}
	return diagnoses, nil
}

func (d *diagnoseRepository) GetById(id string) (model.Diagnoses, error) {
	query := `SELECT * FROM diagnoses WHERE id = ?`
	row, err := d.db.Query(query, id)
	if err != nil {
		return model.Diagnoses{}, err
	}
	var diagnose model.Diagnoses
	defer row.Close()
	for row.Next() {
		err = row.Scan(&diagnose.ID, &diagnose.CreatedAt, &diagnose.UpdatedAt, &diagnose.DeletedAt,
			&diagnose.DoctorId, &diagnose.Diagnose, &diagnose.Prescription, &diagnose.ScheduleId, &diagnose.Status)
		if err != nil {
			return model.Diagnoses{}, err
		}
	}
	return diagnose, err
}

func (d *diagnoseRepository) GetBySchedule(scheduleId string) (model.Diagnoses, error) {
	query := `SELECT * FROM diagnoses WHERE schedule_id = ?`
	row, err := d.db.Query(query, scheduleId)
	if err != nil {
		return model.Diagnoses{}, err
	}
	var diagnose model.Diagnoses
	defer row.Close()
	for row.Next() {
		err = row.Scan(&diagnose.ID, &diagnose.CreatedAt, &diagnose.UpdatedAt, &diagnose.DeletedAt,
			&diagnose.DoctorId, &diagnose.Diagnose, &diagnose.Prescription, &diagnose.ScheduleId, &diagnose.Status)
		if err != nil {
			return model.Diagnoses{}, err
		}
	}
	return diagnose, err
}
