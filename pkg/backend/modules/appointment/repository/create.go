package appointment_repository

import "context"

func (repo *Repository) Create(ctx context.Context, appointment *AppointmentEntity) error {
	_, err := repo.Collection.InsertOne(ctx, appointment)

	if err != nil {
		// TEMP, TODO: return a proper error
		return nil
	}

	return nil
}
