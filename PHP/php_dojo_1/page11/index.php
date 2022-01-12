<?php
$menus = array(
  array('name' => 'GULAI', 'price' => 9),
  array('name' => 'PASTA', 'price' => 12),
  array('name' => 'KOPI', 'price' => 6)
);

$totalPrice = 0;
$maxPrice = 0; // Variable untuk harga termahal
$maxPriceMenuName = ''; // Variable untuk menyimpan harga item termahal
foreach ($menus as $menu) {
  $name = $menu['name'];
  $price = $menu['price'];
  echo $name.' berharga $'.$price;
  echo '<br>';
  $totalPrice += $price;
  // Ketika lebih besar dari nilai yang disimpan di $maxPrice
  if ($price > $maxPrice) {
    // Perbarui $maxPrice dengan nilai tertinggi
    $maxPrice = $price;
    // Perbarui $maxPriceMenuName dengan nama item termahal yang baru
    $maxPriceMenuName = $name;
  }
}
echo 'Harga total adalah $'.$totalPrice;
echo '<br>';
echo 'Harga item termahal adalah'.$maxPriceMenuName.' dengan harga $'.$maxPrice;

?>