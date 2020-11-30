const express = require("express");
const app = express();
const delay = require("delay");

app.post("/api/games", async (req, res) => {
    await delay(3000)

    res.json({
        id: "04eb6ffa-c1f0-4a43-8d7a-3535e591d471"
    })
})

app.listen(5000, () => {
    console.log("Listening on port 5000");
});
