package consumption

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReservationRecommendationsClient is the consumption management client provides access to consumption resources for
// Azure Enterprise Subscriptions.
type ReservationRecommendationsClient struct {
	BaseClient
}

// NewReservationRecommendationsClient creates an instance of the ReservationRecommendationsClient client.
func NewReservationRecommendationsClient(subscriptionID string) ReservationRecommendationsClient {
	return NewReservationRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReservationRecommendationsClientWithBaseURI creates an instance of the ReservationRecommendationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewReservationRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationsClient {
	return ReservationRecommendationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list of recommendations for purchasing reserved instances.
// Parameters:
// resourceScope - the scope associated with reservation recommendations operations. This includes
// '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resource group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for BillingAccount scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope
// filter - may be used to filter reservationRecommendations by: properties/scope with allowed values
// ['Single', 'Shared'] and default value 'Single'; properties/resourceType with allowed values
// ['VirtualMachines', 'SQLDatabases', 'PostgreSQL', 'ManagedDisk', 'MySQL', 'RedHat', 'MariaDB', 'RedisCache',
// 'CosmosDB', 'SqlDataWarehouse', 'SUSELinux', 'AppService', 'BlockBlob', 'AzureDataExplorer',
// 'VMwareCloudSimple'] and default value 'VirtualMachines'; and properties/lookBackPeriod with allowed values
// ['Last7Days', 'Last30Days', 'Last60Days'] and default value 'Last7Days'.
func (client ReservationRecommendationsClient) List(ctx context.Context, resourceScope string, filter string) (result ReservationRecommendationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationRecommendationsClient.List")
		defer func() {
			sc := -1
			if result.rrlr.Response.Response != nil {
				sc = result.rrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceScope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rrlr.hasNextLink() && result.rrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReservationRecommendationsClient) ListPreparer(ctx context.Context, resourceScope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceScope": resourceScope,
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Consumption/reservationRecommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReservationRecommendationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReservationRecommendationsClient) ListResponder(resp *http.Response) (result ReservationRecommendationsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReservationRecommendationsClient) listNextResults(ctx context.Context, lastResults ReservationRecommendationsListResult) (result ReservationRecommendationsListResult, err error) {
	req, err := lastResults.reservationRecommendationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationRecommendationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReservationRecommendationsClient) ListComplete(ctx context.Context, resourceScope string, filter string) (result ReservationRecommendationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationRecommendationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceScope, filter)
	return
}
