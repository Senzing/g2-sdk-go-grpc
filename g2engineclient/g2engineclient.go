/*
 *
 */

// Package g2configclient implements a client for the service.
package g2engineclient

import (
	"context"

	pb "github.com/senzing/g2-sdk-proto/go/g2engine"
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddRecord method adds a record into the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
func (client *G2engineClient) AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	request := pb.AddRecordRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		JsonData:       jsonData,
		LoadID:         loadID,
	}
	_, err := client.GrpcClient.AddRecord(ctx, &request)
	return err
}

/*
The AddRecordWithInfo method adds a record into the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	request := pb.AddRecordWithInfoRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		JsonData:       jsonData,
		LoadID:         loadID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.AddRecordWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The AddRecordWithInfoWithReturnedRecordID method adds a record into the Senzing repository and returns information on the affected entities and the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing the AFFECTED_ENTITIES, INTERESTING_ENTITIES, and RECORD_ID.
    Example: `{"DATA_SOURCE":"TEST","RECORD_ID":"2D4DABB3FAEAFBD452E9487D06FABC22DC69C846","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
  - The record identifier.
    Example: `2D4DABB3FAEAFBD452E9487D06FABC22DC69C846`
*/
func (client *G2engineClient) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	request := pb.AddRecordWithInfoWithReturnedRecordIDRequest{
		DataSourceCode: dataSourceCode,
		JsonData:       jsonData,
		LoadID:         loadID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.AddRecordWithInfoWithReturnedRecordID(ctx, &request)
	return response.GetWithInfo(), response.GetRecordID(), err
}

/*
The AddRecordWithReturnedRecordID method adds a record into the Senzing repository and returns the record identifier.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.

Output
  - The record identifier.
    Example: `2D4DABB3FAEAFBD452E9487D06FABC22DC69C846`
*/
func (client *G2engineClient) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	request := pb.AddRecordWithReturnedRecordIDRequest{
		DataSourceCode: dataSourceCode,
		JsonData:       jsonData,
		LoadID:         loadID,
	}
	response, err := client.GrpcClient.AddRecordWithReturnedRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The CheckRecord method FIXME:.

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document with the attribute data for the record to check with the "DATA_SOURCE" field.
  - recordQueryList: A JSON document with the datasource codes and recordID's of the records to check against.

Output

  - A JSON document that FIXME:
    See the example output.
*/
func (client *G2engineClient) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	request := pb.CheckRecordRequest{
		Record:          record,
		RecordQueryList: recordQueryList,
	}
	response, err := client.GrpcClient.CheckRecord(ctx, &request)
	return response.GetResult(), err
}

/*
The CloseExport method closes the exported document created by ExportJSONEntityReport().
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of sized entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport() or ExportCSVEntityReport().
*/
func (client *G2engineClient) CloseExport(ctx context.Context, responseHandle uintptr) error {
	request := pb.CloseExportRequest{
		ResponseHandle: int64(responseHandle),
	}
	_, err := client.GrpcClient.CloseExport(ctx, &request)
	return err
}

/*
The CountRedoRecords method returns the number of records in need of redo-ing.

Input
  - ctx: A context to control lifecycle.

Output
  - The number of redo records in Senzing's redo queue.
*/
func (client *G2engineClient) CountRedoRecords(ctx context.Context) (int64, error) {
	request := pb.CountRedoRecordsRequest{}
	response, err := client.GrpcClient.CountRedoRecords(ctx, &request)
	return response.GetResult(), err
}

/*
The DeleteRecord method deletes a record from the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
*/
func (client *G2engineClient) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	request := pb.DeleteRecordRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		LoadID:         loadID,
	}
	_, err := client.GrpcClient.DeleteRecord(ctx, &request)
	return err
}

/*
The DeleteRecordWithInfo method deletes a record from the Senzing repository and returns information on the affected entities.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
    FIXME: How does the "loadID" affect what is deleted?
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	request := pb.DeleteRecordWithInfoRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		LoadID:         loadID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.DeleteRecordWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2 object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engineClient) Destroy(ctx context.Context) error {
	request := pb.DestroyRequest{}
	_, err := client.GrpcClient.Destroy(ctx, &request)
	return err
}

/*
The ExportConfig method returns the Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing the current Senzing Engine configuration.
*/
func (client *G2engineClient) ExportConfig(ctx context.Context) (string, error) {
	request := pb.ExportConfigRequest{}
	response, err := client.GrpcClient.ExportConfig(ctx, &request)
	return response.GetResult(), err
}

/*
Similar to ExportConfig(), the ExportConfigAndConfigID method returns the Senzing engine configuration and it's identifier.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing the current Senzing Engine configuration.
  - The unique identifier of the Senzing Engine configuration.
*/
func (client *G2engineClient) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	request := pb.ExportConfigAndConfigIDRequest{}
	response, err := client.GrpcClient.ExportConfigAndConfigID(ctx, &request)
	return response.GetConfig(), response.GetConfigID(), err
}

/*
The ExportCSVEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - csvColumnList: A comma-separated list of column names for the CSV export.
  - flags: Flags used to control information returned.

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
func (client *G2engineClient) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	request := pb.ExportCSVEntityReportRequest{
		CsvColumnList: csvColumnList,
		Flags:         flags,
	}
	response, err := client.GrpcClient.ExportCSVEntityReport(ctx, &request)
	return uintptr(response.GetResult()), err
}

/*
The ExportJSONEntityReport method initializes a cursor over a document of exported entities.
It is part of the ExportJSONEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.“

Output
  - A handle that identifies the document to be scrolled through using FetchNext().
*/
func (client *G2engineClient) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	request := pb.ExportJSONEntityReportRequest{
		Flags: flags,
	}
	response, err := client.GrpcClient.ExportJSONEntityReport(ctx, &request)
	return uintptr(response.GetResult()), err
}

/*
The FetchNext method is used to scroll through an exported document.
It is part of the ExportJSONEntityReport() or ExportCSVEntityReport(), FetchNext(), CloseExport()
lifecycle of a list of exported entities.

Input
  - ctx: A context to control lifecycle.
  - responseHandle: A handle created by ExportJSONEntityReport() or ExportCSVEntityReport().

Output
  - FIXME:
*/
func (client *G2engineClient) FetchNext(ctx context.Context, responseHandle uintptr) (string, error) {
	request := pb.FetchNextRequest{
		ResponseHandle: int64(responseHandle),
	}
	response, err := client.GrpcClient.FetchNext(ctx, &request)
	return response.GetResult(), err
}

