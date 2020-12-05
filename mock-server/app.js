const express = require("express");
const delay = require("delay");
const { v4: uuidv4 } = require('uuid');
const bodyParser = require('body-parser');

const app = express();
app.use(function (req, res, next) {
    console.log(`${req.method} ${req.url}`);
    next()
});

app.use(bodyParser.json())

const games = {};

app.post("/api/games", async (req, res) => {
    await delay(3000)

    const id = uuidv4();
    games[id] = {
        id: id,
        title: req.body.title
    }

    res.json({
        id: id
    })
});

app.get(`/api/games/:id`, async (req, res) => {
    await delay(3000)

    const game = games[req.params.id];

    if (game) {
        res.json(game)
    } else {
        res.status(404);
    }
});

app.listen(5000, () => {
    console.log("Listening on port 5000");
});
