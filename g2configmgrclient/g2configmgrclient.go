/*
 *
 */

// Package g2configmgrclient implements a client for the service.
package g2configmgrclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	pb "github.com/senzing/g2-sdk-proto/go/g2configmgr"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

// Get the Logger singleton.
func (client *G2configmgrClient) getLogger() messagelogger.MessageLoggerInterface {
	if client.logger == nil {
		client.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return client.logger
}

func (client *G2configmgrClient) notify(ctx context.Context, messageId int, err error, details map[string]string) {
	now := time.Now()
	details["subjectId"] = strconv.Itoa(ProductId)
	details["messageId"] = strconv.Itoa(messageId)
	details["messageTime"] = strconv.FormatInt(now.UnixNano(), 10)
	if err != nil {
		details["error"] = err.Error()
	}
	message, err := json.Marshal(details)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		client.observers.NotifyObservers(ctx, string(message))
	}
}

// Trace method entry.
func (client *G2configmgrClient) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *G2configmgrClient) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddConfig method adds a Senzing configuration JSON document to the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configStr: The Senzing configuration JSON document.
  - configComments: A free-form string of comments describing the configuration document.

Output
  - A configuration identifier.
*/
func (client *G2configmgrClient) AddConfig(ctx context.Context, configStr string, configComments string) (int64, error) {
	if client.isTrace {
		client.traceEntry(1, configStr, configComments)
	}
	entryTime := time.Now()
	request := pb.AddConfigRequest{
		ConfigStr:      configStr,
		ConfigComments: configComments,
	}
	response, err := client.GrpcClient.AddConfig(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configComments": configComments,
			}
			client.notify(ctx, 8001, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(2, configStr, configComments, response.GetResult(), err, time.Since(entryTime))
	}
	return response.GetResult(), err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2ConfigMgr object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2configmgrClient) Destroy(ctx context.Context) error {
	if client.isTrace {
		client.traceEntry(5)
	}
	entryTime := time.Now()
	request := pb.DestroyRequest{}
	_, err := client.GrpcClient.Destroy(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8002, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(6, err, time.Since(entryTime))
	}
	return err
}

/*
The GetConfig method retrieves a specific Senzing configuration JSON document from the Senzing database.

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the desired Senzing Engine configuration JSON document to retrieve.

Output
  - A JSON document containing the Senzing configuration.
    See the example output.
*/
func (client *G2configmgrClient) GetConfig(ctx context.Context, configID int64) (string, error) {
	if client.isTrace {
		client.traceEntry(7, configID)
	}
	entryTime := time.Now()
	request := pb.GetConfigRequest{
		ConfigID: configID,
	}
	response, err := client.GrpcClient.GetConfig(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8003, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(8, configID, response.GetResult(), err, time.Since(entryTime))
	}
	return response.GetResult(), err
}

/*
The GetConfigList method retrieves a list of Senzing configurations from the Senzing database.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document containing Senzing configurations.
    See the example output.
*/
func (client *G2configmgrClient) GetConfigList(ctx context.Context) (string, error) {
	if client.isTrace {
		client.traceEntry(9)
	}
	entryTime := time.Now()
	request := pb.GetConfigListRequest{}
	response, err := client.GrpcClient.GetConfigList(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8004, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(10, response.GetResult(), err, time.Since(entryTime))
	}
	return response.GetResult(), err
}

/*
The GetDefaultConfigID method retrieves from the Senzing database the configuration identifier of the default Senzing configuration.

Input
  - ctx: A context to control lifecycle.

Output
  - A configuration identifier which identifies the current configuration in use.
*/
func (client *G2configmgrClient) GetDefaultConfigID(ctx context.Context) (int64, error) {
	if client.isTrace {
		client.traceEntry(11)
	}
	entryTime := time.Now()
	request := pb.GetDefaultConfigIDRequest{}
	response, err := client.GrpcClient.GetDefaultConfigID(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			client.notify(ctx, 8005, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(12, response.GetConfigID(), err, time.Since(entryTime))
	}
	return response.GetConfigID(), err
}

/*
The Init method initializes the Senzing G2ConfigMgr object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2configmgrClient) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	if client.isTrace {
		client.traceEntry(17, moduleName, iniParams, verboseLogging)
	}
	entryTime := time.Now()
	request := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.GrpcClient.Init(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"iniParams":      iniParams,
				"moduleName":     moduleName,
				"verboseLogging": strconv.Itoa(verboseLogging),
			}
			client.notify(ctx, 8006, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(18, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	return err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2configmgrClient) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	if client.observers == nil {
		client.observers = &subject.SubjectImpl{}
	}
	return client.observers.RegisterObserver(ctx, observer)
}

/*
The ReplaceDefaultConfigID method replaces the old configuration identifier with a new configuration identifier in the Senzing database.
It is like a "compare-and-swap" instruction to serialize concurrent editing of configuration.
If oldConfigID is no longer the "old configuration identifier", the operation will fail.
To simply set the default configuration ID, use SetDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - oldConfigID: The configuration identifier to replace.
  - newConfigID: The configuration identifier to use as the default.
*/
func (client *G2configmgrClient) ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error {
	if client.isTrace {
		client.traceEntry(19, oldConfigID, newConfigID)
	}
	entryTime := time.Now()
	request := pb.ReplaceDefaultConfigIDRequest{
		OldConfigID: oldConfigID,
		NewConfigID: newConfigID,
	}
	_, err := client.GrpcClient.ReplaceDefaultConfigID(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"newConfigID": strconv.FormatInt(newConfigID, 10),
			}
			client.notify(ctx, 8007, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(20, oldConfigID, newConfigID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetDefaultConfigID method replaces the sets a new configuration identifier in the Senzing database.
To serialize modifying of the configuration identifier, see ReplaceDefaultConfigID().

Input
  - ctx: A context to control lifecycle.
  - configID: The configuration identifier of the Senzing Engine configuration to use as the default.
*/
func (client *G2configmgrClient) SetDefaultConfigID(ctx context.Context, configID int64) error {
	if client.isTrace {
		client.traceEntry(21, configID)
	}
	entryTime := time.Now()
	request := pb.SetDefaultConfigIDRequest{
		ConfigID: configID,
	}
	_, err := client.GrpcClient.SetDefaultConfigID(ctx, &request)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configID": strconv.FormatInt(configID, 10),
			}
			client.notify(ctx, 8008, err, details)
		}()
	}
	if client.isTrace {
		defer client.traceExit(22, configID, err, time.Since(entryTime))
	}
	return err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *G2configmgrClient) SetLogLevel(ctx context.Context, logLevel logger.Level) error {
	if client.isTrace {
		client.traceEntry(23, logLevel)
	}
	entryTime := time.Now()
	var err error = nil
	if client.isTrace {
		defer client.traceExit(24, logLevel, err, time.Since(entryTime))
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2configmgrClient) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	err := client.observers.UnregisterObserver(ctx, observer)
	if err != nil {
		return err
	}
	if !client.observers.HasObservers(ctx) {
		client.observers = nil
	}
	return err
}