/*
The FindInterestingEntitiesByEntityID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	request := pb.FindInterestingEntitiesByEntityIDRequest{
		EntityID: entityID,
		Flags:    flags,
	}
	response, err := client.GrpcClient.FindInterestingEntitiesByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindInterestingEntitiesByRecordID method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	request := pb.FindInterestingEntitiesByRecordIDRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.FindInterestingEntitiesByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindNetworkByEntityID method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"SEAMAN","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-11-29 22:25:18.997","LAST_SEEN_DT":"2022-11-29 22:25:19.005"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.005"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-11-29 22:25:19.009","LAST_SEEN_DT":"2022-11-29 22:25:19.009"}],"LAST_SEEN_DT":"2022-11-29 22:25:19.009"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	request := pb.FindNetworkByEntityIDRequest{
		EntityList:     entityList,
		MaxDegree:      int32(maxDegree),
		BuildOutDegree: int32(buildOutDegree),
		MaxEntities:    int32(maxEntities),
	}
	response, err := client.GrpcClient.FindNetworkByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindNetworkByEntityID_V2 method finds all entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	request := pb.FindNetworkByEntityID_V2Request{
		EntityList:     entityList,
		MaxDegree:      int32(maxDegree),
		BuildOutDegree: int32(buildOutDegree),
		MaxEntities:    int32(maxEntities),
		Flags:          flags,
	}
	response, err := client.GrpcClient.FindNetworkByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindNetworkByRecordID method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
To control output, use FindNetworkByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:40:34.285","LAST_SEEN_DT":"2022-12-06 14:40:34.420"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.420"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:40:34.359","LAST_SEEN_DT":"2022-12-06 14:40:34.359"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.359"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":3,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:40:34.424","LAST_SEEN_DT":"2022-12-06 14:40:34.424"}],"LAST_SEEN_DT":"2022-12-06 14:40:34.424"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	request := pb.FindNetworkByRecordIDRequest{
		RecordList:     recordList,
		MaxDegree:      int32(maxDegree),
		BuildOutDegree: int32(buildOutDegree),
		MaxEntities:    int32(maxEntities),
	}
	response, err := client.GrpcClient.FindNetworkByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindNetworkByRecordID_V2 method finds all entities surrounding a requested set of entities identified by record identifiers.
This includes the requested entities, paths between them, and relations to other nearby entities.
It extends FindNetworkByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityList: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegree: The maximum number of degrees in paths between search entities.
  - buildOutDegree: The number of degrees of relationships to show around each search entity.
  - maxEntities: The maximum number of entities to return in the discovered network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	request := pb.FindNetworkByRecordID_V2Request{
		RecordList:     recordList,
		MaxDegree:      int32(maxDegree),
		BuildOutDegree: int32(buildOutDegree),
		MaxEntities:    int32(maxEntities),
		Flags:          flags,
	}
	response, err := client.GrpcClient.FindNetworkByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
To control output, use FindPathByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:43:49.024","LAST_SEEN_DT":"2022-12-06 14:43:49.164"}],"LAST_SEEN_DT":"2022-12-06 14:43:49.164"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:43:49.104","LAST_SEEN_DT":"2022-12-06 14:43:49.104"}],"LAST_SEEN_DT":"2022-12-06 14:43:49.104"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	request := pb.FindPathByEntityIDRequest{
		EntityID1: entityID1,
		EntityID2: entityID2,
		MaxDegree: int32(maxDegree),
	}
	response, err := client.GrpcClient.FindPathByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
It extends FindPathByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	request := pb.FindPathByEntityID_V2Request{
		EntityID1: entityID1,
		EntityID2: entityID2,
		MaxDegree: int32(maxDegree),
		Flags:     flags,
	}
	response, err := client.GrpcClient.FindPathByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathByRecordID method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
To control output, use FindPathByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.

Output

  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:48:19.522","LAST_SEEN_DT":"2022-12-06 14:48:19.667"}],"LAST_SEEN_DT":"2022-12-06 14:48:19.667"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:48:19.593","LAST_SEEN_DT":"2022-12-06 14:48:19.593"}],"LAST_SEEN_DT":"2022-12-06 14:48:19.593"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	request := pb.FindPathByRecordIDRequest{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
	}
	response, err := client.GrpcClient.FindPathByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathByRecordID_V2 method finds single relationship paths between two entities.
The entities are identified by starting and ending records.
Paths are found using known relationships with other entities.
It extends FindPathByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	request := pb.FindPathByRecordID_V2Request{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
		Flags:           flags,
	}
	response, err := client.GrpcClient.FindPathByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathExcludingByEntityID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:50:49.222","LAST_SEEN_DT":"2022-12-06 14:50:49.356"}],"LAST_SEEN_DT":"2022-12-06 14:50:49.356"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:50:49.295","LAST_SEEN_DT":"2022-12-06 14:50:49.295"}],"LAST_SEEN_DT":"2022-12-06 14:50:49.295"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	request := pb.FindPathExcludingByEntityIDRequest{
		EntityID1:        entityID1,
		EntityID2:        entityID2,
		MaxDegree:        int32(maxDegree),
		ExcludedEntities: excludedEntities,
	}
	response, err := client.GrpcClient.FindPathExcludingByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathExcludingByEntityID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByEntityID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	request := pb.FindPathExcludingByEntityID_V2Request{
		EntityID1:        entityID1,
		EntityID2:        entityID2,
		MaxDegree:        int32(maxDegree),
		ExcludedEntities: excludedEntities,
		Flags:            flags,
	}
	response, err := client.GrpcClient.FindPathExcludingByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathExcludingByRecordID method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
To control output, use FindPathExcludingByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[1,2]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 14:55:02.577","LAST_SEEN_DT":"2022-12-06 14:55:02.711"}],"LAST_SEEN_DT":"2022-12-06 14:55:02.711"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 14:55:02.649","LAST_SEEN_DT":"2022-12-06 14:55:02.649"}],"LAST_SEEN_DT":"2022-12-06 14:55:02.649"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	request := pb.FindPathExcludingByRecordIDRequest{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
	}
	response, err := client.GrpcClient.FindPathExcludingByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathExcludingByRecordID_V2 method finds single relationship paths between two entities.
Paths are found using known relationships with other entities.
In addition, it will find paths that exclude certain entities from being on the path.
It extends FindPathExcludingByRecordID() by adding output control flags.

When excluding entities, the user may choose to either strictly exclude the entities,
or prefer to exclude the entities but still include them if no other path is found.
By default, entities will be strictly excluded.
A "preferred exclude" may be done by specifying the G2_FIND_PATH_PREFER_EXCLUDE control flag.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	request := pb.FindPathExcludingByRecordID_V2Request{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
		Flags:           flags,
	}
	response, err := client.GrpcClient.FindPathExcludingByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathIncludingSourceByEntityID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:00:30.268","LAST_SEEN_DT":"2022-12-06 15:00:30.429"}],"LAST_SEEN_DT":"2022-12-06 15:00:30.429"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:00:30.339","LAST_SEEN_DT":"2022-12-06 15:00:30.339"}],"LAST_SEEN_DT":"2022-12-06 15:00:30.339"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error) {
	request := pb.FindPathIncludingSourceByEntityIDRequest{
		EntityID1:        entityID1,
		EntityID2:        entityID2,
		MaxDegree:        int32(maxDegree),
		ExcludedEntities: excludedEntities,
		RequiredDsrcs:    requiredDsrcs,
	}
	response, err := client.GrpcClient.FindPathIncludingSourceByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathIncludingSourceByEntityID_V2 method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedEntities: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	request := pb.FindPathIncludingSourceByEntityID_V2Request{
		EntityID1:        entityID1,
		EntityID2:        entityID2,
		MaxDegree:        int32(maxDegree),
		ExcludedEntities: excludedEntities,
		RequiredDsrcs:    requiredDsrcs,
		Flags:            flags,
	}
	response, err := client.GrpcClient.FindPathIncludingSourceByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID() and FindPathExcludingByRecordID().
To control output, use FindPathIncludingSourceByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:

Output
  - A JSON document.
    Example: `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":2,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:03:52.805","LAST_SEEN_DT":"2022-12-06 15:03:52.947"}],"LAST_SEEN_DT":"2022-12-06 15:03:52.947"},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:03:52.876","LAST_SEEN_DT":"2022-12-06 15:03:52.876"}],"LAST_SEEN_DT":"2022-12-06 15:03:52.876"},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0}]}]}`
*/
func (client *G2engineClient) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	request := pb.FindPathIncludingSourceByRecordIDRequest{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
		ExcludedRecords: excludedRecords,
		RequiredDsrcs:   requiredDsrcs,
	}
	response, err := client.GrpcClient.FindPathIncludingSourceByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The FindPathIncludingSourceByRecordID method finds single relationship paths between two entities.
