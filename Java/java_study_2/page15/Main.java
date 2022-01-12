class Main {
  public static void main(String[] args) {
    // Tetapkan daftar bilangan ke variable numbers
    int[] numbers = {1, 4, 6, 9, 13, 16};
    
    int oddSum = 0;
    int evenSum = 0;
    
    // Kalkulasikan jumlah bilangan ganjil dan jumlah bilangan genap
    for (int number: numbers) {
      if (number % 2 == 0) {
        evenSum += number;
      } else {
        oddSum += number;
      }
    }
 
    // ---- Jika ingin menggunakan loop for standar ----
    // for (int i = 0; i < numbers.length; i++) {
    //   if (numbers[i] % 2 == 0) {
    //     evenSum += numbers[i];
    //   } else {
    //     oddSum += numbers[i];
    //   }
    // }

    System.out.println("Jumlah bilangan ganjil: " + oddSum);
    System.out.println("Jumlah bilangan genap: " + evenSum);
  }
}
