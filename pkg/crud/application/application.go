package application

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/application"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplication(row *ent.Application) *npool.ApplicationInfo {
	return &npool.ApplicationInfo{
		ID:               row.ID,
		ApplicationName:  row.ApplicationName,
		ApplicationOwner: row.ApplicationOwner.String(),
		ApplicationLogo:  row.ApplicationLogo,
		HomepageUrl:      row.HomepageURL,
		RedirectUrl:      row.RedirectURL,
		CreateAT:         int32(row.CreateAt),
		UpdateAT:         int32(row.UpdateAt),
		ClientSecret:     row.ClientSecret,
	}
}

func Create(ctx context.Context, in *npool.CreateApplicationRequest) (*npool.CreateApplicationResponse, error) {
	owner, err := uuid.Parse(in.Request.ApplicationOwner)
	if err != nil {
		return nil, xerrors.Errorf("invalid owner id: %v", err)
	}
	info, err := db.Client().
		Application.
		Create().
		SetApplicationName(in.Request.ApplicationName).
		SetApplicationOwner(owner).
		SetApplicationLogo(in.Request.ApplicationLogo).
		SetHomepageURL(in.Request.HomepageUrl).
		SetRedirectURL(in.Request.RedirectUrl).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create application: %v", err)
	}

	return &npool.CreateApplicationResponse{
		Info: dbRowToApplication(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetApplicationRequest) (*npool.GetApplicationResponse, error) {
	info, err := db.Client().
		Application.
		Query().
		Where(
			application.And(
				application.ID(in.AppID),
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
	infos, err := db.Client().
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

	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty database")
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
	info, err := db.Client().
		Application.
		UpdateOneID(in.Request.ID).
		SetApplicationName(in.Request.ApplicationOwner).
		SetApplicationLogo(in.Request.ApplicationLogo).
		SetHomepageURL(in.Request.HomepageUrl).
		SetRedirectURL(in.Request.RedirectUrl).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update application: %v", err)
	}

	return &npool.UpdateApplicationResponse{
		Info: dbRowToApplication(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteApplicationRequest) (*npool.DeleteApplicationResponse, error) {
	_, err := db.Client().
		Application.
		UpdateOneID(in.AppId).
		SetDeleteAt(time.Now().UnixNano()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete application: %v", err)
	}

	return &npool.DeleteApplicationResponse{
		Info: "delete application successfully",
	}, nil
}

func GetApplicationsByOwner(ctx context.Context, owner string) ([]*npool.ApplicationInfo, error) {
	ownerID, err := uuid.Parse(owner)
	if err != nil {
		return nil, xerrors.Errorf("invalid owner id: %v", err)
	}

	infos, err := db.Client().
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
	return response, nil
}