In addition, one of the enties along the path must include a specified data source.
Specific entities may also be excluded,
using the same methodology as the FindPathExcludingByEntityID_V2() and FindPathExcludingByRecordID_V2().
It extends FindPathIncludingSourceByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the record for the starting entity of the search path.
  - recordID1: The unique identifier within the records of the same data source for the starting entity of the search path.
  - dataSourceCode2: Identifies the provenance of the record for the ending entity of the search path.
  - recordID2: The unique identifier within the records of the same data source for the ending entity of the search path.
  - maxDegree: The maximum number of degrees in paths between search entities.
  - excludedRecords: A JSON document listing entities that should be avoided on the path.
  - requiredDsrcs: A JSON document listing data sources that should be included on the path. FIXME:
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	request := pb.FindPathIncludingSourceByRecordID_V2Request{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		MaxDegree:       int32(maxDegree),
		ExcludedRecords: excludedRecords,
		RequiredDsrcs:   requiredDsrcs,
		Flags:           flags,
	}
	response, err := client.GrpcClient.FindPathIncludingSourceByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The GetActiveConfigID method returns the identifier of the loaded Senzing engine configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - The identifier of the active Senzing Engine configuration.
*/
func (client *G2engineClient) GetActiveConfigID(ctx context.Context) (int64, error) {
	request := pb.GetActiveConfigIDRequest{}
	response, err := client.GrpcClient.GetActiveConfigID(ctx, &request)
	return response.GetResult(), err
}

/*
The GetEntityByEntityID method returns entity data based on the ID of a resolved identity.
To control output, use GetEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.

Output

  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:09:48.577","LAST_SEEN_DT":"2022-12-06 15:09:48.705"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.705","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:09:48.577"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:09:48.705"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:09:48.647","LAST_SEEN_DT":"2022-12-06 15:09:48.647"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.647"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:09:48.709","LAST_SEEN_DT":"2022-12-06 15:09:48.709"}],"LAST_SEEN_DT":"2022-12-06 15:09:48.709"}]}`
*/
func (client *G2engineClient) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	request := pb.GetEntityByEntityIDRequest{
		EntityID: entityID,
	}
	response, err := client.GrpcClient.GetEntityByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The GetEntityByEntityID_V2 method returns entity data based on the ID of a resolved identity.
It extends GetEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	request := pb.GetEntityByEntityID_V2Request{
		EntityID: entityID,
		Flags:    flags,
	}
	response, err := client.GrpcClient.GetEntityByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The GetEntityByRecordID method returns entity data based on the ID of a record which is a member of the entity.
