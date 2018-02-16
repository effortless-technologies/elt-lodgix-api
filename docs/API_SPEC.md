# API Specification

## Response Models

### Property Response Model
```
{  
    name: <string>,
    lodgix_id: <int>,
    address: <Address>,
    bedrooms: <int>,
    bathrooms: <float64>,
    sleeps: <int>,
    rates: [<Rate>],
    taxes: [<TaxFeeDeposit>],
    fees: [<TaxFeeDeposit>],
    deposits: [<TaxFeeDeposit>],
    check_in: <string>,
    check_out: <string>,
    description: <string>,
    marketing_title: <string>,
    marketing_teaser: <string>,
    amenities: [<string>],
    type: <string>,
    images: [<Image>],
    reviews: [<Review]
}
```

### Address Response Model

```
{
    street_1: <string>,
    street_2: <string>,
    city: <string>,
    state: <string>,
    country: <string>,
    zip_code: <string>
}
```

### Rate Response Model

```
{
    name: <string>,
    start_date: <string>,
    end_date: <string>,
    weekday_rate: <string>,
    weekend_rate: <string>
}
```

### TaxFeeDeposit Response Model

```
{
    type: <string>,
    name: <string>,
    frequency: <string>,
    is_flat: <bool>,
    value: <float64>
}
```

### Image Response Model

```
{
    title: <string>,
    url: <string>,
    thumb: <string>,
    preview_url: <string>
}
```


### Review Response Model

```
{
    title: <string>,
    date: <string>,
    reviewer: <string>,
    stars: <int>,
    description: <string>
}
```

### Filter Response Model
```
{  
    name: <string>,
    options: [<string>]
}
```

## Requests

### GET/properties

This endpoint returns a list of properties

#### Optional Query Parameters

```
state=<state_code>
```

#### Response

```
{  
    count: <int>,
    filters: [<filter>]
    properties: [<property>]
}
```
