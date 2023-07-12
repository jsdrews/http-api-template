# {{classname}}

All URIs are relative to *http://localhost:8888/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUsers**](UserApi.md#AddUsers) | **Post** /users | Add Users
[**GetUsers**](UserApi.md#GetUsers) | **Get** /users | Get Users

# **AddUsers**
> UsersAdded AddUsers(ctx, body)
Add Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]User**](User.md)| Create a new pet in the store | 

### Return type

[**UsersAdded**](UsersAdded.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> []User GetUsers(ctx, optional)
Get Users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | 
 **limit** | **optional.Int32**| The numbers of items to return | 
 **sort** | **optional.String**| Sort by this field | 

### Return type

[**[]User**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

