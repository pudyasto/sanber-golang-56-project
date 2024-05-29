## Dokumentasi API

Berikut adalah dokumentasi API:

1. Auth :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	POST	| 	{{base_url}}/auth/register	| 	body : { name, phone, username, password }	 | registrasi canvasser baru |
| 	POST	| 	{{base_url}}/auth/login	| 	body : { username, password }	 | login canvasser jika berhasil maka akan muncul token |


2. Master Data :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	1.	| 	Canvasser	|  |  |
| 	GET	| 	{{base_url}}/canvasser	| header : { token } | menampilkan semua data canvasser |
| 	POST	| 	{{base_url}}/canvasser	| header : { token } <br/> body : { name, phone, username, password }	 | menambah canvasser baru |
| 	PUT	| 	{{base_url}}/canvasser/:id	| 	header : { token } <br/> body : { name, phone, username, password }	 | mengubah data canvasser |
| 	DELETE	| 	{{base_url}}/canvasser/:id	| header : { token } | menghapus data canvasser |
| 	2.	| 	Item	|  |  |
| 	GET	| 	{{base_url}}/item	| header : { token } | menampilkan semua data item |
| 	POST	| 	{{base_url}}/item	| 	header : { token } <br/> body : { name, price }	 | menambah item baru |
| 	PUT	| 	{{base_url}}/item/:id	| 	header : { token } <br/> body : { name, price }	 | mengubah data item |
| 	DELETE	| 	{{base_url}}/item/:id	| header : { token } | menghapus data item |
| 	3.	| 	Customer	|  |  |
| 	GET	| 	{{base_url}}/customer	| header : { token } | menampilkan semua data customer |
| 	POST	| 	{{base_url}}/customer	| 	header : { token } <br/> body : { name, address, email }	 | menambah customer baru |
| 	PUT	| 	{{base_url}}/customer/:id	| 	header : { token } <br/> body : { name, address, email }	 | mengubah data customer |
| 	DELETE	| 	{{base_url}}/customer/:id	| header : { token } | menghapus data customer |
| 	4.	| 	Stock	|  |  |
| 	GET	| 	{{base_url}}/stock	| header : { token } | menampilkan semua data stock |
| 	POST	| 	{{base_url}}/stock	| 	header : { token } <br/> body : { item_id, canvasser_id, qty }	 | menambah stock baru |
| 	PUT	| 	{{base_url}}/stock/:item_id/:canvasser_id	| 	header : { token } <br/> body : { item_id, canvasser_id, qty }	 | mengubah data stock |
| 	DELETE	| 	{{base_url}}/stock/:item_id/:canvasser_id   | header : { token } | menghapus data stock |

3. Transaksi :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	1.	| 	Transaksi Sales Header	|  |  |
| 	GET	| 	{{base_url}}/trnsales	| header : { token } | menampilkan semua data transaksi sales header |
| 	POST	| 	{{base_url}}/trnsales	| 	header : { token } <br/> body : { customer_id, canvasser_id, date_sales, description }	 | menambah transaksi sales header baru |
| 	PUT	| 	{{base_url}}/trnsales/:id	| 	header : { token } <br/> body : { customer_id, canvasser_id, date_sales, description }	 | mengubah data transaksi sales header |
| 	DELETE	| 	{{base_url}}/trnsales/:id   | header : { token } | menghapus data transaksi sales header |
| 	2.	| 	Transaksi Sales Detail	|  |  |
| 	GET	| 	{{base_url}}/trnsalesdetail	| header : { token } | menampilkan semua data transaksi sales detail |
| 	POST	| 	{{base_url}}/trnsalesdetail	| 	header : { token } <br/> body : { trn_sales_id, item_id, qty, price }	 | menambah transaksi sales detail baru |
| 	PUT	| 	{{base_url}}/trnsalesdetail/:id	| 	header : { token } <br/> body : { trn_sales_id, item_id, qty, price }	 | mengubah data transaksi sales detail |
| 	DELETE	| 	{{base_url}}/trnsalesdetail/:id   | header : { token } | menghapus data transaksi sales detail |

4. Laporan :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	1.	| 	Report Canvasser Stock	|  |  |
| 	GET	| 	{{base_url}}/report/stock	| header : { token } | menampilkan semua data stok per canvasser |
| 	2.	| 	Transaksi Sales Detail	|  |  |
| 	GET	| 	{{base_url}}/report/sales	| header : { token } | menampilkan semua data transaksi per sales dengan total transaksi |