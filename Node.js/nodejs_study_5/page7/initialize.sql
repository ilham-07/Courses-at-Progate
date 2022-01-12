DROP DATABASE IF EXISTS blog;
CREATE DATABASE blog CHARACTER SET utf8 COLLATE utf8_general_ci;
USE blog;
CREATE TABLE articles (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(100),
  summary varchar(200),
  content text,
  PRIMARY KEY (id)
);
INSERT INTO articles(
  id, 
  title, 
  summary, 
  content
) VALUES (
  '1',
  'Mengapa disebut dengan Bug?',
  'Anda mungkin sering mendengar kata "bug" ketika sedang melakukan pembuatan sebuah program, kan? Saya telah melakukan riset tentang asal muasal dari kenapa kata "bug" digunakan.',
  'Pernahkan Anda berujar "Oh ada bug!" ketika sesuatu yang aneh terjadi ketika memainkan game atau menggunakan beberapa aplikasi? Saya jadi sangat penasaran perihal "bug" ini, jadi saya melakukan beberapa riset dan inilah yang kutemukan! \r\n\r\nPertama-tama, kata "bug" itu sendiri mengarah ke sebuah glitch yang terjadinya tidak disengaja. Bagaimana ceritanya kata "bug" menjadi rujukan ke arah sebuah glitch? Yaa, Saya pun sudah melakukan riset tentang hal ini! \r\n\r\nKetika melakukan riset tentang asal muasal kata "bug", membuat saya kembali ke masa sebelum komputer diciptakan. \r\nKomputer-komputer kala itu berukuran sebuah ruang kelas dan tampak seperti mesin pabrik. Suatu hari, komputer besar ini tiba-tiba berhenti bekerja dikarenakan sebuah glitch. Team yang menemukan komputer ini mencari penyebab glitch tersebut kesana-kemari dan menemukan sebuah ngengat tersangkut didalam komputer tersebut. Hanya karena satu ngengat kecil ini, mengakibatkan komputer sebesar itu berhenti bekerja. Sejak saat itu kata "bug" telah difahami sebagai maksud dari glitch-glitch yang disebabkan oleh sebuah ngengat (bug). \r\n\r\nSulit dibayangkan saat ini bahwa alasan dari terjadinya glitch pada komputer besaar itu dikarenakan oleh sebuah serangga yang benar-benar berada didalam komputer! Jika Anda memiliki seorang teman yang berkata, "Oh ada bug!", Anda bisa menceritakan padanya cerita ini!'
);
INSERT INTO articles(
  id, 
  title, 
  summary, 
  content
) VALUES (
  '2',
  'Apakah Anda tahu asal usul dari Cookies?',
  'Kita mungkin telah mengenal cookies sebagai sebuah makanan, tetapi tidak berhubungan sama sekali. Sangat penasaran dengan hal itu, sehingga saya mencarinya di internet dan hasilnya menarik sekali!',
  'Ketika pertama kali mendengar tentang Cookies dari Guru Domba, apakah Anda berfikir itu sejenis makanan yang bisa dimakan? Yaa, saya memang menyangka itu adalah sebuah makanan! Tetapi saya kecewa setelah mendengar penjelasan tentang Cookies sebab tidak ada hubungannya sama sekali dengan makanan. Disebabkan kekecewaan dan penasaran bagaimana ceritanya bisa dinamai Cookies, saya mencarinya di online! \r\n\r\nBanyak yang menyebutkan asal-muasal dari penamaan cookies, tetapi saya akan berbagi beberapa saja dari yang saya temukan.\r\n\r\n1. Cookies digunakan untuk menyimpan data, sehingga penamaan ini berasal dari kue yang digunakan sebagai makanan pengawet.\r\n\r\n2. Dinamai cookies sebab developer yang menemukanya itu menyukai kue (cookies).\r\n\r\n3. Dinamai sebab data yang berperilaku sama, disebut "Kue Ajaib" (Magic Cookie).\r\n\r\nSaya sendiri, menyukai poin pertama yang berhubungan dengan pengawetan makanan!  Anda akan menemukan sesuatu yang menarik ketika Anda mulai mencari tahu asal-muasal dari sebuah nama!'  
);
INSERT INTO articles(
  id, 
  title, 
  summary, 
  content
) VALUES (
  '3',
  'Tikus yang terhubung pada sebuah Komputer',
  'Pernah dengar bahwa setiap komputer ada tikus-nya? Beneran, ini kisah nyata!',
  'Hai, Saya punya mouse (tikus) yang tersambung ke komputer saya! Kamu harus punya juga!  \r\n\r\nMau sih.., tapi saya sudah punya hamster! \r\n\r\nItu maksudnya mouse jenis baru, kan? Apa itu hamster? \r\n\r\nHaha! Milikku tidak ada "buntut"-nya. Tikus tanpa "buntut" itu namanya hamster, kan? \r\n\r\nCukup bercandanya, mari kembali ke topik\r\n\r\nAkhir-akhir ini, teknologi sudah berkembang pesat dan nirkabel (wireless) sudah menjadi biasa (mainstream). Sebelum hadirnya teknologi wireless, kita terbiasa menyambungkan semuanya dengan perkabelan. Hal yang sama dengan mouse komputer. Disebutkan katanya nama "mouse" itu dikarenakan bentuknya dan caranya kabel memanjang darinya, yang memang membuatnya tampak seperti tikus yang sebenarnya. \r\n\r\nSekarang ini, meskipun masih bisa disebut "mouse", banyak tombol-tombol, tidak ada kabel-kabel, dan berbagai bentuk, yang membuatnya semakin tidak menyerupai seekor tikus!\r\n\r\nOh, iya! Satu hal lagi. Lakukan pencarian untuk "unit gerakan mouse" dan Anda mungkin terkejut dengan yang akan Anda temukan!'
);
INSERT INTO articles(
  id, 
  title, 
  summary, 
  content
) VALUES (
  '4',
  'Apa sih makna "My" dalam "MySQL"?',
  'Pernahkan Anda berfikir tentang "My" dalam MySQL merujuk pada "My" seperti dalam "Myself"? Cobalah mencarinya dan Anda akan terkejut menemukan bahwa sebenarnya itu hal yang samasekali berbeda!',
  'Kali ini, kita akan membicarakan perihal MySQL, yang juga digunakan dalam kursus Node.js.\r\nPada suatu hari, saya tiba-tiba menemukan diri ini mempertanyakan tentang "My" pada MySQL, jadi saya melakukan riset lagi!\r\nSebelum saya mencarinya, saya menduga itu artinya SQL "milik saya" karena dinamakan MySQL, tapi ternyata itu berarti hal yang jauh berbeda!\r\n\r\nJadi idzinkan saya membagikan hasil riset saya pada Anda!\r\n\r\nMySQL didirikan oleh Michael Widenius dan teman-temannya. Michael memiliki tiga orang anak, masing-masing bernama My, Max dan Maria.\r\n\r\nBeberapa dari kalian mungkin telah menduga bahwa ada anak bernama My, dan benar, MySQL dinamai dari nama anak pendirinya!\r\n\r\nNgomong-ngomong, Michael Widenius, pendiri MySQL, juga membuat database lain bernama MaxDB dan MariaDB. Senang juga mengetahui bahwa bisa menamai sistem Anda dengan nama anak Anda!'
);

CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  username varchar(20) DEFAULT NULL,
  email varchar(100) DEFAULT NULL,
  password varchar(60) DEFAULT NULL,
  PRIMARY KEY (id)
);

INSERT INTO users (username, email, password) VALUES ('Ninja Ken','ken@progate.com','ninja');
INSERT INTO users (username, email, password) VALUES ('Guru Domba','guru@progate.com','gurudomba');
INSERT INTO users (username, email, password) VALUES ('Baby Ben','baby@progate.com','baby');
INSERT INTO users (username, email, password) VALUES ('Birdie','birdie@progate.com','birdie');