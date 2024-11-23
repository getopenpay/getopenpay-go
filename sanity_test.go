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
    cfg.DefaultHeader["Authorization"] = "Bearer st_stg_8i4HhDfwNh4mgpIfzUb9rA"
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
        Currency:                   currency, // Use the pointer created with NewCurrencyEnumFromValue
        BillingInterval:            *billingInterval, // Dereference the pointer
        BillingIntervalCount:       *billingIntervalCount, // Dereference the pointer
        TrialPeriodDays:            getopenpay.PtrInt32(7),
        ProductId:                  productResult.GetId(),
        UnitAmountAtom:             *unitAmountAtom, // Dereference the pointer
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
}