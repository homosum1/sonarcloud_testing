const express = require('express');
const cors = require('cors');

const port = 3000;

const app = express();

app.use(cors());
app.use(express.json());


let ItemList = [
    { Name: "apple", ID: 1, Price: 4.50, Quantity: 5, Icon: "🍎" },
    { Name: "banana", ID: 2, Price: 49.99, Quantity: 3, Icon: "🍌" },
    { Name: "pineapple", ID: 3, Price: 8.50, Quantity: 10, Icon: "🍍" }
];

app.get('/', (req, res) => {
  res.json(ItemList);
});

app.get('/getAll', (req, res) => {
    res.json(ItemList);
});


app.post('/purchase', (req, res) => {

    const { itemID, quantity } = req.body;

    if (itemID === undefined || quantity === undefined) {
        return res.status(400).send('Nie podano wymaganych danych');
    }

    const item = ItemList.find(i => i.ID === itemID);

    if (!item) {
        return res.status(404).send('Nie znaleziono przedmiotu');
    }

    if (quantity > item.Quantity) {
        return res.status(400).send('Na magazynie nie ma wystarczającej ilości towaru');
    }
    else {
        item.Quantity -= quantity;
    }

    res.send(`Zakupiono: ${quantity} przedmiotu: ${item.Name}`);
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
