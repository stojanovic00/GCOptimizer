package repo

import (
	"application_service/core/domain"
	"application_service/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DelegationMemberRepoPg struct {
	dbClient *gorm.DB
}

func NewDelegationMemberRepoPg(dbClient *gorm.DB) *DelegationMemberRepoPg {
	return &DelegationMemberRepoPg{dbClient: dbClient}
}

func (r *DelegationMemberRepoPg) GetPositionByName(name string) (*domain.DelegationMemberPosition, error) {
	var position domain.DelegationMemberPosition

	result := r.dbClient.Where("name = ?", name).First(&position)
	if result.Error != nil {
		return &domain.DelegationMemberPosition{}, result.Error
	}

	if &position == nil {
		return &domain.DelegationMemberPosition{}, errors.ErrNotFound{Message: "Sports organisation with given email not found"}
	}
	return &position, nil
}

func (r *DelegationMemberRepoPg) RegisterJudge(judge *domain.Judge) (uuid.UUID, error) {
	_, err := r.GetJudgeByEmail(judge.Email)
	if err == nil {
		return uuid.UUID{}, errors.ErrEmailTaken{}
	}

	position, err := r.GetPositionByName("judge")
	if err != nil {
		return uuid.UUID{}, err
	}

	judge.Position = *position

	if judge.ID == uuid.Nil {
		judge.ID, _ = uuid.NewUUID()
	}

	result := r.dbClient.Create(judge)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return judge.ID, nil
}
func (r *DelegationMemberRepoPg) GetSportsOrganisationJudges(soId uuid.UUID) ([]*domain.Judge, error) {
	var judges []*domain.Judge
	if err := r.dbClient.Where("sports_organization_id = ?", soId).Preload("Position").Find(&judges).Error; err != nil {
		return nil, nil
	}

	return judges, nil
}
func (r *DelegationMemberRepoPg) GetJudgeByEmail(email string) (*domain.Judge, error) {
	var judge domain.Judge

	result := r.dbClient.Where("email = ?", email).First(&judge)
	if result.Error != nil {
		return &domain.Judge{}, result.Error
	}

	if &judge == nil {
		return &domain.Judge{}, errors.ErrNotFound{Message: "Judge with given email not found"}
	}
	return &judge, nil
}

func (r *DelegationMemberRepoPg) RegisterContestant(contestant *domain.Contestant) (uuid.UUID, error) {
	_, err := r.GetContestantByEmail(contestant.Email)
	if err == nil {
		return uuid.UUID{}, errors.ErrEmailTaken{}
	}

	position, err := r.GetPositionByName("contestant")
	if err != nil {
		return uuid.UUID{}, err
	}

	contestant.Position = *position

	if contestant.ID == uuid.Nil {
		contestant.ID, _ = uuid.NewUUID()
	}

	result := r.dbClient.Create(contestant)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return contestant.ID, nil
}

func (r *DelegationMemberRepoPg) GetSportsOrganisationContestants(soId uuid.UUID) ([]*domain.Contestant, error) {
	var contestant []*domain.Contestant
	if err := r.dbClient.Where("sports_organization_id = ?", soId).Preload("Position").Find(&contestant).Error; err != nil {
		return nil, nil
	}

	return contestant, nil
}

func (r *DelegationMemberRepoPg) GetContestantByEmail(email string) (*domain.Contestant, error) {
	var contestant domain.Contestant

	result := r.dbClient.Where("email = ?", email).First(&contestant)
	if result.Error != nil {
		return &domain.Contestant{}, result.Error
	}

	if &contestant == nil {
		return &domain.Contestant{}, errors.ErrNotFound{Message: "Contestant with given email not found"}
	}
	return &contestant, nil
}
