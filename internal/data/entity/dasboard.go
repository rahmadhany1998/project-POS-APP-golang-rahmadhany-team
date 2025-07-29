package entity

//m membuatkan data yang akan digunakan pada dashboard
type Dashboard struct {
	Model
	TotalOrders      int `json:"total_orders"`
	TotalProducts    int `json:"total_products"`
	TotalUsers       int `json:"total_users"`
	TotalOrderItems  int `json:"total_order_items"`
	TotalTables      int `json:"total_tables"`
	TotalCategories  int `json:"total_categories"`
	TotalUserAccess  int `json:"total_user_access"`
	TotalRevenue     int `json:"total_revenue"`
	TotalStock       int `json:"total_stock"`
	TotalAvailable   int `json:"total_available"`
	TotalRetailPrice int `json:"total_retail_price"`
	TotalQuantity    int `json:"total_quantity"`		
	TotalUnit        int `json:"total_unit"`
	TotalStatus      int `json:"total_status"`
}
// Dashboard menyimpan informasi ringkasan untuk tampilan dashboard
// yang mencakup total pesanan, produk, pengguna, item pesanan, tabel,
// kategori, akses pengguna, pendapatan, stok, ketersediaan, harga eceran,
// kuantitas, unit, dan status.
// Ini digunakan untuk memberikan gambaran umum tentang performa dan status sistem.	
// Setiap atribut menyimpan jumlah total dari masing-masing entitas yang relevan,
// memungkinkan analisis cepat dan pengambilan keputusan yang lebih baik.
// Model adalah struct dasar yang berisi ID dan timestamp untuk entitas ini.
// Atribut-atribut ini akan diisi dengan data yang relevan dari database				
// untuk memberikan informasi yang berguna bagi pengguna sistem.