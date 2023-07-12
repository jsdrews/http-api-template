# UsersAdded


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **str** |  | 
**num_requested** | **int** |  | 
**num_existed** | **int** |  | 
**num_added** | **int** |  | 

## Example

```python
from openapi_client.models.users_added import UsersAdded

# TODO update the JSON string below
json = "{}"
# create an instance of UsersAdded from a JSON string
users_added_instance = UsersAdded.from_json(json)
# print the JSON string representation of the object
print UsersAdded.to_json()

# convert the object into a dict
users_added_dict = users_added_instance.to_dict()
# create an instance of UsersAdded from a dict
users_added_form_dict = users_added.from_dict(users_added_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


