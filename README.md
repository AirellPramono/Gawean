# Gawean

Project yang berkonsep untuk memfalisitasikan para penyedia jasa dalam operasi atau gawean (kerjaan) mereka. 

Tentunya para penyedia jasa, contohnya dokter umum, tukang rawat taman, tukang servis AC, reparasi mesin-mesin secara berkala, sangat berhubungan dengan client mereka. Client mereka dapat berjumlah sangat banyak dalam rentang waktu yang singkat, dan juga masing-masing memiliki problematika atau masalah berbeda-beda. Bayangkan miskomunikasi yang bisa terjadi bila mereka lupa akan klien mereka.

Gawean dapat membantu proses-proses ini, mendatakan informasi-informasi yang bejibun atau bertaburan, dan meningkatkan kinerja serta akurasi para penyedia servis tersebut. Gawean menyediakan cara untuk mendata detail klien, alamat klien terkait, kategori servis, dan gawean atau urusan yang terkait dengan klien tersebut.


Use BasicAuth with username Arif & password Bugaresa or username Saya & password Ganteng


ROUTES


/clients 


GET /= Get all clients data

GET /:id = Get specific client by ID

GET /city/:id = Get specific client by which city he lives

POST /= Basic posting of 1 client data using JSON format

PUT /:id = Update client data by ID

DELETE /:id = Delete client data by ID

DELETE /:id/:name = Delete client data by ID and also it’s name


/address 


GET / = Get all address data

GET /:id = Get address specifically by ID

POST / = Basic posting of 1 address data using JSON format

PUT /:id = Update address by ID

DELETE /:id = Delete address data by ID


/category 


GET / = Get all category data

POST / = Basic posting of 1 category data using JSON format

PUT /:id = Update category by ID

DELETE /:id = Delete category data by ID


/gawean 


GET / = Get all gawean data

GET /client/:id = Get gawean data by specific client

GET /category/:id = Get gawean data by their category

POST / = Basic posting of 1 gawean data using JSON format

PUT /:id = Update gawean by ID

DELETE /:id = Delete gawean data by ID


/data


Super functionality

Fill the URL parameter with d (days) in dd format, m (month) in mm format

And y (year) in yyyy format

And it will display the list of gaweans with next appointment on that particular date if exist

Example 

/data?d=01&m=04&y=2023 => Show gaweans with next appointment at 1st of April 2023

So workers can use these functionality, especially the last one, to check on their dealings with the necessary clients



