const express = require('express');
const mysql = require('mysql');
const session = require('express-session');
const app = express();

app.use(express.static('public'));
app.use(express.urlencoded({extended: false}));

const connection = mysql.createConnection({
  host: 'localhost',
  user: 'progate',
  password: 'password',	
  database: 'blog'
});

app.use(
  session({
    secret: 'my_secret_key',
    resave: false,
    saveUninitialized: false,
  })
);

app.use((req, res, next) => {
  if (req.session.userId === undefined) {
    res.locals.username = 'Tamu';
    res.locals.isLoggedIn = false;
  } else {
    res.locals.username = req.session.username;
    res.locals.isLoggedIn = true;
  }
  next();
});

app.get('/', (req, res) => {
  res.render('top.ejs');
});

app.get('/list', (req, res) => {
  connection.query(
    'SELECT * FROM articles',
    (error, results) => {
      res.render('list.ejs', { articles: results });
    }
  );
});

app.get('/article/:id', (req, res) => {
  const id = req.params.id;
  connection.query(
    'SELECT * FROM articles WHERE id = ?',
    [id],
    (error, results) => {
      res.render('article.ejs', { article: results[0] });
    }
  );
});

app.get('/signup', (req, res) => {
  res.render('signup.ejs');
});

app.post('/signup', 
  (req, res, next) => {
    console.log('Pemeriksaan nilai input kosong');
    // Deklarasikan tiga constant dan tetapkan nilai dari formulir pendaftaran pada masing-masingnya
    const username = req.body.username;
    const email = req.body.email;
    const password = req.body.password;
    
    // Definisikan sebuah array kosong bernama `errors`
    const errors = [];
    
    // Buatlah sebuah pernyataan if untuk masing-masing dari tiga constant untuk memeriksa apakah mereka merupakan sebuah string kosong
    // Jika memang kosong, tambahkan sebuah pesan error pada array `errors`
    if (username === '') {
      errors.push('Nama Pengguna kosong');
    }
    
    if (email === '') {
      errors.push('Email kosong');
    }
    
    if (password === '') {
      errors.push('Kata Sandi kosong');
    }
    
    // Gunakan console.log untuk mencetak array `errors` pada terminal
    console.log(errors);
    
    // Gunakan sebuah pernyataan if untuk memeriksa apakah jumlah element dalam array `errors` lebih dari nol
    // Juga tuliskan code berdasarkan hasil expression bersyarat
    if (errors.length > 0) {
      res.redirect('/signup');
    } else {
      next();
    }
    
    // Hapus 1 baris code dibawah ini
    
  },
  (req, res) => {
    console.log('Pendaftaran');
    const username = req.body.username;
    const email = req.body.email;
    const password = req.body.password;
    connection.query(
      'INSERT INTO users (username, email, password) VALUES (?, ?, ?)',
      [username, email, password],
      (error, results) => {
        req.session.userId = results.insertId;
        req.session.username = username;
        res.redirect('/list');
      }
    );
  }
);

app.get('/login', (req, res) => {
  res.render('login.ejs');
});

app.post('/login', (req, res) => {
  const email = req.body.email;
  connection.query(
    'SELECT * FROM users WHERE email = ?',
    [email],
    (error, results) => {
      if (results.length > 0) {
        if (req.body.password === results[0].password){
          req.session.userId = results[0].id;
          req.session.username = results[0].username;
          res.redirect('/list');
        } else {
          res.redirect('/login');
        }    
      } else {
        res.redirect('/login');
      }
    }
  );
});

app.get('/logout', (req, res) => {
  req.session.destroy((error) => {
    res.redirect('/list');
  });
});

app.listen(3000);