To control output, use GetEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:12:25.464","LAST_SEEN_DT":"2022-12-06 15:12:25.597"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.597","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:12:25.464"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+EXACTLY_SAME","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:12:25.597"}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:12:25.536","LAST_SEEN_DT":"2022-12-06 15:12:25.536"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.536"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:12:25.603","LAST_SEEN_DT":"2022-12-06 15:12:25.603"}],"LAST_SEEN_DT":"2022-12-06 15:12:25.603"}]}`
*/
func (client *G2engineClient) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	request := pb.GetEntityByRecordIDRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
	}
	response, err := client.GrpcClient.GetEntityByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The GetEntityByRecordID_V2 method returns entity data based on the ID of a record which is a member of the entity.
It extends GetEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	request := pb.GetEntityByRecordID_V2Request{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.GetEntityByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The GetRecord method returns a JSON document of a single record from the Senzing repository.
To control output, use GetRecord_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	request := pb.GetRecordRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
	}
	response, err := client.GrpcClient.GetRecord(ctx, &request)
	return response.GetResult(), err
}

/*
The GetRecord_V2 method returns a JSON document of a single record from the Senzing repository.
It extends GetRecord() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	request := pb.GetRecord_V2Request{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.GetRecord_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The GetRedoRecord method returns the next internally queued maintenance record from the Senzing repository.
Usually, the ProcessRedoRecord() or ProcessRedoRecordWithInfo() method is called to process the maintenance record
retrieved by GetRedoRecord().

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *G2engineClient) GetRedoRecord(ctx context.Context) (string, error) {
	request := pb.GetRedoRecordRequest{}
	response, err := client.GrpcClient.GetRedoRecord(ctx, &request)
	return response.GetResult(), err
}

/*
The GetRepositoryLastModifiedTime method retrieves the last modified time of the Senzing repository,
measured in the number of seconds between the last modified time and January 1, 1970 12:00am GMT (epoch time).

Input
  - ctx: A context to control lifecycle.

Output
  - A Unix Timestamp.
*/
func (client *G2engineClient) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	request := pb.GetRepositoryLastModifiedTimeRequest{}
	response, err := client.GrpcClient.GetRepositoryLastModifiedTime(ctx, &request)
	return response.GetResult(), err
}

/*
The GetVirtualEntityByRecordID method FIXME:
To control output, use GetVirtualEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":2,"FIRST_SEEN_DT":"2022-12-06 15:20:17.088","LAST_SEEN_DT":"2022-12-06 15:20:17.161"}],"LAST_SEEN_DT":"2022-12-06 15:20:17.161","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","LAST_SEEN_DT":"2022-12-06 15:20:17.088","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","LAST_SEEN_DT":"2022-12-06 15:20:17.161","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]}}`
*/
func (client *G2engineClient) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error) {
	request := pb.GetVirtualEntityByRecordIDRequest{
		RecordList: recordList,
	}
	response, err := client.GrpcClient.GetVirtualEntityByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The GetVirtualEntityByRecordID_V2 method FIXME:
It extends GetVirtualEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - recordList: A JSON document.
    Example: `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "111"},{"DATA_SOURCE": "TEST","RECORD_ID": "222"}]}`
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	request := pb.GetVirtualEntityByRecordID_V2Request{
		RecordList: recordList,
		Flags:      flags,
	}
	response, err := client.GrpcClient.GetVirtualEntityByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The HowEntityByEntityID method FIXME:
To control output, use HowEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	request := pb.HowEntityByEntityIDRequest{
		EntityID: entityID,
	}
	response, err := client.GrpcClient.HowEntityByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The HowEntityByEntityID_V2 method FIXME:
It extends HowEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	request := pb.HowEntityByEntityID_V2Request{
		EntityID: entityID,
		Flags:    flags,
	}
	response, err := client.GrpcClient.HowEntityByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The Init method initializes the Senzing G2 object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2engineClient) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	request := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.GrpcClient.Init(ctx, &request)
	return err
}

/*
The InitWithConfigID method initializes the Senzing G2 object with a non-default configuration ID.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - initConfigID: The configuration ID used for the initialization.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2engineClient) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	request := pb.InitWithConfigIDRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		InitConfigID:   initConfigID,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.GrpcClient.InitWithConfigID(ctx, &request)
	return err
}

/*
The PrimeEngine method pre-initializes some of the heavier weight internal resources of the G2 engine.
The G2 Engine uses "lazy initialization".
PrimeEngine() forces initialization.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engineClient) PrimeEngine(ctx context.Context) error {
	request := pb.PrimeEngineRequest{}
	_, err := client.GrpcClient.PrimeEngine(ctx, &request)
	return err
}

/*
The Process method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
*/
func (client *G2engineClient) Process(ctx context.Context, record string) error {
	request := pb.ProcessRequest{}
	_, err := client.GrpcClient.Process(ctx, &request)
	return err
}

/*
The ProcessRedoRecord method processes the next redo record and returns it.
Calling ProcessRedoRecord() has the potential to create more redo records in certain situations.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *G2engineClient) ProcessRedoRecord(ctx context.Context) (string, error) {
	request := pb.ProcessRedoRecordRequest{}
	response, err := client.GrpcClient.ProcessRedoRecord(ctx, &request)
	return response.GetResult(), err
}

/*
The ProcessRedoRecordWithInfo method processes the next redo record and returns it and affected entities.
Calling ProcessRedoRecordWithInfo() has the potential to create more redo records in certain situations.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.

Output
  - A JSON document with the record that was re-done.
  - A JSON document with affected entities.
*/
func (client *G2engineClient) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	request := pb.ProcessRedoRecordWithInfoRequest{
		Flags: flags,
	}
	response, err := client.GrpcClient.ProcessRedoRecordWithInfo(ctx, &request)
	return response.GetResult(), response.GetWithInfo(), err
}

/*
The ProcessWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	request := pb.ProcessWithInfoRequest{
		Record: record,
		Flags:  flags,
	}
	response, err := client.GrpcClient.ProcessWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The ProcessWithResponse method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	request := pb.ProcessWithResponseRequest{
		Record: record,
	}
	response, err := client.GrpcClient.ProcessWithResponse(ctx, &request)
	return response.GetResult(), err
}

/*
The ProcessWithResponseResize method FIXME:

Input
  - ctx: A context to control lifecycle.
  - record: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	request := pb.ProcessWithResponseResizeRequest{
		Record: record,
	}
	response, err := client.GrpcClient.ProcessWithResponseResize(ctx, &request)
	return response.GetResult(), err
}

/*
The PurgeRepository method removes every record in the Senzing repository.

Before calling purgeRepository() all other instances of the Senzing API
(whether in custom code, REST API, stream-loader, redoer, G2Loader, etc)
MUST be destroyed or shutdown.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2engineClient) PurgeRepository(ctx context.Context) error {
	request := pb.PurgeRepositoryRequest{}
	_, err := client.GrpcClient.PurgeRepository(ctx, &request)
	return err
}

/*
The ReevaluateEntity method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.
*/
func (client *G2engineClient) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	request := pb.ReevaluateEntityRequest{
		EntityID: entityID,
		Flags:    flags,
	}
	_, err := client.GrpcClient.ReevaluateEntity(ctx, &request)
	return err
}

