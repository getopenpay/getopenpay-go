package getopenpay_test

import (
    "context"
    "testing"

    "github.com/getopenpay/getopenpay-go"
)

func TestClientUsage(t *testing.T) {
    // Create a new API client
    cfg := getopenpay.NewConfiguration()
    cfg.Servers = getopenpay.ServerConfigurations{
        {URL: "https://connto.openpaystaging.com"},
    }
    cfg.DefaultHeader["Authorization"] = "Bearer ADD_SECRET_TOKEN"
    client := getopenpay.NewAPIClient(cfg)

    // Test a simple API call (e.g., list customers)
    ctx := context.Background()
    customerParams := getopenpay.CustomerQueryParams{
        PageNumber: getopenpay.PtrInt32(1),
        PageSize:   getopenpay.PtrInt32(10),
    }

    result, resp, err := client.CustomersAPI.ListCustomers(ctx).CustomerQueryParams(customerParams).Execute()
    if err != nil {
        t.Fatalf("Error calling ListCustomers: %v", err)
    }

    if resp.StatusCode != 200 {
        t.Errorf("Expected status code 200, got %d", resp.StatusCode)
    }

    if result.Data == nil || len(result.Data) == 0 {
        t.Logf("No customers found, but request was successful")
    } else {
        t.Logf("Found %d customers", len(result.Data))
        for i, customer := range result.Data {
            t.Logf("Customer %d:", i+1)
            t.Logf("  ID: %s", customer.GetId())
            t.Logf("  Email: %s", customer.GetEmail())
            t.Logf("  Created At: %s", customer.GetCreatedAt().Format("2006-01-02 15:04:05"))
            t.Logf("  Updated At: %s", customer.GetUpdatedAt().Format("2006-01-02 15:04:05"))
            // Add more fields as needed
        }
    }

    // Create a new product
    productParams := getopenpay.CreateProductRequest{
        Name:        "influencer pic",
        Description: "description",
        Features:    []string{"optional description of what this will include"},
        AccountSku:  *getopenpay.NewNullableString(getopenpay.PtrString("optional SKU")),
    }

    productResult, productResp, productErr := client.ProductsAPI.CreateProduct(ctx).CreateProductRequest(productParams).Execute()
    if productErr != nil {
        t.Fatalf("Error calling CreateProduct: %v", productErr)
    }

    if productResp.StatusCode != 201 {
        t.Errorf("Expected status code 201, got %d", productResp.StatusCode)
    }

    t.Logf("Created product:")
    t.Logf("  ID: %s", productResult.GetId())
    t.Logf("  Name: %s", productResult.GetName())
    t.Logf("  Description: %s", productResult.GetDescription())
    t.Logf("  Account SKU: %s", productResult.GetAccountSku())

    // Create a new price
    currency, _ := getopenpay.NewCurrencyEnumFromValue("usd")
    billingInterval := getopenpay.NewNullableCalendarIntervalEnum(getopenpay.CALENDARINTERVALENUM_MONTH.Ptr())
    billingIntervalCount := getopenpay.NewNullableInt32(getopenpay.PtrInt32(1))
    unitAmountAtom := getopenpay.NewNullableInt32(getopenpay.PtrInt32(3000))
    
    priceParams := getopenpay.CreatePriceRequest{
        IsActive:                   true,
        PricingModel:               getopenpay.PRICINGMODEL_STANDARD,
        PriceType:                  getopenpay.PRICETYPEENUM_RECURRING,
        TransformQuantityDivideBy:  getopenpay.PtrFloat32(1.0),
        Currency:                   currency,
        BillingInterval:            *billingInterval,
        BillingIntervalCount:       *billingIntervalCount,
        TrialPeriodDays:            getopenpay.PtrInt32(7),
        ProductId:                  productResult.GetId(),
        UnitAmountAtom:             *unitAmountAtom,
    }
    
  
    priceResult, priceResp, priceErr := client.PricesAPI.CreatePriceForProduct(ctx).CreatePriceRequest(priceParams).Execute()

    if priceErr != nil {
        t.Fatalf("Error calling CreatePrice: %v", priceErr)
    }

    if priceResp.StatusCode != 201 {
        t.Errorf("Expected status code 201, got %d", priceResp.StatusCode)
    }

    t.Logf("Created price:")
    t.Logf("  ID: %s", priceResult.GetId())
    t.Logf("  Product ID: %s", priceResult.GetProductId())
    t.Logf("  Unit Amount: %d", priceResult.GetUnitAmountAtom())
    t.Logf("  Currency: %s", priceResult.GetCurrency())

		// Create a new customer request
		customerRequest := getopenpay.NewCreateCustomerRequest("lanz+1@xyz.com")

		// Set optional fields
		customerRequest.SetFirstName("Lanz")
		customerRequest.SetLastName("Brown")
		// Add any additional fields you want to set here, for example:
		// customerRequest.SetLine1("123 Main St")
		// customerRequest.SetCity("Metropolis")

		customerResult, customerResp, customerErr := client.CustomersAPI.CreateCustomer(ctx).CreateCustomerRequest(*customerRequest).Execute()
		if customerErr != nil {
				t.Fatalf("Error calling CreateCustomer: %v", customerErr)
		}

		if customerResp.StatusCode != 201 {
				t.Errorf("Expected status code 201, got %d", customerResp.StatusCode)
		}

		t.Logf("Created customer:")
		t.Logf("  ID: %s", customerResult.GetId())
		t.Logf("  Email: %s", customerResult.GetEmail())
		t.Logf("  First Name: %s", customerResult.GetFirstName())
		t.Logf("  Last Name: %s", customerResult.GetLastName())

		// I've attached a test payment method to customer cus_stg_xtF4pntkBSa7yM3k
		//   you can attach a payment method too via payment links or openpayjs
		// Create Subscription
		subscriptionRequest := getopenpay.NewCreateSubscriptionRequest(
			"cus_stg_1dKX71Ozm5P0hfLN",
			[]getopenpay.SelectedPriceQuantity{
				{
					PriceId:  priceResult.GetId(),
					Quantity: 1,
				},
			},
			int32(priceResult.GetUnitAmountAtom()),
		)

		// Set additional fields
		subscriptionRequest.SetPaymentMethodId("cc_stg_lpBeHksJbLEnjLve")
		subscriptionRequest.SetCancelAtEnd(false)
		subscriptionRequest.SetTrialFromPrice(true)
		subscriptionRequest.SetCollectionMethod(getopenpay.COLLECTIONMETHODENUM_CHARGE_AUTOMATICALLY)
		subscriptionRequest.SetCustomFields(map[string]interface{}{
			"custom_field_1": "value1",
			"custom_field_2": "value2",
		})

		// Set optional fields
		subscriptionRequest.SetDescription("Test Subscription")
		subscriptionRequest.SetTrialPeriodDays(7)

		// Create the subscription
		subscriptionResult, subscriptionResp, subscriptionErr := client.SubscriptionsAPI.CreateSubscriptions(ctx).CreateSubscriptionRequest(*subscriptionRequest).Execute()

		if subscriptionErr != nil {
			t.Fatalf("Error calling CreateSubscriptions: %v", subscriptionErr)
		}

		if subscriptionResp.StatusCode != 200 {
			t.Errorf("Expected status code 200, got %d", subscriptionResp.StatusCode)
		}

		if len(subscriptionResult.Created) > 0 {
			createdSubscription := subscriptionResult.Created[0]
			t.Logf("Created subscription:")
			t.Logf("  ID: %s", createdSubscription.GetId())
			t.Logf("  Customer ID: %s", createdSubscription.GetCustomerId())
			t.Logf("  Status: %s", createdSubscription.GetStatus())
			t.Logf("  Current Period Start: %s", createdSubscription.GetCurrentPeriodStart().String())
			t.Logf("  Current Period End: %s", createdSubscription.GetCurrentPeriodEnd().String())
		} else {
			t.Errorf("No subscriptions were created")
		}

		// Parse invoice details
		if len(subscriptionResult.Invoices) > 0 {
			t.Logf("Created invoices:")
			for _, invoice := range subscriptionResult.Invoices {
				t.Logf("  Invoice ID: %s", invoice.GetId())
				t.Logf("  Amount Due: %d", invoice.GetTotalAmountAtom())
				t.Logf("  Status: %s", invoice.GetStatus())
				t.Logf("  Created At: %s", invoice.GetCreatedAt().String())
				t.Logf("  Due Date: %s", invoice.GetDueDate().String())
			}
		} else {
			t.Logf("No invoices were created with this subscription")
		}

}