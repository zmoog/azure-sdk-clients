# azure-sdk-clients

This azure-sdk-clients repo contains the [Previous Versions](https://github.com/Azure/azure-sdk-for-go#management-previous-versions) of the Azure SDK client generated from the [OpenAPI spec](https://github.com/Azure/azure-rest-api-specs).


## Copy OpenAPI Specification

Clone the https://github.com/Azure/azure-rest-api-specs repo locally; you will have all the specs easy accessibile on your machine.

```shell
$ cd ~/code/projects/azure

$ gh repo clone Azure/azure-rest-api-specs
```


## Install the required tools

I used the following tools:

- nodejs 12.22.12
- Ubuntu 20.04.4 LTS

After installing nodejs, create a new folder and install the required packages:

```shell
$ mkdir ~/code/projects/zmoog/autorest
$ npm install autorest
```

I am intentionally installing autorest locally and using `./node_modules/.bin/autorest` to not mess up with the global nodejs install.


## Generate the clients

Let's generate the Go client from Consumptions API version 2021-05-01.

```shell
$ ~/code/projects/zmoog/autorest

$ ./node_modules/.bin/autorest \
        --input-file=/home/zmoog/code/projects/azure/azure-rest-api-specs/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-05-01/consumption.json \
        --output-folder=generated/2021-05-01 \
        --go \
        --namespace=consumption
AutoRest code generation utility [cli version: 3.6.1; node: v12.22.12]
(C) 2018 Microsoft Corporation.
https://aka.ms/autorest
info    | AutoRest core version selected from configuration: ~2.0.4413.
   Loading AutoRest core      '/home/zmoog/.autorest/@microsoft.azureautorest-core@2.0.4421/nodemodules/@microsoft.azure/autorest-core/dist' (2.0.4421)
   Loading AutoRest extension '@microsoft.azure/autorest.go' (~2.1.47->2.1.188)
   Loading AutoRest extension '@microsoft.azure/autorest.modeler' (2.3.38->2.3.38)
