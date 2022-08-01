package costmanagementapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/zmoog/azure-sdk-clients/services/costmanagement/mgmt/2021-10-01/costmanagement"
)

// ViewsClientAPI contains the set of methods on the ViewsClient type.
type ViewsClientAPI interface {
	CreateOrUpdate(ctx context.Context, viewName string, parameters costmanagement.View) (result costmanagement.View, err error)
	CreateOrUpdateByScope(ctx context.Context, scope string, viewName string, parameters costmanagement.View) (result costmanagement.View, err error)
	Delete(ctx context.Context, viewName string) (result autorest.Response, err error)
	DeleteByScope(ctx context.Context, scope string, viewName string) (result autorest.Response, err error)
	Get(ctx context.Context, viewName string) (result costmanagement.View, err error)
	GetByScope(ctx context.Context, scope string, viewName string) (result costmanagement.View, err error)
	List(ctx context.Context) (result costmanagement.ViewListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.ViewListResultIterator, err error)
	ListByScope(ctx context.Context, scope string) (result costmanagement.ViewListResultPage, err error)
	ListByScopeComplete(ctx context.Context, scope string) (result costmanagement.ViewListResultIterator, err error)
}

var _ ViewsClientAPI = (*costmanagement.ViewsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	Dismiss(ctx context.Context, scope string, alertID string, parameters costmanagement.DismissAlertPayload) (result costmanagement.Alert, err error)
	Get(ctx context.Context, scope string, alertID string) (result costmanagement.Alert, err error)
	List(ctx context.Context, scope string) (result costmanagement.AlertsResult, err error)
	ListExternal(ctx context.Context, externalCloudProviderType costmanagement.ExternalCloudProviderType, externalCloudProviderID string) (result costmanagement.AlertsResult, err error)
}

var _ AlertsClientAPI = (*costmanagement.AlertsClient)(nil)

// ForecastClientAPI contains the set of methods on the ForecastClient type.
type ForecastClientAPI interface {
	ExternalCloudProviderUsage(ctx context.Context, externalCloudProviderType costmanagement.ExternalCloudProviderType, externalCloudProviderID string, parameters costmanagement.ForecastDefinition, filter string) (result costmanagement.QueryResult, err error)
	Usage(ctx context.Context, scope string, parameters costmanagement.ForecastDefinition, filter string) (result costmanagement.QueryResult, err error)
}

var _ ForecastClientAPI = (*costmanagement.ForecastClient)(nil)

// DimensionsClientAPI contains the set of methods on the DimensionsClient type.
type DimensionsClientAPI interface {
	ByExternalCloudProviderType(ctx context.Context, externalCloudProviderType costmanagement.ExternalCloudProviderType, externalCloudProviderID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	List(ctx context.Context, scope string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ DimensionsClientAPI = (*costmanagement.DimensionsClient)(nil)

// QueryClientAPI contains the set of methods on the QueryClient type.
type QueryClientAPI interface {
	Usage(ctx context.Context, scope string, parameters costmanagement.QueryDefinition) (result costmanagement.QueryResult, err error)
	UsageByExternalCloudProviderType(ctx context.Context, externalCloudProviderType costmanagement.ExternalCloudProviderType, externalCloudProviderID string, parameters costmanagement.QueryDefinition) (result costmanagement.QueryResult, err error)
}

var _ QueryClientAPI = (*costmanagement.QueryClient)(nil)

// GenerateReservationDetailsReportClientAPI contains the set of methods on the GenerateReservationDetailsReportClient type.
type GenerateReservationDetailsReportClientAPI interface {
	ByBillingAccountID(ctx context.Context, billingAccountID string, startDate string, endDate string) (result costmanagement.GenerateReservationDetailsReportByBillingAccountIDFuture, err error)
	ByBillingProfileID(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string) (result costmanagement.GenerateReservationDetailsReportByBillingProfileIDFuture, err error)
}

var _ GenerateReservationDetailsReportClientAPI = (*costmanagement.GenerateReservationDetailsReportClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)
