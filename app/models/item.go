package models


type Item struct {
    Name     string  `json:"name"`
    ID       int     `json:"id"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
    Icon     string  `json:"icon"`
}

// must start with capitalized letter to be globally avaiable
var ItemList = []Item{
    {Name: "apple", ID: 1, Price: 4.50, Quantity: 5, Icon: "🍎"},
    {Name: "banana", ID: 2, Price: 49.99, Quantity: 3, Icon: "🍌"},
    {Name: "pineapple", ID: 3, Price: 8.50, Quantity: 10, Icon: "🍍"},
}