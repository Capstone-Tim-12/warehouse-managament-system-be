create table provinces(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255)
);

create table regencies(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    province_id int(11),
    name varchar(255),
    FOREIGN KEY (province_id) REFERENCES provinces(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table districts(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    regency_id int(11),
    name varchar(255),
    FOREIGN KEY (regency_id) REFERENCES regencies(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table villages(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    district_id int(11),
    name varchar(255),
    FOREIGN KEY (district_id) REFERENCES districts(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table users(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username varchar(35),
    email varchar(100),
    is_verify_acount bool,
    is_verify_identity bool,
    password text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table user_details(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT
    address longtext,
    country longtext,
    photo longtext,
    nik longtext,
    full_name longtext,
    gender longtext,
    place_of_birth longtext,
    date_birth datetime(3) DEFAULT NULL,
    work longtext,
    citizenship longtext,
    user_id bigint DEFAULT NULL,
    province_id varchar(12) DEFAULT NULL,
    regency_id varchar(12) DEFAULT NULL,
    district_id varchar(12) DEFAULT NULL,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (province_id) REFERENCES provinces(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (regency_id) REFERENCES regencies(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (district_id) REFERENCES districts(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table warehouse(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
    province_id int(10),
    regency_id int(10),
    wide varchar(200),
    owner varchar(100),
    phone_number varchar(15),
    price bigint,
    status enum('tesedia', 'disewa', 'dalam pemeliharaan'),
    description text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
	FOREIGN KEY (province_id) REFERENCES province(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (regency_id) REFERENCES regency(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table capacity(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table warehouse_picture(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    picture varchar(255),
    warehouse_id int(12),
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE ON UPDATE CASCADE
);