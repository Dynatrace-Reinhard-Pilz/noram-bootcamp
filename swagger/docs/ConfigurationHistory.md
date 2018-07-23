# ConfigurationHistory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [***ConfigurationHistoryMetadata**](ConfigurationHistoryMetadata.md) | Metadata useful for debugging | [optional] [default to null]
**From** | **int64** | Start timestamp of the report timeframe, in UTC milliseconds. | [optional] [default to null]
**To** | **int64** | End timestamp of the report timeframe, in UTC milliseconds. | [optional] [default to null]
**ResultsTruncated** | **bool** | If true, the queried timespan contained more results than could be delivered at once so only the most recent results were delivered, older ones truncated. | [optional] [default to null]
**Results** | [**[]ConfigurationInstanceDto**](ConfigurationInstanceDto.md) | List of all configuration changes during the specified timeframe. | [optional] [default to null]
**TimestampOfFirstTruncatedResult** | **int64** | If the result limit was exceeded, the timestamp of the first result that was truncated. Use this as your query&#39;s \&quot;to\&quot; to get the next batch of results. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


