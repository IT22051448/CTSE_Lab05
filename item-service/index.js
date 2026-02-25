const express = require("express");

const app = express();
app.use(express.json());

let items = ["Book", "Laptop", "Phone"];

// Getting All Items

app.get("/items", (req, res) => {
  res.json(items);
});

// Adding Items

app.post("/items", (req, res) => {
  const item = req.body?.name;

  if (!item || typeof item !== "string" || !item.trim()) {
    return res.status(400).json({ error: "name is required" });
  }

  items.push(item.trim());
  res.status(201).send(`Item added: ${item.trim()}`);
});

// Getting Items by ID

app.get("/items/:id", (req, res) => {
  const id = Number(req.params.id);

  if (id < 0 || id >= items.length) {
    return res.status(404).end();
  }

  res.send(items[id]);
});

app.listen(8081, "0.0.0.0", () => {
  console.log("Item Service running on port 8081");
});