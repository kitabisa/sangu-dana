# Dana Sangu

## Usage blueprint

1. There is a type named `Client` (`dana.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used. The gateway type need a Client instance.
3. All Header field is handled by this library
4. There's also VerifySignature to verify whether the signature response/request is valid.
5. Replace `.sample` files to your own credential.

## Example

```go
    danaClient := dana.NewClient()
    danaClient.BaseUrl = "DANA_BASE_URL",
    ---
    ---

    coreGateway := dana.CoreGateway{
        Client: danaClient,
    }

    body := &dana.RequestBody{
        Order: {},
        MerchantId: "MERCHANT_ID",
        ---
        ---
        ---
    }

    res, _ := coreGateway.Order(req)
```