/*
The ReevaluateEntityWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output

  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	request := pb.ReevaluateEntityWithInfoRequest{
		EntityID: entityID,
		Flags:    flags,
	}
	response, err := client.GrpcClient.ReevaluateEntityWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The ReevaluateRecord method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.
*/
func (client *G2engineClient) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	request := pb.ReevaluateRecordRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	_, err := client.GrpcClient.ReevaluateRecord(ctx, &request)
	return err
}

/*
The ReevaluateRecordWithInfo method FIXME:

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output

  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	request := pb.ReevaluateRecordWithInfoRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.ReevaluateRecordWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The Reinit method re-initializes the Senzing G2Engine object using a specified configuration identifier.

Input
  - ctx: A context to control lifecycle.
  - initConfigID: The configuration ID used for the initialization.
*/
func (client *G2engineClient) Reinit(ctx context.Context, initConfigID int64) error {
	request := pb.ReinitRequest{
		InitConfigID: initConfigID,
	}
	_, err := client.GrpcClient.Reinit(ctx, &request)
	return err
}

/*
The ReplaceRecord method updates/replaces a record in the Senzing repository.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
*/
func (client *G2engineClient) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	request := pb.ReplaceRecordRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		JsonData:       jsonData,
		LoadID:         loadID,
	}
	_, err := client.GrpcClient.ReplaceRecord(ctx, &request)
	return err
}

/*
The ReplaceRecordWithInfo method updates/replaces a record in the Senzing repository and returns information on the affected entities.
If record doesn't exist, a new record is added to the data repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - loadID: An identifier used to distinguish different load batches/sessions. An empty string is acceptable.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	request := pb.ReplaceRecordWithInfoRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		JsonData:       jsonData,
		LoadID:         loadID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.ReplaceRecordWithInfo(ctx, &request)
	return response.GetResult(), err
}

/*
The SearchByAttributes method retrieves entity data based on a user-specified set of entity attributes.
To control output, use SearchByAttributes_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.

Output
  - A JSON document.
    Example: `{"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","MATCH_KEY":"+NAME+SSN","ERRULE_CODE":"SF1_PNAME_CSTAB","FEATURE_SCORES":{"NAME":[{"INBOUND_FEAT":"JOHNSON","CANDIDATE_FEAT":"JOHNSON","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1}],"SSN":[{"INBOUND_FEAT":"053-39-3251","CANDIDATE_FEAT":"053-39-3251","FULL_SCORE":100}]}},"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:38:06.175","LAST_SEEN_DT":"2022-12-06 15:38:06.957"}],"LAST_SEEN_DT":"2022-12-06 15:38:06.957"}}}]}`
*/
func (client *G2engineClient) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	request := pb.SearchByAttributesRequest{
		JsonData: jsonData,
	}
	response, err := client.GrpcClient.SearchByAttributes(ctx, &request)
	return response.GetResult(), err
}

/*
The SearchByAttributes_V2 method retrieves entity data based on a user-specified set of entity attributes.
It extends SearchByAttributes() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - jsonData: A JSON document containing the record to be added to the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	request := pb.SearchByAttributes_V2Request{
		JsonData: jsonData,
		Flags:    flags,
	}
	response, err := client.GrpcClient.SearchByAttributes_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *G2engineClient) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	var err error = nil
	return err
}

/*
The Stats method retrieves workload statistics for the current process.
These statistics will automatically reset after retrieval.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
    Example: `{"workload":{"loadedRecords":5,"addedRecords":2,"deletedRecords":0,"reevaluations":0,"repairedEntities":0,"duration":56,"retries":0,"candidates":19,"actualAmbiguousTest":0,"cachedAmbiguousTest":0,"libFeatCacheHit":219,"libFeatCacheMiss":73,"unresolveTest":1,"abortedUnresolve":0,"gnrScorersUsed":1,"unresolveTriggers":{"normalResolve":0,"update":0,"relLink":0,"extensiveResolve":0,"ambiguousNoResolve":1,"ambiguousMultiResolve":0},"reresolveTriggers":{"abortRetry":0,"unresolveMovement":0,"multipleResolvableCandidates":0,"resolveNewFeatures":1,"newFeatureFTypes":[{"DOB":1}]},"reresolveSkipped":0,"filteredObsFeat":0,"expressedFeatureCalls":[{"EFCALL_ID":1,"EFUNC_CODE":"PHONE_HASHER","numCalls":1},{"EFCALL_ID":2,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":3,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":5,"EFUNC_CODE":"EXPRESS_BOM","numCalls":1},{"EFCALL_ID":7,"EFUNC_CODE":"NAME_HASHER","numCalls":4},{"EFCALL_ID":9,"EFUNC_CODE":"ADDR_HASHER","numCalls":1},{"EFCALL_ID":10,"EFUNC_CODE":"EXPRESS_BOM","numCalls":1},{"EFCALL_ID":14,"EFUNC_CODE":"EXPRESS_ID","numCalls":1},{"EFCALL_ID":16,"EFUNC_CODE":"EXPRESS_ID","numCalls":4}],"expressedFeaturesCreated":[{"ADDR_KEY":2},{"ID_KEY":7},{"NAME_KEY":14},{"PHONE_KEY":1},{"SEARCH_KEY":2}],"scoredPairs":[{"ACCT_NUM":16},{"ADDRESS":16},{"DOB":25},{"GENDER":16},{"LOGIN_ID":16},{"NAME":19},{"PHONE":16},{"SSN":19}],"cacheHit":[{"ADDRESS":12},{"DOB":18},{"NAME":13},{"PHONE":15}],"cacheMiss":[{"ADDRESS":4},{"DOB":7},{"NAME":6},{"PHONE":1}],"redoTriggers":[],"latchContention":[],"highContentionFeat":[],"highContentionResEnt":[],"genericDetect":[],"candidateBuilders":[{"ACCT_NUM":7},{"ADDR_KEY":7},{"DOB":7},{"ID_KEY":9},{"LOGIN_ID":7},{"NAME_KEY":9},{"PHONE":7},{"PHONE_KEY":7},{"SEARCH_KEY":7},{"SSN":9}],"suppressedCandidateBuilders":[],"suppressedScoredFeatureType":[],"reducedScoredFeatureType":[],"suppressedDisclosedRelationshipDomainCount":0,"CorruptEntityTestDiagnosis":{},"threadState":{"active":0,"idle":4,"sqlExecuting":0,"loader":0,"resolver":0,"scoring":0,"dataLatchContention":0,"obsEntContention":0,"resEntContention":0},"systemResources":{"initResources":[{"physicalCores":16},{"logicalCores":16},{"totalMemory":"62.6GB"},{"availableMemory":"49.5GB"}],"currResources":[{"availableMemory":"47.4GB"},{"activeThreads":0},{"workerThreads":4},{"systemLoad":[{"cpuUser":13.442277},{"cpuSystem":2.635741},{"cpuIdle":82.024246},{"cpuWait":1.634159},{"cpuSoftIrq":0.263574}]}]}}}`
*/
func (client *G2engineClient) Stats(ctx context.Context) (string, error) {
	request := pb.StatsRequest{}
	response, err := client.GrpcClient.Stats(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntities method explains why records belong to their resolved entities.
WhyEntities() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
To control output, use WhyEntities_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":2,"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}]},"DISCLOSED_RELATIONS":{},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":26,"CANDIDATE_FEAT":"772 Armstrong RD Delhi WI 53543","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":81,"SCORE_BUCKET":"LIKELY","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":79,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"FMES"},{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":86,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":28,"CANDIDATE_FEAT":"flavorh2","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":24,"CANDIDATE_FEAT":"OCEANGUY","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":33,"GNR_SN":32,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":27,"CANDIDATE_FEAT":"153-33-5185","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:58:57.129","LAST_SEEN_DT":"2022-12-06 15:58:57.906"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.906","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:58:57.906","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.400","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.404","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.407","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.410","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 15:58:57.259","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.201","LAST_SEEN_DT":"2022-12-06 15:58:57.201"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.201"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.263","LAST_SEEN_DT":"2022-12-06 15:58:57.263"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.263"}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.201","LAST_SEEN_DT":"2022-12-06 15:58:57.201"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.201","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 15:58:57.201","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 15:58:57.129","LAST_SEEN_DT":"2022-12-06 15:58:57.906"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.906"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 15:58:57.263","LAST_SEEN_DT":"2022-12-06 15:58:57.263"}],"LAST_SEEN_DT":"2022-12-06 15:58:57.263"}]}]}`
*/
func (client *G2engineClient) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	request := pb.WhyEntitiesRequest{
		EntityID1: entityID1,
		EntityID2: entityID2,
	}
	response, err := client.GrpcClient.WhyEntities(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntities_V2 method explains why records belong to their resolved entities.
WhyEntities_V2() will compare the record data within an entity
against the rest of the entity data and show why they are connected.
This is calculated based on the features that record data represents.
It extends WhyEntities() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The entity ID for the starting entity of the search path.
  - entityID2: The entity ID for the ending entity of the search path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	request := pb.WhyEntities_V2Request{
		EntityID1: entityID1,
		EntityID2: entityID2,
		Flags:     flags,
	}
	response, err := client.GrpcClient.WhyEntities_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntityByEntityID method explains why records belong to their resolved entities.
To control output, use WhyEntityByEntityID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity for the starting entity of the search path.

Output

  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":100001,"CANDIDATE_FEAT":"4/8/1985","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"4/8/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:02:35.306","LAST_SEEN_DT":"2022-12-06 16:02:36.083"}],"LAST_SEEN_DT":"2022-12-06 16:02:36.083","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:02:36.083","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.572","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.575","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.579","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.582","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:02:35.432","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:02:35.373","LAST_SEEN_DT":"2022-12-06 16:02:35.373"}],"LAST_SEEN_DT":"2022-12-06 16:02:35.373"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:02:35.436","LAST_SEEN_DT":"2022-12-06 16:02:35.436"}],"LAST_SEEN_DT":"2022-12-06 16:02:35.436"}]}]}`
