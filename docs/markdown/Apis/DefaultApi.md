# DefaultApi

All URIs are relative to *http://localhost:8080*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PlayersActiveGet**](DefaultApi.md#apiV1PlayersActiveGet) | **GET** /api/v1/players/active | Get active players by Career year |


<a name="apiV1PlayersActiveGet"></a>
# **apiV1PlayersActiveGet**
> ActivePlayers apiV1PlayersActiveGet(careerYear)

Get active players by Career year

    Returns a list of player names by Career year

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **careerYear** | **Integer**| year that the player is still in career | [default to null] |

### Return type

[**ActivePlayers**](../Models/ActivePlayers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

