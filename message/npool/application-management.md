# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/application-management.proto](#npool/application-management.proto)
    - [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest)
    - [AddGroupUsersResponse](#application.management.v1.AddGroupUsersResponse)
    - [AddUsersToApplicationRequest](#application.management.v1.AddUsersToApplicationRequest)
    - [AddUsersToApplicationResponse](#application.management.v1.AddUsersToApplicationResponse)
    - [ApplicationInfo](#application.management.v1.ApplicationInfo)
    - [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo)
    - [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest)
    - [CreateApplicationResponse](#application.management.v1.CreateApplicationResponse)
    - [CreateGroupRequest](#application.management.v1.CreateGroupRequest)
    - [CreateGroupResponse](#application.management.v1.CreateGroupResponse)
    - [CreateResourceRequest](#application.management.v1.CreateResourceRequest)
    - [CreateResourceResponse](#application.management.v1.CreateResourceResponse)
    - [CreateRoleRequest](#application.management.v1.CreateRoleRequest)
    - [CreateRoleResponse](#application.management.v1.CreateRoleResponse)
    - [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest)
    - [DeleteApplicationResponse](#application.management.v1.DeleteApplicationResponse)
    - [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest)
    - [DeleteGroupResponse](#application.management.v1.DeleteGroupResponse)
    - [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest)
    - [DeleteResourceResponse](#application.management.v1.DeleteResourceResponse)
    - [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest)
    - [DeleteRoleResponse](#application.management.v1.DeleteRoleResponse)
    - [GetAllGroupsRequest](#application.management.v1.GetAllGroupsRequest)
    - [GetAllGroupsResponse](#application.management.v1.GetAllGroupsResponse)
    - [GetApplicationRequest](#application.management.v1.GetApplicationRequest)
    - [GetApplicationResponse](#application.management.v1.GetApplicationResponse)
    - [GetApplicationsRequest](#application.management.v1.GetApplicationsRequest)
    - [GetApplicationsResponse](#application.management.v1.GetApplicationsResponse)
    - [GetGroupRequest](#application.management.v1.GetGroupRequest)
    - [GetGroupResponse](#application.management.v1.GetGroupResponse)
    - [GetGroupUsersRequest](#application.management.v1.GetGroupUsersRequest)
    - [GetGroupUsersResponse](#application.management.v1.GetGroupUsersResponse)
    - [GetResourceRequest](#application.management.v1.GetResourceRequest)
    - [GetResourceResponse](#application.management.v1.GetResourceResponse)
    - [GetResourcesRequest](#application.management.v1.GetResourcesRequest)
    - [GetResourcesResponse](#application.management.v1.GetResourcesResponse)
    - [GetRoleRequest](#application.management.v1.GetRoleRequest)
    - [GetRoleResponse](#application.management.v1.GetRoleResponse)
    - [GetRoleUsersRequest](#application.management.v1.GetRoleUsersRequest)
    - [GetRoleUsersResponse](#application.management.v1.GetRoleUsersResponse)
    - [GetRolesRequest](#application.management.v1.GetRolesRequest)
    - [GetRolesResponse](#application.management.v1.GetRolesResponse)
    - [GetUserFromApplicationRequest](#application.management.v1.GetUserFromApplicationRequest)
    - [GetUserFromApplicationResponse](#application.management.v1.GetUserFromApplicationResponse)
    - [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest)
    - [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse)
    - [GetUsersRequest](#application.management.v1.GetUsersRequest)
    - [GetUsersResponse](#application.management.v1.GetUsersResponse)
    - [GroupInfo](#application.management.v1.GroupInfo)
    - [GroupUserInfo](#application.management.v1.GroupUserInfo)
    - [PageInfo](#application.management.v1.PageInfo)
    - [RemoveGroupUsersNoDataResponse](#application.management.v1.RemoveGroupUsersNoDataResponse)
    - [RemoveGroupUsersRequest](#application.management.v1.RemoveGroupUsersRequest)
    - [RemoveUsersFromApplicationRequest](#application.management.v1.RemoveUsersFromApplicationRequest)
    - [RemoveUsersFromApplicationResponse](#application.management.v1.RemoveUsersFromApplicationResponse)
    - [ResourceInfo](#application.management.v1.ResourceInfo)
    - [RoleInfo](#application.management.v1.RoleInfo)
    - [RoleUserInfo](#application.management.v1.RoleUserInfo)
    - [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest)
    - [SetUserRoleResponse](#application.management.v1.SetUserRoleResponse)
    - [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest)
    - [UnSetUserRoleResponse](#application.management.v1.UnSetUserRoleResponse)
    - [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest)
    - [UpdateApplicationResponse](#application.management.v1.UpdateApplicationResponse)
    - [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest)
    - [UpdateGroupResponse](#application.management.v1.UpdateGroupResponse)
    - [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest)
    - [UpdateResourceResponse](#application.management.v1.UpdateResourceResponse)
    - [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest)
    - [UpdateRoleResponse](#application.management.v1.UpdateRoleResponse)
    - [VersionResponse](#application.management.v1.VersionResponse)
  
    - [ApplicationManagement](#application.management.v1.ApplicationManagement)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/application-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/application-management.proto



<a name="application.management.v1.AddGroupUsersRequest"></a>

### AddGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.AddGroupUsersResponse"></a>

### AddGroupUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |






<a name="application.management.v1.AddUsersToApplicationRequest"></a>

### AddUsersToApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| AppID | [string](#string) |  |  |
| Original | [bool](#bool) |  |  |






<a name="application.management.v1.AddUsersToApplicationResponse"></a>

### AddUsersToApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) | repeated |  |






<a name="application.management.v1.ApplicationInfo"></a>

### ApplicationInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ApplicationName | [string](#string) |  |  |
| ApplicationOwner | [string](#string) |  |  |
| HomepageUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |
| ApplicationLogo | [string](#string) |  |  |
| CreateAT | [int32](#int32) |  |  |
| UpdateAT | [int32](#int32) |  |  |
| ClientSecret | [string](#string) |  |  |






<a name="application.management.v1.ApplicationUserInfo"></a>

### ApplicationUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Original | [bool](#bool) |  |  |
| CreateAT | [int32](#int32) |  |  |






<a name="application.management.v1.CreateApplicationRequest"></a>

### CreateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.CreateApplicationResponse"></a>

### CreateApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.CreateGroupRequest"></a>

### CreateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.CreateResourceResponse"></a>

### CreateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.CreateRoleResponse"></a>

### CreateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.DeleteApplicationRequest"></a>

### DeleteApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.DeleteApplicationResponse"></a>

### DeleteApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupRequest"></a>

### DeleteGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupResponse"></a>

### DeleteGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteResourceResponse"></a>

### DeleteResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteRoleResponse"></a>

### DeleteRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.GetAllGroupsRequest"></a>

### GetAllGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetAllGroupsResponse"></a>

### GetAllGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupInfo](#application.management.v1.GroupInfo) | repeated |  |






<a name="application.management.v1.GetApplicationRequest"></a>

### GetApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetApplicationResponse"></a>

### GetApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.GetApplicationsRequest"></a>

### GetApplicationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [PageInfo](#application.management.v1.PageInfo) |  |  |






<a name="application.management.v1.GetApplicationsResponse"></a>

### GetApplicationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationInfo](#application.management.v1.ApplicationInfo) | repeated |  |






<a name="application.management.v1.GetGroupRequest"></a>

### GetGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetGroupResponse"></a>

### GetGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.GetGroupUsersRequest"></a>

### GetGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetGroupUsersResponse"></a>

### GetGroupUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |






<a name="application.management.v1.GetResourceRequest"></a>

### GetResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetResourceResponse"></a>

### GetResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.GetResourcesRequest"></a>

### GetResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetResourcesResponse"></a>

### GetResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ResourceInfo](#application.management.v1.ResourceInfo) | repeated |  |






<a name="application.management.v1.GetRoleRequest"></a>

### GetRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="application.management.v1.GetRoleResponse"></a>

### GetRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.GetRoleUsersRequest"></a>

### GetRoleUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="application.management.v1.GetRoleUsersResponse"></a>

### GetRoleUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleUserInfo](#application.management.v1.RoleUserInfo) | repeated |  |






<a name="application.management.v1.GetRolesRequest"></a>

### GetRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleInfo](#application.management.v1.RoleInfo) | repeated |  |






<a name="application.management.v1.GetUserFromApplicationRequest"></a>

### GetUserFromApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="application.management.v1.GetUserFromApplicationResponse"></a>

### GetUserFromApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) |  |  |






<a name="application.management.v1.GetUserRoleRequest"></a>

### GetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetUserRoleResponse"></a>

### GetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleInfo](#application.management.v1.RoleInfo) | repeated |  |






<a name="application.management.v1.GetUsersRequest"></a>

### GetUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetUsersResponse"></a>

### GetUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) | repeated |  |






<a name="application.management.v1.GroupInfo"></a>

### GroupInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GroupName | [string](#string) |  |  |
| GroupOwner | [string](#string) |  |  |
| GroupLogo | [string](#string) |  |  |
| Annotation | [string](#string) |  |  |
| CreateAT | [int32](#int32) |  |  |
| UpdateAT | [int32](#int32) |  |  |






<a name="application.management.v1.GroupUserInfo"></a>

### GroupUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Annotation | [string](#string) |  |  |
| CreateAT | [string](#string) |  |  |






<a name="application.management.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageSize | [int32](#int32) |  |  |






<a name="application.management.v1.RemoveGroupUsersNoDataResponse"></a>

### RemoveGroupUsersNoDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.RemoveGroupUsersRequest"></a>

### RemoveGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.RemoveUsersFromApplicationRequest"></a>

### RemoveUsersFromApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.RemoveUsersFromApplicationResponse"></a>

### RemoveUsersFromApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ResourceName | [string](#string) |  |  |
| ResourceDescription | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| CreateAT | [int32](#int32) |  |  |
| UpdateAT | [int32](#int32) |  |  |






<a name="application.management.v1.RoleInfo"></a>

### RoleInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleName | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| CreateAT | [int32](#int32) |  |  |
| UpdateAT | [int32](#int32) |  |  |
| Annotation | [string](#string) |  |  |






<a name="application.management.v1.RoleUserInfo"></a>

### RoleUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CreateAT | [int32](#int32) |  |  |






<a name="application.management.v1.SetUserRoleRequest"></a>

### SetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserId | [string](#string) | repeated |  |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="application.management.v1.SetUserRoleResponse"></a>

### SetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleUserInfo](#application.management.v1.RoleUserInfo) | repeated |  |






<a name="application.management.v1.UnSetUserRoleRequest"></a>

### UnSetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| RoleID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.UnSetUserRoleResponse"></a>

### UnSetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.UpdateApplicationRequest"></a>

### UpdateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.UpdateApplicationResponse"></a>

### UpdateApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.UpdateGroupRequest"></a>

### UpdateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.UpdateGroupResponse"></a>

### UpdateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.UpdateResourceResponse"></a>

### UpdateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Request | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.UpdateRoleResponse"></a>

### UpdateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="application.management.v1.ApplicationManagement"></a>

### ApplicationManagement
rpc Echo (StringMessage) returns (StringMessage){
    option (google.api.http) = {
        post: &#34;/v1/echo&#34;
        body: &#34;*&#34;
    };
}

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#application.management.v1.VersionResponse) | Method Version |
| CreateApplication | [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest) | [CreateApplicationResponse](#application.management.v1.CreateApplicationResponse) | Create an application. |
| UpdateApplication | [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest) | [UpdateApplicationResponse](#application.management.v1.UpdateApplicationResponse) | Update an application&#39;s basic info. |
| GetApplication | [GetApplicationRequest](#application.management.v1.GetApplicationRequest) | [GetApplicationResponse](#application.management.v1.GetApplicationResponse) | Get application. |
| GetApplications | [GetApplicationsRequest](#application.management.v1.GetApplicationsRequest) | [GetApplicationsResponse](#application.management.v1.GetApplicationsResponse) | Get all applications. |
| DeleteApplication | [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest) | [DeleteApplicationResponse](#application.management.v1.DeleteApplicationResponse) | Delete an application. |
| CreateRole | [CreateRoleRequest](#application.management.v1.CreateRoleRequest) | [CreateRoleResponse](#application.management.v1.CreateRoleResponse) | Create a role in app. |
| UpdateRole | [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest) | [UpdateRoleResponse](#application.management.v1.UpdateRoleResponse) | Update role&#39;s basic info. |
| GetRole | [GetRoleRequest](#application.management.v1.GetRoleRequest) | [GetRoleResponse](#application.management.v1.GetRoleResponse) | Get Role. |
| GetRoles | [GetRolesRequest](#application.management.v1.GetRolesRequest) | [GetRolesResponse](#application.management.v1.GetRolesResponse) | Get Roles. |
| DeleteRole | [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest) | [DeleteRoleResponse](#application.management.v1.DeleteRoleResponse) | Delete role from app. |
| SetUserRole | [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest) | [SetUserRoleResponse](#application.management.v1.SetUserRoleResponse) | Set role to user. |
| GetUserRole | [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest) | [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse) | Get user role. |
| GetRoleUsers | [GetRoleUsersRequest](#application.management.v1.GetRoleUsersRequest) | [GetRoleUsersResponse](#application.management.v1.GetRoleUsersResponse) | Get role users. |
| UnSetUserRole | [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest) | [UnSetUserRoleResponse](#application.management.v1.UnSetUserRoleResponse) | Unset user role. |
| AddUsersToApplication | [AddUsersToApplicationRequest](#application.management.v1.AddUsersToApplicationRequest) | [AddUsersToApplicationResponse](#application.management.v1.AddUsersToApplicationResponse) | Add users to app. |
| GetUserFromApplication | [GetUserFromApplicationRequest](#application.management.v1.GetUserFromApplicationRequest) | [GetUserFromApplicationResponse](#application.management.v1.GetUserFromApplicationResponse) | Get user from app. |
| GetUsersFromApplication | [GetUsersRequest](#application.management.v1.GetUsersRequest) | [GetUsersResponse](#application.management.v1.GetUsersResponse) | Get users from app. |
| RemoveUsersFromApplication | [RemoveUsersFromApplicationRequest](#application.management.v1.RemoveUsersFromApplicationRequest) | [RemoveUsersFromApplicationResponse](#application.management.v1.RemoveUsersFromApplicationResponse) | Remove users from app. |
| CreateGroup | [CreateGroupRequest](#application.management.v1.CreateGroupRequest) | [CreateGroupResponse](#application.management.v1.CreateGroupResponse) | Create group in an application. |
| GetGroup | [GetGroupRequest](#application.management.v1.GetGroupRequest) | [GetGroupResponse](#application.management.v1.GetGroupResponse) | Get group info. |
| GetAllGroups | [GetAllGroupsRequest](#application.management.v1.GetAllGroupsRequest) | [GetAllGroupsResponse](#application.management.v1.GetAllGroupsResponse) | Get all groups. |
| UpdateGroup | [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest) | [UpdateGroupResponse](#application.management.v1.UpdateGroupResponse) | Update group info. |
| DeleteGroup | [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest) | [DeleteGroupResponse](#application.management.v1.DeleteGroupResponse) | Delete group. |
| AddGroupUsers | [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest) | [AddGroupUsersResponse](#application.management.v1.AddGroupUsersResponse) | Add users into group. |
| GetGroupUsers | [GetGroupUsersRequest](#application.management.v1.GetGroupUsersRequest) | [GetGroupUsersResponse](#application.management.v1.GetGroupUsersResponse) | Get group users. |
| RemoveGroupUsers | [RemoveGroupUsersRequest](#application.management.v1.RemoveGroupUsersRequest) | [RemoveGroupUsersNoDataResponse](#application.management.v1.RemoveGroupUsersNoDataResponse) | Remove users from group. |
| CreateResource | [CreateResourceRequest](#application.management.v1.CreateResourceRequest) | [CreateResourceResponse](#application.management.v1.CreateResourceResponse) | Create resource for app. |
| UpdateResource | [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest) | [UpdateResourceResponse](#application.management.v1.UpdateResourceResponse) | Update resource of app. |
| GetResource | [GetResourceRequest](#application.management.v1.GetResourceRequest) | [GetResourceResponse](#application.management.v1.GetResourceResponse) | Get resource. |
| GetResources | [GetResourcesRequest](#application.management.v1.GetResourcesRequest) | [GetResourcesResponse](#application.management.v1.GetResourcesResponse) | Get all resources from app. |
| DeleteResource | [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest) | [DeleteResourceResponse](#application.management.v1.DeleteResourceResponse) | Delete resource from app. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
