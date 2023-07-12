# openapi_client.HealthApi

All URIs are relative to *http://localhost:8888/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_health**](HealthApi.md#get_health) | **GET** /health | Server Health Check


# **get_health**
> HealthCheck get_health()

Server Health Check

### Example

```python
import time
import os
import openapi_client
from openapi_client.models.health_check import HealthCheck
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
    api_instance = openapi_client.HealthApi(api_client)

    try:
        # Server Health Check
        api_response = api_instance.get_health()
        print("The response of HealthApi->get_health:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling HealthApi->get_health: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheck**](HealthCheck.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**503** | Server Unavailable |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

