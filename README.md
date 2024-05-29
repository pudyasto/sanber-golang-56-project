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