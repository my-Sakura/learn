const express = require("express");

const app = express();
const port = 1010;

app.get("/", (req, res) => {
    let data = req.query.data;
    res.send(data);
});

app.listen(port, () => {
    console.log(`listen ${port}`);
});