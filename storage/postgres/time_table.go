package postgres

import (
	"backend_course/lms/api/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type timeRepo struct {
	db *pgxpool.Pool
}

func NewTime(db *pgxpool.Pool) timeRepo {
	return timeRepo{
		db: db,
	}
}

func (s *timeRepo) Create(ctx context.Context, time models.Time) (string, error) {
	id := uuid.New()

	query := `
	INSERT INTO
		time_table (id, teacher_id, student_id, subject_id, from_date, to_date, room_name) VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := s.db.Exec(ctx, query, id, time.TeacherId, time.StudentId, time.SubjectId, time.FromDate, time.ToDate, time.RoomName)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *timeRepo) Update(ctx context.Context, time models.Time) (string, error) {
	query := `
	UPDATE
		time_table
	SET
		teacher_id = $2, student_id = $3, subject_id = $4, from_date = $5, to_date = $6, room_name = $7
	WHERE 
		id = $1; `

	_, err := s.db.Exec(ctx, query, time.Id, time.TeacherId, time.StudentId, time.SubjectId, time.FromDate, time.ToDate)
	if err != nil {
		return "", err
	}
	return time.Id, nil
}

func (s *timeRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE
	FROM
		time_table
	WHERE 
		id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *timeRepo) GetAll(ctx context.Context, req models.GetAllTimeRequest) (models.GetAllTimeResponse, error) {
	resp := models.GetAllTimeResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `
	SELECT
		id,
		teacher_id,
		student_id,
		subject_id,
		TO_CHAR(from_date,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(to_date,'YYYY-MM-DD HH:MM:SS'),
		room_name
	FROM 
		time_table
	WHERE 
		TRUE ` + filter + `
	OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			time models.Time
		)
		if err := rows.Scan(
			&time.Id,
			&time.TeacherId,
			&time.StudentId,
			&time.SubjectId,
			&time.FromDate,
			&time.ToDate,
			&time.RoomName); err != nil {
			return resp, err
		}

		resp.Time = append(resp.Time, time)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) from time_table WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *timeRepo) GetTime(ctx context.Context, id string) (models.Time, error) {

	query := `
	SELECT
		id,
		teacher_id,
		student_id,
		subject_id,
		TO_CHAR(from_date,'YYYY-MM-DD HH:MM:SS'),
		TO_CHAR(to_date,'YYYY-MM-DD HH:MM:SS'),
		room_name
	FROM
		time_table
	WHERE
		id = $1;`
	row := s.db.QueryRow(ctx, query, id)

	var time models.Time

	err := row.Scan(&time.Id, &time.TeacherId, &time.StudentId, &time.SubjectId, &time.FromDate, &time.ToDate, &time.RoomName)

	if err != nil {
		return time, err
	}
	return time, nil
}

func (s *timeRepo) GetAllStudentsAttandenceReport(ctx context.Context, req models.GetAllStudentsAttandenceReportRequest) (models.GetAllStudentsAttandenceReportResponse, error) {
	resp := models.GetAllStudentsAttandenceReportResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.StudentId != "" {
		filter = ` AND s.id =` + req.StudentId + ` `
	}

	if req.TeacherId != "" {
		filter += ` AND t.id =` + req.TeacherId + ` `
	}

	if req.StartDate != "" && req.EndDate != "" {
		filter += ` AND tt.from_date BETWEEN '` + req.StartDate + `' AND '` + req.EndDate + `' `
	}

	// 	1. Student name,
	// 	2. student createdAt,
	// 	3. o’qituvchi name,
	// 	4. studying_time,
	// 	5. avg_studying_time,
	query := `
	SELECT
		s.id,
		s.first_name || ' ' || s.last_name AS student_name,
		TO_CHAR(s.created_at,'YYYY-MM-DD HH:MM:SS'),
		t.first_name || ' ' || t.last_name AS teacher_name,
		EXTRACT(tt.to_date - tt.from_date) / 60 AS studying_time
		
	FROM 
		time_table tt
		JOIN students s on tt.student_id = s.id
		JOIN teachers t on tt.teacher_id = t.id

	WHERE 
		TRUE ` + filter + `
	OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	studentAttandence := models.StudentAttandenceReport{}
	for rows.Next() {

		if err := rows.Scan(
			&studentAttandence.StudentId,
			&studentAttandence.StudentName,
			&studentAttandence.StudentCreatedAt,
			&studentAttandence.TeacherName,
			&studentAttandence.StudyTime); err != nil {
			return resp, err
		}

		resp.Students = append(resp.Students, studentAttandence)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) from time_table time_table tt
	JOIN students s on tt.student_id = s.id
	JOIN teachers t on tt.teacher_id = t.id WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
