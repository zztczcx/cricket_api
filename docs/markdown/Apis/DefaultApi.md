# DefaultApi

All URIs are relative to *http://localhost:8080*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiV1PlayersActiveGet**](DefaultApi.md#apiV1PlayersActiveGet) | **GET** /api/v1/players/active | Get active players by Career year |
| [**apiV1PlayersMostRunsGet**](DefaultApi.md#apiV1PlayersMostRunsGet) | **GET** /api/v1/players/most_runs | Get player who has the most_runs |


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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="apiV1PlayersMostRunsGet"></a>
# **apiV1PlayersMostRunsGet**
> MostRuns apiV1PlayersMostRunsGet(careerEndYear)

Get player who has the most_runs

    Returns a player name and his most_runs

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **careerEndYear** | **Integer**| year that the player ends his career | [optional] [default to null] |

### Return type

[**MostRuns**](../Models/MostRuns.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

