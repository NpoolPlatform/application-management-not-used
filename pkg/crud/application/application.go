package application

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/application"
	npool "github.com/NpoolPlatform/message/npool/application"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplication(row *ent.Application) *npool.ApplicationInfo {
	return &npool.ApplicationInfo{
		ID:                 row.ID.String(),
		ApplicationName:    row.ApplicationName,
		ApplicationOwner:   row.ApplicationOwner.String(),
		ApplicationLogo:    row.ApplicationLogo,
		HomepageUrl:        row.HomepageURL,
		RedirectUrl:        row.RedirectURL,
		CreateAT:           row.CreateAt,
		UpdateAT:           row.UpdateAt,
		ClientSecret:       row.ClientSecret,
		SmsLogin:           row.SmsLogin,
		GoogleRecaptcha:    row.GoogleRecaptcha,
		InvitationCodeMust: row.InvitationCodeMust,
	}
}

func Create(ctx context.Context, in *npool.CreateApplicationRequest) (*npool.CreateApplicationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	owner, err := uuid.Parse(in.Info.ApplicationOwner)
	if err != nil {
		return nil, xerrors.Errorf("invalid owner id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	info, err := cli.
		Application.
		Create().
		SetApplicationName(in.Info.ApplicationName).
		SetApplicationOwner(owner).
		SetApplicationLogo(in.Info.ApplicationLogo).
		SetHomepageURL(in.Info.HomepageUrl).
		SetRedirectURL(in.Info.RedirectUrl).
		SetGoogleRecaptcha(in.Info.GoogleRecaptcha).
		SetSmsLogin(in.Info.SmsLogin).
		SetInvitationCodeMust(in.Info.InvitationCodeMust).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create application: %v", err)
	}

	return &npool.CreateApplicationResponse{
		Info: dbRowToApplication(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetApplicationRequest) (*npool.GetApplicationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	info, err := cli.
		Application.
		Query().
		Where(
			application.And(
				application.ID(id),
				application.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get application: %v", err)
	}

	if len(info) == 0 {
		return nil, xerrors.Errorf("application is not exist")
	}

	return &npool.GetApplicationResponse{
		Info: dbRowToApplication(info[0]),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetApplicationsRequest) (*npool.GetApplicationsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	infos, err := cli.
		Application.
		Query().
		Where(
			application.And(
				application.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get all applications: %v", err)
	}

	resp := []*npool.ApplicationInfo{}
	for _, info := range infos {
		resp = append(resp, dbRowToApplication(info))
	}

	return &npool.GetApplicationsResponse{
		Infos: resp,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateApplicationRequest) (*npool.UpdateApplicationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	info, err := cli.
		Application.
		UpdateOneID(id).
		SetApplicationName(in.Info.ApplicationName).
		SetApplicationLogo(in.Info.ApplicationLogo).
		SetHomepageURL(in.Info.HomepageUrl).
		SetRedirectURL(in.Info.RedirectUrl).
		SetGoogleRecaptcha(in.Info.GoogleRecaptcha).
		SetSmsLogin(in.Info.SmsLogin).
		SetInvitationCodeMust(in.Info.InvitationCodeMust).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update application: %v", err)
	}

	return &npool.UpdateApplicationResponse{
		Info: dbRowToApplication(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteApplicationRequest) (*npool.DeleteApplicationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	_, err = cli.
		Application.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete application: %v", err)
	}

	return &npool.DeleteApplicationResponse{
		Info: "delete application successfully",
	}, nil
}

func GetApplicationByOwner(ctx context.Context, in *npool.GetApplicationByOwnerRequest) (*npool.GetApplicationByOwnerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ownerID, err := uuid.Parse(in.Owner)
	if err != nil {
		return nil, xerrors.Errorf("invalid owner id: %v", err)
	}
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}
	infos, err := cli.
		Application.
		Query().
		Where(
			application.And(
				application.ApplicationOwner(ownerID),
				application.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get owner's application: %v", err)
	}

	response := []*npool.ApplicationInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplication(info))
	}
	return &npool.GetApplicationByOwnerResponse{
		Info: &npool.OwnerApplication{
			Infos: response,
			Owner: in.Owner,
		},
	}, nil
}
