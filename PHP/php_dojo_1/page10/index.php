<?php
$menus = array(
  array('name' => 'GULAI', 'price' => 9),
  array('name' => 'PASTA', 'price' => 12),
  array('name' => 'KOPI', 'price' => 6)
);

$totalPrice = 0; // Variable untuk menyimpan harga total
foreach ($menus as $menu) {
  $price = $menu['price'];
  echo $menu['name'].' berharga $'.$price;
  echo '<br>';
  // Tambahkan setiap harga ke $totalPrice
  $totalPrice += $price;
}
echo 'Harga total adalah $'.$totalPrice;

?>