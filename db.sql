create table users(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    photo varchar(255),
    username varchar(35),
    email varchar(100),
    nik varchar(16),
    full_name varchar(100),
    gender varchar(200),
    place_of_bird date,
    work varchar(100),
    citizenship varchar(100),
    is_verify_acount bool,
    is_verify_identity bool,
    password text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table city(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table province(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table warehouse(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
    city_id int(10),
    province_id int(10),
    cityId int(10),
    wide varchar(200),
    owner varchar(100),
    phone_number varchar(15),
    price bigint,
    status enum('tesedia', 'disewa', 'dalam pemeliharaan'),
    description text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
    
	FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (province_id) REFERENCES province(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table capacity(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table warehouse_picture(
	id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    picture varchar(255),
    warehouse_id int(12),
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE ON UPDATE CASCADE
);