*/
func (client *G2engineClient) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	request := pb.WhyEntityByEntityIDRequest{
		EntityID: entityID,
	}
	response, err := client.GrpcClient.WhyEntityByEntityID(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntityByEntityID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByEntityID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity for the starting entity of the search path.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
*/
func (client *G2engineClient) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	request := pb.WhyEntityByEntityID_V2Request{
		EntityID: entityID,
		Flags:    flags,
	}
	response, err := client.GrpcClient.WhyEntityByEntityID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntityByRecordID method explains why records belong to their resolved entities.
To control output, use WhyEntityByRecordID_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"4/8/1983","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":100001,"CANDIDATE_FEAT":"4/8/1985","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"},{"FEAT_ID":18,"FEAT_DESC":"772|ARMSTRNK||71232"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"},{"FEAT_ID":20,"FEAT_DESC":"SSN=053-39-3251"}],"LOGIN_ID":[{"FEAT_ID":7,"FEAT_DESC":"flavorh"}],"NAME_KEY":[{"FEAT_ID":9,"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804"},{"FEAT_ID":11,"FEAT_DESC":"JNSN"},{"FEAT_ID":12,"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL"},{"FEAT_ID":13,"FEAT_DESC":"JNSN|DOB=80804"},{"FEAT_ID":14,"FEAT_DESC":"JNSN|POST=71232"},{"FEAT_ID":15,"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796"},{"FEAT_ID":16,"FEAT_DESC":"JNSN|SSN=3251"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}],"SEARCH_KEY":[{"FEAT_ID":22,"FEAT_DESC":"LOGIN_ID:FLAVORH|"},{"FEAT_ID":23,"FEAT_DESC":"SSN:3251|80804|"}],"SSN":[{"FEAT_ID":6,"FEAT_DESC":"053-39-3251"}]},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"772 Armstrong RD Delhi LA 71232","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"4/8/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":88,"SCORE_BUCKET":"PLAUSIBLE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":7,"CANDIDATE_FEAT":"flavorh","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"JOHNSON","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":100,"GNR_SN":100,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":6,"CANDIDATE_FEAT":"053-39-3251","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:09:05.626","LAST_SEEN_DT":"2022-12-06 16:09:06.399"}],"LAST_SEEN_DT":"2022-12-06 16:09:06.399","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:09:06.399","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.954","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.957","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.960","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.963","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:09:05.789","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:09:05.724","LAST_SEEN_DT":"2022-12-06 16:09:05.724"}],"LAST_SEEN_DT":"2022-12-06 16:09:05.724"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:09:05.794","LAST_SEEN_DT":"2022-12-06 16:09:05.794"}],"LAST_SEEN_DT":"2022-12-06 16:09:05.794"}]}]}`
*/
func (client *G2engineClient) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	request := pb.WhyEntityByRecordIDRequest{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
	}
	response, err := client.GrpcClient.WhyEntityByRecordID(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyEntityByRecordID_V2 method explains why records belong to their resolved entities.
It extends WhyEntityByRecordID() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"444"},{"DATA_SOURCE":"TEST","RECORD_ID":"555"},{"DATA_SOURCE":"TEST","RECORD_ID":"666"},{"DATA_SOURCE":"TEST","RECORD_ID":"777"},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}},{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"MATCH_INFO":{"WHY_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","WHY_ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
*/
func (client *G2engineClient) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	request := pb.WhyEntityByRecordID_V2Request{
		DataSourceCode: dataSourceCode,
		RecordID:       recordID,
		Flags:          flags,
	}
	response, err := client.GrpcClient.WhyEntityByRecordID_V2(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyRecords method explains why records belong to their resolved entities.
To control output, use WhyRecords_V2() instead.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.

Output

  - A JSON document.
    Example: `{"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":2,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"TEST","RECORD_ID":"222"}],"MATCH_INFO":{"WHY_KEY":"+PHONE+ACCT_NUM-DOB-SSN","WHY_ERRULE_CODE":"SF1","MATCH_LEVEL_CODE":"POSSIBLY_RELATED","CANDIDATE_KEYS":{"ACCT_NUM":[{"FEAT_ID":8,"FEAT_DESC":"5534202208773608"}],"ADDR_KEY":[{"FEAT_ID":17,"FEAT_DESC":"772|ARMSTRNK||TL"}],"ID_KEY":[{"FEAT_ID":19,"FEAT_DESC":"ACCT_NUM=5534202208773608"}],"PHONE":[{"FEAT_ID":5,"FEAT_DESC":"225-671-0796"}],"PHONE_KEY":[{"FEAT_ID":21,"FEAT_DESC":"2256710796"}]},"DISCLOSED_RELATIONS":{},"FEATURE_SCORES":{"ACCT_NUM":[{"INBOUND_FEAT_ID":8,"INBOUND_FEAT":"5534202208773608","INBOUND_FEAT_USAGE_TYPE":"CC","CANDIDATE_FEAT_ID":8,"CANDIDATE_FEAT":"5534202208773608","CANDIDATE_FEAT_USAGE_TYPE":"CC","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"ADDRESS":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"772 Armstrong RD Delhi LA 71232","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":26,"CANDIDATE_FEAT":"772 Armstrong RD Delhi WI 53543","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":81,"SCORE_BUCKET":"LIKELY","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":100001,"INBOUND_FEAT":"4/8/1985","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":25,"CANDIDATE_FEAT":"6/9/1983","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":79,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"FMES"}],"GENDER":[{"INBOUND_FEAT_ID":3,"INBOUND_FEAT":"F","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"F","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}],"LOGIN_ID":[{"INBOUND_FEAT_ID":7,"INBOUND_FEAT":"flavorh","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":28,"CANDIDATE_FEAT":"flavorh2","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"JOHNSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":24,"CANDIDATE_FEAT":"OCEANGUY","CANDIDATE_FEAT_USAGE_TYPE":"","GNR_FN":33,"GNR_SN":32,"GNR_GN":70,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"225-671-0796","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"225-671-0796","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"SSN":[{"INBOUND_FEAT_ID":6,"INBOUND_FEAT":"053-39-3251","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":27,"CANDIDATE_FEAT":"153-33-5185","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":0,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"F1ES"}]}}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"JOHNSON","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi LA 71232","LIB_FEAT_ID":4,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||71232","LIB_FEAT_ID":18,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1983","LIB_FEAT_ID":2,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"FEAT_DESC_VALUES":[{"FEAT_DESC":"4/8/1985","LIB_FEAT_ID":100001,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=053-39-3251","LIB_FEAT_ID":20,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh","LIB_FEAT_ID":7,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JOHNSON","LIB_FEAT_ID":1,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN","LIB_FEAT_ID":11,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":12,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMDD_HASH=0804","LIB_FEAT_ID":9,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0483","LIB_FEAT_ID":10,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB.MMYY_HASH=0485","LIB_FEAT_ID":100002,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|DOB=80804","LIB_FEAT_ID":13,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":15,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|POST=71232","LIB_FEAT_ID":14,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"JNSN|SSN=3251","LIB_FEAT_ID":16,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH|","LIB_FEAT_ID":22,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:3251|80804|","LIB_FEAT_ID":23,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"FEAT_DESC_VALUES":[{"FEAT_DESC":"053-39-3251","LIB_FEAT_ID":6,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:13:27.135","LAST_SEEN_DT":"2022-12-06 16:13:27.916"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.916","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"111","ENTITY_TYPE":"TEST","INTERNAL_ID":100001,"ENTITY_KEY":"A6C927986DF7329D1D2CDE0E8F34328AE640FB7E","ENTITY_DESC":"JOHNSON","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:13:27.916","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23},{"LIB_FEAT_ID":100001},{"LIB_FEAT_ID":100002}]},{"DATA_SOURCE":"TEST","RECORD_ID":"444","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.405","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"555","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.408","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"666","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.411","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"777","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.418","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]},{"DATA_SOURCE":"TEST","RECORD_ID":"FCCE9793DAAD23159DBCCEB97FF2745B92CE7919","ENTITY_TYPE":"TEST","INTERNAL_ID":1,"ENTITY_KEY":"C6063D4396612FBA7324DB0739273BA1FE815C43","ENTITY_DESC":"JOHNSON","MATCH_KEY":"+NAME+ADDRESS+PHONE+SSN+LOGIN_ID+ACCT_NUM","MATCH_LEVEL":1,"MATCH_LEVEL_CODE":"RESOLVED","ERRULE_CODE":"SF1_PNAME_CFF_CSTAB","LAST_SEEN_DT":"2022-12-06 16:13:27.265","FEATURES":[{"LIB_FEAT_ID":1},{"LIB_FEAT_ID":2},{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":4},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":6},{"LIB_FEAT_ID":7},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":9},{"LIB_FEAT_ID":10},{"LIB_FEAT_ID":11},{"LIB_FEAT_ID":12},{"LIB_FEAT_ID":13},{"LIB_FEAT_ID":14},{"LIB_FEAT_ID":15},{"LIB_FEAT_ID":16},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":18},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":20},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":22},{"LIB_FEAT_ID":23}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":2,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"OCEANGUY","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.208","LAST_SEEN_DT":"2022-12-06 16:13:27.208"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.208"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.272","LAST_SEEN_DT":"2022-12-06 16:13:27.272"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.272"}]},{"RESOLVED_ENTITY":{"ENTITY_ID":2,"ENTITY_NAME":"OCEANGUY","FEATURES":{"ACCT_NUM":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USAGE_TYPE":"CC","FEAT_DESC_VALUES":[{"FEAT_DESC":"5534202208773608","LIB_FEAT_ID":8,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDRESS":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772 Armstrong RD Delhi WI 53543","LIB_FEAT_ID":26,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ADDR_KEY":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||53543","LIB_FEAT_ID":37,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"FEAT_DESC_VALUES":[{"FEAT_DESC":"772|ARMSTRNK||TL","LIB_FEAT_ID":17,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"DOB":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"FEAT_DESC_VALUES":[{"FEAT_DESC":"6/9/1983","LIB_FEAT_ID":25,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"GENDER":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"FEAT_DESC_VALUES":[{"FEAT_DESC":"F","LIB_FEAT_ID":3,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"ID_KEY":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ACCT_NUM=5534202208773608","LIB_FEAT_ID":19,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN=153-33-5185","LIB_FEAT_ID":38,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"LOGIN_ID":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"FEAT_DESC_VALUES":[{"FEAT_DESC":"flavorh2","LIB_FEAT_ID":28,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"FEAT_DESC_VALUES":[{"FEAT_DESC":"OCEANGUY","LIB_FEAT_ID":24,"USED_FOR_CAND":"N","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"NAME_KEY":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK","LIB_FEAT_ID":29,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|ADDRESS.CITY_STD=TL","LIB_FEAT_ID":34,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMDD_HASH=0906","LIB_FEAT_ID":32,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB.MMYY_HASH=0683","LIB_FEAT_ID":30,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|DOB=80906","LIB_FEAT_ID":31,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|PHONE.PHONE_LAST_5=10796","LIB_FEAT_ID":33,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|POST=53543","LIB_FEAT_ID":36,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"FEAT_DESC_VALUES":[{"FEAT_DESC":"ASNK|SSN=5185","LIB_FEAT_ID":35,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"225-671-0796","LIB_FEAT_ID":5,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"PHONE_KEY":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"FEAT_DESC_VALUES":[{"FEAT_DESC":"2256710796","LIB_FEAT_ID":21,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":3,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SEARCH_KEY":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"FEAT_DESC_VALUES":[{"FEAT_DESC":"LOGIN_ID:FLAVORH2|","LIB_FEAT_ID":40,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]},{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"FEAT_DESC_VALUES":[{"FEAT_DESC":"SSN:5185|80906|","LIB_FEAT_ID":39,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"N","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}],"SSN":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"FEAT_DESC_VALUES":[{"FEAT_DESC":"153-33-5185","LIB_FEAT_ID":27,"USED_FOR_CAND":"Y","USED_FOR_SCORING":"Y","ENTITY_COUNT":1,"CANDIDATE_CAP_REACHED":"N","SCORING_CAP_REACHED":"N","SUPPRESSED":"N"}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.208","LAST_SEEN_DT":"2022-12-06 16:13:27.208"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.208","RECORDS":[{"DATA_SOURCE":"TEST","RECORD_ID":"222","ENTITY_TYPE":"TEST","INTERNAL_ID":2,"ENTITY_KEY":"740BA22D15CA88462A930AF8A7C904FF5E48226C","ENTITY_DESC":"OCEANGUY","MATCH_KEY":"","MATCH_LEVEL":0,"MATCH_LEVEL_CODE":"","ERRULE_CODE":"","LAST_SEEN_DT":"2022-12-06 16:13:27.208","FEATURES":[{"LIB_FEAT_ID":3},{"LIB_FEAT_ID":5},{"LIB_FEAT_ID":8,"USAGE_TYPE":"CC"},{"LIB_FEAT_ID":17},{"LIB_FEAT_ID":19},{"LIB_FEAT_ID":21},{"LIB_FEAT_ID":24},{"LIB_FEAT_ID":25},{"LIB_FEAT_ID":26},{"LIB_FEAT_ID":27},{"LIB_FEAT_ID":28},{"LIB_FEAT_ID":29},{"LIB_FEAT_ID":30},{"LIB_FEAT_ID":31},{"LIB_FEAT_ID":32},{"LIB_FEAT_ID":33},{"LIB_FEAT_ID":34},{"LIB_FEAT_ID":35},{"LIB_FEAT_ID":36},{"LIB_FEAT_ID":37},{"LIB_FEAT_ID":38},{"LIB_FEAT_ID":39},{"LIB_FEAT_ID":40}]}]},"RELATED_ENTITIES":[{"ENTITY_ID":1,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PHONE+ACCT_NUM-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"JOHNSON","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":6,"FIRST_SEEN_DT":"2022-12-06 16:13:27.135","LAST_SEEN_DT":"2022-12-06 16:13:27.916"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.916"},{"ENTITY_ID":3,"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+ADDRESS+PHONE+ACCT_NUM-DOB-SSN","ERRULE_CODE":"SF1","IS_DISCLOSED":0,"IS_AMBIGUOUS":0,"ENTITY_NAME":"Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"TEST","RECORD_COUNT":1,"FIRST_SEEN_DT":"2022-12-06 16:13:27.272","LAST_SEEN_DT":"2022-12-06 16:13:27.272"}],"LAST_SEEN_DT":"2022-12-06 16:13:27.272"}]}]}`
*/
func (client *G2engineClient) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	request := pb.WhyRecordsRequest{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
	}
	response, err := client.GrpcClient.WhyRecords(ctx, &request)
	return response.GetResult(), err
}

/*
The WhyRecords_V2 method explains why records belong to their resolved entities.
It extends WhyRecords() by adding output control flags.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
    See the example output.
*/
func (client *G2engineClient) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	request := pb.WhyRecords_V2Request{
		DataSourceCode1: dataSourceCode1,
		RecordID1:       recordID1,
		DataSourceCode2: dataSourceCode2,
		RecordID2:       recordID2,
		Flags:           flags,
	}
	response, err := client.GrpcClient.WhyRecords_V2(ctx, &request)
	return response.GetResult(), err
}