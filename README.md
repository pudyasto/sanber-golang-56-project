## Dokumentasi API

Berikut adalah dokumentasi API:

1. Auth :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	POST	| 	{{base_url}}/auth/register	| 	name, phone, username, password	 | registrasi canvasser baru |
| 	POST	| 	{{base_url}}/auth/login	| 	username, password	 | login canvasser jika berhasil maka akan muncul token |


2. Master Data :

| 	Method	 | 	Path	 | 	Parameter	 | 	Keterangan	 | 
| 	:-----:	 | 	:-----:	 | 	:-----:	 | 	:-----:	 |  
| 	1.	| 	Canvasser	|  |  |
| 	GET	| 	{{base_url}}/canvasser	|  | menampilkan semua data canvasser |
| 	POST	| 	{{base_url}}/canvasser	| 	name, phone, username, password	 | menambah canvasser baru |
| 	PUT	| 	{{base_url}}/canvasser/:id	| 	name, phone, username, password	 | mengubah data canvasser |
| 	DELETE	| 	{{base_url}}/canvasser/:id	|  | menghapus data canvasser |
| 	2.	| 	Item	|  |  |
| 	GET	| 	{{base_url}}/item	|  | menampilkan semua data item |
| 	POST	| 	{{base_url}}/item	| 	name, price	 | menambah item baru |
| 	PUT	| 	{{base_url}}/item/:id	| 	name, price	 | mengubah data item |
| 	DELETE	| 	{{base_url}}/item/:id	|  | menghapus data item |
| 	3.	| 	Customer	|  |  |
| 	GET	| 	{{base_url}}/customer	|  | menampilkan semua data customer |
| 	POST	| 	{{base_url}}/customer	| 	name, address, email	 | menambah customer baru |
| 	PUT	| 	{{base_url}}/customer/:id	| 	name, address, email	 | mengubah data customer |
| 	DELETE	| 	{{base_url}}/customer/:id	|  | menghapus data customer |
| 	4.	| 	Stock	|  |  |
| 	GET	| 	{{base_url}}/stock	|  | menampilkan semua data stock |
| 	POST	| 	{{base_url}}/stock	| 	item_id, canvasser_id, qty	 | menambah stock baru |
| 	PUT	| 	{{base_url}}/stock/:item_id/:canvasser_id	| 	item_id, canvasser_id, qty	 | mengubah data stock |
| 	DELETE	| 	{{base_url}}/stock/:item_id/:canvasser_id   |  | menghapus data stock |