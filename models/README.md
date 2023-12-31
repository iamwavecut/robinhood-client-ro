# Go API client for openapi

Robinhood API Documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.robinhood.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAPI* | [**GetAccounts**](docs/AccountsAPI.md#getaccounts) | **Get** /accounts/ | getAccounts
*AccountsAPI* | [**GetPortfolio**](docs/AccountsAPI.md#getportfolio) | **Get** /accounts/{accountId}/portfolio/ | getPortfolio
*AccountsAPI* | [**GetPosition**](docs/AccountsAPI.md#getposition) | **Get** /accounts/{accountId}/positions/{positionId}/ | getPosition
*AccountsAPI* | [**GetPositions**](docs/AccountsAPI.md#getpositions) | **Get** /accounts/{accountId}/positions/ | getPositions
*AuthenticationAPI* | [**Login**](docs/AuthenticationAPI.md#login) | **Post** /api-token-auth/ | login
*AuthenticationAPI* | [**Logout**](docs/AuthenticationAPI.md#logout) | **Post** /api-token-logout/ | logout
*DividendsAPI* | [**GetDividend**](docs/DividendsAPI.md#getdividend) | **Get** /dividends/{id}/ | getDividend
*FundamentalsAPI* | [**GetFundamentals**](docs/FundamentalsAPI.md#getfundamentals) | **Get** /fundamentals/ | getFundamentals
*FundamentalsAPI* | [**GetSymbolFundamentals**](docs/FundamentalsAPI.md#getsymbolfundamentals) | **Get** /fundamentals/{symbol}/ | getSymbolFundamentals
*InstrumentsAPI* | [**GetInstrument**](docs/InstrumentsAPI.md#getinstrument) | **Get** /instruments/{instrument_id}/ | getInstrument
*InstrumentsAPI* | [**GetInstrumentSplits**](docs/InstrumentsAPI.md#getinstrumentsplits) | **Get** /instruments/{instrument_id}/splits/ | getInstrumentSplits
*InstrumentsAPI* | [**GetInstruments**](docs/InstrumentsAPI.md#getinstruments) | **Get** /instruments/ | getInstruments
*MarketsAPI* | [**GetMArketHours**](docs/MarketsAPI.md#getmarkethours) | **Get** /markets/{mic}/hours/{date}/ | getMArketHours
*MarketsAPI* | [**GetMarkets**](docs/MarketsAPI.md#getmarkets) | **Get** /markets | getMarkets
*MoversAPI* | [**GetMovers**](docs/MoversAPI.md#getmovers) | **Get** /midlands/movers/sp500/ | getMovers
*OrderAPI* | [**CancelOrder**](docs/OrderAPI.md#cancelorder) | **Post** /orders/{order_id}/cancel/ | cancelOrder
*OrderAPI* | [**GetOptionsOrders**](docs/OrderAPI.md#getoptionsorders) | **Get** /options/orders | getOptionsOrders
*OrderAPI* | [**GetOrder**](docs/OrderAPI.md#getorder) | **Get** /orders/{order_id}/ | getOrder
*OrderAPI* | [**GetOrders**](docs/OrderAPI.md#getorders) | **Get** /orders/ | getOrders
*OrderAPI* | [**PlaceOrder**](docs/OrderAPI.md#placeorder) | **Post** /orders/ | placeOrder
*QuoteAPI* | [**GetQuotes**](docs/QuoteAPI.md#getquotes) | **Get** /quotes/ | getQuotes
*QuoteAPI* | [**GetSymbolQuote**](docs/QuoteAPI.md#getsymbolquote) | **Get** /quotes/{symbol}/ | getSymbolQuote
*UserAPI* | [**GetBasicUserInfo**](docs/UserAPI.md#getbasicuserinfo) | **Get** /user/basic_info/ | getBasicUserInfo
*UserAPI* | [**GetInvestmentProfile**](docs/UserAPI.md#getinvestmentprofile) | **Get** /user/investment_profile/ | getInvestmentProfile
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | **Get** /user/ | getUser
*UserAPI* | [**GetUserID**](docs/UserAPI.md#getuserid) | **Get** /user/id/ | getUserID
*WatchlistsAPI* | [**BulkAddWatchlists**](docs/WatchlistsAPI.md#bulkaddwatchlists) | **Post** /watchlists/Default/bulk_add/ | bulkAddWatchlists
*WatchlistsAPI* | [**CreateWatchlist**](docs/WatchlistsAPI.md#createwatchlist) | **Post** /watchlists/ | createWatchlist
*WatchlistsAPI* | [**DeleteInstrumentFromWatchlist**](docs/WatchlistsAPI.md#deleteinstrumentfromwatchlist) | **Delete** /watchlists/{name}/{instrumentId} | deleteInstrumentFromWatchlist
*WatchlistsAPI* | [**GetWatchlistByName**](docs/WatchlistsAPI.md#getwatchlistbyname) | **Get** /watchlists/{name}/ | getWatchlistByName
*WatchlistsAPI* | [**GetWatchlists**](docs/WatchlistsAPI.md#getwatchlists) | **Get** /watchlists/ | getWatchlists


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountInfo](docs/AccountInfo.md)
 - [AccountType](docs/AccountType.md)
 - [Auth](docs/Auth.md)
 - [BasicInfo](docs/BasicInfo.md)
 - [CashBalances](docs/CashBalances.md)
 - [CryptoAccount](docs/CryptoAccount.md)
 - [CryptoAssetCurrency](docs/CryptoAssetCurrency.md)
 - [CryptoCurrencyPair](docs/CryptoCurrencyPair.md)
 - [CryptoOrder](docs/CryptoOrder.md)
 - [CryptoOrderOptions](docs/CryptoOrderOptions.md)
 - [CryptoOrderOutput](docs/CryptoOrderOutput.md)
 - [CryptoPortfolio](docs/CryptoPortfolio.md)
 - [Direction](docs/Direction.md)
 - [Execution](docs/Execution.md)
 - [ExecutionType](docs/ExecutionType.md)
 - [FundamentalsData](docs/FundamentalsData.md)
 - [GetOptionOrdersResponse](docs/GetOptionOrdersResponse.md)
 - [InstrumentData](docs/InstrumentData.md)
 - [InstrumentSplit](docs/InstrumentSplit.md)
 - [InstrumentState](docs/InstrumentState.md)
 - [InvestmentExperience](docs/InvestmentExperience.md)
 - [InvestmentObjective](docs/InvestmentObjective.md)
 - [InvestmentProfile](docs/InvestmentProfile.md)
 - [Leg](docs/Leg.md)
 - [LiquidityNeeds](docs/LiquidityNeeds.md)
 - [MarginBalances](docs/MarginBalances.md)
 - [MaritalStatus](docs/MaritalStatus.md)
 - [MarketData](docs/MarketData.md)
 - [MarketHours](docs/MarketHours.md)
 - [MinTicks](docs/MinTicks.md)
 - [OpenCloseStrategy](docs/OpenCloseStrategy.md)
 - [OptionChain](docs/OptionChain.md)
 - [OptionInstrument](docs/OptionInstrument.md)
 - [OptionOrder](docs/OptionOrder.md)
 - [OptionOrderInput](docs/OptionOrderInput.md)
 - [OptionOrderLeg](docs/OptionOrderLeg.md)
 - [OptionOrderLegExecutionsInner](docs/OptionOrderLegExecutionsInner.md)
 - [OptionsMarketData](docs/OptionsMarketData.md)
 - [Order](docs/Order.md)
 - [OrderAction](docs/OrderAction.md)
 - [OrderState](docs/OrderState.md)
 - [OrderTotalNotional](docs/OrderTotalNotional.md)
 - [PaginatedAccountInfo](docs/PaginatedAccountInfo.md)
 - [PaginatedFundamentalsData](docs/PaginatedFundamentalsData.md)
 - [PaginatedInstrumentData](docs/PaginatedInstrumentData.md)
 - [PaginatedInstrumentSplit](docs/PaginatedInstrumentSplit.md)
 - [PaginatedMarketData](docs/PaginatedMarketData.md)
 - [PaginatedMovers](docs/PaginatedMovers.md)
 - [PaginatedOptionChain](docs/PaginatedOptionChain.md)
 - [PaginatedOptionInstrument](docs/PaginatedOptionInstrument.md)
 - [PaginatedOrder](docs/PaginatedOrder.md)
 - [PaginatedPosition](docs/PaginatedPosition.md)
 - [PaginatedQuoteData](docs/PaginatedQuoteData.md)
 - [PaginatedWatchListCreateResponse](docs/PaginatedWatchListCreateResponse.md)
 - [PaginatedWatchListsData](docs/PaginatedWatchListsData.md)
 - [Portfolio](docs/Portfolio.md)
 - [Position](docs/Position.md)
 - [PositionEffect](docs/PositionEffect.md)
 - [Quote](docs/Quote.md)
 - [QuoteCurrency](docs/QuoteCurrency.md)
 - [QuoteData](docs/QuoteData.md)
 - [RiskTolerance](docs/RiskTolerance.md)
 - [Side](docs/Side.md)
 - [SourceOfFunds](docs/SourceOfFunds.md)
 - [TaxBracket](docs/TaxBracket.md)
 - [TimeHorizon](docs/TimeHorizon.md)
 - [TimeInForce](docs/TimeInForce.md)
 - [TotalNetWorth](docs/TotalNetWorth.md)
 - [Trigger](docs/Trigger.md)
 - [UnderlyingInstrument](docs/UnderlyingInstrument.md)
 - [UserId](docs/UserId.md)
 - [UserInfo](docs/UserInfo.md)
 - [WatchListCreateResponse](docs/WatchListCreateResponse.md)
 - [WatchListsData](docs/WatchListsData.md)
 - [Watchlist](docs/Watchlist.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

austin.millan@protonmail.com

