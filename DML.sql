-- INSERT service
"INSERT INTO mst_service (serviceid,nama,satuan,harga,indate,inby) VALUES (1,'CUCI + SETRIKA','PCS',7000,'2024-01-12 21:00:00','rofi');"

-- UPDATE service
"UPDATE mst_service SET nama = 'CUCI KERING + GOSOK',satuan = 'PCS', harga = 8000, updateddate = '2024-01-12 21:00:00', updatedby = 'rofi' WHERE serviceid = 1;"

-- DELETE service
"DELETE FROM mst_service WHERE serviceid = 1;"

-- INSERT satuan
"INSERT INTO mst_satuan (satuanid,nama,indate,inby) VALUES ('PCS','PCS','2024-01-12 21:00:00','rofi')"
-- UPDATE satuan
"UPDATE mst_satuan SET nama = 'PCS', UpdatedDate = '2024-01-12 21:00:00', UpdatedBy = 'rofi' WHERE satuanid = 'PCS'; "
-- DELETE satuan
"DELETE FROM mst_satuan WHERE satuanid = 'PCS'; "

-- INSERT Transaksi Header
"INSERT INTO trx_laundry (trsid,customer,contact,totalqty,totaltagihan,indate,inby) VALUES ('TRX202401200001','Dio',8765431228,4,32000,'2024-01-12 21:00:00','rofi')"
-- INSERT Transaksi Detail
"INSERT INTO trx_laundry_detail (trsid,serviceid,harga,qty) VALUES (1,1,8000,2);"