# openapi_client.UserApi

All URIs are relative to *http://localhost:8888/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_users**](UserApi.md#add_users) | **POST** /users | Add Users
[**get_users**](UserApi.md#get_users) | **GET** /users | Get Users


# **add_users**
> UsersAdded add_users(user)

Add Users

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.user import User
from openapi_client.models.users_added import UsersAdded
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8888/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8888/api/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    user = [openapi_client.User()] # List[User] | Create a new pet in the store

    try:
        # Add Users
        api_response = api_instance.add_users(user)
        print("The response of UserApi->add_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->add_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**List[User]**](User.md)| Create a new pet in the store | 

### Return type

[**UsersAdded**](UsersAdded.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful operation |  -  |
**400** | Invalid input |  -  |
**401** | Unauthorized |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_users**
> List[User] get_users(offset=offset, limit=limit, sort=sort)

Get Users

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.user import User
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8888/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8888/api/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    offset = 56 # int | The number of items to skip before starting to collect the result set (optional)
    limit = 56 # int | The numbers of items to return (optional)
    sort = 'sort_example' # str | Sort by this field (optional)

    try:
        # Get Users
        api_response = api_instance.get_users(offset=offset, limit=limit, sort=sort)
        print("The response of UserApi->get_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->get_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int**| The number of items to skip before starting to collect the result set | [optional] 
 **limit** | **int**| The numbers of items to return | [optional] 
 **sort** | **str**| Sort by this field | [optional] 

### Return type

[**List[User]**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

