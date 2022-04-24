CREATE USER FILLERAPPLICATION IDENTIFIED BY "123kb123svbhd123";
GRANT DBA TO FILLERAPPLICATION;

ALTER TABLE FILLERAPPLICATION.ENVIO ADD SUPPLEMENTAL LOG DATA (ALL) COLUMNS;
ALTER TABLE FILLERAPPLICATION.ESTADO_ENVIO ADD SUPPLEMENTAL LOG DATA (ALL) COLUMNS;

INSERT INTO COURIER (NOMBRE, COSTO_POR_PESO, REGION_REPARTE)
values ('Blue', 1000, 'c');
INSERT INTO COURIER (NOMBRE, COSTO_POR_PESO, REGION_REPARTE)
values ('Starken', 1500, 'b');
INSERT INTO COURIER (NOMBRE, COSTO_POR_PESO, REGION_REPARTE)
values ('ChileExpress', 1200, 'a');
COMMIT;