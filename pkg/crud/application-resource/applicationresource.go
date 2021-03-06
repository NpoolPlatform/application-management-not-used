package applicationresource

import (
	"context"
	"time"

	"github.com/NpoolPlatform/application-management/pkg/db"
	"github.com/NpoolPlatform/application-management/pkg/db/ent"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/NpoolPlatform/application-management/pkg/exist"
	npool "github.com/NpoolPlatform/message/npool/application"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func dbRowToApplicationResource(row *ent.ApplicationResource) *npool.ResourceInfo {
	return &npool.ResourceInfo{
		ID:                  row.ID.String(),
		AppID:               row.AppID.String(),
		ResourceName:        row.ResourceName,
		Type:                row.Type,
		ResourceDescription: row.ResourceDescription,
		Creator:             row.Creator.String(),
		CreateAT:            row.CreateAt,
		UpdateAT:            row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateResourceRequest) (*npool.CreateResourceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	creator, err := uuid.Parse(in.Info.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		ApplicationResource.
		Create().
		SetAppID(id).
		SetResourceName(in.Info.ResourceName).
		SetType(in.Info.Type).
		SetResourceDescription(in.Info.ResourceDescription).
		SetCreator(creator).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create resource: %v", err)
	}
	return &npool.CreateResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetResourceRequest) (*npool.GetResourceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	resourceID, err := uuid.Parse(in.ResourceID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.ID(resourceID),
				applicationresource.AppID(id),
				applicationresource.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource: %v", err)
	}
	return &npool.GetResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func GetResourceByCreator(ctx context.Context, in *npool.GetResourceByCreatorRequest) (*npool.GetResourceByCreatorResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	creatorID, err := uuid.Parse(in.Creator)
	if err != nil {
		return nil, xerrors.Errorf("invalid creator id: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.Creator(creatorID),
				applicationresource.AppID(id),
				applicationresource.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource by creator: %v", err)
	}

	response := []*npool.ResourceInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationResource(info))
	}

	return &npool.GetResourceByCreatorResponse{
		Info: &npool.CreatorResource{
			Infos:   response,
			AppID:   in.AppID,
			Creator: in.Creator,
		},
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetResourcesRequest) (*npool.GetResourcesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	id, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.AppID(id),
				applicationresource.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get resources of app: %v", err)
	}

	response := []*npool.ResourceInfo{}
	for _, info := range infos {
		response = append(response, dbRowToApplicationResource(info))
	}
	return &npool.GetResourcesResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateResourceRequest) (*npool.UpdateResourceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.Info.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	resourceID, err := uuid.Parse(in.Info.ID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	id, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	query, err := cli.
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.ID(resourceID),
				applicationresource.AppID(id),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to query resource: %v", err)
	}

	if query.DeleteAt != 0 {
		return nil, xerrors.Errorf("resource has already been delete")
	}

	info, err := cli.
		ApplicationResource.
		UpdateOneID(resourceID).
		SetResourceName(in.Info.ResourceName).
		SetType(in.Info.Type).
		SetResourceDescription(in.Info.ResourceDescription).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update resource: %v", err)
	}
	return &npool.UpdateResourceResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	resourceID, err := uuid.Parse(in.ResourceID)
	if err != nil {
		return nil, xerrors.Errorf("invalid resource id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		ApplicationResource.
		UpdateOneID(resourceID).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to delete resource: %v", err)
	}
	return &npool.DeleteResourceResponse{
		Info: "delete resource successfully",
	}, nil
}

func GetResourceByName(ctx context.Context, in *npool.GetResourceByNameRequest) (*npool.GetResourceByNameResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if existApp, err := exist.Application(ctx, in.AppID); err != nil || !existApp {
		return nil, xerrors.Errorf("application does not exist: %v", err)
	}

	appID, err := uuid.Parse(in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		ApplicationResource.
		Query().
		Where(
			applicationresource.And(
				applicationresource.AppID(appID),
				applicationresource.DeleteAt(0),
				applicationresource.ResourceName(in.ResourceName),
			),
		).Only(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to get resource by resource name: %v", err)
	}

	return &npool.GetResourceByNameResponse{
		Info: dbRowToApplicationResource(info),
	}, nil
}
