# Go API client for getopenpay

super charge your subscription management.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
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
import getopenpay "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `getopenpay.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), getopenpay.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `getopenpay.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), getopenpay.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `getopenpay.ContextOperationServerIndices` and `getopenpay.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), getopenpay.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), getopenpay.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://connto.getopenpay.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BillingPortalAPI* | [**CreatePortalSession**](docs/BillingPortalAPI.md#createportalsession) | **Post** /billing_portal/sessions | Create Portal Session
*ChargesAPI* | [**GetCharge**](docs/ChargesAPI.md#getcharge) | **Get** /charges/{charge_id} | Get Charge
*ChargesAPI* | [**ListCharges**](docs/ChargesAPI.md#listcharges) | **Post** /charges/list | List Charges
*ChargesAPI* | [**SearchCharges**](docs/ChargesAPI.md#searchcharges) | **Post** /charges/search | Search Charges
*ChargesAPI* | [**UpdateCharge**](docs/ChargesAPI.md#updatecharge) | **Put** /charges/{charge_id} | Update Charge
*CheckoutAPI* | [**CreateCheckoutSession**](docs/CheckoutAPI.md#createcheckoutsession) | **Post** /checkout/sessions | Create Checkout Session
*CheckoutAPI* | [**GetCheckoutSession**](docs/CheckoutAPI.md#getcheckoutsession) | **Get** /checkout/sessions/{session_id} | Get Checkout Session
*CheckoutAPI* | [**ListCheckoutSessions**](docs/CheckoutAPI.md#listcheckoutsessions) | **Post** /checkout/list | List Checkout Sessions
*CouponsAPI* | [**CreateCoupon**](docs/CouponsAPI.md#createcoupon) | **Post** /coupons/ | Create Coupon
*CouponsAPI* | [**DeleteCoupon**](docs/CouponsAPI.md#deletecoupon) | **Delete** /coupons/{coupon_id} | Delete Coupon
*CouponsAPI* | [**GetCoupon**](docs/CouponsAPI.md#getcoupon) | **Get** /coupons/{coupon_id} | Get Coupon
*CouponsAPI* | [**ListCoupons**](docs/CouponsAPI.md#listcoupons) | **Post** /coupons/list | List Coupons
*CouponsAPI* | [**UpdateCoupon**](docs/CouponsAPI.md#updatecoupon) | **Put** /coupons/{coupon_id} | Update Coupon
*CreditNotesAPI* | [**CreateCreditNote**](docs/CreditNotesAPI.md#createcreditnote) | **Post** /credit-notes/ | Create Credit Note
*CreditNotesAPI* | [**GetCreditNote**](docs/CreditNotesAPI.md#getcreditnote) | **Get** /credit-notes/{credit_note_id} | Get Credit Note
*CreditNotesAPI* | [**ListCreditNotes**](docs/CreditNotesAPI.md#listcreditnotes) | **Post** /credit-notes/list | List Credit Notes
*CustomersAPI* | [**CreateCustomer**](docs/CustomersAPI.md#createcustomer) | **Post** /customers/ | Create Customer
*CustomersAPI* | [**CreateCustomerBalanceTransaction**](docs/CustomersAPI.md#createcustomerbalancetransaction) | **Post** /customers/{customer_id}/balance-transactions | Create Customer Balance Transaction
*CustomersAPI* | [**DeleteCustomerDiscount**](docs/CustomersAPI.md#deletecustomerdiscount) | **Delete** /customers/{customer_id}/discount | Delete Customer Discount
*CustomersAPI* | [**GetCustomer**](docs/CustomersAPI.md#getcustomer) | **Get** /customers/{customer_id} | Get Customer
*CustomersAPI* | [**GetCustomerBalanceTransaction**](docs/CustomersAPI.md#getcustomerbalancetransaction) | **Get** /customers/{customer_id}/balance-transactions/{transaction_id} | Get Customer Balance Transaction
*CustomersAPI* | [**GetCustomerBalanceTransactions**](docs/CustomersAPI.md#getcustomerbalancetransactions) | **Get** /customers/{customer_id}/balance-transactions | Get Customer Balance Transactions
*CustomersAPI* | [**ListCustomerPaymentMethods**](docs/CustomersAPI.md#listcustomerpaymentmethods) | **Post** /customers/{customer_id}/payment-methods | List Customer Payment Methods
*CustomersAPI* | [**ListCustomers**](docs/CustomersAPI.md#listcustomers) | **Post** /customers/list | List Customers
*CustomersAPI* | [**ListValidSubscriptions**](docs/CustomersAPI.md#listvalidsubscriptions) | **Post** /customers{customer_id}/list_valid_subscriptions | List Valid Subscriptions
*CustomersAPI* | [**SearchCustomers**](docs/CustomersAPI.md#searchcustomers) | **Post** /customers/search | Search Customers
*CustomersAPI* | [**UpdateCustomer**](docs/CustomersAPI.md#updatecustomer) | **Put** /customers/{customer_external_id} | Update Customer
*DisputesAPI* | [**CreateDispute**](docs/DisputesAPI.md#createdispute) | **Post** /dispute/ | Create Dispute
*DisputesAPI* | [**GetDispute**](docs/DisputesAPI.md#getdispute) | **Get** /dispute/{dispute_id} | Get Dispute
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | **Get** /events/{event_id} | Get Event
*EventsAPI* | [**ListEvents**](docs/EventsAPI.md#listevents) | **Post** /events/list | List Events
*EventsAPI* | [**SearchEvents**](docs/EventsAPI.md#searchevents) | **Post** /events/search | Search Events
*InvoiceItemsAPI* | [**DeleteInvoiceItemById**](docs/InvoiceItemsAPI.md#deleteinvoiceitembyid) | **Delete** /invoice-items/{invoice_item_id} | Delete Invoice Item By Id
*InvoiceItemsAPI* | [**GetInvoiceItem**](docs/InvoiceItemsAPI.md#getinvoiceitem) | **Get** /invoice-items/{invoice_item_id} | Get Invoice Item
*InvoiceItemsAPI* | [**ListInvoiceItems**](docs/InvoiceItemsAPI.md#listinvoiceitems) | **Post** /invoice-items/list | List Invoice Items
*InvoicesAPI* | [**AddComment**](docs/InvoicesAPI.md#addcomment) | **Post** /invoices/{invoice_external_id}/comment | Add Comment
*InvoicesAPI* | [**CreateInvoice**](docs/InvoicesAPI.md#createinvoice) | **Post** /invoices/ | Create Invoice
*InvoicesAPI* | [**FinalizeInvoice**](docs/InvoicesAPI.md#finalizeinvoice) | **Post** /invoices/{invoice_external_id}/finalize | Finalize Invoice
*InvoicesAPI* | [**GetInvoice**](docs/InvoicesAPI.md#getinvoice) | **Get** /invoices/{invoice_external_id} | Get Invoice
*InvoicesAPI* | [**GetInvoicePublic**](docs/InvoicesAPI.md#getinvoicepublic) | **Get** /invoices/public/{public_permanent_token} | Get Invoice Public
*InvoicesAPI* | [**ListInvoices**](docs/InvoicesAPI.md#listinvoices) | **Post** /invoices/list | List Invoices
*InvoicesAPI* | [**MarkInvoiceAsUncollectible**](docs/InvoicesAPI.md#markinvoiceasuncollectible) | **Post** /invoices/{invoice_external_id}/mark_uncollectible | Mark Invoice As Uncollectible
*InvoicesAPI* | [**MarkInvoiceAsVoid**](docs/InvoicesAPI.md#markinvoiceasvoid) | **Post** /invoices/{invoice_external_id}/void | Mark Invoice As Void
*InvoicesAPI* | [**PayInvoice**](docs/InvoicesAPI.md#payinvoice) | **Post** /invoices/{invoice_external_id}/pay | Pay Invoice
*InvoicesAPI* | [**SearchInvoices**](docs/InvoicesAPI.md#searchinvoices) | **Post** /invoices/search | Search Invoices
*InvoicesAPI* | [**UpdateInvoice**](docs/InvoicesAPI.md#updateinvoice) | **Put** /invoices/{invoice_id} | Update Invoice
*PaymentIntentsAPI* | [**GetPaymentIntent**](docs/PaymentIntentsAPI.md#getpaymentintent) | **Get** /payment-intents/{payment_intent_id} | Get Payment Intent
*PaymentIntentsAPI* | [**ListPaymentIntents**](docs/PaymentIntentsAPI.md#listpaymentintents) | **Post** /payment-intents/list | List Payment Intents
*PaymentIntentsAPI* | [**SearchPaymentIntents**](docs/PaymentIntentsAPI.md#searchpaymentintents) | **Post** /payment-intents/search | Search Payment Intents
*PaymentIntentsAPI* | [**UpdatePaymentIntent**](docs/PaymentIntentsAPI.md#updatepaymentintent) | **Put** /payment-intents/{payment_intent_id} | Update Payment Intent
*PaymentLinksAPI* | [**CreatePaymentLink**](docs/PaymentLinksAPI.md#createpaymentlink) | **Post** /payment-link/ | Create Payment Link
*PaymentLinksAPI* | [**GetPaymentLink**](docs/PaymentLinksAPI.md#getpaymentlink) | **Get** /payment-link/{plink_id} | Get Payment Link
*PaymentLinksAPI* | [**ListPaymentLinks**](docs/PaymentLinksAPI.md#listpaymentlinks) | **Post** /payment-link/list | List Payment Links
*PaymentLinksAPI* | [**OpenPaymentLinkPagePublic**](docs/PaymentLinksAPI.md#openpaymentlinkpagepublic) | **Get** /payment-link/public/{secure_token} | Open Payment Link Page Public
*PaymentMethodsAPI* | [**AttachPaymentMethod**](docs/PaymentMethodsAPI.md#attachpaymentmethod) | **Post** /payment-methods/{payment_method_id}/attach | Attach Payment Method
*PaymentMethodsAPI* | [**AuthorizePaymentMethod**](docs/PaymentMethodsAPI.md#authorizepaymentmethod) | **Post** /payment-methods/authorize | Authorize Payment Method
*PaymentMethodsAPI* | [**DetachPaymentMethod**](docs/PaymentMethodsAPI.md#detachpaymentmethod) | **Post** /payment-methods/{payment_method_id}/detach | Detach Payment Method
*PaymentMethodsAPI* | [**GetPaymentMethod**](docs/PaymentMethodsAPI.md#getpaymentmethod) | **Get** /payment-methods/{payment_method_id} | Get Payment Method
*PricesAPI* | [**CreatePriceForProduct**](docs/PricesAPI.md#createpriceforproduct) | **Post** /prices/ | Create Price For Product
*PricesAPI* | [**DeletePriceForProduct**](docs/PricesAPI.md#deletepriceforproduct) | **Delete** /prices/{price_id} | Delete Price For Product
*PricesAPI* | [**GetPriceForProduct**](docs/PricesAPI.md#getpriceforproduct) | **Get** /prices/{price_id} | Get Price For Product
*PricesAPI* | [**ListPrices**](docs/PricesAPI.md#listprices) | **Post** /prices/list | List Prices
*PricesAPI* | [**SearchPrices**](docs/PricesAPI.md#searchprices) | **Post** /prices/search | Search Prices
*PricesAPI* | [**UpdatePriceForProduct**](docs/PricesAPI.md#updatepriceforproduct) | **Put** /prices/{price_id} | Update Price For Product
*ProductFamilyAPI* | [**CreateProductFamily**](docs/ProductFamilyAPI.md#createproductfamily) | **Post** /product-family/ | Create Product Family
*ProductFamilyAPI* | [**DeleteProductFamily**](docs/ProductFamilyAPI.md#deleteproductfamily) | **Delete** /product-family/{id} | Delete Product Family
*ProductFamilyAPI* | [**GetProductFamily**](docs/ProductFamilyAPI.md#getproductfamily) | **Get** /product-family/{id} | Get Product Family
*ProductFamilyAPI* | [**ListProductFamilies**](docs/ProductFamilyAPI.md#listproductfamilies) | **Post** /product-family/list | List Product Families
*ProductFamilyAPI* | [**UpdateProductFamily**](docs/ProductFamilyAPI.md#updateproductfamily) | **Put** /product-family/{id} | Update Product Family
*ProductsAPI* | [**CreateProduct**](docs/ProductsAPI.md#createproduct) | **Post** /products/ | Create Product
*ProductsAPI* | [**DeleteProduct**](docs/ProductsAPI.md#deleteproduct) | **Delete** /products/{product_id} | Delete Product
*ProductsAPI* | [**GetProduct**](docs/ProductsAPI.md#getproduct) | **Get** /products/{product_id} | Get Product
*ProductsAPI* | [**ListProducts**](docs/ProductsAPI.md#listproducts) | **Post** /products/list | List Products
*ProductsAPI* | [**SearchProducts**](docs/ProductsAPI.md#searchproducts) | **Post** /products/search | Search Products
*ProductsAPI* | [**UpdateProduct**](docs/ProductsAPI.md#updateproduct) | **Put** /products/{product_id} | Update Product
*PromotionCodesAPI* | [**CreatePromoCode**](docs/PromotionCodesAPI.md#createpromocode) | **Post** /promotion-codes/ | Create Promo Code
*PromotionCodesAPI* | [**GetPromoCode**](docs/PromotionCodesAPI.md#getpromocode) | **Get** /promotion-codes/{promo_code_id} | Get Promo Code
*PromotionCodesAPI* | [**GetPromoCodeByCode**](docs/PromotionCodesAPI.md#getpromocodebycode) | **Get** /promotion-codes/code/{promo_code} | Get Promo Code By Code
*PromotionCodesAPI* | [**ListPromoCodes**](docs/PromotionCodesAPI.md#listpromocodes) | **Post** /promotion-codes/list | List Promo Codes
*PromotionCodesAPI* | [**UpdatePromoCode**](docs/PromotionCodesAPI.md#updatepromocode) | **Put** /promotion-codes/{promo_code_id} | Update Promo Code
*RefundsAPI* | [**CreateRefund**](docs/RefundsAPI.md#createrefund) | **Post** /refunds/ | Create Refund
*RefundsAPI* | [**ListRefunds**](docs/RefundsAPI.md#listrefunds) | **Post** /refunds/list | List Refunds
*SubscriptionItemsAPI* | [**CreateSubscriptionItem**](docs/SubscriptionItemsAPI.md#createsubscriptionitem) | **Post** /subscription-items/ | Create Subscription Item
*SubscriptionItemsAPI* | [**DeleteSubscriptionItem**](docs/SubscriptionItemsAPI.md#deletesubscriptionitem) | **Delete** /subscription-items/{subscription_item_id} | Delete Subscription Item
*SubscriptionItemsAPI* | [**GetSubscriptionItem**](docs/SubscriptionItemsAPI.md#getsubscriptionitem) | **Get** /subscription-items/{subscription_item_id} | Get Subscription Item
*SubscriptionItemsAPI* | [**ListSubscriptionItems**](docs/SubscriptionItemsAPI.md#listsubscriptionitems) | **Post** /subscription-items/list | List Subscription Items
*SubscriptionItemsAPI* | [**UpdateSubscriptionItem**](docs/SubscriptionItemsAPI.md#updatesubscriptionitem) | **Put** /subscription-items/{subscription_item_id} | Update Subscription Item
*SubscriptionsAPI* | [**CancelSubscriptionTrial**](docs/SubscriptionsAPI.md#cancelsubscriptiontrial) | **Post** /subscriptions/{subscription_id}/cancel-subscription-trial | Cancel Subscription Trial
*SubscriptionsAPI* | [**CreateSubscriptions**](docs/SubscriptionsAPI.md#createsubscriptions) | **Post** /subscriptions/ | Create Subscriptions
*SubscriptionsAPI* | [**DeleteSubscription**](docs/SubscriptionsAPI.md#deletesubscription) | **Delete** /subscriptions/{subscription_id} | Delete Subscription
*SubscriptionsAPI* | [**DeleteSubscriptionDiscount**](docs/SubscriptionsAPI.md#deletesubscriptiondiscount) | **Delete** /subscriptions/{subscription_id}/discount | Delete Subscription Discount
*SubscriptionsAPI* | [**GetSubscription**](docs/SubscriptionsAPI.md#getsubscription) | **Get** /subscriptions/{subscription_id} | Get Subscription
*SubscriptionsAPI* | [**ListSubscriptions**](docs/SubscriptionsAPI.md#listsubscriptions) | **Post** /subscriptions/list | List Subscriptions
*SubscriptionsAPI* | [**PauseSubscription**](docs/SubscriptionsAPI.md#pausesubscription) | **Put** /subscriptions/{subscription_id}/pause | Pause Subscription
*SubscriptionsAPI* | [**RefreshSubscriptionStatus**](docs/SubscriptionsAPI.md#refreshsubscriptionstatus) | **Post** /subscriptions/{subscription_id}/refresh-status | Refresh Subscription Status
*SubscriptionsAPI* | [**ResumeSubscription**](docs/SubscriptionsAPI.md#resumesubscription) | **Put** /subscriptions/{subscription_id}/resume | Resume Subscription
*SubscriptionsAPI* | [**SearchSubscriptions**](docs/SubscriptionsAPI.md#searchsubscriptions) | **Post** /subscriptions/search | Search Subscriptions
*SubscriptionsAPI* | [**UpdateSubscription**](docs/SubscriptionsAPI.md#updatesubscription) | **Put** /subscriptions/{subscription_id} | Update Subscription
*TaxIntegrationsAPI* | [**CreateTaxIntegration**](docs/TaxIntegrationsAPI.md#createtaxintegration) | **Post** /integrations/tax/ | Create Tax Integration


## Documentation For Models

 - [ActiveSubResponse](docs/ActiveSubResponse.md)
 - [AddCommentRequest](docs/AddCommentRequest.md)
 - [AttachPaymentMethodRequest](docs/AttachPaymentMethodRequest.md)
 - [AuthorizePaymentMethodRequest](docs/AuthorizePaymentMethodRequest.md)
 - [BillingReasonEnum](docs/BillingReasonEnum.md)
 - [BillingSchemeEnum](docs/BillingSchemeEnum.md)
 - [BooleanConfigField](docs/BooleanConfigField.md)
 - [CalendarIntervalEnum](docs/CalendarIntervalEnum.md)
 - [ChargeExternal](docs/ChargeExternal.md)
 - [ChargeQueryParams](docs/ChargeQueryParams.md)
 - [ChargeStatusEnum](docs/ChargeStatusEnum.md)
 - [CheckoutMode](docs/CheckoutMode.md)
 - [CheckoutSessionExternal](docs/CheckoutSessionExternal.md)
 - [CheckoutSessionLineItemExternal](docs/CheckoutSessionLineItemExternal.md)
 - [CheckoutSessionQueryParams](docs/CheckoutSessionQueryParams.md)
 - [CheckoutSessionStatus](docs/CheckoutSessionStatus.md)
 - [CollectionMethodEnum](docs/CollectionMethodEnum.md)
 - [CompleteAddress](docs/CompleteAddress.md)
 - [CouponDuration](docs/CouponDuration.md)
 - [CouponExternal](docs/CouponExternal.md)
 - [CouponQueryParams](docs/CouponQueryParams.md)
 - [CreateCheckoutLineItem](docs/CreateCheckoutLineItem.md)
 - [CreateCheckoutSessionRequest](docs/CreateCheckoutSessionRequest.md)
 - [CreateCouponRequest](docs/CreateCouponRequest.md)
 - [CreateCreditNoteLine](docs/CreateCreditNoteLine.md)
 - [CreateCreditNoteRequest](docs/CreateCreditNoteRequest.md)
 - [CreateCustomerBalanceTransactionRequest](docs/CreateCustomerBalanceTransactionRequest.md)
 - [CreateCustomerRequest](docs/CreateCustomerRequest.md)
 - [CreateDisputeRequest](docs/CreateDisputeRequest.md)
 - [CreateInvoiceRequest](docs/CreateInvoiceRequest.md)
 - [CreatePaymentLinkRequest](docs/CreatePaymentLinkRequest.md)
 - [CreatePortalSessionRequest](docs/CreatePortalSessionRequest.md)
 - [CreatePriceRequest](docs/CreatePriceRequest.md)
 - [CreateProductFamilyRequest](docs/CreateProductFamilyRequest.md)
 - [CreateProductRequest](docs/CreateProductRequest.md)
 - [CreatePromoCodeRequest](docs/CreatePromoCodeRequest.md)
 - [CreateRefundRequest](docs/CreateRefundRequest.md)
 - [CreateSubscriptionItemRequest](docs/CreateSubscriptionItemRequest.md)
 - [CreateSubscriptionRequest](docs/CreateSubscriptionRequest.md)
 - [CreateSubscriptionResponse](docs/CreateSubscriptionResponse.md)
 - [CreateTaxIntegrationRequest](docs/CreateTaxIntegrationRequest.md)
 - [CreditNoteExternal](docs/CreditNoteExternal.md)
 - [CreditNoteItemExternal](docs/CreditNoteItemExternal.md)
 - [CreditNoteLineType](docs/CreditNoteLineType.md)
 - [CreditNoteQueryParams](docs/CreditNoteQueryParams.md)
 - [CreditNoteReason](docs/CreditNoteReason.md)
 - [CurrencyEnum](docs/CurrencyEnum.md)
 - [CustomerBalanceExternal](docs/CustomerBalanceExternal.md)
 - [CustomerBalanceTransactionExternal](docs/CustomerBalanceTransactionExternal.md)
 - [CustomerBalanceTransactionType](docs/CustomerBalanceTransactionType.md)
 - [CustomerExternal](docs/CustomerExternal.md)
 - [CustomerPaymentMethodQueryParams](docs/CustomerPaymentMethodQueryParams.md)
 - [CustomerQueryParams](docs/CustomerQueryParams.md)
 - [DateTimeFilter](docs/DateTimeFilter.md)
 - [DeleteInvoiceItemResponse](docs/DeleteInvoiceItemResponse.md)
 - [DeletePriceResponse](docs/DeletePriceResponse.md)
 - [DeleteProductResponse](docs/DeleteProductResponse.md)
 - [DeleteSubscriptionItemRequest](docs/DeleteSubscriptionItemRequest.md)
 - [DeleteSubscriptionItemResponse](docs/DeleteSubscriptionItemResponse.md)
 - [DeleteSubscriptionRequest](docs/DeleteSubscriptionRequest.md)
 - [DiscountExternal](docs/DiscountExternal.md)
 - [Discounts](docs/Discounts.md)
 - [DisputeExternal](docs/DisputeExternal.md)
 - [EnumConfigField](docs/EnumConfigField.md)
 - [EventExternal](docs/EventExternal.md)
 - [EventSearchParams](docs/EventSearchParams.md)
 - [EventType](docs/EventType.md)
 - [EventsQueryParams](docs/EventsQueryParams.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [InlineSubscriptionItemUpdate](docs/InlineSubscriptionItemUpdate.md)
 - [IntRangeFilter](docs/IntRangeFilter.md)
 - [InvoiceComment](docs/InvoiceComment.md)
 - [InvoiceDiscountAmountsExternal](docs/InvoiceDiscountAmountsExternal.md)
 - [InvoiceDiscountOptions](docs/InvoiceDiscountOptions.md)
 - [InvoiceExternal](docs/InvoiceExternal.md)
 - [InvoiceItemDiscountAmountsExternal](docs/InvoiceItemDiscountAmountsExternal.md)
 - [InvoiceItemDiscountAmountsPublic](docs/InvoiceItemDiscountAmountsPublic.md)
 - [InvoiceItemExternal](docs/InvoiceItemExternal.md)
 - [InvoiceItemPublic](docs/InvoiceItemPublic.md)
 - [InvoiceItemsQueryParams](docs/InvoiceItemsQueryParams.md)
 - [InvoicePublic](docs/InvoicePublic.md)
 - [InvoiceQueryParams](docs/InvoiceQueryParams.md)
 - [InvoiceSettings](docs/InvoiceSettings.md)
 - [InvoiceStatusEnum](docs/InvoiceStatusEnum.md)
 - [ListActiveSubParams](docs/ListActiveSubParams.md)
 - [ListResponseChargeExternal](docs/ListResponseChargeExternal.md)
 - [ListResponseCheckoutSessionExternal](docs/ListResponseCheckoutSessionExternal.md)
 - [ListResponseCouponExternal](docs/ListResponseCouponExternal.md)
 - [ListResponseCreditNoteExternal](docs/ListResponseCreditNoteExternal.md)
 - [ListResponseCustomerExternal](docs/ListResponseCustomerExternal.md)
 - [ListResponseEventExternal](docs/ListResponseEventExternal.md)
 - [ListResponseInvoiceExternal](docs/ListResponseInvoiceExternal.md)
 - [ListResponseInvoiceItemExternal](docs/ListResponseInvoiceItemExternal.md)
 - [ListResponsePaymentIntentExternal](docs/ListResponsePaymentIntentExternal.md)
 - [ListResponsePaymentLinkExternal](docs/ListResponsePaymentLinkExternal.md)
 - [ListResponsePriceExternal](docs/ListResponsePriceExternal.md)
 - [ListResponseProductExternal](docs/ListResponseProductExternal.md)
 - [ListResponseProductFamilyExternal](docs/ListResponseProductFamilyExternal.md)
 - [ListResponsePromotionCodeExternal](docs/ListResponsePromotionCodeExternal.md)
 - [ListResponseRefundExternal](docs/ListResponseRefundExternal.md)
 - [ListResponseSubscriptionExternal](docs/ListResponseSubscriptionExternal.md)
 - [ListResponseSubscriptionItemExternal](docs/ListResponseSubscriptionItemExternal.md)
 - [ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal](docs/ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal.md)
 - [ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner](docs/ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner.md)
 - [MarkUncollectibleRequest](docs/MarkUncollectibleRequest.md)
 - [MarkVoidRequest](docs/MarkVoidRequest.md)
 - [NonPciIntegration](docs/NonPciIntegration.md)
 - [NonPciIntegrationConfigurationInner](docs/NonPciIntegrationConfigurationInner.md)
 - [NonPciIntegrationEnum](docs/NonPciIntegrationEnum.md)
 - [NumberConfigField](docs/NumberConfigField.md)
 - [ObjectName](docs/ObjectName.md)
 - [PayInvoiceRequest](docs/PayInvoiceRequest.md)
 - [PaymentIntentExternal](docs/PaymentIntentExternal.md)
 - [PaymentIntentQueryParams](docs/PaymentIntentQueryParams.md)
 - [PaymentIntentStatus](docs/PaymentIntentStatus.md)
 - [PaymentLinkExternal](docs/PaymentLinkExternal.md)
 - [PaymentLinkLineItemExternal](docs/PaymentLinkLineItemExternal.md)
 - [PaymentLinkQueryParams](docs/PaymentLinkQueryParams.md)
 - [PaymentMethodExternal](docs/PaymentMethodExternal.md)
 - [PaymentProcessorName](docs/PaymentProcessorName.md)
 - [PaymentProviderType](docs/PaymentProviderType.md)
 - [PortalSessionExternal](docs/PortalSessionExternal.md)
 - [PriceExternal](docs/PriceExternal.md)
 - [PriceQueryParams](docs/PriceQueryParams.md)
 - [PriceTierExternal](docs/PriceTierExternal.md)
 - [PriceTierParams](docs/PriceTierParams.md)
 - [PriceTierPublic](docs/PriceTierPublic.md)
 - [PriceTypeEnum](docs/PriceTypeEnum.md)
 - [PricingModel](docs/PricingModel.md)
 - [PricingTiersEnum](docs/PricingTiersEnum.md)
 - [ProductExternal](docs/ProductExternal.md)
 - [ProductFamilyExternal](docs/ProductFamilyExternal.md)
 - [ProductFamilyQueryParams](docs/ProductFamilyQueryParams.md)
 - [ProductQueryParams](docs/ProductQueryParams.md)
 - [PromoCodeQueryParams](docs/PromoCodeQueryParams.md)
 - [PromoRestrictions](docs/PromoRestrictions.md)
 - [PromotionCodeExternal](docs/PromotionCodeExternal.md)
 - [ProrationEnum](docs/ProrationEnum.md)
 - [RecurringDetails](docs/RecurringDetails.md)
 - [RefundExternal](docs/RefundExternal.md)
 - [RefundQueryParams](docs/RefundQueryParams.md)
 - [RefundReasonEnum](docs/RefundReasonEnum.md)
 - [RefundStatusEnum](docs/RefundStatusEnum.md)
 - [ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet](docs/ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet.md)
 - [SearchChargeRequest](docs/SearchChargeRequest.md)
 - [SearchCustomerRequest](docs/SearchCustomerRequest.md)
 - [SearchFilter](docs/SearchFilter.md)
 - [SearchInvoiceRequest](docs/SearchInvoiceRequest.md)
 - [SearchPaymentIntentRequest](docs/SearchPaymentIntentRequest.md)
 - [SearchPriceRequest](docs/SearchPriceRequest.md)
 - [SearchProductRequest](docs/SearchProductRequest.md)
 - [SearchSubscriptionRequest](docs/SearchSubscriptionRequest.md)
 - [SelectedPriceQuantity](docs/SelectedPriceQuantity.md)
 - [SetupIntentExternal](docs/SetupIntentExternal.md)
 - [StringConfigField](docs/StringConfigField.md)
 - [SubscriptionCancelFeedbackEnum](docs/SubscriptionCancelFeedbackEnum.md)
 - [SubscriptionCancellationDetails](docs/SubscriptionCancellationDetails.md)
 - [SubscriptionExternal](docs/SubscriptionExternal.md)
 - [SubscriptionItemExternal](docs/SubscriptionItemExternal.md)
 - [SubscriptionItemQueryParams](docs/SubscriptionItemQueryParams.md)
 - [SubscriptionPauseRequest](docs/SubscriptionPauseRequest.md)
 - [SubscriptionQueryParams](docs/SubscriptionQueryParams.md)
 - [SubscriptionResumeRequest](docs/SubscriptionResumeRequest.md)
 - [SubscriptionStatusEnum](docs/SubscriptionStatusEnum.md)
 - [TaxIntegrationApiName](docs/TaxIntegrationApiName.md)
 - [TaxIntegrationExternal](docs/TaxIntegrationExternal.md)
 - [TokenizedCreditCardExternal](docs/TokenizedCreditCardExternal.md)
 - [UpdateChargeRequest](docs/UpdateChargeRequest.md)
 - [UpdateCouponRequest](docs/UpdateCouponRequest.md)
 - [UpdateCustomerRequest](docs/UpdateCustomerRequest.md)
 - [UpdateInvoiceRequest](docs/UpdateInvoiceRequest.md)
 - [UpdatePaymentIntent](docs/UpdatePaymentIntent.md)
 - [UpdatePriceRequest](docs/UpdatePriceRequest.md)
 - [UpdateProductFamilyRequest](docs/UpdateProductFamilyRequest.md)
 - [UpdateProductRequest](docs/UpdateProductRequest.md)
 - [UpdatePromoCodeRequest](docs/UpdatePromoCodeRequest.md)
 - [UpdateSubscriptionItemRequest](docs/UpdateSubscriptionItemRequest.md)
 - [UpdateSubscriptionRequest](docs/UpdateSubscriptionRequest.md)
 - [UpdateSubscriptionResponse](docs/UpdateSubscriptionResponse.md)
 - [UsageAggMethodEnum](docs/UsageAggMethodEnum.md)
 - [UsageTypeEnum](docs/UsageTypeEnum.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### HTTPBearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), getopenpay.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


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



