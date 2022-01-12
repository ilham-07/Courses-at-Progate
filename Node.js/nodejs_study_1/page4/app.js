const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.render('hello.ejs');
});

// Tambahkan route untuk halaman top
app.get('/top', (req, res) => {
  res.render('top.ejs');
});

app.listen(3000);
